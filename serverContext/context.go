package serverContext

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/authentication"
	"github.com/mailpiggy/MailPiggy/config"
	"github.com/mailpiggy/MailPiggy/storage"
)

type Context struct {
	Authentication authentication.Authentication
	Storage        storage.Storage
	Config         config.AppConfig
}

func (context *Context) SmtpBindAddr() string {
	return context.Config.Smtp.Host + ":" + fmt.Sprint(context.Config.Smtp.Port)
}

func (context *Context) HttpBindAddr() string {
	return context.Config.Http.Host + ":" + fmt.Sprint(context.Config.Http.Port)
}
