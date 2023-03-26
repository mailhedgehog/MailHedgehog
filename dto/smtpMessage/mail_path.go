package smtpMessage

import "strings"

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

// MailPathFromString parses a forward-path or reverse-path into its parts
func MailPathFromString(path string) *MailPath {
	var relays []string
	userEmail := path
	if strings.Contains(path, ":") {
		x := strings.SplitN(path, ":", 2)
		r, e := x[0], x[1]
		userEmail = e
		relays = strings.Split(r, ",")
	}
	mailbox, domain := "", ""
	if strings.Contains(userEmail, "@") {
		x := strings.SplitN(userEmail, "@", 2)
		mailbox, domain = x[0], x[1]
	} else {
		mailbox = userEmail
	}

	return &MailPath{
		Relays:  relays,
		Mailbox: mailbox,
		Domain:  domain,
		Params:  "", // TODO: add params config
	}
}
