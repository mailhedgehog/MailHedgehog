package smtpServer

import (
	"github.com/mailhedgehog/MailHedgehog/gounit"
	"testing"
)

func TestReplyOk(t *testing.T) {
	reply := ReplyOk()

	(*gounit.T)(t).AssertEqualsInt(CODE_ACTION_OK, reply.Status)
	(*gounit.T)(t).AssertEqualsInt(1, len(reply.lines))
	(*gounit.T)(t).AssertEqualsString("Ok", reply.lines[0])

	reply = ReplyOk("foo", "BAR", "baz")

	(*gounit.T)(t).AssertEqualsInt(CODE_ACTION_OK, reply.Status)
	(*gounit.T)(t).AssertEqualsInt(3, len(reply.lines))
	(*gounit.T)(t).AssertEqualsString("foo", reply.lines[0])
	(*gounit.T)(t).AssertEqualsString("BAR", reply.lines[1])
	(*gounit.T)(t).AssertEqualsString("baz", reply.lines[2])
}

func TestLinesOfReply(t *testing.T) {
	reply := ReplyOk("foo", "BAR", "baz")
	lines := reply.Lines()

	(*gounit.T)(t).AssertEqualsInt(3, len(lines))
	(*gounit.T)(t).AssertEqualsString("250-foo"+COMMAND_END_SYMBOL, lines[0])
	(*gounit.T)(t).AssertEqualsString("250-BAR"+COMMAND_END_SYMBOL, lines[1])
	(*gounit.T)(t).AssertEqualsString("250 baz"+COMMAND_END_SYMBOL, lines[2])
}
