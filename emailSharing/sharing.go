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
	Find(id string) (*EmailSharingRecord, error)
}
