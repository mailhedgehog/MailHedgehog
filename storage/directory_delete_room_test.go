package storage

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/dto"
	"testing"
	"time"
)

func TestDeleteRoom(t *testing.T) {
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

	if storage.RoomsCount() != 15 {
		t.Errorf("storage.RoomsCount() expected: %d, got: %d", 15, storage.RoomsCount())
	}
	err := storage.DeleteRoom(room + fmt.Sprint(5))
	if err != nil {
		t.Error(err)
	}
	err = storage.DeleteRoom(room + fmt.Sprint(10))
	if err != nil {
		t.Error(err)
	}
	if storage.RoomsCount() != 13 {
		t.Errorf("storage.RoomsCount() expected: %d, got: %d", 13, storage.RoomsCount())
	}
}
