package authorisation

import (
	"github.com/mailpiggy/MailPiggy/logger"
	"os"
	"testing"
)

var filePath = ""

func init() {
	dir, err := os.MkdirTemp("", "mailpiggy_")
	logger.PanicIfError(err)

	filePath = dir + string(os.PathSeparator) + "auth.file"
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
	authorisation := CreateFileAuthorisation(filePath)

	if len(authorisation.users) != 3 {
		t.Errorf("Invalid users count expected: %d, got: %d", 3, len(authorisation.users))
	}

	authorisation.AuthFile(filePath + "2")

	if len(authorisation.users) > 0 {
		t.Errorf("Invalid users count expected: %d, got: %d", 0, len(authorisation.users))
	}

	countUsers := authorisation.AuthFile(filePath)

	if len(authorisation.users) != 3 {
		t.Errorf("Invalid users count expected: %d, got: %d", 3, len(authorisation.users))
	}
	if countUsers != 3 {
		t.Errorf("Invalid users count expected: %d, got: %d", 3, countUsers)
	}
}

func TestAuthorised(t *testing.T) {
	authorisation := CreateFileAuthorisation(filePath)

	if !authorisation.Authorised(HTTP, "user1", "test1") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", true, false)
	}
	if !authorisation.Authorised(SMTP, "user1", "test2") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", true, false)
	}

	if !authorisation.Authorised(HTTP, "user2", "test1") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", true, false)
	}
	if !authorisation.Authorised(SMTP, "user2", "test1") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", true, false)
	}

	if !authorisation.Authorised(HTTP, "user4", "test1") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", true, false)
	}
	if !authorisation.Authorised(SMTP, "user4", "test1") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", true, false)
	}

	if authorisation.Authorised(HTTP, "user4", "foo") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", false, true)
	}
	if authorisation.Authorised(SMTP, "user4", "bar") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", false, true)
	}
	if authorisation.Authorised(HTTP, "userX", "test1") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", false, true)
	}
	if authorisation.Authorised(SMTP, "userX", "test1") {
		t.Errorf("authorisation.Authorised expected: %t, got: %t", false, true)
	}
}
