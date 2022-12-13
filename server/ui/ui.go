package ui

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"net/http"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server.ui")
	}
	return configuredLogger
}

//go:embed static/*
var EmbedDirStatic embed.FS

func CreateUIRoutes(context *serverContext.Context, httpApp *fiber.App) {
	ui := httpApp.Group(context.PathWithPrefix(""), func(c *fiber.Ctx) error {
		return c.Next()
	})

	ui.Get("/mh-configuration.json", indexHandler(context))

	if len(context.Config.Http.AssetsRoot) > 0 {
		ui.Static("/", context.Config.Http.AssetsRoot)
	} else {
		ui.Use("/", filesystem.New(filesystem.Config{
			Root:       http.FS(EmbedDirStatic),
			PathPrefix: "static",
			Browse:     false,
		}))
	}
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
