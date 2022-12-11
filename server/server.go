package server

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/mailpiggy/MailPiggy/authentication"
	"github.com/mailpiggy/MailPiggy/config"
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/mailpiggy/MailPiggy/server/api"
	"github.com/mailpiggy/MailPiggy/server/smtp"
	"github.com/mailpiggy/MailPiggy/server/ui"
	"github.com/mailpiggy/MailPiggy/server/websocket"
	"github.com/mailpiggy/MailPiggy/serverContext"
	"github.com/mailpiggy/MailPiggy/storage"
	"os"
	"strings"
	"time"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server")
	}
	return configuredLogger
}

var exitChannel chan int

func Configure(config *config.AppConfig) *serverContext.Context {
	context := &serverContext.Context{
		Config: *config,
	}

	logger.MinLogLevel = config.Log.Level

	switch config.Storage.Use {
	case "directory":
		context.Storage = storage.CreateDirectoryStorage(config.Storage.Directory.Path)
		if config.Storage.PerRoomLimit > 0 {
			storage.SetPerRoomLimit(config.Storage.PerRoomLimit)
		}
	default:
		panic("Incorrect storage type, Supports: directory")
	}

	switch config.Authentication.Use {
	case "file":
		context.Authentication = authentication.CreateFileAuthentication(config.Authentication.File.Path)
	default:
		panic("Incorrect authentication type, Supports: file")
	}

	context.HttpSession = session.New(session.Config{Expiration: 10 * time.Minute})

	return context
}

func Start(context *serverContext.Context) {
	exitChannel = make(chan int)

	go smtp.Listen(context, exitChannel)

	httpApp := fiber.New()
	if context.Authentication.RequiresAuthentication() {
		httpApp.Use(httpAuthentication(context))
	}

	api.CreateAPIRoutes(context, httpApp)
	ui.CreateUIRoutes(context, httpApp)
	websocket.CreateWebsocket(context, httpApp)

	logManager().Debug(fmt.Sprintf("HTTP Binding to address %s", context.HttpBindAddr()))
	go logger.PanicIfError(httpApp.Listen(context.HttpBindAddr()))

	for {
		select {
		case <-exitChannel:
			logManager().Debug("Received exit signal")
			os.Exit(0)
		}
	}
}

func httpAuthentication(context *serverContext.Context) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		Unauthorized := func() error {
			ctx.Set(fiber.HeaderWWWAuthenticate, "basic realm=Restricted")
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		username, err := context.GetHttpAuthenticatedUser(ctx)

		if err != nil {
			// Set a custom header on all responses:
			auth := ctx.Get(fiber.HeaderAuthorization)

			// Check if the header contains content besides "basic".
			if len(auth) <= 6 || strings.ToLower(auth[:5]) != "basic" {
				return Unauthorized()
			}

			// Decode the header contents
			raw, err := base64.StdEncoding.DecodeString(auth[6:])
			if err != nil {
				return Unauthorized()
			}

			// Get the credentials
			credentials := utils.UnsafeString(raw)

			// Check if the credentials are in the correct form
			// which is "username:password".
			index := strings.Index(credentials, ":")
			if index == -1 {
				return Unauthorized()
			}

			// Get the username and password
			username = credentials[:index]
			password := credentials[index+1:]

			if !context.Authentication.Authenticate(authentication.HTTP, username, password) {
				return Unauthorized()
			}

			if context.SetHttpAuthenticatedUser(ctx, username) != nil {
				logManager().Error(fmt.Sprintf("Error on saving session %s", err.Error()))
				return Unauthorized()
			}
		}

		// Go to next middleware:
		return ctx.Next()
	}
}
