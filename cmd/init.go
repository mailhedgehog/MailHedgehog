package cmd

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

//go:embed publish/.mh-config.yml
var configFileContent string

func initApplicationArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.RangeArgs(0, 1)(cmd, args); err != nil {
		return err
	}

	return nil
}

func initApplication(cmd *cobra.Command, args []string) {
	configFileName := ".mh-config.yml"
	if len(args) > 0 {
		configFileName = args[0]
	}
	if _, err := os.Stat(configFileName); err == nil {
		if flagForce {
			logManager().Warning(fmt.Sprintf("Rewriting file `%s`", configFileName))
		} else {
			logManager().Error(fmt.Sprintf("File `%s` already exists, use flag `-F` to rewrite file", configFileName))
			return
		}
	}

	f, err := os.Create(configFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(configFileContent)
	if err != nil {
		panic(err)
	}

	logManager().Info(fmt.Sprintf("Config file `%s` created.", configFileName))
}
