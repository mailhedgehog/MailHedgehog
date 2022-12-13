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

	ui.Get("/mh-configuration.json", indexHandler(context))

	ui.Static("/", context.Config.Http.AssetsRoot)
}

func indexHandler(context *serverContext.Context) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"http": fiber.Map{
				"baseUrl": "//" + context.HttpBindAddr() + context.PathWithPrefix(""),
			},
		})
	}
}
