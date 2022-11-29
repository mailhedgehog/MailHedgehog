package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/mailpiggy/MailPiggy/serverContext"
	"github.com/mailpiggy/MailPiggy/storage"
	"math"
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

	if context.Authentication.RequiresAuthentication() {
		v1.Get("/logout", func(c *fiber.Ctx) error {
			context.GetHttpSession(c).Destroy()
			logManager().Error("Session destroyed")
			return c.Redirect("//logout:logout@" + context.HttpBindAddr() + context.PathWithPrefix("/"))
		})
	}

	v1.Get("/emails", func(ctx *fiber.Ctx) error {
		username, _ := context.GetHttpAuthenticatedUser(ctx)

		page := 1
		perPage := 50
		from := (page - 1) * perPage
		messages, totalCount, err := context.Storage.List(username, storage.SearchQuery{}, from, perPage)
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		to := from + perPage
		if totalCount < to {
			to = totalCount
		}

		return ctx.Status(200).JSON(fiber.Map{
			"meta": fiber.Map{
				"current_page": page,
				"last_page":    int(math.Ceil(float64(totalCount) / float64(perPage))),
				"from":         from,
				"to":           to,
				"total":        totalCount,
			},
			"data": messages,
		})
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
