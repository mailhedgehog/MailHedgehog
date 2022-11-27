package smtpServer

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"github.com/mailpiggy/MailPiggy/logger"
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

type ConversationState int

const (
	STATE_INVALID   = ConversationState(-1)
	STATE_ESTABLISH = ConversationState(iota)
	STATE_AUTH      = ConversationState(3)
	STATE_MAIL      = ConversationState(4)
	STATE_RCPT      = ConversationState(5)
	STATE_DATA      = ConversationState(6)
	STATE_DONE      = ConversationState(7)
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
}

func CreateProtocol(hostname string, validation *Validation) *Protocol {
	if validation == nil {
		validation = &Validation{
			MaximumLineLength: 0,
			MaximumReceivers:  0,
		}
	}
	return &Protocol{
		Hostname:   hostname,
		validation: validation,

		State:   STATE_ESTABLISH,
		Message: &dto.SMTPMessage{},
	}
}

func (protocol *Protocol) SayHi(identification string) *Reply {
	identification = strings.TrimSpace(identification)
	if len(identification) > 0 {
		identification = " " + identification
	}
	protocol.State = STATE_ESTABLISH
	return ReplyServiceReady(protocol.Hostname + identification + " Service ready")
}

func (protocol *Protocol) HandleReceivedLine(receivedLine string) *Reply {
	if protocol.validation.MaximumLineLength > 0 && len(receivedLine) > 0 {
		if len(receivedLine) > protocol.validation.MaximumLineLength {
			return ReplyLineTooLong()
		}
	}

	if protocol.State == STATE_DATA {
		return protocol.handleMailContent(receivedLine)
	}

	return protocol.handleCommand(receivedLine)
}

func (protocol *Protocol) handleMailContent(receivedLine string) *Reply {
	fmt.Println(receivedLine)
	fmt.Println("---handleMailContent----")
	return ReplyLineTooLong()
}

func (protocol *Protocol) handleCommand(receivedLine string) *Reply {
	receivedLine = strings.Trim(receivedLine, "\r\n")
	command := CommandFromLine(receivedLine)

	logManager().Debug(fmt.Sprintf("Handle command: '%s', with args: '%s'", command.verb, command.args))

	switch command.verb {
	case COMMAND_HELO:
		return protocol.HELO(command)
	case COMMAND_EHLO:
		return protocol.EHLO(command)
	case COMMAND_AUTH:
		return protocol.AUTH(command)
	case COMMAND_QUIT:
		return ReplyBye()
	default:
		return ReplyUnrecognisedCommand()
	}
}

func (protocol *Protocol) HELO(command *Command) *Reply {
	protocol.Message.Helo = command.args
	logManager().Debug(fmt.Sprintf("Got %s command", command.verb))

	return ReplyOk("Hello " + command.args)
}

func (protocol *Protocol) EHLO(command *Command) *Reply {
	protocol.Message.Helo = command.args
	logManager().Debug(fmt.Sprintf("Got %s command", command.verb))

	replyArgs := []string{"Hello " + command.args, "PIPELINING"}

	logManager().Warning("TODO: add tls support") // TODO

	if protocol.AuthenticationMechanismsCallback != nil {
		if mechanisms := protocol.AuthenticationMechanismsCallback(); len(mechanisms) > 0 {
			replyArgs = append(replyArgs, string(COMMAND_AUTH)+" "+strings.Join(mechanisms, " "))
		}
	}

	return ReplyOk(replyArgs...)
}

func (protocol *Protocol) AUTH(command *Command) *Reply {
	protocol.Message.Helo = command.args
	logManager().Debug(fmt.Sprintf("Got %s command", command.verb))

	return ReplyOk()
}
