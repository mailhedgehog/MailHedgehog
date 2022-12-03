package dto

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/mailpiggy/MailPiggy/logger"
	"io"
	"strings"
	"time"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("dto")
	}
	return configuredLogger
}

// MessageID represents the ID of an SMTP message including the hostname part
type MessageID string

// Message represents a parsed SMTP message
type Message struct {
	ID      MessageID
	From    *Path
	To      []*Path
	Content *Content
	Created time.Time
	Raw     *SMTPMessage
}

// Path represents an SMTP forward-path or return-path
type Path struct {
	Relays  []string
	Mailbox string
	Domain  string
	Params  string
}

// Content represents the body content of an SMTP message
type Content struct {
	Headers ContentHeaders
	Body    string
	Size    int
	MIME    *MIMEBody
}

// SMTPMessage represents a raw SMTP message
type SMTPMessage struct {
	From string
	To   []string
	Data string
	Helo string
}

// MIMEBody represents a collection of MIME parts
type MIMEBody struct {
	Parts []*Content
}

// Parse converts a raw SMTP message to a parsed MIME message
func (message *SMTPMessage) Parse() *Message {
	var arr []*Path
	for _, path := range message.To {
		arr = append(arr, PathFromString(path))
	}

	id := NewMessageID()
	msg := &Message{
		ID:      id,
		From:    PathFromString(message.From),
		To:      arr,
		Content: ContentFromString(message.Data),
		Created: time.Now(),
		Raw:     message,
	}

	_, err := msg.Content.Headers.GetOne("Message-ID")
	if err != nil {
		msg.Content.Headers.Set("Message-ID", []string{string(id)})
	}

	received, _ := msg.Content.Headers.Get("Received")
	msg.Content.Headers.Set("Received", append(received, "from "+message.Helo+" by (MailHedgehog)\r\n          id "+string(id)+"; "+time.Now().Format(time.RFC1123Z)))

	returnPath, _ := msg.Content.Headers.Get("Return-Path")
	msg.Content.Headers.Set("Return-Path", append(returnPath, "<"+message.From+">"))

	return msg
}

// Bytes returns an io.Reader containing the raw message data
func (message *SMTPMessage) Bytes() io.Reader {
	var bufferReader = new(bytes.Buffer)

	if message != nil {
		bufferReader.WriteString("HELO:<" + message.Helo + ">\r\n")
		bufferReader.WriteString("FROM:<" + message.From + ">\r\n")
		for _, t := range message.To {
			bufferReader.WriteString("TO:<" + t + ">\r\n")
		}
		bufferReader.WriteString("\r\n")
		bufferReader.WriteString(message.Data)
	}

	return bufferReader
}

// FromBytes returns a SMTPMessage from raw message bytes (as output by SMTPMessage.Bytes())
func FromBytes(b []byte) *SMTPMessage {
	msg := &SMTPMessage{}
	var headerDone bool
	for _, l := range strings.Split(string(b), "\n") {
		if !headerDone {
			if strings.HasPrefix(l, "HELO:<") {
				l = strings.TrimPrefix(l, "HELO:<")
				l = strings.TrimSuffix(l, ">\r")
				msg.Helo = l
				continue
			}
			if strings.HasPrefix(l, "FROM:<") {
				l = strings.TrimPrefix(l, "FROM:<")
				l = strings.TrimSuffix(l, ">\r")
				msg.From = l
				continue
			}
			if strings.HasPrefix(l, "TO:<") {
				l = strings.TrimPrefix(l, "TO:<")
				l = strings.TrimSuffix(l, ">\r")
				msg.To = append(msg.To, l)
				continue
			}
			if strings.TrimSpace(l) == "" {
				headerDone = true
				continue
			}
		}
		msg.Data += l + "\n"
	}
	return msg
}

// PathFromString parses a forward-path or reverse-path into its parts
func PathFromString(path string) *Path {
	var relays []string
	email := path
	if strings.Contains(path, ":") {
		x := strings.SplitN(path, ":", 2)
		r, e := x[0], x[1]
		email = e
		relays = strings.Split(r, ",")
	}
	mailbox, domain := "", ""
	if strings.Contains(email, "@") {
		x := strings.SplitN(email, "@", 2)
		mailbox, domain = x[0], x[1]
	} else {
		mailbox = email
	}

	return &Path{
		Relays:  relays,
		Mailbox: mailbox,
		Domain:  domain,
		Params:  "", // FIXME?
	}
}

// ContentFromString parses SMTP content into separate headers and body
func ContentFromString(data string) *Content {
	// logManager().Debug(fmt.Sprintf("Parsing Content from string: '%s'", data))

	x := strings.SplitN(data, "\r\n\r\n", 2)
	h := make(map[string][]string, 0)

	// FIXME this fails if the message content has no headers - specifically,
	// if it doesn't contain \r\n\r\n

	if len(x) == 2 {
		headers, body := x[0], x[1]
		hdrs := strings.Split(headers, "\r\n")
		var lastHdr = ""
		for _, hdr := range hdrs {
			if lastHdr != "" && (strings.HasPrefix(hdr, " ") || strings.HasPrefix(hdr, "\t")) {
				h[lastHdr][len(h[lastHdr])-1] = h[lastHdr][len(h[lastHdr])-1] + hdr
			} else if strings.Contains(hdr, ": ") {
				y := strings.SplitN(hdr, ": ", 2)
				key, value := y[0], y[1]
				// TODO: multiple header fields
				h[key] = []string{value}
				lastHdr = key
			} else if len(hdr) > 0 {
				logManager().Error(fmt.Sprintf("Found invalid header: '%s'", hdr))
			}
		}
		return &Content{
			Size:    len(data),
			Headers: ContentHeaders{h},
			Body:    body,
		}
	}
	return &Content{
		Size:    len(data),
		Headers: ContentHeaders{h},
		Body:    x[0],
	}
}

// NewMessageID generates a new message ID
func NewMessageID() MessageID {
	return MessageID(uuid.New().String())
}
