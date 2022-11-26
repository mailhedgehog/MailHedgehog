package config

import (
	"github.com/mailpiggy/MailPiggy/gounit"
	"testing"
)

func TestParseConfigsWithDefault(t *testing.T) {
	config := ParseConfig("")

	(*gounit.T)(t).AssertEqualsString("0.0.0.0", config.Smtp.Host)
	(*gounit.T)(t).AssertEqualsInt(1025, config.Smtp.Port)

	(*gounit.T)(t).AssertEqualsString("0.0.0.0", config.Http.Host)
	(*gounit.T)(t).AssertEqualsInt(8025, config.Http.Port)
	(*gounit.T)(t).AssertEqualsString("", config.Http.Path)

	(*gounit.T)(t).AssertEqualsString("directory", config.Storage.Use)
	(*gounit.T)(t).AssertEqualsString("", config.Storage.Directory.Path)

	(*gounit.T)(t).AssertEqualsString("file", config.Authorisation.Use)
	(*gounit.T)(t).AssertEqualsString("", config.Authorisation.File.Path)

}

func TestParseConfigs(t *testing.T) {
	config := ParseConfig(".config.yml")

	(*gounit.T)(t).AssertEqualsInt(1025, config.Smtp.Port)

	(*gounit.T)(t).AssertEqualsInt(8025, config.Http.Port)
	(*gounit.T)(t).AssertEqualsString("box", config.Http.Path)

	(*gounit.T)(t).AssertEqualsString("directory", config.Storage.Use)
	(*gounit.T)(t).AssertEqualsString("", config.Storage.Directory.Path)

	(*gounit.T)(t).AssertEqualsString("file", config.Authorisation.Use)
	(*gounit.T)(t).AssertEqualsString("auth.file", config.Authorisation.File.Path)
}
