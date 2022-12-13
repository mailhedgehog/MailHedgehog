package storage

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto"
	"github.com/mailhedgehog/MailHedgehog/gounit"
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

	(*gounit.T)(t).AssertEqualsInt(15, storage.RoomsCount())

	(*gounit.T)(t).AssertNotError(storage.DeleteRoom(room + fmt.Sprint(5)))
	(*gounit.T)(t).AssertNotError(storage.DeleteRoom(room + fmt.Sprint(10)))

	(*gounit.T)(t).AssertEqualsInt(13, storage.RoomsCount())
}
