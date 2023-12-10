package smtpMessage

import (
	"bytes"
	"github.com/mailhedgehog/logger"
	"io"
	"strings"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("dto.smtpMessage")
	}
	return configuredLogger
}

// SMTPMessage represents a raw SMTP message
type SMTPMessage struct {
	From string
	To   []string
	Data string
	Helo string
}

// ToReader returns an io.Reader containing the raw message data
func (message *SMTPMessage) ToReader() io.Reader {
	var bufferReader = new(bytes.Buffer)

	if message != nil {
		bufferReader.WriteString("HELO:<" + message.Helo + ">\r\n")
		bufferReader.WriteString("FROM:<" + message.From + ">\r\n")
		for _, to := range message.To {
			bufferReader.WriteString("TO:<" + to + ">\r\n")
		}
		bufferReader.WriteString("\r\n")
		bufferReader.WriteString(message.Data)
	}

	return bufferReader
}

// FromString returns a SMTPMessage from raw message bytes (as output by SMTPMessage.ToReader())
func FromString(messageString string) *SMTPMessage {
	msg := &SMTPMessage{}
	var headerDone bool
	for _, l := range strings.Split(messageString, "\n") {
		if !headerDone {
			if strings.HasPrefix(l, "HELO:<") {
				l = strings.TrimPrefix(l, "HELO:<")
				l = strings.TrimSuffix(l, "\r")
				l = strings.TrimSuffix(l, ">")
				msg.Helo = l
				continue
			}
			if strings.HasPrefix(l, "FROM:<") {
				l = strings.TrimPrefix(l, "FROM:<")
				l = strings.TrimSuffix(l, "\r")
				l = strings.TrimSuffix(l, ">")
				msg.From = l
				continue
			}
			if strings.HasPrefix(l, "TO:<") {
				l = strings.TrimPrefix(l, "TO:<")
				l = strings.TrimSuffix(l, "\r")
				l = strings.TrimSuffix(l, ">")
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
