package authentication

import (
	"github.com/mailhedgehog/MailHedgehog/logger"
	"golang.org/x/crypto/bcrypt"
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
	// UsernamePresent check is username exists
	UsernamePresent(username string) bool
	// AddUser to auth storage
	AddUser(username string, httpPassHash string, smtpPassHash string) error
	// UpdateUser in auth storage
	UpdateUser(username string, httpPassHash string, smtpPassHash string) error
	// DeleteUser from auth storage
	DeleteUser(username string) error
	// ListUsers from auth storage
	ListUsers(searchQuery string, offset, limit int) ([]UserResource, int, error)
}

const (
	HTTP AuthenticationType = "http"
	SMTP                    = "smtp"
)

type UserResource struct {
	Username string
}

func CreatePasswordHash(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
