package emailSharing

import (
	"github.com/mailhedgehog/MailHedgehog/logger"
	"time"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("emailSharing")
	}
	return configuredLogger
}

type EmailSharingRecord struct {
	Id        string
	Room      string
	EmailId   string
	ExpiredAt time.Time
}

type EmailSharing interface {
	Create(emailSharingRecord *EmailSharingRecord) (*EmailSharingRecord, error)
	Find(id string) (*EmailSharingRecord, error)
	DeleteExpired() (bool, error)
}
