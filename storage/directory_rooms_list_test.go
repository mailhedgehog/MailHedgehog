package storage

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"testing"
	"time"
)

func TestRoomsList(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	for i := 0; i < 15; i++ {
		for j := 0; j < 3; j++ {
			msg := &dto.Message{
				ID:      dto.MessageID(fmt.Sprint(i)),
				Created: time.Now(),
			}
			storage.Store(room+fmt.Sprint(i), msg)
		}
	}

	rooms, err := storage.RoomsList(2, 3)

	if err != nil {
		t.Error(err)
	}

	if len(rooms) != 3 {
		t.Errorf("len(rooms) expected: %d, got: %d", 3, len(rooms))
	}
}

func TestRoomsListOutOfRange(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			msg := &dto.Message{
				ID:      dto.MessageID(fmt.Sprint(i)),
				Created: time.Now(),
			}
			storage.Store(room+fmt.Sprint(i), msg)
		}
	}

	rooms, err := storage.RoomsList(20, 3)
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 0 {
		t.Errorf("len(rooms) expected: %d, got: %d", 0, len(rooms))
	}

	rooms, err = storage.RoomsList(1, 1)
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 1 {
		t.Errorf("len(rooms) expected: %d, got: %d", 1, len(rooms))
	}

	rooms, err = storage.RoomsList(1, 10)
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 2 {
		t.Errorf("len(rooms) expected: %d, got: %d", 2, len(rooms))
	}

	rooms, err = storage.RoomsList(0, 10)
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 3 {
		t.Errorf("len(rooms) expected: %d, got: %d", 3, len(rooms))
	}
}

func TestRoomsListWrongOffset(t *testing.T) {
	storage := CreateDirectoryStorage("")

	rooms, err := storage.RoomsList(20, 3)
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 0 {
		t.Errorf("len(rooms) expected: %d, got: %d", 0, len(rooms))
	}

	_, err = storage.RoomsList(-10, 3)
	if err == nil {
		t.Error("Expect error")
	}

	_, err = storage.RoomsList(10, -3)
	if err == nil {
		t.Error("Expect error")
	}

	_, err = storage.RoomsList(-10, -3)
	if err == nil {
		t.Error("Expect error")
	}
}
