package server

import (
	"github.com/mailpiggy/MailPiggy/authorisation"
	"github.com/mailpiggy/MailPiggy/config"
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
