package ui

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"github.com/mailhedgehog/logger"
	mailHedgehogUi "github.com/mailhedgehog/ui"
	"net/http"
	"os"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server.ui")
	}
	return configuredLogger
}

func CreateUIRoutes(context *serverContext.Context, httpApp *fiber.App) {
	uiRouter := httpApp.Group(context.PathWithPrefix(""), func(c *fiber.Ctx) error {
		return c.Next()
	})

	uiRouter.Get("/mh-configuration.json", configurationHandler(context))

	if len(context.Config.Http.AssetsRoot) > 0 {
		uiRouter.Static("/", context.Config.Http.AssetsRoot)
	} else {
		uiRouter.Use("/", filesystem.New(filesystem.Config{
			Root:       http.FS(mailHedgehogUi.DirStaticEmbed),
			PathPrefix: mailHedgehogUi.DirStaticName,
			Browse:     false,
		}))
	}
}

func configurationHandler(context *serverContext.Context) func(c *fiber.Ctx) error {
	baseUrl := context.Config.Http.BaseUrl
	if len(baseUrl) <= 0 {
		baseUrl = "//" + context.HttpBindAddr() + context.PathWithPrefix("")
	}

	var uiData interface{}
	uiConfigFileName := context.Config.UI.File
	if len(uiConfigFileName) > 0 {
		if _, err := os.Stat(uiConfigFileName); err == nil {
			bytes, err := os.ReadFile(uiConfigFileName)
			if err == nil {
				err := json.Unmarshal([]byte(bytes), &uiData)
				if err != nil {
					uiData = nil
				}
			}
		}
	}

	authType := ""
	if context.Authentication.Dashboard().RequiresAuthentication() {
		authType = context.Config.Authentication.Type
	}

	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"http": fiber.Map{
				"baseUrl":   baseUrl,
				"auth":      authType,
				"websocket": context.Config.Http.Websocket,
			},
			"ui": uiData,
			"sharing": fiber.Map{
				"enabled": context.Sharing != nil,
			},
		})
	}
}
