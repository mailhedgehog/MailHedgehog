package smtp

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/MailHedgehog/server/websocket"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"github.com/mailhedgehog/MailHedgehog/smtpServerProtocol"
	"github.com/mailhedgehog/logger"
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

// Protocol interface describe all methods what use server.
// Used to add possibility override/inherit current protocol
type Protocol interface {
	// SetAuthMechanisms sets allowed by app auth mechanisms
	// if empty then no auth check need.
	SetAuthMechanisms(authMechanisms []string)
	// SayWelcome starts conversation with "client"
	SayWelcome(identification string) *smtpServerProtocol.Reply
	// HandleReceivedLine receive sent client message line by line
	// handle it and returns reply.
	HandleReceivedLine(receivedLine string) *smtpServerProtocol.Reply
	// CreateCustomSceneUsing - for some commands, methods good to have logic code in this top level
	// package than in protocol, like authentication code, etc. And there scenes are useful.
	CreateCustomSceneUsing(callback func(sceneName string) smtpServerProtocol.Scene)
	// SetStateCommandsExchange notify protocol to exit custom scene back to commands exchange
	SetStateCommandsExchange()
	// OnMessageReceived function called only when smtp protocol finished and messaage fully received and processed.
	// So package can save this message or send to other services and return string ID of saved message to send
	// this ID to "client"
	OnMessageReceived(callback func(message *smtpMessage.SMTPMessage) (string, error))
}

const (
	AuthMechanismPlain = "PLAIN"
)

var listener net.Listener

// Listen "client" welcome message to initialise session (protocol).
// Start server
func Listen(context *serverContext.Context, exitCh chan int) {
	var err error
	logManager().Debug(fmt.Sprintf("SMTP Binding to address %s", context.SmtpBindAddr()))

	listener, err = net.Listen("tcp", context.SmtpBindAddr())
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

// Close listener
func Close() error {
	if listener == nil {
		return nil
	}
	return listener.Close()
}

// handleSession initialise and process session between "client" and "server"
func handleSession(connection net.Conn, context *serverContext.Context) {
	defer connection.Close()
	logManager().Debug("Initialising session")

	ipAddr, _ := connection.RemoteAddr().(*net.TCPAddr)

	session := &session{
		context: context,
		protocol: smtpServerProtocol.CreateProtocol(
			context.Config.Hostname,
			ipAddr,
			&smtpServerProtocol.Validation{
				MaximumLineLength: context.Config.Smtp.Validation.MaximumLineLength,
				MaximumReceivers:  context.Config.Smtp.Validation.MaximumReceivers,
			},
		),
		reader: io.Reader(connection),
		writer: io.Writer(connection),

		loggedUsername: "",
	}

	// Set conditionally Authentication flow
	if context.Authentication.RequiresAuthentication() {
		// TODO: there can be added other auth mechanisms
		session.protocol.SetAuthMechanisms([]string{AuthMechanismPlain})
	}

	session.protocol.OnMessageReceived(func(message *smtpMessage.SMTPMessage) (string, error) {
		formattedMessage, err := message.ToSMTPMail(smtpMessage.MessageID(""))
		if err != nil {
			return "", err
		}

		id, err := context.Storage.Store(session.loggedUsername, formattedMessage)

		if context.Config.Http.Websocket {
			logManager().Debug("Send to websocket notification about new message received. (commented for now)")
			// Send to websocket notification about new message received.
			websocket.BroadcastToClient <- websocket.BroadcastMessage{
				Room:    session.loggedUsername,
				Message: `{"flow": "system", "type": "new_message"}`,
			}
		}

		messageId := string(id)
		logManager().Debug(fmt.Sprintf("OnMessageReceived callback processed for message id: %s", messageId))

		return messageId, err
	})

	session.protocol.CreateCustomSceneUsing(
		func(sceneName string) smtpServerProtocol.Scene {
			switch sceneName {
			case "AUTH_PLAIN":
				return &AuthPlainScene{
					authentication: context.Authentication,
					authenticated: func(username string) {
						session.loggedUsername = username
					},
				}
			}
			return nil
		},
	)

	session.start("")
	for session.readAndWriteReply() {
		// loop until session works
	}

	// not sure if this is necessary?, because we already call 'defer connection.Close()'
	err := io.Closer(connection).Close()
	if err != nil {
		logManager().Error(fmt.Sprintf("Write session error: %s", err.Error()))
	}

	logManager().Debug("End session")
}
