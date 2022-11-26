package storage

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
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

	message, err := storage.Load("", "1")
	if err != nil {
		t.Error(err)
	}

	if message.ID != "1" {
		t.Errorf("message.ID expected: %s, got: %s", "1", message.ID)
	}
}
