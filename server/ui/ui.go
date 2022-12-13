package ui

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server.ui")
	}
	return configuredLogger
}

func CreateUIRoutes(context *serverContext.Context, httpApp *fiber.App) {
	ui := httpApp.Group(context.PathWithPrefix(""), func(c *fiber.Ctx) error {
		return c.Next()
	})

	ui.Static("/", context.Config.Http.AssetsRoot)
}
