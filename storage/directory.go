package storage

import (
	"errors"
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Directory store messages in local directory
type Directory struct {
	Path string
}

func CreateDirectoryStorage(path string) *Directory {
	if len(path) <= 0 {
		dir, err := os.MkdirTemp("", "mailhedgehog_")
		logger.PanicIfError(err)
		path = dir
	}
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, 0770)
		logger.PanicIfError(err)
	}
	logManager().Debug(fmt.Sprintf("Mail storage directory path is '%s'", path))

	return &Directory{path}
}

// RoomDirectory create correct path to specific `room`
func (directory *Directory) RoomDirectory(room Room) string {
	if len(room) <= 0 {
		room = "_default"
	}
	path := filepath.Join(directory.Path, room)
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, 0770)
		logger.PanicIfError(err)
	}

	return path
}

func (directory *Directory) Store(room Room, message *smtpMessage.SMTPMail) (smtpMessage.MessageID, error) {
	b, err := io.ReadAll(message.Origin.ToReader())
	if err != nil {
		return "", err
	}
	if perRoomLimit > 0 && perRoomLimit <= directory.Count(room) {
		directory.DeleteRoom(room)
	}

	path := filepath.Join(directory.RoomDirectory(room), string(message.ID))
	err = os.WriteFile(path, b, 0660)
	logManager().Debug(fmt.Sprintf("New message saved at %s", path))

	return message.ID, err
}

func (directory *Directory) List(room Room, query SearchQuery, offset, limit int) ([]smtpMessage.SMTPMail, int, error) {
	if offset < 0 || limit < 0 {
		return nil, 0, errors.New("offset and limit should be >= 0")
	}

	dir, err := os.Open(directory.RoomDirectory(room))
	if err != nil {
		return nil, 0, err
	}
	defer dir.Close()

	unfilteredN, err := dir.Readdir(0)
	if err != nil {
		return nil, 0, err
	}

	sort.Slice(unfilteredN, func(i, j int) bool {
		return unfilteredN[i].ModTime().After(unfilteredN[j].ModTime())
	})

	var n []os.FileInfo

	if len(query) > 0 {
	filtrationLoop:
		for i := range unfilteredN {
			msg, err := directory.Load(room, smtpMessage.MessageID(unfilteredN[i].Name()))
			if err != nil {
				continue
			}
			for criteria, queryValue := range query {
				queryValue = strings.ToLower(queryValue)
				switch criteria {
				case "to":
					for _, t := range msg.To {
						if strings.Contains(strings.ToLower(t.Address()), queryValue) {
							n = append(n, unfilteredN[i])
							continue filtrationLoop
						}
					}
				case "from":
					if strings.Contains(strings.ToLower(msg.From.Address()), queryValue) {
						n = append(n, unfilteredN[i])
						continue filtrationLoop
					}
				case "content":
					if strings.Contains(strings.ToLower(msg.Origin.Data), queryValue) {
						n = append(n, unfilteredN[i])
						continue filtrationLoop
					}
				}
			}
		}
	} else {
		n = unfilteredN
	}

	messages, err := directory.parseList(room, n, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return messages, len(n), nil
}

func (directory *Directory) parseList(room string, n []os.FileInfo, offset, limit int) ([]smtpMessage.SMTPMail, error) {
	messages := make([]smtpMessage.SMTPMail, 0)

	if offset >= len(n) {
		return messages, nil
	}

	endIndex := len(n)
	if offset+limit < len(n) {
		endIndex = offset + limit
	}
	n = n[offset:endIndex]

	for _, fileinfo := range n {
		b, err := os.ReadFile(filepath.Join(directory.RoomDirectory(room), fileinfo.Name()))
		if err != nil {
			logManager().Error(err.Error())
			continue
		}
		msg := smtpMessage.FromString(string(b))
		smtpMail, err := msg.ToSMTPMail(smtpMessage.MessageID(fileinfo.Name()))
		if err != nil {
			logManager().Error(err.Error())
			continue
		}
		messages = append(messages, *smtpMail)
	}

	logManager().Debug(fmt.Sprintf("Found %d messages", len(messages)))

	return messages, nil
}

func (directory *Directory) Count(room Room) int {
	dir, err := os.Open(directory.RoomDirectory(room))
	logger.PanicIfError(err)
	defer dir.Close()
	n, _ := dir.Readdirnames(0)
	return len(n)
}

func (directory *Directory) Delete(room Room, messageId smtpMessage.MessageID) error {
	return os.Remove(filepath.Join(directory.RoomDirectory(room), string(messageId)))
}

func (directory *Directory) Load(room Room, messageId smtpMessage.MessageID) (*smtpMessage.SMTPMail, error) {
	b, err := os.ReadFile(filepath.Join(directory.RoomDirectory(room), string(messageId)))
	if err != nil {
		return nil, err
	}

	m, err := smtpMessage.FromString(string(b)).ToSMTPMail(messageId)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (directory *Directory) RoomsList(offset, limit int) ([]Room, error) {
	if offset < 0 || limit < 0 {
		return nil, errors.New("offset and limit should be >= 0")
	}

	dir, err := os.Open(directory.Path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	n, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	rooms := make([]Room, 0)

	sort.Slice(n, func(i, j int) bool {
		return n[i].Name() < n[j].Name()
	})

	if offset >= len(n) {
		return rooms, nil
	}

	endIndex := len(n)
	if offset+limit < len(n) {
		endIndex = offset + limit
	}

	n = n[offset:endIndex]

	for _, fileinfo := range n {
		rooms = append(rooms, fileinfo.Name())
	}

	logManager().Debug(fmt.Sprintf("Found %d rooms", len(rooms)))

	return rooms, nil
}

func (directory *Directory) RoomsCount() int {
	dir, err := os.Open(directory.Path)
	logger.PanicIfError(err)
	defer dir.Close()
	n, _ := dir.Readdirnames(0)
	return len(n)
}

func (directory *Directory) DeleteRoom(room Room) error {
	err := os.RemoveAll(directory.RoomDirectory(room))
	if err != nil {
		return err
	}
	return nil
}
