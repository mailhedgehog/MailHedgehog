package server

import (
	"github.com/mailpiggy/MailPiggy/logger"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server")
	}
	return configuredLogger
}

func Start(context *Context) {
	logManager().Warning("TODO: SMTP Implement", context.Config.Smtp.Port)
	logManager().Warning("TODO: HTTP Implement", context.Config.Http.Port)
	logManager().Warning("TODO: UI Implement", context.Config.Http.Path)
}
