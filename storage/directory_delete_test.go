package storage

import (
    "fmt"
    "github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
    "github.com/mailhedgehog/MailHedgehog/gounit"
    "testing"
)

func TestDelete(t *testing.T) {
    storage := CreateDirectoryStorage("")

    (*gounit.T)(t).AssertEqualsInt(0, storage.Count(""))

    for i := 0; i < 2; i++ {
        id := smtpMessage.MessageID(fmt.Sprint(i))
        msg := &smtpMessage.SMTPMail{
            ID: id,
        }

        storedId, err := storage.Store("", msg)
        (*gounit.T)(t).AssertEqualsString(string(id), string(storedId))
        (*gounit.T)(t).AssertNotError(err)
    }

    (*gounit.T)(t).AssertEqualsInt(2, storage.Count(""))

    (*gounit.T)(t).AssertNotError(storage.Delete("", "1"))
    (*gounit.T)(t).AssertEqualsInt(1, storage.Count(""))

    (*gounit.T)(t).AssertNotError(storage.Delete("", "0"))
    (*gounit.T)(t).AssertEqualsInt(0, storage.Count(""))

    (*gounit.T)(t).ExpectError(storage.Delete("", "1"))
    (*gounit.T)(t).AssertEqualsInt(0, storage.Count(""))
}
