package smtpClient

import (
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"net/smtp"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("smtpClient")
	}
	return configuredLogger
}

type SmtpClient struct {
	smtpAddr   string
	authName   string
	authParams []string
}

func NewClient(smtpAddr string, authName string, authParams []string) *SmtpClient {
	return &SmtpClient{smtpAddr, authName, authParams}
}

func (client *SmtpClient) SendMail(message *smtpMessage.SMTPMail) error {

	fromAddr := message.From.Address()
	to := []string{}

	for _, path := range message.To {
		to = append(to, path.Address())
	}
	msg := []byte(message.Origin.Data)

	var auth smtp.Auth = nil

	switch client.authName {
	case "plain":
		auth = smtp.PlainAuth(client.authParams[0], client.authParams[1], client.authParams[2], client.authParams[3])
	case "linux":
		auth = smtp.CRAMMD5Auth(client.authParams[0], client.authParams[1])
	}

	return smtp.SendMail(client.smtpAddr, auth, fromAddr, to, msg)
}
