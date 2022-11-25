package cmd

import (
	"github.com/mailpiggy/MailPiggy/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "MailPiggy",
	Short: "Mail storage system",
	Long:  ``,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(authAddCmd)
}

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("MailPiggy")
	}
	return configuredLogger
}

var authAddCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run MailPiggy application",
	Long:  `Run MailPiggy client and server`,
	Args:  serveArgs,
	Run:   serve,
}
