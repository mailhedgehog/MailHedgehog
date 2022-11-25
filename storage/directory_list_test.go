package storage

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	if storage.Count(room) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room))
	}

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

	if err != nil {
		t.Error(err)
	}

	if len(messages) != 3 {
		t.Errorf("len(messages) expected: %d, got: %d", 3, len(messages))
	}
	if string(messages[0].ID) != "12" {
		t.Errorf("Message.ID expected: %s, got: %s", "12", string(messages[0].ID))
	}
	if string(messages[1].ID) != "11" {
		t.Errorf("Message.ID expected: %s, got: %s", "11", string(messages[1].ID))
	}
	if string(messages[2].ID) != "10" {
		t.Errorf("Message.ID expected: %s, got: %s", "10", string(messages[2].ID))
	}
}

func TestListOutOfRange(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	if storage.Count(room) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room))
	}

	for i := 0; i < 3; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
	}

	messages, err := storage.List(room, 20, 3)
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 0 {
		t.Errorf("len(messages) expected: %d, got: %d", 0, len(messages))
	}

	messages, err = storage.List(room, 1, 1)
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 1 {
		t.Errorf("len(messages) expected: %d, got: %d", 1, len(messages))
	}

	messages, err = storage.List(room, 1, 10)
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 2 {
		t.Errorf("len(messages) expected: %d, got: %d", 2, len(messages))
	}

	messages, err = storage.List(room, 0, 10)
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 3 {
		t.Errorf("len(messages) expected: %d, got: %d", 3, len(messages))
	}
}

func TestListWrongOffset(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	if storage.Count(room) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room))
	}

	messages, err := storage.List(room, 20, 3)
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 0 {
		t.Errorf("len(messages) expected: %d, got: %d", 0, len(messages))
	}

	_, err = storage.List(room, -10, 3)
	if err == nil {
		t.Error("Expect error")
	}

	_, err = storage.List(room, 10, -3)
	if err == nil {
		t.Error("Expect error")
	}

	_, err = storage.List(room, -10, -3)
	if err == nil {
		t.Error("Expect error")
	}
}
