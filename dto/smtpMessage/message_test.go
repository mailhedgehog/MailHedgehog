package smtpMessage

import (
	"github.com/mailhedgehog/MailHedgehog/gounit"
	"io"
	"testing"
)

var emailString = `HELO:<[127.0.0.1]>
FROM:<sender@example.test>
TO:<your-email@here.test>
TO:<second@here.test>

From: Sender <sender@example.test>
To: ReceiverName <your-email@here.test>
X-Priority: 1 (Highest)
Subject: Mail test 2023-03-25 22:16:52
Message-ID: <3ac479f00c5d9ea7519ade0784ed1060@example.test>
MIME-Version: 1.0
Date: Sat, 25 Mar 2023 22:16:52 +0200
Content-Type: multipart/alternative; boundary=UhEiB9Sb

--UhEiB9Sb
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: quoted-printable

This is an important message!
--UhEiB9Sb
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: quoted-printable

<strong>This is an test message!</strong>
--UhEiB9Sb--`

func TestFromString(t *testing.T) {
	smtpMessage := FromString(emailString)
	(*gounit.T)(t).AssertEqualsString("[127.0.0.1]", smtpMessage.Helo)
	(*gounit.T)(t).AssertEqualsString("sender@example.test", smtpMessage.From)
	(*gounit.T)(t).AssertEqualsString("your-email@here.test", smtpMessage.To[0])
	(*gounit.T)(t).AssertEqualsString("second@here.test", smtpMessage.To[1])
	(*gounit.T)(t).AssertLengthString(586, smtpMessage.Data)

}

func TestToReader(t *testing.T) {
	smtpMessage := FromString(emailString)

	bytes, err := io.ReadAll(smtpMessage.ToReader())
	(*gounit.T)(t).AssertNotError(err)

	(*gounit.T)(t).AssertLengthString(686, string(bytes))

}
