package v1

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mailhedgehog/MailHedgehog/serverContext"
	"github.com/mailhedgehog/MailHedgehog/smtpClient"
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/email"
	"github.com/mailhedgehog/logger"
	"github.com/mailhedgehog/smtpMessage"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
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

	if context.Authentication.Dashboard().RequiresAuthentication() {
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

	if apiV1.context.Sharing != nil {
		v1.Post("/emails/:id/share", apiV1.shareEmail)
		v1.Get("/shared-email/:id", apiV1.showSharedEmail)
		v1.Get("/shared-email/:id/attachment/:fileIndex", apiV1.downloadSharedEmailAttachment)
	}

	users := v1.Group("/users", func(c *fiber.Ctx) error {
		return c.Next()
	})
	users.Use(func(ctx *fiber.Ctx) error {
		username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)

		if username == "" || !apiV1.context.Config.IsAdmin(username) {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		// Go to next middleware:
		return ctx.Next()
	})

	/* Users management */
	users.Get("", apiV1.getUsers)
	users.Post("", apiV1.createUser)
	users.Put("/:username", apiV1.updateUser)
	users.Delete("/:username", apiV1.deleteUser)
}

func (apiV1 *ApiV1) showUser(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)

	if len(username) <= 0 {
		return ctx.Status(http.StatusResetContent).JSON(fiber.Map{})
	}

	var permissions []string
	if apiV1.context.Config.IsAdmin(username) {
		permissions = append(permissions, "manage_users")
	}

	return (&Response{Data: fiber.Map{
		"username":    username,
		"permissions": permissions,
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

	query := contracts.SearchQuery{}

	if len(listQuery.SearchText) > 0 {
		query[contracts.SearchParamFrom] = listQuery.SearchText
		query[contracts.SearchParamTo] = listQuery.SearchText
		query[contracts.SearchParamContent] = listQuery.SearchText
	}

	from := (listQuery.Page - 1) * listQuery.PerPage
	messages, totalCount, err := apiV1.context.Storage.MessagesRepo(contracts.Room(username)).List(query, from, listQuery.PerPage)
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
		if len(message.GetEmail().From) > 0 {
			from = fiber.Map{
				"name":  message.GetEmail().From[0].Name,
				"email": message.GetEmail().From[0].Address,
			}
		}
		to := []fiber.Map{}
		for _, toAddress := range message.GetEmail().To {
			to = append(to, fiber.Map{
				"name":  toAddress.Name,
				"email": toAddress.Address,
			})
		}

		messagesResponse = append(messagesResponse, fiber.Map{
			"id":          message.ID,
			"from":        from,
			"to":          to,
			"subject":     message.GetEmail().Subject,
			"received_at": message.GetEmail().Date.Format("2006-01-02 15:04:05"),
			"size":        len(message.GetOrigin()),
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
	err := apiV1.context.Storage.RoomsRepo().Delete(contracts.Room(username))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	return (&Response{Message: "DashboardAuthEmails cleared"}).Send(ctx)
}

func (apiV1 *ApiV1) showEmail(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	smtpEmail, err := apiV1.context.Storage.MessagesRepo(contracts.Room(username)).Load(smtpMessage.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	attachments := []fiber.Map{}
	for index, a := range smtpEmail.GetEmail().Attachments {
		attachments = append(attachments, fiber.Map{
			"filename":    a.Filename,
			"contentType": a.ContentType,
			"data":        "",
			"index":       index,
		})
	}

	logManager().Debug(fmt.Sprintf("%v", smtpEmail.GetEmail()))

	return (&Response{Data: fiber.Map{
		"id":          smtpEmail.ID,
		"headers":     smtpEmail.GetEmail().Headers,
		"html":        smtpEmail.GetEmail().HTMLBody,
		"plain":       smtpEmail.GetEmail().TextBody,
		"source":      smtpEmail.GetOrigin(),
		"attachments": attachments,
	}}).Send(ctx)
}

func (apiV1 *ApiV1) downloadAttachment(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	smtpEmail, err := apiV1.context.Storage.MessagesRepo(contracts.Room(username)).Load(smtpMessage.MessageID(ctx.Params("id")))
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

	parsedEmail, err := email.Parse(strings.NewReader(smtpEmail.GetOrigin())) // returns email struct and error
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
	ctx.Set("Access-Control-Expose-Headers", "Content-Disposition")

	return ctx.Send(bytes)
}

func (apiV1 *ApiV1) deleteEmail(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	err := apiV1.context.Storage.MessagesRepo(contracts.Room(username)).Delete(smtpMessage.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	return (&Response{Message: "email deleted"}).Send(ctx)
}

func (apiV1 *ApiV1) releaseEmail(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	smtpEmail, err := apiV1.context.Storage.MessagesRepo(contracts.Room(username)).Load(smtpMessage.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	type ReleaseQuery struct {
		Host     string `json:"host" xml:"host" form:"host" validate:""`
		Port     int    `json:"port" xml:"port" form:"port" validate:"min=1,max=9999"`
		Username string `json:"username" xml:"username" form:"username" validate:"omitempty"`
		Password string `json:"password" xml:"password" form:"password" validate:"omitempty"`
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

	return (&Response{Message: "email released"}).Send(ctx)
}

func (apiV1 *ApiV1) shareEmail(ctx *fiber.Ctx) error {
	username, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	smtpEmail, err := apiV1.context.Storage.MessagesRepo(contracts.Room(username)).Load(smtpMessage.MessageID(ctx.Params("id")))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	type ShareQuery struct {
		Hours int `json:"expiration_in" xml:"expiration_in" form:"expiration_in" validate:"min=1,max=800"`
	}

	shareQuery := new(ShareQuery)

	if err := ctx.BodyParser(shareQuery); err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	emailSharingRecord := contracts.NewSharedMessageRecord(contracts.Room(username), smtpEmail.ID)
	emailSharingRecord.SetExpirationInHours(shareQuery.Hours)

	emailSharingRecord, err = apiV1.context.Sharing.Add(emailSharingRecord)

	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	return (&Response{
		Message: "Shared link created.",
		Data: fiber.Map{
			"id": emailSharingRecord.Id,
		},
	}).Send(ctx)
}

func (apiV1 *ApiV1) postInternalLogin(ctx *fiber.Ctx) error {
	type LoginBody struct {
		Username string `json:"username" xml:"username" form:"username" validate:"min=1,max=99999"`
		Password string `json:"password" xml:"password" form:"password" validate:"min=0,max=99999"`
	}

	loginBody := new(LoginBody)

	if err := ctx.BodyParser(loginBody); err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	if !apiV1.context.Authentication.Dashboard().ViaPasswordAuthentication().Authenticate(loginBody.Username, loginBody.Password) {
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

func (apiV1 *ApiV1) getUsers(ctx *fiber.Ctx) error {
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

	from := (listQuery.Page - 1) * listQuery.PerPage
	users, totalCount, err := apiV1.context.Authentication.UsersStorage().List(listQuery.SearchText, from, listQuery.PerPage)
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

	sort.Slice(users, func(i, j int) bool {
		return users[i].Username < users[j].Username
	})

	usersResponse := []fiber.Map{}
	for _, user := range users {
		usersResponse = append(usersResponse, fiber.Map{
			"username":              user.Username,
			"smtp_auth_ips":         user.SmtpAuthIPs,
			"smtp_allow_listed_ips": user.SmtpAllowListedIPs,
			"dashboard_auth_emails": user.DashboardAuthEmails,
		})
	}

	return (&Response{
		Data: usersResponse,
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

func (apiV1 *ApiV1) createUser(ctx *fiber.Ctx) error {
	type CreateUserBody struct {
		Username            string   `json:"username" xml:"username" form:"username" validate:"min=1,max=255,alphanum"`
		DashboardPassword   string   `json:"dashboard_password" xml:"dashboard_password" form:"hub_password" validate:"omitempty,min=1,max=255"`
		SmtpPassword        string   `json:"smtp_password" xml:"smtp_password" form:"smtp_password" validate:"omitempty,min=1,max=255"`
		SmtpAuthIPs         []string `json:"smtp_auth_ips" xml:"smtp_auth_ips" form:"smtp_auth_ips" validate:"omitempty,dive"`
		SmtpAllowListedIPs  []string `json:"smtp_allow_listed_ips" xml:"smtp_allow_listed_ips" form:"smtp_allow_listed_ips" validate:"omitempty,dive"`
		DashboardAuthEmails []string `json:"dashboard_auth_emails" xml:"dashboard_auth_emails" form:"dashboard_auth_emails" validate:"omitempty,dive"`
	}

	createUserBody := new(CreateUserBody)

	if err := ctx.BodyParser(createUserBody); err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	errs := ValidateStruct(createUserBody)
	if errs != nil {
		return UnprocessableEntityResponse(ctx, errs)
	}

	if apiV1.context.Authentication.UsersStorage().Exists(createUserBody.Username) {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", errors.New("user with same username already present")),
		})
	}

	err := apiV1.context.Authentication.UsersStorage().Add(createUserBody.Username)
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	if len(createUserBody.DashboardPassword) > 0 {
		err = apiV1.context.Authentication.Dashboard().ViaPasswordAuthentication().SetPassword(createUserBody.Username, createUserBody.DashboardPassword)
	}
	if len(createUserBody.SmtpPassword) > 0 {
		err = apiV1.context.Authentication.SMTP().ViaPasswordAuthentication().SetPassword(createUserBody.Username, createUserBody.SmtpPassword)
	}

	err = apiV1.context.Authentication.SMTP().ViaIpAuthentication().ClearAllIps(createUserBody.Username)
	if err == nil {
		for i := range createUserBody.SmtpAuthIPs {
			err = apiV1.context.Authentication.SMTP().ViaIpAuthentication().AddIp(createUserBody.Username, strings.TrimSpace(createUserBody.SmtpAuthIPs[i]))
			if err != nil {
				logManager().Warning(err.Error())
			}
		}
	}

	err = apiV1.context.Authentication.SMTP().IpsAllowList().ClearAllIps(createUserBody.Username)
	if err == nil {
		for i := range createUserBody.SmtpAllowListedIPs {
			err = apiV1.context.Authentication.SMTP().IpsAllowList().AddIp(createUserBody.Username, strings.TrimSpace(createUserBody.SmtpAllowListedIPs[i]))
			if err != nil {
				logManager().Warning(err.Error())
			}
		}
	}

	err = apiV1.context.Authentication.Dashboard().ViaEmailAuthentication().ClearAllEmails(createUserBody.Username)
	if err == nil {
		for i := range createUserBody.DashboardAuthEmails {
			err = apiV1.context.Authentication.Dashboard().ViaEmailAuthentication().AddEmail(createUserBody.Username, strings.TrimSpace(createUserBody.DashboardAuthEmails[i]))
			if err != nil {
				logManager().Warning(err.Error())
			}
		}
	}

	return (&Response{Message: "Create User"}).Send(ctx)
}

func (apiV1 *ApiV1) updateUser(ctx *fiber.Ctx) error {
	var err error

	type UpdateUserBody struct {
		DashboardPassword   string   `json:"dashboard_password" xml:"dashboard_password" form:"hub_password" validate:"omitempty,min=1,max=255"`
		SmtpPassword        string   `json:"smtp_password" xml:"smtp_password" form:"smtp_password" validate:"omitempty,min=1,max=255"`
		SmtpAuthIPs         []string `json:"smtp_auth_ips" xml:"smtp_auth_ips" form:"smtp_auth_ips" validate:"omitempty,dive"`
		SmtpAllowListedIPs  []string `json:"smtp_allow_listed_ips" xml:"smtp_allow_listed_ips" form:"smtp_allow_listed_ips" validate:"omitempty,dive"`
		DashboardAuthEmails []string `json:"dashboard_auth_emails" xml:"dashboard_auth_emails" form:"dashboard_auth_emails" validate:"omitempty,dive"`
	}

	updateUserBody := new(UpdateUserBody)

	if err := ctx.BodyParser(updateUserBody); err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	errs := ValidateStruct(updateUserBody)
	if errs != nil {
		return UnprocessableEntityResponse(ctx, errs)
	}

	username := ctx.Params("username")

	if len(updateUserBody.DashboardPassword) > 0 {
		err = apiV1.context.Authentication.Dashboard().ViaPasswordAuthentication().SetPassword(username, updateUserBody.DashboardPassword)
		if err != nil {
			return UnprocessableEntityResponse(ctx, []*ValidationError{
				ValidationErrorFromError("query", err),
			})
		}
	}

	if len(updateUserBody.SmtpPassword) > 0 {
		err = apiV1.context.Authentication.SMTP().ViaPasswordAuthentication().SetPassword(username, updateUserBody.SmtpPassword)
		if err != nil {
			return UnprocessableEntityResponse(ctx, []*ValidationError{
				ValidationErrorFromError("query", err),
			})
		}
	}

	err = apiV1.context.Authentication.SMTP().ViaIpAuthentication().ClearAllIps(username)
	if err == nil {
		for i := range updateUserBody.SmtpAuthIPs {
			err = apiV1.context.Authentication.SMTP().ViaIpAuthentication().AddIp(username, strings.TrimSpace(updateUserBody.SmtpAuthIPs[i]))
			if err != nil {
				logManager().Warning(err.Error())
			}
		}
	}

	err = apiV1.context.Authentication.SMTP().IpsAllowList().ClearAllIps(username)
	if err == nil {
		for i := range updateUserBody.SmtpAllowListedIPs {
			err = apiV1.context.Authentication.SMTP().IpsAllowList().AddIp(username, strings.TrimSpace(updateUserBody.SmtpAllowListedIPs[i]))
			if err != nil {
				logManager().Warning(err.Error())
			}
		}
	}

	err = apiV1.context.Authentication.Dashboard().ViaEmailAuthentication().ClearAllEmails(username)
	if err == nil {
		for i := range updateUserBody.DashboardAuthEmails {
			err = apiV1.context.Authentication.Dashboard().ViaEmailAuthentication().AddEmail(username, strings.TrimSpace(updateUserBody.DashboardAuthEmails[i]))
			if err != nil {
				logManager().Warning(err.Error())
			}
		}
	}

	return (&Response{Message: "User Updated"}).Send(ctx)
}

func (apiV1 *ApiV1) deleteUser(ctx *fiber.Ctx) error {
	loggedUsername, _ := apiV1.context.GetHttpAuthenticatedUser(ctx)
	username, err := url.QueryUnescape(ctx.Params("username"))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	if loggedUsername == username {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", errors.New("you can't delete your own account")),
		})
	}

	err = apiV1.context.Storage.RoomsRepo().Delete(contracts.Room(username))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	err = apiV1.context.Authentication.UsersStorage().Delete(username)
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	return (&Response{Message: "User deleted."}).Send(ctx)
}

func (apiV1 *ApiV1) showSharedEmail(ctx *fiber.Ctx) error {
	// Lottery to remove all expired
	if rand.Intn(10) == 0 {
		apiV1.context.Sharing.DeleteExpired()
	}

	emailSharingRecord, err := apiV1.context.Sharing.Find(ctx.Params("id"))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	smtpEmail, err := apiV1.context.Storage.MessagesRepo(emailSharingRecord.Room).Load(emailSharingRecord.MessageId)
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	attachments := []fiber.Map{}
	for index, a := range smtpEmail.GetEmail().Attachments {
		attachments = append(attachments, fiber.Map{
			"filename":    a.Filename,
			"contentType": a.ContentType,
			"data":        "",
			"index":       index,
		})
	}

	logManager().Debug(fmt.Sprintf("%v", smtpEmail.GetEmail()))

	return (&Response{Data: fiber.Map{
		"id":          smtpEmail.ID,
		"headers":     smtpEmail.GetEmail().Headers,
		"html":        smtpEmail.GetEmail().HTMLBody,
		"plain":       smtpEmail.GetEmail().TextBody,
		"source":      smtpEmail.GetOrigin(),
		"attachments": attachments,
	}}).Send(ctx)
}

func (apiV1 *ApiV1) downloadSharedEmailAttachment(ctx *fiber.Ctx) error {
	emailSharingRecord, err := apiV1.context.Sharing.Find(ctx.Params("id"))
	if err != nil {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("query", err),
		})
	}

	smtpEmail, err := apiV1.context.Storage.MessagesRepo(emailSharingRecord.Room).Load(emailSharingRecord.MessageId)
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

	if fileIndex >= len(smtpEmail.GetEmail().Attachments) || fileIndex < 0 {
		return UnprocessableEntityResponse(ctx, []*ValidationError{
			ValidationErrorFromError("fileIndex", errors.New("incorrect attachment number")),
		})
	}
	attachment := smtpEmail.GetEmail().Attachments[fileIndex]

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
