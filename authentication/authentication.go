package authentication

import (
	"github.com/mailhedgehog/logger"
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
	// AuthenticateSMTPViaIP check is IP cab bypass password auth
	AuthenticateSMTPViaIP(username string, ip string) bool
	// SmtpIpIsWhitelisted check is IP allow to receive mails
	SmtpIpIsWhitelisted(username string, ip string) bool
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
	// AddNoPassSmtpIp to auth storage related to user
	AddNoPassSmtpIp(username string, ip string) error
	// DeleteNoPassSmtpIp from auth storage related to user
	DeleteNoPassSmtpIp(username string, ip string) error
	// ClearAllNoPassSmtpIps in auth storage related to user
	ClearAllNoPassSmtpIps(username string) error
}

const (
	HTTP AuthenticationType = "http"
	SMTP                    = "smtp"
)

type UserResource struct {
	Username      string
	NoPassIPs     []string
	RestrictedIPs []string
	LoginEmails   []string
}

func CreatePasswordHash(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
