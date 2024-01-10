package cmd

import (
	"errors"
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/config"
	"github.com/mailhedgehog/MailHedgehog/dbConnectionMongo"
	"github.com/mailhedgehog/MailHedgehog/userInput"
	"github.com/mailhedgehog/authenticationFile"
	"github.com/mailhedgehog/authenticationMongo"
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/logger"
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
		addUser(authenticationFile.CreateFileAuthentication(&configuration.Authentication.File, &configuration.Authentication.Config))
	case "mongodb":
		conf := configuration.DB.Connections[configuration.Authentication.MongoDB.Connection]
		if conf == nil {
			logger.PanicIfError(errors.New(fmt.Sprintf("Undefined db connection [%s]", configuration.Authentication.MongoDB.Connection)))
		}
		addUser(authenticationMongo.CreateMongoDbAuthentication(
			dbConnectionMongo.MakeCollection(conf, configuration.Authentication.MongoDB.Collection),
			&configuration.Authentication.Config,
		))
	default:
		logManager().Error(fmt.Sprintf("Unsupported auth type [%s]", configuration.Authentication.Use))
	}
}

func addUser(auth contracts.Authentication) {
	roomName, err := userInput.Get("Please input room name:")
	logger.PanicIfError(err)
	err = validateMinMaxLength(roomName, 0, 20)
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}

	if auth.UsersStorage().Exists(roomName) {
		logManager().Critical(fmt.Sprintf("Room [%s] already present in credentials list.", roomName))
		os.Exit(0)
	}

	httpPassword, err := userInput.GetSecret("Please set password for http login:")
	logger.PanicIfError(err)
	err = validateMinMaxLength(httpPassword, 6, 20)
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}

	smtpPassword, err := userInput.GetSecret("Please set password for smtp login(optional, if empty will be used http password):")
	logger.PanicIfError(err)
	err = validateMinMaxLength(httpPassword, 6, 20)
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}

	err = auth.UsersStorage().Add(roomName)
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}

	if len(httpPassword) > 0 {
		err = auth.Dashboard().ViaPasswordAuthentication().SetPassword(roomName, httpPassword)
		if err != nil {
			logManager().Critical(err.Error())
			os.Exit(0)
		}
	}

	if len(smtpPassword) > 0 {
		err = auth.SMTP().ViaPasswordAuthentication().SetPassword(roomName, smtpPassword)
		if err != nil {
			logManager().Critical(err.Error())
			os.Exit(0)
		}
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
