package smtpServer

import (
	"github.com/mailhedgehog/MailHedgehog/gounit"
	"testing"
)

func TestResetState(t *testing.T) {
	protocol := CreateProtocol("", nil)

	protocol.State = STATE_DATA
	protocol.Message.From = "foo bar"

	(*gounit.T)(t).AssertEqualsString(string(STATE_DATA), string(protocol.State))
	(*gounit.T)(t).AssertEqualsString("foo bar", protocol.Message.From)

	protocol.resetState()

	(*gounit.T)(t).AssertEqualsString(string(STATE_CONVERSATION), string(protocol.State))
	(*gounit.T)(t).AssertEqualsString("", protocol.Message.From)
}

func TestSayHi(t *testing.T) {
	protocol := CreateProtocol("", nil)
	reply := protocol.SayHi("   foo bar    ")

	(*gounit.T)(t).AssertEqualsInt(CODE_SERVICE_READY, reply.Status)
	(*gounit.T)(t).AssertEqualsString("foo bar Service ready", reply.lines[0])
}

func TestHandleReceivedLine(t *testing.T) {
	logManager().Warning("TODO: add test for HandleReceivedLine") // TODO
}

func TestHandleMailContent(t *testing.T) {
	protocol := CreateProtocol("", nil)
	(*gounit.T)(t).AssertNil(protocol.handleMailContent("content foo bar baz"))
	(*gounit.T)(t).AssertNil(protocol.handleMailContent("content foo bar"))
	(*gounit.T)(t).AssertNil(protocol.handleMailContent("foo bar baz"))
}

func TestHandleCommandQUIT(t *testing.T) {
	protocol := CreateProtocol("", nil)
	reply := protocol.handleCommand("QUIT")

	(*gounit.T)(t).AssertEqualsInt(CODE_SERVICE_CLOSING, reply.Status)
	(*gounit.T)(t).AssertEqualsInt(1, len(reply.lines))
	(*gounit.T)(t).AssertEqualsString("Bye", reply.lines[0])
}

func TestHandleCommandCommandFake(t *testing.T) {
	protocol := CreateProtocol("", nil)
	reply := protocol.handleCommand("FAKE :)")

	(*gounit.T)(t).AssertEqualsInt(CODE_COMMAND_SYNTAX_ERROR, reply.Status)
}

func TestHELO(t *testing.T) {
	command := CommandFromLine("HELO foo.host.bar")
	protocol := CreateProtocol("", nil)
	reply := protocol.HELO(command)

	(*gounit.T)(t).AssertEqualsInt(CODE_ACTION_OK, reply.Status)
	(*gounit.T)(t).AssertEqualsInt(1, len(reply.lines))
	(*gounit.T)(t).AssertEqualsString("Hello foo.host.bar", reply.lines[0])

	(*gounit.T)(t).AssertEqualsString("foo.host.bar", protocol.Message.Helo)
}

func TestEHLO(t *testing.T) {
	command := CommandFromLine("EHLO foo.host.bar")
	protocol := CreateProtocol("", nil)
	reply := protocol.EHLO(command)

	(*gounit.T)(t).AssertEqualsInt(CODE_ACTION_OK, reply.Status)
	(*gounit.T)(t).AssertEqualsInt(2, len(reply.lines))
	(*gounit.T)(t).AssertEqualsString("Hello foo.host.bar", reply.lines[0])
	(*gounit.T)(t).AssertEqualsString("PIPELINING", reply.lines[1])

	(*gounit.T)(t).AssertEqualsString("foo.host.bar", protocol.Message.Helo)
}

func TestMAIL(t *testing.T) {
	command := CommandFromLine("MAIL FROM:<userx@y.foo.org>")
	protocol := CreateProtocol("", nil)

	(*gounit.T)(t).AssertEqualsString("", protocol.Message.From)

	reply := protocol.MAIL(command)

	(*gounit.T)(t).AssertEqualsInt(CODE_ACTION_OK, reply.Status)
	(*gounit.T)(t).AssertEqualsString("Sender userx@y.foo.org ok", reply.lines[0])

	(*gounit.T)(t).AssertEqualsString("userx@y.foo.org", protocol.Message.From)
}

