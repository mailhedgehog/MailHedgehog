package server

import (
	"github.com/mailpiggy/MailPiggy/authentication"
	"github.com/mailpiggy/MailPiggy/config"
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/mailpiggy/MailPiggy/server/smtp"
	"github.com/mailpiggy/MailPiggy/serverContext"
	"github.com/mailpiggy/MailPiggy/storage"
	"os"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server")
	}
	return configuredLogger
}

var exitChannel chan int

func Configure(config *config.AppConfig) *serverContext.Context {
	context := &serverContext.Context{
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

	switch config.Authentication.Use {
	case "file":
		context.Authentication = authentication.CreateFileAuthentication(config.Authentication.File.Path)
	default:
		panic("Incorrect authentication type, Supports: file")
	}

	return context
}

func Start(context *serverContext.Context) {
	logManager().Warning("TODO: SMTP Implement", context.Config.Smtp.Port)
	logManager().Warning("TODO: HTTP Implement", context.Config.Http.Port)
	logManager().Warning("TODO: UI Implement", context.Config.Http.Path)

	exitChannel = make(chan int)

	go smtp.Listen(context, exitChannel)

	for {
		select {
		case <-exitChannel:
			logManager().Debug("Received exit signal")
			os.Exit(0)
		}
	}
}
