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

	for i := 0; i < 5; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID("a" + fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 7; i++ {
		msg := &dto.Message{
			ID: dto.MessageID("b" + fmt.Sprint(i)),
			Raw: &dto.SMTPMessage{
				To:   []string{"foo&bar.com"},
				From: "foo@baz.com",
			},
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 3; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID("c" + fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	(*gounit.T)(t).AssertEqualsInt(15, storage.Count(room))

	messages, count, err := storage.List(room, SearchQuery{
		"to":   "foo@bar",
		"from": "foo@baz",
	}, 2, 3)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(3, len(messages))
	(*gounit.T)(t).AssertEqualsInt(7, count)
	(*gounit.T)(t).AssertEqualsString("b4", string(messages[0].ID))
	(*gounit.T)(t).AssertEqualsString("b3", string(messages[1].ID))
	(*gounit.T)(t).AssertEqualsString("b2", string(messages[2].ID))

	messages, count, err = storage.List(room, SearchQuery{}, 2, 3)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(3, len(messages))
	(*gounit.T)(t).AssertEqualsInt(15, count)
}

func TestListOutOfRange(t *testing.T) {
	room := "foo_bar"
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room))

	for i := 0; i < 5; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID("a" + fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 7; i++ {
		msg := &dto.Message{
			ID: dto.MessageID("b" + fmt.Sprint(i)),
			Raw: &dto.SMTPMessage{
				To:   []string{"foo&bar.com"},
				From: "foo@baz.com",
			},
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 3; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID("c" + fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	(*gounit.T)(t).AssertEqualsInt(15, storage.Count(room))

	messages, count, err := storage.List(room, SearchQuery{
		"to":   "foo@bar",
		"from": "foo@baz",
	}, 9, 3)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(0, len(messages))
	(*gounit.T)(t).AssertEqualsInt(7, count)

	messages, count, err = storage.List(room, SearchQuery{
		"to":   "foo@bar",
		"from": "foo@baz",
	}, 5, 3)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(2, len(messages))
	(*gounit.T)(t).AssertEqualsInt(7, count)

}

func TestListWrongOffset(t *testing.T) {
	room := "foo_bar"
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room))

	for i := 0; i < 5; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID("a" + fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 7; i++ {
		msg := &dto.Message{
			ID: dto.MessageID("b" + fmt.Sprint(i)),
			Raw: &dto.SMTPMessage{
				To:   []string{"foo&bar.com"},
				From: "foo@baz.com",
			},
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 3; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID("c" + fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	(*gounit.T)(t).AssertEqualsInt(15, storage.Count(room))

	_, _, err := storage.List(room, SearchQuery{}, -22, 3)
	(*gounit.T)(t).ExpectError(err)

	_, _, err = storage.List(room, SearchQuery{}, 0, -33)
	(*gounit.T)(t).ExpectError(err)

	_, _, err = storage.List(room, SearchQuery{}, -22, -3)
	(*gounit.T)(t).ExpectError(err)
}
