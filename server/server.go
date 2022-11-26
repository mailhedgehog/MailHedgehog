package server

import (
	"github.com/mailpiggy/MailPiggy/logger"
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

func Start(context *Context) {
	logManager().Warning("TODO: SMTP Implement", context.Config.Smtp.Port)
	logManager().Warning("TODO: HTTP Implement", context.Config.Http.Port)
	logManager().Warning("TODO: UI Implement", context.Config.Http.Path)

	exitChannel = make(chan int)

	go smtpListen(context, exitChannel)

	for {
		select {
		case <-exitChannel:
			logManager().Debug("Received exit signal")
			os.Exit(0)
		}
	}
}
