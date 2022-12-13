package authentication

import (
	"github.com/mailhedgehog/MailHedgehog/gounit"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"os"
	"testing"
)

var filePath = ""

func init() {
	dir, err := os.MkdirTemp("", "mailhedgehog_")
	logger.PanicIfError(err)

	filePath = dir + string(os.PathSeparator) + ".mh-authfile"
	file, err := os.Create(filePath)
	logger.PanicIfError(err)

	passTest1 := "$2a$12$CV3q6WzQBGEPqrPkh.hYn.HFO6mAxKfLLNxAMWIKx9wF93X6539nS"
	passTest2 := "$2a$12$6aBv1ox1kgMBcS9st4ixdu6HKW77DNdpyJNENN5vVMFqHHcF.q5Ra"

	fileLines := [][]byte{
		[]byte("user1:" + passTest1 + ":" + passTest2 + "\n"),
		[]byte("user2:" + passTest1 + "\n"),
		[]byte("user3::" + passTest2 + "\n"),
		[]byte("user4:" + passTest1 + ":\n"),
		[]byte(":::\n"),
		[]byte(":user5\n"),
	}
	for _, line := range fileLines {
		_, err = file.Write(line)
		logger.PanicIfError(err)
	}
	file.Sync()
	file.Close()

	file, err = os.Create(filePath + "2")
	logger.PanicIfError(err)
	file.Close()
}

func TestAuthFile(t *testing.T) {
	auth := CreateFileAuthentication(filePath)

	(*gounit.T)(t).AssertEqualsInt(3, len(auth.users))

	auth.AuthFile(filePath + "2")

	(*gounit.T)(t).AssertLessOrEqualInt(len(auth.users), 0)

	countUsers := auth.AuthFile(filePath)

	(*gounit.T)(t).AssertEqualsInt(3, len(auth.users))
	(*gounit.T)(t).AssertEqualsInt(3, countUsers)
}

func TestAuthorised(t *testing.T) {
	auth := CreateFileAuthentication(filePath)

	(*gounit.T)(t).AssertTrue(auth.Authenticate(HTTP, "user1", "test1"))
	(*gounit.T)(t).AssertTrue(auth.Authenticate(SMTP, "user1", "test2"))

	(*gounit.T)(t).AssertTrue(auth.Authenticate(HTTP, "user2", "test1"))
	(*gounit.T)(t).AssertTrue(auth.Authenticate(HTTP, "user2", "test1"))

	(*gounit.T)(t).AssertTrue(auth.Authenticate(HTTP, "user4", "test1"))
	(*gounit.T)(t).AssertTrue(auth.Authenticate(HTTP, "user4", "test1"))

	(*gounit.T)(t).AssertFalse(auth.Authenticate(HTTP, "user4", "foo"))
	(*gounit.T)(t).AssertFalse(auth.Authenticate(SMTP, "user4", "bar"))
	(*gounit.T)(t).AssertFalse(auth.Authenticate(HTTP, "userX", "test1"))
	(*gounit.T)(t).AssertFalse(auth.Authenticate(SMTP, "userX", "test1"))
}
