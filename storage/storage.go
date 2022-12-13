package storage

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto"
	"github.com/mailhedgehog/MailHedgehog/logger"
)

type SearchQuery = map[string]string
type Room = string

var configuredLogger *logger.Logger
var perRoomLimit = 0

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("storage")
	}
	return configuredLogger
}

func SetPerRoomLimit(limit int) {
	perRoomLimit = limit
	logManager().Debug(fmt.Sprintf("New per room messages limit: %d", perRoomLimit))
}

// Storage interface represents a backend flow to store or retrieve messages
type Storage interface {
	// Store `message` to specific `room`
	Store(room Room, message *dto.Message) (dto.MessageID, error)
	// List retrieve list of messages based on `query` starting with `offset` index and count limited by `limit`
	// `query` - represents of key->value map, where key is search parameter
	List(room Room, query SearchQuery, offset, limit int) ([]dto.Message, int, error)
	// Count total messages in storage
	Count(room Room) int
	// Delete delete specific message from storage by `messageId`
	Delete(room Room, messageId dto.MessageID) error
	// Load find specific message from storage by `messageId`
	Load(room Room, messageId dto.MessageID) (*dto.Message, error)

	// RoomsList returns list of rooms in system
	RoomsList(offset, limit int) ([]Room, error)
	// RoomsCount total count rooms in storage
	RoomsCount() int
	// DeleteRoom delete all messages in room from storage
	DeleteRoom(room Room) error
}
