package smtpMessage

import (
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestMailPathHasAddress(t *testing.T) {
	mailPath := &MailPath{
		Relays:  []string{},
		Mailbox: "foo",
		Domain:  "bar.com",
		Params:  "",
	}
	(*gounit.T)(t).AssertEqualsString("foo@bar.com", mailPath.Address())
}

func TestMailPathFromString(t *testing.T) {
	from := MailPathFromString("baz@foo.com")
	(*gounit.T)(t).AssertTrue(len(from.Relays) == 0)
	(*gounit.T)(t).AssertEqualsString("baz", from.Mailbox)
	(*gounit.T)(t).AssertEqualsString("foo.com", from.Domain)

	to := MailPathFromString("foo,bar,baz:quix@quib.com")
	(*gounit.T)(t).AssertTrue(len(to.Relays) == 3)
	(*gounit.T)(t).AssertEqualsString("bar", to.Relays[1])
	(*gounit.T)(t).AssertEqualsString("quix", to.Mailbox)
	(*gounit.T)(t).AssertEqualsString("quib.com", to.Domain)
}
