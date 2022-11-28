package http

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/mailpiggy/MailPiggy/serverContext"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server.http")
	}
	return configuredLogger
}

func Listen(context *serverContext.Context, exitCh chan int) {
	logManager().Debug(fmt.Sprintf("HTTP Binding to address %s", context.HttpBindAddr()))

	//router := pat.New()
	// registerCallback

	logManager().Warning("TODO: implement http.")
}
