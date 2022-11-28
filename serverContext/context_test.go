package serverContext

import (
	"github.com/mailpiggy/MailPiggy/config"
	"github.com/mailpiggy/MailPiggy/gounit"
	"testing"
)

var context *Context

func init() {
	context = &Context{
		Config: *config.ParseConfig(""),
	}
}

func TestSmtpBindAddr(t *testing.T) {
	(*gounit.T)(t).AssertEqualsInt(1025, context.Config.Smtp.Port)
	(*gounit.T)(t).AssertEqualsString("0.0.0.0", context.Config.Smtp.Host)
	(*gounit.T)(t).AssertEqualsString("0.0.0.0:1025", context.SmtpBindAddr())

	context.Config.Smtp.Port = 123
	context.Config.Smtp.Host = "foo.bar"
	(*gounit.T)(t).AssertEqualsString("foo.bar:123", context.SmtpBindAddr())
}

func TestHttpBindAddr(t *testing.T) {
	(*gounit.T)(t).AssertEqualsInt(8025, context.Config.Http.Port)
	(*gounit.T)(t).AssertEqualsString("0.0.0.0", context.Config.Http.Host)
	(*gounit.T)(t).AssertEqualsString("0.0.0.0:8025", context.HttpBindAddr())

	context.Config.Http.Port = 123
	context.Config.Http.Host = "foo.bar"
	(*gounit.T)(t).AssertEqualsString("foo.bar:123", context.HttpBindAddr())
}
