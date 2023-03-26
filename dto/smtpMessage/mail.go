package smtpMessage

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mailhedgehog/MailHedgehog/dto/email"
	"strings"
)

// SMTPMail represents parsed SMTPMessage what allow
// easily get and manipulate data
type SMTPMail struct {
	ID     MessageID
	From   *MailPath
	To     []*MailPath
	Email  *email.Email
	Origin *SMTPMessage
}

// MessageID represents the ID of an SMTP message
type MessageID string

// NewMessageID generates a new mail identificatior
func NewMessageID() MessageID {
	return MessageID(uuid.New().String())
}

// MailPath represents an SMTP forward-path or return-path
type MailPath struct {
	Relays  []string
	Mailbox string
	Domain  string
	Params  string
}

func (path *MailPath) Address() string {
	return path.Mailbox + "@" + path.Domain
}

// ToSMTPMail converts SMTPMassage to SMTPMail structure
// In case if `id` is empty will be automatically generated new identificator
func (message *SMTPMessage) ToSMTPMail(id MessageID) (*SMTPMail, error) {
	if len(id) <= 0 {
		id = NewMessageID()
	}

	parsedEmail, err := email.Parse(strings.NewReader(message.Data))
	if err != nil && !errors.Is(err, email.ErrEmptyString) {
		return nil, err
	}

	var receiversList []*MailPath
	for _, path := range message.To {
		receiversList = append(receiversList, MailPathFromString(path))
	}

	smtpMail := &SMTPMail{
		ID:     id,
		From:   MailPathFromString(message.From),
		To:     receiversList,
		Email:  parsedEmail,
		Origin: message,
	}

	return smtpMail, nil
}
