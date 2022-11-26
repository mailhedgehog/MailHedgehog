package server

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/authorisation"
	"github.com/mailpiggy/MailPiggy/config"
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/mailpiggy/MailPiggy/storage"
)

type Context struct {
	Authorisation authorisation.Authorisation
	Storage       storage.Storage
	Config        config.AppConfig
}

func Configure(config *config.AppConfig) *Context {
	context := &Context{
		Config: *config,
	}

	logger.MinLogLevel = config.Log.Level

	switch config.Storage.Use {
	case "directory":
		context.Storage = storage.CreateDirectoryStorage(config.Storage.Directory.Path)
		if config.Storage.PerRoomLimit > 0 {
			storage.SetPerRoomLimit(config.Storage.PerRoomLimit)
		}
	default:
		panic("Incorrect storage type, Supports: directory")
	}

	switch config.Authorisation.Use {
	case "file":
		context.Authorisation = authorisation.CreateFileAuthorisation(config.Authorisation.File.Path)
	default:
		panic("Incorrect authorisation type, Supports: file")
	}

	return context
}

func (context *Context) smtpBindAddr() string {
	return context.Config.Smtp.Host + ":" + fmt.Sprint(context.Config.Smtp.Port)
}

func (context *Context) httpBindAddr() string {
	return context.Config.Smtp.Host + ":" + fmt.Sprint(context.Config.Smtp.Port)
}
