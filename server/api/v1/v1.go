package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/mailpiggy/MailPiggy/serverContext"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server.api.v1")
	}
	return configuredLogger
}

type ApiV1 struct {
}

func CreateAPIV1Routes(context *serverContext.Context, api fiber.Router) {
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	v1.Get("/emails", func(c *fiber.Ctx) error {
		username := context.GetHttpSession(c).Get("CurrentUser")
		logManager().Error("Not implemented")
		return c.SendString("List messages" + fmt.Sprint(username))
	})
	v1.Delete("/emails", func(c *fiber.Ctx) error {
		logManager().Error("Not implemented")
		return c.SendString("Delete messages")
	})
	v1.Get("/emails/:id", func(c *fiber.Ctx) error {
		logManager().Error("Not implemented")
		return c.SendString("Show message")
	})
	v1.Delete("/emails/:id", func(c *fiber.Ctx) error {
		logManager().Error("Not implemented")
		return c.SendString("Delete message")
	})
	v1.Post("/emails/:id/release", func(c *fiber.Ctx) error {
		logManager().Error("Not implemented")
		return c.SendString("Release one message")
	})
	v1.Get("/emails/:id/download", func(c *fiber.Ctx) error {
		logManager().Error("Not implemented")
		return c.SendString("Download one message")
	})
	v1.Get("/emails/:id/attachment/:attachmentId/download", func(c *fiber.Ctx) error {
		logManager().Error("Not implemented")
		return c.SendString("Download attachment")
	})
	v1.Get("/websocket", func(c *fiber.Ctx) error {
		logManager().Error("Not implemented")
		return c.SendString("websocket")
	})
}
