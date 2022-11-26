package storage

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"github.com/mailpiggy/MailPiggy/gounit"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
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

	(*gounit.T)(t).AssertNotError(storage.Delete("", "1"))
	(*gounit.T)(t).AssertEqualsInt(1, storage.Count(""))

	(*gounit.T)(t).AssertNotError(storage.Delete("", "0"))
	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(""))

	(*gounit.T)(t).ExpectError(storage.Delete("", "1"))
	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(""))
}
