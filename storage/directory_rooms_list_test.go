package storage

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/MailHedgehog/gounit"
	"golang.org/x/exp/slices"
	"strings"
	"testing"
)

func TestRoomsList(t *testing.T) {
	room := "foo_bar"

	storage := CreateDirectoryStorage("")

	for i := 0; i < 9; i++ {
		for j := 0; j < 3; j++ {
			msg := &smtpMessage.SMTPMail{
				ID: smtpMessage.MessageID(fmt.Sprint(i)),
			}
			storage.Store(room+fmt.Sprint(i), msg)
		}
	}

	rooms, err := storage.RoomsList(2, 3)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(3, len(rooms))

	if !slices.Contains(rooms, "foo_bar3") {
		t.Errorf("rooms contains expected: %s, got: %s", "foo_bar3", strings.Join(rooms, ", "))
	}
}

func TestRoomsListOutOfRange(t *testing.T) {
	room := "foo_bar"
	storage := CreateDirectoryStorage("")

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			msg := &smtpMessage.SMTPMail{
				ID: smtpMessage.MessageID(fmt.Sprint(i)),
			}
			storage.Store(room+fmt.Sprint(i), msg)
		}
	}

	rooms, err := storage.RoomsList(20, 3)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(0, len(rooms))

	rooms, err = storage.RoomsList(1, 1)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(1, len(rooms))

	rooms, err = storage.RoomsList(1, 10)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(2, len(rooms))

	rooms, err = storage.RoomsList(0, 10)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(3, len(rooms))
}

func TestRoomsListWrongOffset(t *testing.T) {
	storage := CreateDirectoryStorage("")

	rooms, err := storage.RoomsList(20, 3)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(0, len(rooms))

	_, err = storage.RoomsList(-10, 3)
	(*gounit.T)(t).ExpectError(err)

	_, err = storage.RoomsList(10, -3)
	(*gounit.T)(t).ExpectError(err)

	_, err = storage.RoomsList(-10, -3)
	(*gounit.T)(t).ExpectError(err)
}
