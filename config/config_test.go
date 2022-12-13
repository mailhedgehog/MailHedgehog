package config

import (
	"github.com/mailhedgehog/MailHedgehog/gounit"
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

	(*gounit.T)(t).AssertEqualsString("file", config.Authentication.Use)
	(*gounit.T)(t).AssertEqualsString("", config.Authentication.File.Path)

}

func TestParseConfigs(t *testing.T) {
	config := ParseConfig("../cmd/publish/.mh-config.yml")

	(*gounit.T)(t).AssertEqualsInt(1025, config.Smtp.Port)

	(*gounit.T)(t).AssertEqualsInt(8025, config.Http.Port)
	(*gounit.T)(t).AssertEqualsString("box", config.Http.Path)

	(*gounit.T)(t).AssertEqualsString("directory", config.Storage.Use)
	(*gounit.T)(t).AssertEqualsString("", config.Storage.Directory.Path)

	(*gounit.T)(t).AssertEqualsString("file", config.Authentication.Use)
	(*gounit.T)(t).AssertEqualsString(".mh-authfile", config.Authentication.File.Path)
}
