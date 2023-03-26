package storage

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/MailHedgehog/gounit"
	"os"
	"testing"
)

func TestStoreDefaultRoom(t *testing.T) {
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(""))

	for i := 0; i < 2; i++ {
		msg := &smtpMessage.SMTPMail{
			ID: smtpMessage.MessageID(fmt.Sprint(i)),
		}
		storage.Store("", msg)
	}

	(*gounit.T)(t).AssertEqualsInt(2, storage.Count(""))
}

func TestStore(t *testing.T) {
	room1 := "foo_bar"
	room2 := "baz_bar"
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room1))
	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room2))

	for i := 0; i < 3; i++ {
		msg := &smtpMessage.SMTPMail{
			ID: smtpMessage.MessageID(fmt.Sprint(i)),
		}
		storage.Store(room1, msg)
	}

	for i := 0; i < 4; i++ {
		msg := &smtpMessage.SMTPMail{
			ID: smtpMessage.MessageID(fmt.Sprint(i)),
		}
		storage.Store(room2, msg)
	}

	(*gounit.T)(t).AssertEqualsInt(3, storage.Count(room1))
	(*gounit.T)(t).AssertEqualsInt(4, storage.Count(room2))
}

func TestStoreCustomPath(t *testing.T) {
	room := "foo_bar"
	pathToStore := "relative_path_foo"

	_, err := os.Stat(pathToStore)
	(*gounit.T)(t).ExpectError(err)

	storage := CreateDirectoryStorage(pathToStore)

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room))

	for i := 0; i < 2; i++ {
		msg := &smtpMessage.SMTPMail{
			ID: smtpMessage.MessageID(fmt.Sprint(i)),
		}
		storage.Store(room, msg)
	}

	(*gounit.T)(t).AssertEqualsInt(2, storage.Count(room))

	_, err = os.Stat(pathToStore)
	(*gounit.T)(t).AssertNotError(err)

	os.RemoveAll(pathToStore)
}

func TestStoreWithLimit(t *testing.T) {
	room1 := "foo_bar"

	SetPerRoomLimit(3)
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room1))

	for i := 0; i < 15; i++ {
		msg := &smtpMessage.SMTPMail{
			ID: smtpMessage.MessageID(fmt.Sprint(i)),
		}
		storage.Store(room1, msg)
	}

	(*gounit.T)(t).AssertEqualsInt(3, storage.Count(room1))

	for i := 0; i < 2; i++ {
		msg := &smtpMessage.SMTPMail{
			ID: smtpMessage.MessageID(fmt.Sprint(i)),
		}
		storage.Store(room1, msg)
	}

	(*gounit.T)(t).AssertEqualsInt(2, storage.Count(room1))

	SetPerRoomLimit(0)
}
