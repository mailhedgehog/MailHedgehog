package storage

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto"
	"github.com/mailhedgehog/MailHedgehog/gounit"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(""))

	for i := 0; i < 2; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store("", msg)
	}

	(*gounit.T)(t).AssertEqualsInt(2, storage.Count(""))

	message, err := storage.Load("", "1")
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsString("1", string(message.ID))
}
