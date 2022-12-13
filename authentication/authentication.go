package authentication

import (
	"github.com/mailhedgehog/MailHedgehog/logger"
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
	// RequiresAuthentication check is supports auth
	RequiresAuthentication() bool
	// Authenticate check is credentials valid
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
