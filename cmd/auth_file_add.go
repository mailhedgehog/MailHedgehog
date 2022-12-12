package cmd

import (
	"errors"
	"fmt"
	"github.com/mailpiggy/MailPiggy/authentication"
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/mailpiggy/MailPiggy/userInput"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func authFileAddArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.RangeArgs(0, 1)(cmd, args); err != nil {
		return err
	}

	return nil
}

func authFileAdd(cmd *cobra.Command, args []string) {
	authFileName := ".mh-authfile"
	if len(args) > 0 {
		authFileName = args[0]
	}
	
	roomName, err := userInput.Get("Please input room name:")
	logger.PanicIfError(err)
	err = validateMinMaxLength(roomName, 0, 20)
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}

	httpPassword, err := userInput.Get("Please set password for http login:")
	logger.PanicIfError(err)
	err = validateMinMaxLength(httpPassword, 6, 20)
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}
	hashHttpPassword, err := createPasswordHash(httpPassword)
	logger.PanicIfError(err)

	hashSmtpPassword := []byte{}
	smtpPassword, err := userInput.Get("Please set password for smtp login(optional, if empty will be used http password):")
	logger.PanicIfError(err)
	if len(smtpPassword) > 0 {
		err = validateMinMaxLength(smtpPassword, 6, 20)
		if err != nil {
			logManager().Critical(err.Error())
			os.Exit(0)
		}
		hashSmtpPassword, err = createPasswordHash(httpPassword)
		logger.PanicIfError(err)
	}

	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(authFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logger.PanicIfError(err)

	auth := authentication.CreateFileAuthentication(authFileName)
	if auth.UsernamePresent(roomName) {
		logManager().Critical(fmt.Sprintf("Room [%s] already present in credentials list.", roomName))
		os.Exit(0)
	}

	_, err = file.WriteString(fmt.Sprintf("%s:%s:%s\n", roomName, hashHttpPassword, hashSmtpPassword))
	logger.PanicIfError(err)

	logManager().Info(fmt.Sprintf("Room [%s] credentials added.", roomName))
}

func validateMinMaxLength(input string, min int, max int) error {
	if len(input) <= min {
		return errors.New(fmt.Sprintf("Input lingth is less or equal than max: %d", min))
	}

	if len(input) > max {
		return errors.New(fmt.Sprintf("Input lingth is bigger than max: %d", max))
	}

	return nil
}

func createPasswordHash(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
