package smtpServer

import (
	"github.com/mailpiggy/MailPiggy/gounit"
	"testing"
)

func TestCommandFromLine(t *testing.T) {
	command := CommandFromLine("")
	(*gounit.T)(t).AssertEqualsString("", string(command.verb))
	(*gounit.T)(t).AssertEqualsString("", command.args)

	command = CommandFromLine(string(COMMAND_QUIT))
	(*gounit.T)(t).AssertEqualsString(string(COMMAND_QUIT), string(command.verb))
	(*gounit.T)(t).AssertEqualsString("", command.args)

	command = CommandFromLine(string(COMMAND_AUTH) + " PLAIN")
	(*gounit.T)(t).AssertEqualsString(string(COMMAND_AUTH), string(command.verb))
	(*gounit.T)(t).AssertEqualsString("PLAIN", command.args)

	command = CommandFromLine("foo bar baz")
	(*gounit.T)(t).AssertEqualsString("FOO", string(command.verb))
	(*gounit.T)(t).AssertEqualsString("bar baz", command.args)
}

func TestCommandStrings(t *testing.T) {
	(*gounit.T)(t).AssertEqualsString(string(COMMAND_HELO), "HELO")
	(*gounit.T)(t).AssertEqualsString(string(COMMAND_EHLO), "EHLO")
	(*gounit.T)(t).AssertEqualsString(string(COMMAND_AUTH), "AUTH")
	(*gounit.T)(t).AssertEqualsString(string(COMMAND_MAIL), "MAIL")
	(*gounit.T)(t).AssertEqualsString(string(COMMAND_RCPT), "RCPT")
	(*gounit.T)(t).AssertEqualsString(string(COMMAND_DATA), "DATA")
	(*gounit.T)(t).AssertEqualsString(string(COMMAND_QUIT), "QUIT")
}
