package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mailpiggy/MailPiggy/logger"
	v1 "github.com/mailpiggy/MailPiggy/server/api/v1"
	"github.com/mailpiggy/MailPiggy/serverContext"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server.api")
	}
	return configuredLogger
}

func CreateAPIRoutes(context *serverContext.Context, httpApp *fiber.App) {

	api := httpApp.Group(context.PathWithPrefix("api"), func(c *fiber.Ctx) error {
		return c.Next()
	})

	v1.CreateAPIV1Routes(context, api)
}
