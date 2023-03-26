package smtpMessage

import (
	"github.com/mailhedgehog/MailHedgehog/gounit"
	"testing"
)

func TestToSMTPMail(t *testing.T) {
	smtpMessage := &SMTPMessage{
		Helo: "test-X510",
		From: "from@example.com",
		To: []string{
			"joe@example.net",
			"cc@example.com",
		},
		Data: "",
	}

	smtpMail, err := smtpMessage.ToSMTPMail(MessageID("foo-bar"))
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsString("foo-bar", string(smtpMail.ID))
	(*gounit.T)(t).AssertEqualsString(smtpMessage.From, smtpMail.From.Address())
	for i, to := range smtpMessage.To {
		(*gounit.T)(t).AssertEqualsString(to, smtpMail.To[i].Address())
	}
	// nil because content empty
	(*gounit.T)(t).AssertNil(smtpMail.Email)

}

func TestNewMessageIDIsUuid(t *testing.T) {
	id := NewMessageID()
	(*gounit.T)(t).AssertLengthString(36, string(id))
}
