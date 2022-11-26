package server

import (
	"github.com/mailpiggy/MailPiggy/authorisation"
	"github.com/mailpiggy/MailPiggy/config"
	"github.com/mailpiggy/MailPiggy/gounit"
	"github.com/mailpiggy/MailPiggy/storage"
	"testing"
)

func TestConfigure(t *testing.T) {
	context := Configure(config.ParseConfig(""))

	(*gounit.T)(t).AssertEqualsInt(1025, context.Config.Smtp.Port)
	(*gounit.T)(t).AssertEqualsInt(8025, context.Config.Http.Port)
	(*gounit.T)(t).AssertEqualsString("", context.Config.Http.Path)
	(*gounit.T)(t).AssertEqualsString("directory", context.Config.Storage.Use)
	(*gounit.T)(t).AssertEqualsString("", context.Config.Storage.Directory.Path)
	(*gounit.T)(t).AssertEqualsString("file", context.Config.Authorisation.Use)
	(*gounit.T)(t).AssertEqualsString("", context.Config.Authorisation.File.Path)

	(*gounit.T)(t).AssertInstanceOf((*storage.Directory)(nil), context.Storage)
	(*gounit.T)(t).AssertInstanceOf((*authorisation.FileAuth)(nil), context.Authorisation)

}
