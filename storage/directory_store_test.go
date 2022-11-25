package storage

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"os"
	"testing"
	"time"
)

func TestStoreDefaultRoom(t *testing.T) {
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
}

func TestStore(t *testing.T) {
	room1 := "foo_bar"
	room2 := "baz_bar"
	storage := CreateDirectoryStorage("")

	if storage.Count(room1) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room1))
	}
	if storage.Count(room2) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room2))
	}

	for i := 0; i < 3; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room1, msg)
	}

	for i := 0; i < 4; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room2, msg)
	}

	if storage.Count(room1) != 3 {
		t.Errorf("storage.Count() expected: %d, got: %d", 3, storage.Count(room1))
	}
	if storage.Count(room2) != 4 {
		t.Errorf("storage.Count() expected: %d, got: %d", 4, storage.Count(room2))
	}
}

func TestStoreCustomPath(t *testing.T) {
	room := "foo_bar"
	pathToStore := "relative_path_foo"

	if _, err := os.Stat(pathToStore); err == nil {
		t.Errorf("storage.Count() directory already exists: %s", pathToStore)
	}

	storage := CreateDirectoryStorage(pathToStore)

	if storage.Count(room) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room))
	}

	for i := 0; i < 2; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room, msg)
	}

	if storage.Count(room) != 2 {
		t.Errorf("storage.Count() expected: %d, got: %d", 2, storage.Count(room))
	}

	if _, err := os.Stat(pathToStore); err != nil {
		t.Errorf("storage.Count() directory not exists: %s", pathToStore)
	}

	os.RemoveAll(pathToStore)
}

func TestStoreWithLimit(t *testing.T) {
	room1 := "foo_bar"

	SetPerRoomLimit(3)
	storage := CreateDirectoryStorage("")

	if storage.Count(room1) != 0 {
		t.Errorf("storage.Count() expected: %d, got: %d", 0, storage.Count(room1))
	}

	for i := 0; i < 15; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room1, msg)
	}

	if storage.Count(room1) != 3 {
		t.Errorf("storage.Count() expected: %d, got: %d", 3, storage.Count(room1))
	}

	for i := 0; i < 2; i++ {
		msg := &dto.Message{
			ID:      dto.MessageID(fmt.Sprint(i)),
			Created: time.Now(),
		}
		storage.Store(room1, msg)
	}

	if storage.Count(room1) != 2 {
		t.Errorf("storage.Count() expected: %d, got: %d", 2, storage.Count(room1))
	}

	SetPerRoomLimit(0)
}
