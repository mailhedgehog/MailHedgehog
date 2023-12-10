package websocket

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"github.com/mailhedgehog/logger"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server.websocket")
	}
	return configuredLogger
}

type client struct{}
type BroadcastMessage struct {
	Room    string
	Message string
}

// TODO: clear old sockets
var clients = make(map[*websocket.Conn]client)
var registerClient = make(chan *websocket.Conn)
var unregisterClient = make(chan *websocket.Conn)
var BroadcastToClient = make(chan BroadcastMessage)

func runWebsocketHub() {
	for {
		select {
		case connection := <-registerClient:
			clients[connection] = client{}
			logManager().Debug("connection registered")

		case connection := <-unregisterClient:
			// Remove the client from the hub
			delete(clients, connection)
			logManager().Debug("connection unregistered")
		case broadcastMessage := <-BroadcastToClient:
			logManager().Debug(fmt.Sprintf("sending to '%s' message: %s", broadcastMessage.Room, broadcastMessage.Message))

			// Send the message to all clients
			for connection := range clients {
				if fmt.Sprint(connection.Locals("room")) != broadcastMessage.Room {
					return
				}

				defer func() {
					if err := recover(); err != nil {
						logManager().Error(fmt.Sprint("panic occurred:", err))
						go runWebsocketHub()
					}
				}()
				err := connection.WriteMessage(websocket.TextMessage, []byte(broadcastMessage.Message))
				if err != nil {
					logManager().Error(fmt.Sprintf("write error: %s", err.Error()))
					unregisterClient <- connection
					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
				}
			}

		}
	}
}

func CreateWebsocket(context *serverContext.Context, httpApp *fiber.App) {
	httpApp.Use(context.PathWithPrefix("websocket"), func(ctx *fiber.Ctx) error {
		username, _ := context.GetHttpAuthenticatedUser(ctx)
		ctx.Locals("room", username)
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(ctx) {
			ctx.Locals("allowed", true)
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	go runWebsocketHub()

	httpApp.Get(context.PathWithPrefix("websocket"), websocket.New(func(c *websocket.Conn) {
		defer func() {
			unregisterClient <- c
			c.Close()
		}()

		if context.Config.Websocket.MaxConnection <= len(clients) {
			c.WriteMessage(websocket.CloseMessage, []byte("Max limit of connections"))
			unregisterClient <- c
			c.Close()
			return
		}
		// Register the client
		registerClient <- c

		for {
			messageType, msg, err := c.ReadMessage()
			if err != nil {
				logManager().Error(fmt.Sprintf("read error: %s", err.Error()))
				break
			}

			logManager().Debug(fmt.Sprintf("received msg: %s", msg))

			if messageType == websocket.TextMessage {
				// app not need handle received messages
			} else {
				logManager().Error(fmt.Sprintf("unrecognised message type: %d", messageType))
			}

			if err = c.WriteMessage(messageType, msg); err != nil {
				logManager().Error(fmt.Sprintf("write error: %s", msg))
				break
			}
		}

	}))
}
