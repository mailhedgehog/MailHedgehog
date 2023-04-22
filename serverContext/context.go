package serverContext

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mailhedgehog/MailHedgehog/authentication"
	"github.com/mailhedgehog/MailHedgehog/config"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"github.com/mailhedgehog/MailHedgehog/storage"
	"strings"
	"time"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("server.context")
	}
	return configuredLogger
}

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
	username := ""

	switch context.Config.Authentication.Type {
	case "basic":
		httpSession := context.GetHttpSession(ctx)
		usernameValue := httpSession.Get("CurrentUser")
		if usernameValue != nil {
			username = fmt.Sprintf("%v", usernameValue)
		}
	case "internal":
		tokenString := ctx.Get(fiber.HeaderAuthorization)
		if len(tokenString) > 0 {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			if len(tokenString) > 0 {
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return context.Config.Authentication.HmacSecret, nil
				})
				if err == nil {
					if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
						if time.Now().Unix() < int64(claims["exp"].(float64)) {
							username = claims["user"].(string)
						}
					}
				}
			}
		}
	}

	logManager().Debug(fmt.Sprintf("GetHttpAuthenticatedUser: %s", username))

	if len(username) > 0 {
		return username, nil
	}
	return "", errors.New("user not found")
}

func (context *Context) SetHttpAuthenticatedUser(ctx *fiber.Ctx, username string) (string, error) {
	switch context.Config.Authentication.Type {
	case "basic":
		httpSession := context.GetHttpSession(ctx)
		httpSession.Set("CurrentUser", username)

		logManager().Debug(fmt.Sprintf("SetHttpAuthenticatedUser: %s", username))

		return "", httpSession.Save()
	case "internal":
		t := jwt.NewWithClaims(jwt.SigningMethodHS512,
			jwt.MapClaims{
				"iss":  "mailhedgehog",
				"user": username,
				"exp":  time.Now().Add(48 * time.Hour).Unix(),
			})
		tokenString, err := t.SignedString(context.Config.Authentication.HmacSecret)
		logger.PanicIfError(err)
		return tokenString, nil
	}

	return "", errors.New("incorrect auth type")
}
