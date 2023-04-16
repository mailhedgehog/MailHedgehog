package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
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

	if context.Authentication.RequiresAuthentication() {
		switch context.Config.Authentication.Type {
		case "internal":
			api.Use(skip.New(authenticationInternal(context), func(ctx *fiber.Ctx) bool {
				return false
			}))
		}
	}

	v1.CreateAPIV1Routes(context, api)
}

func authenticationInternal(context *serverContext.Context) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		Unauthorized := func() error {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		_, err := context.GetHttpAuthenticatedUser(ctx)

		if err != nil {
			return Unauthorized()
		}

		// Go to next middleware:
		return ctx.Next()
	}
}
