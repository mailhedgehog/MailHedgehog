package authentication

import (
	"github.com/mailpiggy/MailPiggy/logger"
)

type AuthenticationType = string

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("authentication")
	}
	return configuredLogger
}

// Authentication interface represents a backend flow to store or retrieve messages
type Authentication interface {
	// Authenticate check us credentials valid
	Authenticate(authType AuthenticationType, username string, password string) bool
}

const (
	HTTP AuthenticationType = "http"
	SMTP                    = "smtp"
)

type userInfo struct {
	username string
	httpPass string
	smtpPass string
}
