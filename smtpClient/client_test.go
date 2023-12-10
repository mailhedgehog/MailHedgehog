package smtpClient

import (
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/gounit"
	"os"
	"path/filepath"
	"testing"
)

func TestSendMail(t *testing.T) {
	login := ""
	pass := ""
	if len(login) <= 0 || len(pass) <= 0 {
		t.Skip("skipping testing in short mode")
	}

	messageId := smtpMessage.MessageID("example-email")
	b, err := os.ReadFile(filepath.Join("./", string(messageId)))
	(*gounit.T)(t).AssertNotError(err)

	email, err := smtpMessage.FromString(string(b)).ToSMTPMail(messageId)
	(*gounit.T)(t).AssertNotError(err)

	client := NewClient(
		"smtp.mailtrap.io:2525",
		"plain",
		[]string{
			"",
			login,
			pass,
			"smtp.mailtrap.io",
		},
	)

	err = client.SendMail(email)
	(*gounit.T)(t).AssertNotError(err)
}
