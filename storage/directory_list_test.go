package storage

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"github.com/mailpiggy/MailPiggy/gounit"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	room := "foo_bar"
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room))

	for i := 0; i < 15; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	messages, err := storage.List(room, 2, 3)
	(*gounit.T)(t).AssertNotError(err)

	(*gounit.T)(t).AssertEqualsInt(3, len(messages))
	(*gounit.T)(t).AssertEqualsString("12", string(messages[0].ID))
	(*gounit.T)(t).AssertEqualsString("11", string(messages[1].ID))
	(*gounit.T)(t).AssertEqualsString("10", string(messages[2].ID))
}

func TestListOutOfRange(t *testing.T) {
	room := "foo_bar"
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room))

	for i := 0; i < 3; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
	}

	messages, err := storage.List(room, 20, 3)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(0, len(messages))

	messages, err = storage.List(room, 1, 1)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(1, len(messages))

	messages, err = storage.List(room, 1, 10)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(2, len(messages))

	messages, err = storage.List(room, 0, 10)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(3, len(messages))
}

func TestListWrongOffset(t *testing.T) {
	room := "foo_bar"
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room))

	messages, err := storage.List(room, 20, 3)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(0, len(messages))

	_, err = storage.List(room, -10, 3)
	(*gounit.T)(t).ExpectError(err)

	_, err = storage.List(room, 10, -3)
	(*gounit.T)(t).ExpectError(err)

	_, err = storage.List(room, -10, -3)
	(*gounit.T)(t).ExpectError(err)
}
