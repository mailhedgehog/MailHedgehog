package smtpServerProtocol

import (
	"strings"
)

type CommandName string

// RSET
const (
	CommandHelo = CommandName("HELO")
	CommandEhlo = CommandName("EHLO")
	CommandAuth = CommandName("AUTH")
	CommandMail = CommandName("MAIL")
	CommandRset = CommandName("RSET")
	CommandRcpt = CommandName("RCPT")
	CommandData = CommandName("DATA")
	CommandQuit = CommandName("QUIT")
)

// Command is a struct representing an SMTP command (verb + arguments)
type Command struct {
	verb CommandName
	args string
}

func CommandFromLine(line string) *Command {
	parts := strings.SplitN(line, " ", 2)
	args := ""
	if len(parts) > 1 {
		args = parts[1]
	}
	return &Command{
		verb: CommandName(strings.ToUpper(parts[0])),
		args: args,
	}
}
