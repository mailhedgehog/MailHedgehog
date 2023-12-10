package storage

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/gounit"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	room := "foo_bar"
	storage := CreateDirectoryStorage("")

	(*gounit.T)(t).AssertEqualsInt(0, storage.Count(room))

	for i := 0; i < 5; i++ {
		id := smtpMessage.MessageID("a" + fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
		}
		storedId, err := storage.Store(room, msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 7; i++ {
		id := smtpMessage.MessageID("b" + fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
			Origin: &smtpMessage.SMTPMessage{
				To:   []string{"foo&bar.com"},
				From: "foo@baz.com",
			},
		}
		storedId, err := storage.Store(room, msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 3; i++ {
		id := smtpMessage.MessageID("c" + fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
		}
		storedId, err := storage.Store(room, msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
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
		id := smtpMessage.MessageID("a" + fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
		}
		storedId, err := storage.Store(room, msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 7; i++ {
		id := smtpMessage.MessageID("b" + fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
			Origin: &smtpMessage.SMTPMessage{
				To:   []string{"foo&bar.com"},
				From: "foo@baz.com",
			},
		}
		storedId, err := storage.Store(room, msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 3; i++ {
		id := smtpMessage.MessageID("c" + fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
		}
		storedId, err := storage.Store(room, msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
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
		id := smtpMessage.MessageID("a" + fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
		}
		storedId, err := storage.Store(room, msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 7; i++ {
		id := smtpMessage.MessageID("b" + fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
			Origin: &smtpMessage.SMTPMessage{
				To:   []string{"foo&bar.com"},
				From: "foo@baz.com",
			},
		}
		storedId, err := storage.Store(room, msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
		// Fix for correct sorting
		time.Sleep(10000000)
	}
	for i := 0; i < 3; i++ {
		id := smtpMessage.MessageID("c" + fmt.Sprint(i))
		msg := &smtpMessage.SMTPMail{
			ID: id,
		}
		storedId, err := storage.Store(room, msg)
		(*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
		(*gounit.T)(t).AssertNotError(err)
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
