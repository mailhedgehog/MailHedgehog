package storage

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestLoad(t *testing.T) {
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(""))

	for i := 0; i < 2; i++ {
		id := smtpMessage.MessageID(fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
		}

		storedId, err := storage.Store("", msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
	}

	(*gounit.T)(t).AssertEqualsInt(2, storage.Count(""))

	message, err := storage.Load("", "1")
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsString("1", string(message.ID))
}
