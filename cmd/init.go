package cmd

import (
	_ "embed"
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"github.com/mailhedgehog/MailHedgehog/userInput"
	"github.com/spf13/cobra"
	"os"
)

//go:embed publish/.mh-config.yml
var configFileContent string

func initApplicationArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.NoArgs(cmd, args); err != nil {
		return err
	}

	return nil
}

func initApplication(cmd *cobra.Command, args []string) {
	defaultConfigFileName := ".mh-config.yml"
	configFileName, err := userInput.Get(fmt.Sprintf("Config file name [%s]:", defaultConfigFileName))
	logger.PanicIfError(err)
	if len(configFileName) <= 0 {
		configFileName = defaultConfigFileName
	}
	if _, err := os.Stat(configFileName); err == nil {
		if flagForce {
			logManager().Warning(fmt.Sprintf("Rewriting file `%s`", configFileName))
		} else {
			logManager().Error(fmt.Sprintf("File `%s` already exists, use flag `-F` to rewrite file", configFileName))
			return
		}
	}

	file, err := os.Create(configFileName)
	logger.PanicIfError(err)
	defer file.Close()

	_, err = file.WriteString(configFileContent)
	logger.PanicIfError(err)

	logManager().Info(fmt.Sprintf("Config file `%s` created.", configFileName))
}
