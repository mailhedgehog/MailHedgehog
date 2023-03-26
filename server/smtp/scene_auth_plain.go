// Authentication PLAIN flow scene
// Auth using encoded in base64 username and login

package smtp

import (
	"encoding/base64"
	"errors"
	"github.com/mailhedgehog/MailHedgehog/authentication"
	"github.com/mailhedgehog/MailHedgehog/smtpServerProtocol"
	"strings"
)

type AuthPlainScene struct {
	authentication authentication.Authentication
	authenticated  func(username string)
	protocol       *smtpServerProtocol.Protocol
}

func (scene *AuthPlainScene) Start(receivedLine string, protocol *smtpServerProtocol.Protocol) *smtpServerProtocol.Reply {
	scene.protocol = protocol
	receivedLine = strings.TrimSpace(receivedLine)

	// RFC provide 2 types of login:
	// Send credentials in same line "AUTH PLAIN vHRjyADROPsdSDIROu="
	// Or send credentials in second message after reply.
	// So to understand if credentials in same line we check it by prefix and decode immediately
	prefix := string(smtpServerProtocol.CommandAuth) + " " + AuthMechanismPlain + " "
	if strings.HasPrefix(receivedLine, prefix) {
		return scene.checkCredentials(strings.TrimPrefix(receivedLine, prefix))
	}

	return smtpServerProtocol.ReplyAuthCredentials("")
}

func (scene *AuthPlainScene) ReadAndWriteReply(receivedLine string) *smtpServerProtocol.Reply {
	// Auth PLAIN has only one possible message with credentials, so going directly to it
	return scene.checkCredentials(receivedLine)
}

func (scene *AuthPlainScene) Finish() {
	scene.protocol.SetStateCommandsExchange()
}

// checkCredentials by decoding and validate and returns appropriate reply
func (scene *AuthPlainScene) checkCredentials(encodedCredentials string) *smtpServerProtocol.Reply {
	scene.Finish()

	username, password, err := scene.decodeCredentials(encodedCredentials)
	if err != nil || !scene.authentication.Authenticate(authentication.SMTP, username, password) {
		return smtpServerProtocol.ReplyAuthFailed()
	}

	if scene.authenticated != nil {
		scene.authenticated(username)
	}
	return smtpServerProtocol.ReplyAuthOk()
}

// decodeCredentials from base64 encoded string passed by "client".
// Returns decoded username and password on successful decode.
func (scene *AuthPlainScene) decodeCredentials(encodedCredentials string) (string, string, error) {
	encodedCredentials = strings.TrimSpace(encodedCredentials)

	decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return "", "", err
	}

	parts := strings.Split(string(decodedBytes), string(rune(0)))

	if len(parts) < 3 {
		return "", "", errors.New("invalid decoded value")
	}

	return parts[1], parts[2], nil
}
