package storage

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/MailHedgehog/gounit"
	"testing"
)

func TestDeleteRoom(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	for i := 0; i < 15; i++ {
		for j := 0; j < 3; j++ {
			id := smtpMessage.MessageID(fmt.Sprint(i))
			msg := &smtpMessage.SMTPMail{
				ID: id,
			}
			storedId, err := storage.Store(room+fmt.Sprint(i), msg)
			(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
			(*gounit.T)(t).AssertNotError(err)
		}
	}

	(*gounit.T)(t).AssertEqualsInt(15, storage.RoomsCount())

	(*gounit.T)(t).AssertNotError(storage.DeleteRoom(room + fmt.Sprint(5)))
	(*gounit.T)(t).AssertNotError(storage.DeleteRoom(room + fmt.Sprint(10)))

	(*gounit.T)(t).AssertEqualsInt(13, storage.RoomsCount())
}
