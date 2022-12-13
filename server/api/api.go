package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mailhedgehog/MailHedgehog/logger"
	v1 "github.com/mailhedgehog/MailHedgehog/server/api/v1"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
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

	logManager().Debug(context.Config.Http.AllowOrigins)
	api.Use(cors.New(cors.Config{
		AllowOrigins: context.Config.Http.AllowOrigins,
	}))

	v1.CreateAPIV1Routes(context, api)
}
