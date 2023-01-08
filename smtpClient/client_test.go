package smtpClient

import (
	"github.com/mailhedgehog/MailHedgehog/dto"
	"github.com/mailhedgehog/MailHedgehog/gounit"
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

	messageId := dto.MessageID("example-email")
	b, err := os.ReadFile(filepath.Join("./", string(messageId)))
	(*gounit.T)(t).AssertNotError(err)

	email := dto.FromBytes(b).Parse()
	email.ID = messageId

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
