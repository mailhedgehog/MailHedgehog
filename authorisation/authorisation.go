package authorisation

import (
	"github.com/mailpiggy/MailPiggy/logger"
)

type AuthorisationType = string

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("authorisation")
	}
	return configuredLogger
}

// Authorisation interface represents a backend flow to store or retrieve messages
type Authorisation interface {
	// Authorised check us credentials valid
	Authorised(authType AuthorisationType, username string, password string) bool
}

const (
	HTTP AuthorisationType = "http"
	SMTP                   = "smtp"
)

type userInfo struct {
	username string
	httpPass string
	smtpPass string
}