func TestMAILFails(t *testing.T) {
	command := CommandFromLine("MAIL fake data")
	protocol := CreateProtocol("", nil)

	reply := protocol.MAIL(command)

	(*gounit.T)(t).AssertEqualsInt(CODE_MAILBOX_404, reply.Status)
	(*gounit.T)(t).AssertEqualsString("Invalid syntax in MAIL command", reply.lines[0])
}

func TestRCPT(t *testing.T) {
	protocol := CreateProtocol("", nil)

	(*gounit.T)(t).AssertEqualsInt(0, len(protocol.Message.To))

	command := CommandFromLine("RCPT TO:<userx@y.foo.org>")
	reply := protocol.RCPT(command)
	(*gounit.T)(t).AssertEqualsInt(CODE_ACTION_OK, reply.Status)
	(*gounit.T)(t).AssertEqualsString("Receiver userx@y.foo.org ok", reply.lines[0])

	command = CommandFromLine("RCPT TO:<user2@y.foo.org>")
	reply = protocol.RCPT(command)
	(*gounit.T)(t).AssertEqualsInt(CODE_ACTION_OK, reply.Status)
	(*gounit.T)(t).AssertEqualsString("Receiver user2@y.foo.org ok", reply.lines[0])

	(*gounit.T)(t).AssertEqualsInt(2, len(protocol.Message.To))
	(*gounit.T)(t).AssertEqualsString("user2@y.foo.org", protocol.Message.To[1])
}

func TestRCPTFails(t *testing.T) {
	protocol := CreateProtocol("", nil)
	command := CommandFromLine("RCPT fake")
	reply := protocol.RCPT(command)
	(*gounit.T)(t).AssertEqualsInt(CODE_MAILBOX_404, reply.Status)
	(*gounit.T)(t).AssertEqualsString("Invalid syntax in MAIL command", reply.lines[0])
}

func TestGetAuthMechanisms(t *testing.T) {
	protocol := CreateProtocol("", nil)
	(*gounit.T)(t).AssertEqualsInt(0, len(protocol.authMechanisms()))

	protocol.AuthenticationMechanismsCallback = func() []string {
		return []string{"PLAIN", "foo", "BAR"}
	}
	mechanisms := protocol.authMechanisms()
	(*gounit.T)(t).AssertEqualsInt(3, len(mechanisms))
	(*gounit.T)(t).AssertEqualsString("PLAIN", mechanisms[0])
	(*gounit.T)(t).AssertEqualsString("foo", mechanisms[1])
	(*gounit.T)(t).AssertEqualsString("BAR", mechanisms[2])
}

func TestParseAuthMechanism(t *testing.T) {
	protocol := CreateProtocol("", nil)

	(*gounit.T)(t).AssertEqualsString("", protocol.parseAuthMechanism(""))
	(*gounit.T)(t).AssertEqualsString("foo", protocol.parseAuthMechanism("foo"))
	(*gounit.T)(t).AssertEqualsString("BAR", protocol.parseAuthMechanism("BAR"))
	(*gounit.T)(t).AssertEqualsString("foo", protocol.parseAuthMechanism("foo baz"))
}

func TestParseFROM(t *testing.T) {
	protocol := CreateProtocol("", nil)

	res, err := protocol.ParseFROM("FROM:<userx@y.foo.org>")
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsString("userx@y.foo.org", res)

	res, err = protocol.ParseFROM("from:<userx@y.foo.org>")
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsString("userx@y.foo.org", res)
}

func TestParseRCPT(t *testing.T) {
	protocol := &Protocol{}

	res, err := protocol.ParseRCPT("TO:<@hosta.int,@jkl.org:userc@d.bar.org>")
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsString("@hosta.int,@jkl.org:userc@d.bar.org", res)

	res, err = protocol.ParseRCPT("to:<userc@d.bar.org>")
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsString("userc@d.bar.org", res)
}
