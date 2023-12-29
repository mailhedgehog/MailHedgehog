package smtp

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"github.com/mailhedgehog/smtpServerProtocol"
	"io"
	"strings"
)

// session between "client" and "sever"
type session struct {
	context  *serverContext.Context
	protocol Protocol
	reader   io.Reader
	writer   io.Writer

	receivedLine string

	loggedUsername string
}

// start session by sending to "client" welcome message
func (session *session) start(identification string) {
	logManager().Debug("Session started")
	session.writeReply(session.protocol.SayWelcome(identification))
}

// readAndWriteReply function reads client message
// and based on it writeReply.
// Returns boolean flag is server should wait next
// message from client.
func (session *session) readAndWriteReply() bool {
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
	for strings.Contains(session.receivedLine, smtpServerProtocol.CommandEndSymbol) {
		parts := strings.SplitN(session.receivedLine, smtpServerProtocol.CommandEndSymbol, 2)
		session.receivedLine = parts[1]

		reply := session.protocol.HandleReceivedLine(parts[0])
		if reply != nil {
			session.writeReply(reply)
			if reply.Status == smtpServerProtocol.CODE_SERVICE_CLOSING {
				logManager().Debug("Server connection closing")
				return false
			}
		}
	}

	return true
}

// writeReply a reply to the specific connection
func (session *session) writeReply(reply *smtpServerProtocol.Reply) {
	lines := reply.FormattedLines()
	replacer := strings.NewReplacer("\n", "\\n", "\r", "\\r")
	for _, l := range lines {
		logManager().Info(fmt.Sprintf("SERVER -> CLIENT: %s", replacer.Replace(l)))

		_, err := session.writer.Write([]byte(l))
		if err != nil {
			logManager().Error(fmt.Sprintf("Write session error: %s", err.Error()))
		}
	}
}
