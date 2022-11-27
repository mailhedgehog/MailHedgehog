package smtpServer

import "strconv"

// Reply is a struct representing an SMTP reply (status code + lines)
type Reply struct {
	Status int
	lines  []string
	Done   func()
}

const (
	CODE_SYSTEM_STATUS          = 211
	CODE_HELP_MESSAGE           = 214
	CODE_SERVICE_READY          = 220
	CODE_SERVICE_CLOSING        = 221
	CODE_AUTHENTICATION_SUCCESS = 235
	CODE_ACTION_OK              = 250
	CODE_SYNTAX_ERROR           = 500
)

/*
      251 User not local; will forward to <forward-path>
         (See section 3.4)
      252 Cannot VRFY user, but will accept message and attempt
         delivery
         (See section 3.5.3)

      354 Start mail input; end with <CRLF>.<CRLF>
421 <domain> Service not available, closing transmission channel
         (This may be a reply to any command if the service knows it
         must shut down)
      450 Requested mail action not taken: mailbox unavailable
         (e.g., mailbox busy)
      451 Requested action aborted: local error in processing
      452 Requested action not taken: insufficient system storage
      500 Syntax error, command unrecognized
         (This may include errors such as command line too long)
      501 Syntax error in parameters or arguments
      502 Command not implemented (see section 4.2.4)
      503 Bad sequence of commands
      504 Command parameter not implemented
      550 Requested action not taken: mailbox unavailable
         (e.g., mailbox not found, no access, or command rejected
         for policy reasons)
      551 User not local; please try <forward-path>
         (See section 3.4)
      552 Requested mail action aborted: exceeded storage allocation
      553 Requested action not taken: mailbox name not allowed
         (e.g., mailbox syntax incorrect)
      554 Transaction failed  (Or, in the case of a connection-opening
          response, "No SMTP service here")
*/

// Lines returns the formatted SMTP reply
func (r Reply) Lines() []string {
	var lines []string

	if len(r.lines) == 0 {
		l := strconv.Itoa(r.Status)
		lines = append(lines, l+"\n")
		return lines
	}

	for i, line := range r.lines {
		l := ""
		if i == len(r.lines)-1 {
			l = strconv.Itoa(r.Status) + " " + line + "\r\n"
		} else {
			l = strconv.Itoa(r.Status) + "-" + line + "\r\n"
		}
		lines = append(lines, l)
	}

	return lines
}

// ReplyServiceReady creates a welcome reply
func ReplyServiceReady(identification string) *Reply {
	return &Reply{CODE_SERVICE_READY, []string{identification}, nil}
}

// ReplyBye creates a Bye reply
func ReplyBye() *Reply { return &Reply{CODE_SERVICE_CLOSING, []string{"Bye"}, nil} }

// ReplyAuthOk creates a authentication successful reply
func ReplyAuthOk() *Reply {
	return &Reply{CODE_AUTHENTICATION_SUCCESS, []string{"Authentication successful"}, nil}
}

// ReplyOk creates a OK reply
func ReplyOk(message ...string) *Reply {
	if len(message) == 0 {
		message = []string{"Ok"}
	}
	return &Reply{CODE_ACTION_OK, message, nil}
}

// ReplyUnrecognisedCommand creates Unrecognised command reply
func ReplyUnrecognisedCommand() *Reply {
	return &Reply{CODE_SYNTAX_ERROR, []string{"Unrecognised command"}, nil}
}

// ReplyLineTooLong due to exceeding these limits
func ReplyLineTooLong() *Reply { return &Reply{CODE_SYNTAX_ERROR, []string{"Line too long."}, nil} }
