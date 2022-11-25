package storage

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	if storage.Count(room) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room))
	}

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
	if storage.Count(room) != 15 {
		t.Errorf("storage.Count() expected: %d, got: %d", 15, storage.Count(room))
	}

	messages, count, err := storage.Search(room, SearchQuery{
		"to":   "foo@bar",
		"from": "foo@baz",
	}, 2, 3)
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 3 {
		t.Errorf("len(messages) expected: %d, got: %d", 3, len(messages))
	}
	if count != 7 {
		t.Errorf("count expected: %d, got: %d", 7, count)
	}
	if string(messages[0].ID) != "b4" {
		t.Errorf("Message.ID expected: %s, got: %s", "b4", string(messages[0].ID))
	}
	if string(messages[1].ID) != "b3" {
		t.Errorf("Message.ID expected: %s, got: %s", "b3", string(messages[1].ID))
	}
	if string(messages[2].ID) != "b2" {
		t.Errorf("Message.ID expected: %s, got: %s", "b2", string(messages[2].ID))
	}

	messages, count, err = storage.Search(room, SearchQuery{}, 2, 3)
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 3 {
		t.Errorf("len(messages) expected: %d, got: %d", 3, len(messages))
	}
	if count != 15 {
		t.Errorf("count expected: %d, got: %d", 5, count)
	}
}

func TestSearchOutOfRange(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	if storage.Count(room) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room))
	}

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
	if storage.Count(room) != 15 {
		t.Errorf("storage.Count() expected: %d, got: %d", 15, storage.Count(room))
	}

	messages, count, err := storage.Search(room, SearchQuery{
		"to":   "foo@bar",
		"from": "foo@baz",
	}, 9, 3)
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 0 {
		t.Errorf("len(messages) expected: %d, got: %d", 0, len(messages))
	}
	if count != 7 {
		t.Errorf("count expected: %d, got: %d", 7, count)
	}

	messages, count, err = storage.Search(room, SearchQuery{
		"to":   "foo@bar",
		"from": "foo@baz",
	}, 5, 3)
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 2 {
		t.Errorf("len(messages) expected: %d, got: %d", 2, len(messages))
	}
	if count != 7 {
		t.Errorf("count expected: %d, got: %d", 7, count)
	}

}

func TestSearchWrongOffset(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	if storage.Count(room) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room))
	}

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
	if storage.Count(room) != 15 {
		t.Errorf("storage.Count() expected: %d, got: %d", 15, storage.Count(room))
	}

	_, _, err := storage.Search(room, SearchQuery{}, -22, 3)
	if err == nil {
		t.Error("Expect error")
	}

	_, _, err = storage.Search(room, SearchQuery{}, 0, -33)
	if err == nil {
		t.Error("Expect error")
	}

	_, _, err = storage.Search(room, SearchQuery{}, -22, -3)
	if err == nil {
		t.Error("Expect error")
	}
}
