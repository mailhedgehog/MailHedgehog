package server

import (
	"github.com/mailhedgehog/MailHedgehog/authentication"
	"github.com/mailhedgehog/MailHedgehog/config"
	"github.com/mailhedgehog/MailHedgehog/storage"
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestConfigure(t *testing.T) {
	context := Configure(config.ParseConfig(""))

	(*gounit.T)(t).AssertEqualsInt(1025, context.Config.Smtp.Port)
	(*gounit.T)(t).AssertEqualsInt(8025, context.Config.Http.Port)
	(*gounit.T)(t).AssertEqualsString("", context.Config.Http.Path)
	(*gounit.T)(t).AssertEqualsString("directory", context.Config.Storage.Use)
	(*gounit.T)(t).AssertEqualsString("", context.Config.Storage.Directory.Path)
	(*gounit.T)(t).AssertEqualsString("file", context.Config.Authentication.Use)
	(*gounit.T)(t).AssertEqualsString("", context.Config.Authentication.File.Path)

	(*gounit.T)(t).AssertInstanceOf((*storage.Directory)(nil), context.Storage)
	(*gounit.T)(t).AssertInstanceOf((*authentication.FileAuth)(nil), context.Authentication)

}
