package storage

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
	storage := CreateDirectoryStorage("")

	if storage.Count("") != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(""))
	}

	for i := 0; i < 2; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store("", msg)
	}

	if storage.Count("") != 2 {
		t.Errorf("storage.Count() expected: %d, got: %d", 2, storage.Count(""))
	}

	storage.Delete("", "2")
	if storage.Count("") != 1 {
		t.Errorf("storage.Count() expected: %d, got: %d", 1, storage.Count(""))
	}

	storage.Delete("", "1")
	if storage.Count("") != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(""))
	}

	storage.Delete("", "1")
	if storage.Count("") != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(""))
	}
}
