package smtpServerProtocol

import (
	"errors"
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"golang.org/x/exp/slices"
	"regexp"
	"strings"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("smtpServerProtocol")
	}
	return configuredLogger
}

// Scene represents custom logic flow (scene) for some specific
// set of commands, for example authentication.
type Scene interface {
	// Start scene by send specific message (reply) to client.
	Start(receivedLine string, protocol *Protocol) *Reply
	// ReadAndWriteReply reads client message and write reply
	ReadAndWriteReply(receivedLine string) *Reply
	// Finish scene, by notifying protocol to finish this scene
	Finish()
}

const CommandEndSymbol = "\r\n"

type ConversationState string

const (
	StateCommandsExchange = ConversationState("commands_exchange")
	StateWaitingAuth      = ConversationState("waiting_auth")
	StateData             = ConversationState("data")
	StateCustomScene      = ConversationState("custom_scene")
)

type Validation struct {
	MaximumLineLength int
	MaximumReceivers  int
}

type Protocol struct {
	Hostname   string
	validation *Validation

	state   ConversationState
	message *smtpMessage.SMTPMessage

	// supportedAuthMechanisms can be empty, if empty client will not go through auth flow
	supportedAuthMechanisms []string
	messageReceivedCallback func(message *smtpMessage.SMTPMessage) (string, error)

	createCustomSceneCallback func(sceneName string) Scene
	currentScene              Scene
}

func CreateProtocol(hostname string, validation *Validation) *Protocol {
	if validation == nil {
		validation = &Validation{
			MaximumLineLength: 0,
			MaximumReceivers:  0,
		}
	}

	protocol := &Protocol{
		Hostname:   hostname,
		validation: validation,
	}
	protocol.resetState()

	return protocol
}

func (protocol *Protocol) SetAuthMechanisms(authMechanisms []string) {
	protocol.supportedAuthMechanisms = authMechanisms
}

func (protocol *Protocol) OnMessageReceived(callback func(message *smtpMessage.SMTPMessage) (string, error)) {
	protocol.messageReceivedCallback = callback
}

func (protocol *Protocol) CreateCustomSceneUsing(callback func(sceneName string) Scene) {
	protocol.createCustomSceneCallback = callback
}

func (protocol *Protocol) SetStateCommandsExchange() {
	protocol.state = StateCommandsExchange
}

func (protocol *Protocol) SayWelcome(identification string) *Reply {
	identification = strings.TrimSpace(identification)
	if len(identification) > 0 {
		identification = identification + " "
	}
	hostname := protocol.Hostname
	if len(hostname) > 0 {
		hostname = hostname + " "
	}
	protocol.state = StateCommandsExchange
	return ReplyServiceReady(hostname + identification + "Service ready")
}

func (protocol *Protocol) HandleReceivedLine(receivedLine string) *Reply {
	if protocol.validation.MaximumLineLength > 0 && len(receivedLine) > 0 {
		if len(receivedLine) > protocol.validation.MaximumLineLength {
			return ReplyLineTooLong()
		}
	}

	if protocol.state == StateCustomScene {
		if protocol.currentScene != nil {
			return protocol.currentScene.ReadAndWriteReply(receivedLine)
		}
		return ReplyCommandNotImplemented()
	}

	if protocol.state == StateData {
		return protocol.handleMailContent(receivedLine)
	}

	return protocol.handleCommand(receivedLine)
}

func (protocol *Protocol) resetState() {
	protocol.message = &smtpMessage.SMTPMessage{}
	protocol.SetStateCommandsExchange()
}

func (protocol *Protocol) handleMailContent(receivedLine string) *Reply {
	protocol.message.Data += receivedLine + "\r\n"
	if strings.HasSuffix(protocol.message.Data, "\r\n.\r\n") {
		protocol.message.Data = strings.ReplaceAll(protocol.message.Data, "\r\n..", "\r\n.")

		logManager().Debug("Got EOF, storing message and reset state.")
		protocol.message.Data = strings.TrimSuffix(protocol.message.Data, "\r\n.\r\n")
		protocol.state = StateCommandsExchange

		defer protocol.resetState()

		if protocol.messageReceivedCallback == nil {
			return ReplyExceededStorage("No storage backend")
		}

		messageId, err := protocol.messageReceivedCallback(protocol.message)
		if err != nil {
			logManager().Error(fmt.Sprintf("Error storing message: %s", err.Error()))
			return ReplyExceededStorage("Unable to store message")
		}
		return ReplyOk("Ok: queued as " + messageId)
	}

	return nil
}

