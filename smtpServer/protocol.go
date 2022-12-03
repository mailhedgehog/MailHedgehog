package smtpServer

import (
	"errors"
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"github.com/mailpiggy/MailPiggy/logger"
	"golang.org/x/exp/slices"
	"regexp"
	"strings"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("smtpServer")
	}
	return configuredLogger
}

const COMMAND_END_SYMBOL = "\r\n"

type ConversationState string

const (
	STATE_CONVERSATION = ConversationState("conversation")
	STATE_WAITING_AUTH = ConversationState("waiting_auth")
	STATE_DATA         = ConversationState("data")
	STATE_CUSTOM_SCENE = ConversationState("custom_scene")
)

type Validation struct {
	MaximumLineLength int
	MaximumReceivers  int
}

type Protocol struct {
	Hostname   string
	validation *Validation

	State   ConversationState
	Message *dto.SMTPMessage

	AuthenticationMechanismsCallback func() []string
	MessageReceivedCallback          func(message *dto.SMTPMessage) (string, error)

	CreateCustomSceneCallback func(sceneName string) Scene
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

func (protocol *Protocol) resetState() {
	protocol.State = STATE_CONVERSATION
	protocol.Message = &dto.SMTPMessage{}
}

func (protocol *Protocol) SayHi(identification string) *Reply {
	identification = strings.TrimSpace(identification)
	if len(identification) > 0 {
		identification = identification + " "
	}
	hostname := protocol.Hostname
	if len(hostname) > 0 {
		hostname = hostname + " "
	}
	protocol.State = STATE_CONVERSATION
	return ReplyServiceReady(hostname + identification + "Service ready")
}

func (protocol *Protocol) HandleReceivedLine(receivedLine string) *Reply {
	if protocol.validation.MaximumLineLength > 0 && len(receivedLine) > 0 {
		if len(receivedLine) > protocol.validation.MaximumLineLength {
			return ReplyLineTooLong()
		}
	}

	if protocol.State == STATE_CUSTOM_SCENE {
		if protocol.currentScene != nil {
			return protocol.currentScene.HandleLine(receivedLine)
		}
		return ReplyCommandNotImplemented()
	}

	if protocol.State == STATE_DATA {
		return protocol.handleMailContent(receivedLine)
	}

	return protocol.handleCommand(receivedLine)
}

func (protocol *Protocol) handleMailContent(receivedLine string) *Reply {
	protocol.Message.Data += receivedLine + "\r\n"
	if strings.HasSuffix(protocol.Message.Data, "\r\n.\r\n") {
		protocol.Message.Data = strings.ReplaceAll(protocol.Message.Data, "\r\n..", "\r\n.")

		logManager().Debug("Got EOF, storing message and reset state.")
		protocol.Message.Data = strings.TrimSuffix(protocol.Message.Data, "\r\n.\r\n")
		protocol.State = STATE_CONVERSATION

		defer protocol.resetState()

		if protocol.MessageReceivedCallback == nil {
			return ReplyExceededStorage("No storage backend")
		}

		messageId, err := protocol.MessageReceivedCallback(protocol.Message)
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

	if protocol.State == STATE_WAITING_AUTH && command.verb != COMMAND_AUTH {
		return ReplyAuthFailed()
	}

	switch command.verb {
	case COMMAND_HELO:
		return protocol.HELO(command)
	case COMMAND_EHLO:
		return protocol.EHLO(command)
	case COMMAND_AUTH:
		logManager().Debug(fmt.Sprintf("Got %s command", command.verb))
		authMechanism := protocol.parseAuthMechanism(command.args)
		if slices.Contains(protocol.authMechanisms(), authMechanism) && protocol.CreateCustomSceneCallback != nil {
			protocol.currentScene = protocol.CreateCustomSceneCallback(string(command.verb) + "_" + authMechanism)
			if protocol.currentScene != nil {
				protocol.State = STATE_CUSTOM_SCENE
				return protocol.currentScene.Init(receivedLine, protocol)
			}
		}
		return ReplyCommandNotImplemented()
	case COMMAND_MAIL:
		return protocol.MAIL(command)
	case COMMAND_RCPT:
		return protocol.RCPT(command)
	case COMMAND_DATA:
		protocol.State = STATE_DATA
		return ReplyMailData()
	case COMMAND_QUIT:
		return ReplyBye()
	default:
		return ReplyUnrecognisedCommand()
	}
}

func (protocol *Protocol) HELO(command *Command) *Reply {
	protocol.Message.Helo = command.args

	if mechanisms := protocol.authMechanisms(); len(mechanisms) > 0 {
		protocol.State = STATE_WAITING_AUTH
	}

	return ReplyOk("Hello " + command.args)
}

func (protocol *Protocol) EHLO(command *Command) *Reply {
	protocol.Message.Helo = command.args
	replyArgs := []string{"Hello " + command.args, "PIPELINING"}

	logManager().Warning("TODO: add tls support") // TODO

	if mechanisms := protocol.authMechanisms(); len(mechanisms) > 0 {
		protocol.State = STATE_WAITING_AUTH
		replyArgs = append(replyArgs, string(COMMAND_AUTH)+" "+strings.Join(mechanisms, " "))
	}

	return ReplyOk(replyArgs...)
}

func (protocol *Protocol) MAIL(command *Command) *Reply {
	from, err := protocol.ParseFROM(command.args)
	if err != nil {
		return ReplyMailbox404(err.Error())
	}
	protocol.Message.From = from

	return ReplyOk("Sender " + protocol.Message.From + " ok")
}

func (protocol *Protocol) RCPT(command *Command) *Reply {
	if protocol.validation.MaximumReceivers > 0 && len(protocol.Message.To) >= protocol.validation.MaximumReceivers {
		return ReplyExceededStorage("Maximum receivers extended")
	}
	to, err := protocol.ParseRCPT(command.args)
	if err != nil {
		return ReplyMailbox404(err.Error())
	}
	protocol.Message.To = append(protocol.Message.To, to)

	return ReplyOk("Receiver " + to + " ok")
}

func (protocol *Protocol) authMechanisms() []string {
	if protocol.AuthenticationMechanismsCallback != nil {
		return protocol.AuthenticationMechanismsCallback()
	}
	return []string{}
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
	match := regexp.MustCompile(`(?i:EmailInfo):\s*<([^>]+)>`).FindStringSubmatch(mail)

	if len(match) != 2 {
		return "", errors.New("Invalid syntax in MAIL command")
	}

	return match[1], nil
}
