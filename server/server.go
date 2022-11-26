package server

import (
	"github.com/mailpiggy/MailPiggy/authorisation"
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/mailpiggy/MailPiggy/storage"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server")
	}
	return configuredLogger
}

type Context struct {
	Authorisation authorisation.Authorisation
	Storage       storage.Storage
}

func Configure(filePath string) *Context {

	return &Context{}
}
