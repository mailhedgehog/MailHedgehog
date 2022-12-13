package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mailhedgehog/MailHedgehog/dto"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"github.com/mailhedgehog/MailHedgehog/parseMail"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"github.com/mailhedgehog/MailHedgehog/storage"
	"io"
	"math"
	"net/http"
	"strings"
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

	v1.Get("/user", apiV1.showUser)

	if context.Authentication.RequiresAuthentication() {
		v1.Post("/logout", func(c *fiber.Ctx) error {
			context.GetHttpSession(c).Destroy()
			logManager().Error("Session destroyed")
			return c.Redirect("//logout:logout@" + context.HttpBindAddr() + context.PathWithPrefix("/"))
		})
	}

	v1.Get("/emails", apiV1.getEmails)
	v1.Delete("/emails", apiV1.deleteEmails)

	v1.Get("/emails/:id", apiV1.showEmail)
	v1.Delete("/emails/:id", apiV1.deleteEmail)
}

func (apiV1 *ApiV1) showUser(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)

	if len(username) <= 0 {
		return ctx.Status(http.StatusResetContent).JSON(fiber.Map{})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"username": username,
		},
	})
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
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
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
		return UnprocessableEntityResponse(ctx, errors)

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
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	if totalCount > 0 {
		from += 1
	}
	to := from + listQuery.PerPage - 1
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

		subject, _ := message.Content.Headers.GetOne("Subject")

		messagesResponse = append(messagesResponse, fiber.Map{
			"id":          message.ID,
			"from":        from,
			"to":          to,
			"subject":     subject,
			"received_at": receivedAt,
			"size":        len(message.Raw.Data),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
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

func (apiV1 *ApiV1) deleteEmails(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	err := apiV1.context.Storage.DeleteRoom(username)
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Emails cleared",
	})
}

func (apiV1 *ApiV1) showEmail(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	email, err := apiV1.context.Storage.Load(username, dto.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	parsedEmail, err := parseMail.Parse(strings.NewReader(email.Raw.Data)) // returns Email struct and error
	if err != nil {
		// handle error
	}

	attachments := []fiber.Map{}
	for _, a := range parsedEmail.Attachments {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, a.Data)
		if err != nil {
			// handle error
		}
		fmt.Println()
		attachments = append(attachments, fiber.Map{
			"filename":    a.Filename,
			"contentType": a.ContentType,
			"data":        buf.String(),
		})
	}

	logManager().Debug(fmt.Sprintf("%v", parsedEmail))

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"id":          email.ID,
			"headers":     email.Content.Headers.All(),
			"html":        parsedEmail.HTMLBody,
			"plain":       parsedEmail.TextBody,
			"source":      email.Raw.Data,
			"attachments": attachments,
		},
	})
}

func (apiV1 *ApiV1) deleteEmail(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	err := apiV1.context.Storage.Delete(username, dto.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Email deleted",
	})
}
