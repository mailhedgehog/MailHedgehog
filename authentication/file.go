package authentication

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mailpiggy/MailPiggy/logger"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

// FileAuth represents the authentication handler using file
type FileAuth struct {
	users             map[string]userInfo
	authenticatedUser *AuthenticatedUser
}

func CreateFileAuthentication(authFilePath string) *FileAuth {
	authFile := &FileAuth{}
	authFile.AuthFile(authFilePath)

	return authFile
}

func (fileAuth *FileAuth) AuthenticatedUser() *AuthenticatedUser {
	return fileAuth.authenticatedUser
}

func (fileAuth *FileAuth) Authenticate(authType AuthenticationType, username string, password string) bool {
	fileAuth.setAuthenticatedUser("")
	if fileAuth.users == nil {
		fileAuth.setAuthenticatedUser(username)
		return true
	}

	user, ok := fileAuth.users[username]
	if !ok {
		return false
	}

	passwordHashToCheck := user.httpPass
	if authType == SMTP {
		passwordHashToCheck = user.smtpPass
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHashToCheck), []byte(password)); err != nil {
		return false
	}

	fileAuth.setAuthenticatedUser(username)
	return true
}

func (fileAuth *FileAuth) setAuthenticatedUser(username string) {
	fileAuth.authenticatedUser = nil
	if len(username) > 0 {
		fileAuth.authenticatedUser = &AuthenticatedUser{Username: username}
	}
}

// AuthFile scan file and add users to memory
func (fileAuth *FileAuth) AuthFile(path string) int {
	fileAuth.users = nil

	if len(path) <= 0 {
		logManager().Debug("File auth empty.")
		return 0
	}

	file, err := os.Open(path)
	logger.PanicIfError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		err := fileAuth.addUserFromFileLine(scanner.Text())
		if err != nil {
			logManager().Notice(err.Error())
		} else {
		}
	}

	if fileAuth.users == nil {
		return 0
	}
	return len(fileAuth.users)
}

func (fileAuth *FileAuth) addUserFromFileLine(line string) error {
	line = strings.TrimSpace(line)
	infoSlice := strings.Split(line, ":")
	for i := range infoSlice {
		infoSlice[i] = strings.TrimSpace(infoSlice[i])
	}

	if len(infoSlice) < 2 {
		return errors.New("at least should be present username and password")
	}

	if len(infoSlice[0]) <= 0 || len(infoSlice[1]) <= 0 {
		return errors.New("username and password can't be empty")
	}

	smtpPass := infoSlice[1]
	if len(infoSlice) > 2 && len(infoSlice[2]) > 0 {
		smtpPass = infoSlice[2]
	}

	if fileAuth.users == nil {
		fileAuth.users = make(map[string]userInfo)
	}

	fileAuth.users[infoSlice[0]] = userInfo{
		username: infoSlice[0],
		httpPass: infoSlice[1],
		smtpPass: smtpPass,
	}

	logManager().Debug(fmt.Sprintf("Processes users: '%s'", infoSlice[0]))

	return nil
}