func (protocol *Protocol) handleCommand(receivedLine string) *Reply {
	receivedLine = strings.Trim(receivedLine, "\r\n")
	command := CommandFromLine(receivedLine)

	logManager().Debug(fmt.Sprintf("Handle command: '%s', with args: '%s'", command.verb, command.args))

	if protocol.state == StateWaitingAuth && command.verb != CommandAuth {
		return ReplyAuthFailed()
	}

	switch command.verb {
	case CommandHelo:
		return protocol.HELO(command)
	case CommandEhlo:
		return protocol.EHLO(command)
	case CommandAuth:
		logManager().Debug(fmt.Sprintf("Got %s command", command.verb))
		authMechanism := protocol.parseAuthMechanism(command.args)
		if slices.Contains(protocol.supportedAuthMechanisms, authMechanism) && protocol.createCustomSceneCallback != nil {
			protocol.currentScene = protocol.createCustomSceneCallback(string(command.verb) + "_" + authMechanism)
			if protocol.currentScene != nil {
				protocol.state = StateCustomScene
				return protocol.currentScene.Start(receivedLine, protocol)
			}
		}
		return ReplyCommandNotImplemented()
	case CommandRset:
		return protocol.RSET(command)
	case CommandMail:
		return protocol.MAIL(command)
	case CommandRcpt:
		return protocol.RCPT(command)
	case CommandData:
		protocol.state = StateData
		return ReplyMailData()
	case CommandQuit:
		return ReplyBye()
	default:
		return ReplyUnrecognisedCommand()
	}
}

func (protocol *Protocol) HELO(command *Command) *Reply {
	protocol.message.Helo = command.args

	if len(protocol.supportedAuthMechanisms) > 0 {
		protocol.state = StateWaitingAuth
	}

	return ReplyOk("Hello " + command.args)
}

func (protocol *Protocol) EHLO(command *Command) *Reply {
	protocol.message.Helo = command.args
	replyArgs := []string{"Hello " + command.args, "PIPELINING"}

	logManager().Warning("TODO: add tls support") // TODO

	if len(protocol.supportedAuthMechanisms) > 0 {
		protocol.state = StateWaitingAuth
		replyArgs = append(replyArgs, string(CommandAuth)+" "+strings.Join(protocol.supportedAuthMechanisms, " "))
	}

	return ReplyOk(replyArgs...)
}

func (protocol *Protocol) RSET(command *Command) *Reply {
	protocol.resetState()

	return ReplyOk("")
}

func (protocol *Protocol) MAIL(command *Command) *Reply {
	from, err := protocol.ParseFROM(command.args)
	if err != nil {
		return ReplyMailbox404(err.Error())
	}
	protocol.message.From = from

	return ReplyOk("Sender " + protocol.message.From + " ok")
}

func (protocol *Protocol) RCPT(command *Command) *Reply {
	if protocol.validation.MaximumReceivers > 0 && len(protocol.message.To) >= protocol.validation.MaximumReceivers {
		return ReplyExceededStorage("Maximum receivers extended")
	}
	to, err := protocol.ParseRCPT(command.args)
	if err != nil {
		return ReplyMailbox404(err.Error())
	}
	protocol.message.To = append(protocol.message.To, to)

	return ReplyOk("Receiver " + to + " ok")
}

func (protocol *Protocol) parseAuthMechanism(args string) string {
	parts := strings.SplitN(args, " ", 2)

	return parts[0]
}

func (protocol *Protocol) ParseFROM(mail string) (string, error) {
	match := regexp.MustCompile(`(?i:From):\s*<([^>]+)>`).FindStringSubmatch(mail)

	if len(match) != 2 {
		return "", errors.New("Invalid syntax in MAIL command")
	}

	return match[1], nil
}

func (protocol *Protocol) ParseRCPT(mail string) (string, error) {
	match := regexp.MustCompile(`(?i:To):\s*<([^>]+)>`).FindStringSubmatch(mail)

	if len(match) != 2 {
		return "", errors.New("Invalid syntax in MAIL command")
	}

	return match[1], nil
}
