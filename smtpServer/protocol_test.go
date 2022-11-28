package smtpServer

import (
	"github.com/mailpiggy/MailPiggy/gounit"
	"testing"
)

func TestParseAuthMechanism(t *testing.T) {
	protocol := &Protocol{}

	(*gounit.T)(t).AssertEqualsString("", protocol.parseAuthMechanism(""))
	(*gounit.T)(t).AssertEqualsString("foo", protocol.parseAuthMechanism("foo"))
	(*gounit.T)(t).AssertEqualsString("BAR", protocol.parseAuthMechanism("BAR"))
	(*gounit.T)(t).AssertEqualsString("foo", protocol.parseAuthMechanism("foo baz"))
}

func TestParseFROM(t *testing.T) {
	protocol := &Protocol{}

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
