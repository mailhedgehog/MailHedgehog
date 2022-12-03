package v1

import (
	"github.com/go-playground/validator/v10"
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
	context *serverContext.Context
}

func CreateAPIV1Routes(context *serverContext.Context, api fiber.Router) {

	apiV1 := &ApiV1{context: context}

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

	v1.Get("/emails", apiV1.getEmails)

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

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(structToValidate interface{}) []*ValidationError {
	var errors []*ValidationError
	err := validate.Struct(structToValidate)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (apiV1 *ApiV1) getEmails(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)

	type ListQuery struct {
		Page       int    `query:"page" validate:"omitempty,min=1,max=99999"`
		PerPage    int    `query:"per_page" validate:"omitempty"`
		SearchText string `query:"search" validate:"omitempty,max=255"`
	}

	listQuery := new(ListQuery)

	if err := ctx.QueryParser(listQuery); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"messages": err.Error(),
		})

	}

	if listQuery.Page == 0 {
		listQuery.Page = 1
	}
	if listQuery.PerPage == 0 {
		listQuery.PerPage = 25
	}

	errors := ValidateStruct(*listQuery)
	if errors != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"messages": errors,
		})

	}

	query := storage.SearchQuery{}

	if len(listQuery.SearchText) > 0 {
		query["from"] = listQuery.SearchText
		query["to"] = listQuery.SearchText
		query["content"] = listQuery.SearchText
	}

	from := (listQuery.Page - 1) * listQuery.PerPage
	messages, totalCount, err := apiV1.context.Storage.List(username, query, from, listQuery.PerPage)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if totalCount > 0 {
		from += 1
	}
	to := from + listQuery.PerPage
	if totalCount < to {
		to = totalCount
	}
	lastPage := int(math.Ceil(float64(totalCount) / float64(listQuery.PerPage)))
	if lastPage < listQuery.Page {
		lastPage = listQuery.Page
	}

	messagesResponse := []fiber.Map{}
	for _, message := range messages {
		var from fiber.Map
		fromInfo, err := message.Content.Headers.From()
		if err == nil {
			from = fiber.Map{
				"name":  fromInfo.Name,
				"email": fromInfo.Email,
			}
		}
		to := []fiber.Map{}
		for _, info := range message.Content.Headers.To() {
			to = append(to, fiber.Map{
				"name":  info.Name,
				"email": info.Email,
			})
		}

		var receivedAt string
		receivedAtInfo, err := message.Content.Headers.DateUTC()
		if err == nil {
			receivedAt = receivedAtInfo.Format("2006-01-02 15:04:05")
		}
		messagesResponse = append(messagesResponse, fiber.Map{
			"id":          message.ID,
			"from":        from,
			"to":          to,
			"subject":     "",
			"received_at": receivedAt,
			"size":        len(message.Raw.Data),
			"yyy":         message.Content.Headers.All(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": messagesResponse,
		"meta": fiber.Map{
			"pagination": fiber.Map{
				"current_page": listQuery.Page,
				"per_page":     listQuery.PerPage,
				"last_page":    lastPage,
				"from":         from,
				"to":           to,
				"total":        totalCount,
			},
		},
	})
}
