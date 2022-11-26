package storage

import (
	"errors"
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"github.com/mailpiggy/MailPiggy/logger"
	"io/ioutil"
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
		dir, err := os.MkdirTemp("", "mailpiggy_")
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

func (directory *Directory) Store(room Room, message *dto.Message) (dto.MessageID, error) {
	b, err := ioutil.ReadAll(message.Raw.Bytes())
	if err != nil {
		return "", err
	}
	if perRoomLimit > 0 && perRoomLimit <= directory.Count(room) {
		directory.DeleteRoom(room)
	}

	path := filepath.Join(directory.RoomDirectory(room), string(message.ID))
	err = ioutil.WriteFile(path, b, 0660)
	logManager().Debug(fmt.Sprintf("New message saved at %s", path))

	return message.ID, err
}

func (directory *Directory) List(room Room, offset, limit int) ([]dto.Message, error) {
	if offset < 0 || limit < 0 {
		return nil, errors.New("offset and limit should be >= 0")
	}

	dir, err := os.Open(directory.RoomDirectory(room))
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	n, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	sort.Slice(n, func(i, j int) bool {
		return n[i].ModTime().After(n[j].ModTime())
	})

	return directory.parseList(room, n, offset, limit)
}

func (directory *Directory) Search(room Room, query SearchQuery, offset, limit int) ([]dto.Message, int, error) {
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
			msg, err := directory.Load(room, dto.MessageID(unfilteredN[i].Name()))
			if err != nil {
				continue
			}
			for criteria, queryValue := range query {
				switch criteria {
				case "to":
					for _, t := range msg.To {
						if strings.Contains(strings.ToLower(t.Mailbox+"@"+t.Domain), queryValue) {
							n = append(n, unfilteredN[i])
							continue filtrationLoop
						}
					}
				case "from":
					if strings.Contains(strings.ToLower(msg.From.Mailbox+"@"+msg.From.Domain), queryValue) {
						n = append(n, unfilteredN[i])
						continue filtrationLoop
					}
				case "content":
					if strings.Contains(strings.ToLower(msg.Raw.Data), queryValue) {
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

func (directory *Directory) parseList(room string, n []os.FileInfo, offset, limit int) ([]dto.Message, error) {
	messages := make([]dto.Message, 0)

	if offset >= len(n) {
		return messages, nil
	}

	endIndex := len(n)
	if offset+limit < len(n) {
		endIndex = offset + limit
	}
	n = n[offset:endIndex]

	for _, fileinfo := range n {
		b, err := ioutil.ReadFile(filepath.Join(directory.RoomDirectory(room), fileinfo.Name()))
		if err != nil {
			return nil, err
		}
		msg := dto.FromBytes(b)
		m := *msg.Parse()
		m.ID = dto.MessageID(fileinfo.Name())
		m.Created = fileinfo.ModTime()
		messages = append(messages, m)
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

func (directory *Directory) Delete(room Room, messageId dto.MessageID) error {
	return os.Remove(filepath.Join(directory.RoomDirectory(room), string(messageId)))
}

func (directory *Directory) Load(room Room, messageId dto.MessageID) (*dto.Message, error) {
	b, err := ioutil.ReadFile(filepath.Join(directory.RoomDirectory(room), string(messageId)))
	if err != nil {
		return nil, err
	}
	m := dto.FromBytes(b).Parse()
	m.ID = messageId
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
