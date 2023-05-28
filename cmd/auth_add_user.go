package cmd

import (
	"errors"
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/authentication"
	"github.com/mailhedgehog/MailHedgehog/config"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"github.com/mailhedgehog/MailHedgehog/userInput"
	"github.com/spf13/cobra"
	"os"
)

func authAddUserArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.RangeArgs(0, 1)(cmd, args); err != nil {
		return err
	}

	return nil
}

func authAddUser(cmd *cobra.Command, args []string) {
	filePath := ""
	if len(args) > 0 {
		filePath = args[0]
	}
	configuration := config.ParseConfig(filePath)
	switch configuration.Authentication.Use {
	case "file":
		addUser(authentication.CreateFileAuthentication(configuration.Authentication.File.Path))
	case "mongodb":
		addUser(authentication.CreateMongoDbAuthentication(
			configuration.DB.GetMongoDBConnection(
				configuration.Authentication.MongoDB.Connection,
			).Collection(
				configuration.Authentication.MongoDB.Collection,
			),
		))
	default:
		logManager().Error(fmt.Sprintf("Unsupported auth type [%s]", configuration.Authentication.Use))
	}
}

func addUser(auth authentication.Authentication) {
	roomName, err := userInput.Get("Please input room name:")
	logger.PanicIfError(err)
	err = validateMinMaxLength(roomName, 0, 20)
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}

	if auth.UsernamePresent(roomName) {
		logManager().Critical(fmt.Sprintf("Room [%s] already present in credentials list.", roomName))
		os.Exit(0)
	}

	httpPassword, err := userInput.GetSilent("Please set password for http login:")
	logger.PanicIfError(err)
	err = validateMinMaxLength(httpPassword, 6, 20)
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}
	hashHttpPassword, err := authentication.CreatePasswordHash(httpPassword)
	logger.PanicIfError(err)

	hashSmtpPassword := []byte{}
	smtpPassword, err := userInput.GetSilent("Please set password for smtp login(optional, if empty will be used http password):")
	logger.PanicIfError(err)
	if len(smtpPassword) > 0 {
		err = validateMinMaxLength(smtpPassword, 6, 20)
		if err != nil {
			logManager().Critical(err.Error())
			os.Exit(0)
		}
		hashSmtpPassword, err = authentication.CreatePasswordHash(smtpPassword)
		logger.PanicIfError(err)
	}

	err = auth.AddUser(roomName, string(hashHttpPassword), string(hashSmtpPassword))
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}

	logManager().Info(fmt.Sprintf("Room [%s] credentials added.", roomName))
}

func validateMinMaxLength(input string, min int, max int) error {
	if len(input) <= min {
		return errors.New(fmt.Sprintf("Input length is less or equal than max: %d", min))
	}

	if len(input) > max {
		return errors.New(fmt.Sprintf("Input length is bigger than max: %d", max))
	}

	return nil
}
