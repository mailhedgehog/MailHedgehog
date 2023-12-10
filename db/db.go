package db

import "github.com/mailhedgehog/logger"

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("db")
	}
	return configuredLogger
}
