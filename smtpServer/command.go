package smtpServer

import (
	"strings"
)

type CommandName string

// RSET
const (
	COMMAND_HELO = CommandName("HELO")
	COMMAND_EHLO = CommandName("EHLO")
	COMMAND_AUTH = CommandName("AUTH")
	COMMAND_MAIL = CommandName("MAIL")
	COMMAND_RSET = CommandName("RSET")
	COMMAND_RCPT = CommandName("RCPT")
	COMMAND_DATA = CommandName("DATA")
	COMMAND_QUIT = CommandName("QUIT")
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
