package serverContext

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/mailpiggy/MailPiggy/authentication"
	"github.com/mailpiggy/MailPiggy/config"
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/mailpiggy/MailPiggy/storage"
	"strings"
)

type Context struct {
	Authentication authentication.Authentication
	Storage        storage.Storage
	Config         config.AppConfig
	HttpSession    *session.Store
}

func (context *Context) SmtpBindAddr() string {
	return context.Config.Smtp.Host + ":" + fmt.Sprint(context.Config.Smtp.Port)
}

func (context *Context) HttpBindAddr() string {
	return context.Config.Http.Host + ":" + fmt.Sprint(context.Config.Http.Port)
}

func (context *Context) PathWithPrefix(path string) string {
	path = strings.TrimPrefix(path, "/")
	prefix := strings.Trim(context.Config.Http.Path, "/")
	if len(prefix) > 0 {
		prefix = "/" + prefix
	}

	return prefix + "/" + path
}

func (context *Context) GetHttpSession(c *fiber.Ctx) *session.Session {
	sess, err := context.HttpSession.Get(c)
	logger.PanicIfError(err)

	return sess
}

func (context *Context) GetHttpAuthenticatedUser(ctx *fiber.Ctx) (string, error) {
	httpSession := context.GetHttpSession(ctx)
	username := ""
	usernameValue := httpSession.Get("CurrentUser")
	if usernameValue != nil {
		username = fmt.Sprintf("%v", usernameValue)
	}
	if len(username) > 0 {
		return username, nil
	}
	return "", errors.New("user not found")
}

func (context *Context) SetHttpAuthenticatedUser(ctx *fiber.Ctx, username string) error {
	httpSession := context.GetHttpSession(ctx)
	httpSession.Set("CurrentUser", username)
	return httpSession.Save()
}
