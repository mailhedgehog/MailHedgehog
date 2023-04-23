package v1

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mailhedgehog/MailHedgehog/authentication"
	"github.com/mailhedgehog/MailHedgehog/dto/email"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"github.com/mailhedgehog/MailHedgehog/smtpClient"
	"github.com/mailhedgehog/MailHedgehog/storage"
	"io"
	"math"
	"net/http"
	"strconv"
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
		switch context.Config.Authentication.Type {
		case "internal":
			v1.Post("/login", apiV1.postInternalLogin)
			v1.Post("/logout", apiV1.postInternalLogout)
		case "basic":
			v1.Post("/logout", func(c *fiber.Ctx) error {
				context.GetHttpSession(c).Destroy()
				logManager().Error("Session destroyed")
				return c.Redirect("//logout:logout@" + context.HttpBindAddr() + context.PathWithPrefix("/"))
			})
		}
	}

	v1.Get("/emails", apiV1.getEmails)
	v1.Delete("/emails", apiV1.deleteEmails)

	v1.Get("/emails/:id", apiV1.showEmail)
	v1.Delete("/emails/:id", apiV1.deleteEmail)

	v1.Post("/emails/:id/release", apiV1.releaseEmail)

	v1.Get("/emails/:id/attachment/:fileIndex", apiV1.downloadAttachment)
}

func (apiV1 *ApiV1) showUser(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)

	if len(username) <= 0 {
		return ctx.Status(http.StatusResetContent).JSON(fiber.Map{})
	}

	return (&Response{Data: fiber.Map{
		"username": username,
	}}).Send(ctx)
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

	errs := ValidateStruct(*listQuery)
	if errs != nil {
		return UnprocessableEntityResponse(ctx, errs)
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
		if len(message.Email.From) > 0 {
			from = fiber.Map{
				"name":  message.Email.From[0].Name,
				"email": message.Email.From[0].Address,
			}
		}
		to := []fiber.Map{}
		for _, toAddress := range message.Email.To {
			to = append(to, fiber.Map{
				"name":  toAddress.Name,
				"email": toAddress.Address,
			})
		}

		messagesResponse = append(messagesResponse, fiber.Map{
			"id":          message.ID,
			"from":        from,
			"to":          to,
			"subject":     message.Email.Subject,
			"received_at": message.Email.Date.Format("2006-01-02 15:04:05"),
			"size":        len(message.Origin.Data),
		})
	}

	return (&Response{
		Data: messagesResponse,
		Meta: fiber.Map{
			"pagination": fiber.Map{
				"current_page": listQuery.Page,
				"per_page":     listQuery.PerPage,
				"last_page":    lastPage,
				"from":         from,
				"to":           to,
				"total":        totalCount,
			},
		},
	}).Send(ctx)
}

func (apiV1 *ApiV1) deleteEmails(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	err := apiV1.context.Storage.DeleteRoom(username)
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	return (&Response{Message: "Emails cleared"}).Send(ctx)
}

func (apiV1 *ApiV1) showEmail(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	smtpEmail, err := apiV1.context.Storage.Load(username, smtpMessage.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	parsedEmail, err := email.Parse(strings.NewReader(smtpEmail.Origin.Data)) // returns Email struct and error
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("parsedEmail", errors.New("email can't be parsed")),
		})
	}

	attachments := []fiber.Map{}
	for index, a := range parsedEmail.Attachments {
		attachments = append(attachments, fiber.Map{
			"filename":    a.Filename,
			"contentType": a.ContentType,
			"data":        "",
			"index":       index,
		})
	}

	logManager().Debug(fmt.Sprintf("%v", parsedEmail))

	return (&Response{Data: fiber.Map{
		"id":          smtpEmail.ID,
		"headers":     smtpEmail.Email.Headers,
		"html":        parsedEmail.HTMLBody,
		"plain":       parsedEmail.TextBody,
		"source":      smtpEmail.Origin.Data,
		"attachments": attachments,
	}}).Send(ctx)
}

func (apiV1 *ApiV1) downloadAttachment(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	smtpEmail, err := apiV1.context.Storage.Load(username, smtpMessage.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}
	fileIndex, err := strconv.Atoi(ctx.Params("fileIndex"))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("fileIndex", err),
		})
	}

	parsedEmail, err := email.Parse(strings.NewReader(smtpEmail.Origin.Data)) // returns Email struct and error
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("parsedEmail", errors.New("email can't be parsed")),
		})
	}

	if fileIndex >= len(parsedEmail.Attachments) || fileIndex < 0 {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("fileIndex", errors.New("incorrect attachment number")),
		})
	}
	attachment := parsedEmail.Attachments[fileIndex]

	bytes, err := io.ReadAll(attachment.Data)
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("fileIndex", errors.New("we can't get content of attachment")),
		})
	}
	ctx.Attachment(attachment.Filename)
	ctx.Type(attachment.ContentType)

	return ctx.Send(bytes)
}

func (apiV1 *ApiV1) deleteEmail(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	err := apiV1.context.Storage.Delete(username, smtpMessage.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	return (&Response{Message: "Email deleted"}).Send(ctx)
}

func (apiV1 *ApiV1) releaseEmail(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	smtpEmail, err := apiV1.context.Storage.Load(username, smtpMessage.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	type ReleaseQuery struct {
		Host     string `query:"host" validate:""`
		Port     int    `query:"port" validate:"min=1,max=9999"`
		Username string `query:"username" validate:"omitempty"`
		Password string `query:"password" validate:"omitempty"`
	}

	releaseQuery := new(ReleaseQuery)

	if err := ctx.BodyParser(releaseQuery); err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	smtpHost := fmt.Sprintf("%s:%d", releaseQuery.Host, releaseQuery.Port)
	var client *smtpClient.SmtpClient
	if len(releaseQuery.Username) > 0 {
		client = smtpClient.NewClient(smtpHost, "plain", []string{
			"",
			releaseQuery.Username,
			releaseQuery.Password,
			releaseQuery.Host,
		})
	} else {
		client = smtpClient.NewClient(smtpHost, "", []string{})
	}

	err = client.SendMail(smtpEmail)
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	return (&Response{Message: "Email released"}).Send(ctx)
}

func (apiV1 *ApiV1) postInternalLogin(ctx *fiber.Ctx) error {
	type LoginBody struct {
		Username string `query:"username" validate:"min=1,max=99999"`
		Password string `query:"password" validate:"min=0,max=99999"`
	}

	loginBody := new(LoginBody)

	if err := ctx.BodyParser(loginBody); err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	if !apiV1.context.Authentication.Authenticate(authentication.HTTP, loginBody.Username, loginBody.Password) {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			{FailedField: "username", Value: "Incorrect credentials"},
		})
	}

	token, err := apiV1.context.SetHttpAuthenticatedUser(ctx, loginBody.Username)
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	fmt.Println(token)

	return (&Response{
		Message: "You are logged",
		Data: fiber.Map{
			"token": token,
		},
	}).Send(ctx)
}

func (apiV1 *ApiV1) postInternalLogout(ctx *fiber.Ctx) error {
	apiV1.context.GetHttpSession(ctx).Destroy()

	// server has not blocklist or any other mechanism to invalidate token,
	// so just remove token from client and keep expiration time shortly as possible

	return (&Response{Message: "Session destroyed"}).Send(ctx)
}
