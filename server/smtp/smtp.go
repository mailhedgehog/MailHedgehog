package smtp

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"github.com/mailhedgehog/MailHedgehog/server/websocket"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"github.com/mailhedgehog/MailHedgehog/smtpServer"
	"io"
	"net"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server.smtp")
	}
	return configuredLogger
}

const (
	AUTH_MECHANISM_PLAIN = "PLAIN"
)

func Listen(context *serverContext.Context, exitCh chan int) {
	logManager().Debug(fmt.Sprintf("SMTP Binding to address %s", context.SmtpBindAddr()))
	listener, err := net.Listen("tcp", context.SmtpBindAddr())
	logger.PanicIfError(err)
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			logManager().Error(fmt.Sprintf("SMTP Error accepting connection: %s", err.Error()))
			continue
		}

		logManager().Warning("TODO: SMTP add Monkey")

		go handleSession(connection, context)
	}
}

func handleSession(connection net.Conn, context *serverContext.Context) {
	defer connection.Close()
	logManager().Debug("Initialising session")

	session := &Session{
		context: context,
		protocol: smtpServer.CreateProtocol(
			context.Config.Hostname,
			&smtpServer.Validation{
				MaximumLineLength: context.Config.Smtp.Validation.MaximumLineLength,
				MaximumReceivers:  context.Config.Smtp.Validation.MaximumReceivers,
			},
		),
		reader: io.Reader(connection),
		writer: io.Writer(connection),

		LoggedUsername: "",
	}

	if context.Authentication.RequiresAuthentication() {
		session.protocol.AuthenticationMechanismsCallback = func() []string { return []string{AUTH_MECHANISM_PLAIN} }
	}
	session.protocol.MessageReceivedCallback = func(message *dto.SMTPMessage) (string, error) {
		formattedMessage := message.Parse()

		id, err := context.Storage.Store(session.LoggedUsername, formattedMessage)

		websocket.BroadcastToClient <- websocket.BroadcastMessage{
			Room:    session.LoggedUsername,
			Message: `{"flow": "system", "type": "new_message"}`,
		}

		return string(id), err
	}
	session.protocol.CreateCustomSceneCallback = func(sceneName string) smtpServer.Scene {
		switch sceneName {
		case "AUTH_PLAIN":
			return &AuthPlainScene{
				authentication: context.Authentication,
				authenticated: func(username string) {
					session.LoggedUsername = username
				},
			}

		}
		return nil
	}

	session.Start("")
	for session.IsConversation() {
		// loop
	}
	io.Closer(connection).Close() // not sure if this is necessary?

	logManager().Debug("End session")
}
