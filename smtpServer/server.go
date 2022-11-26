package smtpServer

import "github.com/mailpiggy/MailPiggy/logger"

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("smtpServer")
	}
	return configuredLogger
}
