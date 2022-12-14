package smtp

import (
	"encoding/base64"
	"errors"
	"github.com/mailhedgehog/MailHedgehog/authentication"
	"github.com/mailhedgehog/MailHedgehog/smtpServer"
	"strings"
)

type AuthPlainScene struct {
	authentication authentication.Authentication
	authenticated  func(username string)
	protocol       *smtpServer.Protocol
}

func (scene *AuthPlainScene) Init(receivedLine string, protocol *smtpServer.Protocol) *smtpServer.Reply {
	scene.protocol = protocol
	prefix := string(smtpServer.COMMAND_AUTH) + " " + AUTH_MECHANISM_PLAIN + " "
	if strings.HasPrefix(receivedLine, prefix) {
		return scene.replyAfterCheckCredentials(strings.TrimPrefix(receivedLine, prefix))
	}
	return smtpServer.ReplyAuthCredentials("")
}

func (scene *AuthPlainScene) HandleLine(receivedLine string) *smtpServer.Reply {
	return scene.replyAfterCheckCredentials(receivedLine)
}

func (scene *AuthPlainScene) replyAfterCheckCredentials(encodedCredentials string) *smtpServer.Reply {
	scene.protocol.State = smtpServer.STATE_CONVERSATION

	username, password, err := scene.decodeCredentials(encodedCredentials)
	if err != nil || !scene.authentication.Authenticate(authentication.SMTP, username, password) {
		return smtpServer.ReplyAuthFailed()
	}

	if scene.authenticated != nil {
		scene.authenticated(username)
	}
	return smtpServer.ReplyAuthOk()
}

func (scene *AuthPlainScene) decodeCredentials(encodedCredentials string) (string, string, error) {
	encodedCredentials = strings.TrimSpace(encodedCredentials)
	val, _ := base64.StdEncoding.DecodeString(encodedCredentials)
	parts := strings.Split(string(val), string(rune(0)))

	if len(parts) < 3 {
		return "", "", errors.New("invalid decoded value")
	}

	return parts[1], parts[2], nil
}
