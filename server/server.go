package server

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/mailhedgehog/MailHedgehog/config"
	"github.com/mailhedgehog/MailHedgehog/emailSharing"
	"github.com/mailhedgehog/MailHedgehog/server/api"
	"github.com/mailhedgehog/MailHedgehog/server/smtp"
	"github.com/mailhedgehog/MailHedgehog/server/ui"
	"github.com/mailhedgehog/MailHedgehog/server/websocket"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"github.com/mailhedgehog/MailHedgehog/storage"
	"github.com/mailhedgehog/authenticationFile"
	"github.com/mailhedgehog/authenticationMongo"
	"github.com/mailhedgehog/logger"
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
	case "mongodb":
		context.Storage = storage.CreateMongoDbStorage(
			config.DB.GetMongoDBConnection(
				config.Storage.MongoDB.Connection,
			).Collection(
				config.Storage.MongoDB.Collection,
			),
		)
		if config.Storage.PerRoomLimit > 0 {
			storage.SetPerRoomLimit(config.Storage.PerRoomLimit)
		}
	default:
		panic("Incorrect storage type, Supports: directory")
	}

	switch config.Authentication.Use {
	case "file":
		context.Authentication = authenticationFile.CreateFileAuthentication(&config.Authentication.File, &config.Authentication.Config)
	case "mongodb":
		context.Authentication = authenticationMongo.CreateMongoDbAuthentication(
			config.DB.GetMongoDBConnection(
				config.Authentication.MongoDB.Connection,
			).Collection(
				config.Authentication.MongoDB.Collection,
			),
			&config.Authentication.Config,
		)
	default:
		panic("Incorrect authentication type, Supports: file, mongodb")
	}

	switch config.Sharing.Use {
	case "csv":
		context.Sharing = emailSharing.CreateSharingEmailUsingCSV(config.Sharing.CSV.Path)
	case "mongodb":
		// TODO: add implementation
	default:
		// Nothing to do
	}

	context.HttpSession = session.New(session.Config{Expiration: 10 * time.Minute})

	return context
}

func Start(context *serverContext.Context) {
	exitChannel = make(chan int)

	go smtp.Listen(context, exitChannel)

	httpApp := fiber.New()

	if len(context.Config.Http.AllowOrigins) > 0 {
		logManager().Debug(fmt.Sprintf("Allow origins: %s", context.Config.Http.AllowOrigins))
		httpApp.Use(cors.New(cors.Config{
			AllowOrigins: context.Config.Http.AllowOrigins,
		}))
	}

	if context.Authentication.Dashboard().RequiresAuthentication() {
		switch context.Config.Authentication.Type {
		case "basic":
			httpApp.Use(authenticationBasic(context))
		}
	}

	api.CreateAPIRoutes(context, httpApp)
	ui.CreateUIRoutes(context, httpApp)
	if context.Config.Http.Websocket {
		websocket.CreateWebsocket(context, httpApp)
	}

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

func authenticationBasic(context *serverContext.Context) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		Unauthorized := func() error {
			ctx.Set(fiber.HeaderWWWAuthenticate, "basic realm=Restricted")
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		_, err := context.GetHttpAuthenticatedUser(ctx)

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
			username := credentials[:index]
			password := credentials[index+1:]

			if !context.Authentication.Dashboard().ViaPasswordAuthentication().Authenticate(username, password) {
				return Unauthorized()
			}

			if _, err := context.SetHttpAuthenticatedUser(ctx, username); err != nil {
				logManager().Error(fmt.Sprintf("Error on saving session %s", err.Error()))
				return Unauthorized()
			}
		}

		// Go to next middleware:
		return ctx.Next()
	}
}
