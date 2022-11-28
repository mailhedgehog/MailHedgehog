package smtp

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/serverContext"
	"github.com/mailpiggy/MailPiggy/smtpServer"
	"io"
	"strings"
)

type Session struct {
	context  *serverContext.Context
	protocol *smtpServer.Protocol
	reader   io.Reader
	writer   io.Writer

	receivedLine string

	LoggedUsername string
}

func (session *Session) Start(identification string) {
	logManager().Debug("Session started")
	session.Write(session.protocol.SayHi(identification))
}

func (session *Session) IsConversation() bool {
	buffer := make([]byte, 1024)
	numberOfBytes, err := session.reader.Read(buffer)
	if numberOfBytes == 0 {
		logManager().Info("Connection closed by remote host")
		return false
	}
	if err != nil {
		logManager().Info(fmt.Sprintf("Error reading from socket: %s", err))
		return false
	}

	receivedText := string(buffer[0:numberOfBytes])
	logManager().Info(fmt.Sprintf(
		"CLIENT -> SERVER (%d bytes): %s",
		numberOfBytes,
		strings.ReplaceAll(strings.ReplaceAll(receivedText, "\n", "\\n"), "\r", "\\r"),
	))

	session.receivedLine += receivedText
	for strings.Contains(session.receivedLine, smtpServer.COMMAND_END_SYMBOL) {
		parts := strings.SplitN(session.receivedLine, smtpServer.COMMAND_END_SYMBOL, 2)
		session.receivedLine = parts[1]

		reply := session.protocol.HandleReceivedLine(parts[0])
		if reply != nil {
			session.Write(reply)
			if reply.Status == smtpServer.CODE_SERVICE_CLOSING {
				logManager().Debug("Server connection closing")
				return false
			}
		}
	}

	return true
}

// Write writes a reply to the specific connection
func (session *Session) Write(reply *smtpServer.Reply) {
	lines := reply.Lines()
	for _, l := range lines {
		logManager().Info(fmt.Sprintf(
			"SERVER -> CLIENT: %s",
			strings.ReplaceAll(strings.ReplaceAll(l, "\n", "\\n"), "\r", "\\r"),
		))
		_, err := session.writer.Write([]byte(l))
		if err != nil {
			logManager().Error(fmt.Sprintf("Write session error: %s", err.Error()))
		}
	}
}
