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
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(authFileAddCmd)
}

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("MailPiggy")
	}
	return configuredLogger
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run MailHedgehog application",
	Long:  `Run MailHedgehog client and server`,
	Args:  serveArgs,
	Run:   serve,
}

var authFileAddCmd = &cobra.Command{
	Use:   "auth:file:add",
	Short: "Add auth credentials",
	Long:  `Add new authentication credentials`,
	Args:  authFileAddArgs,
	Run:   authFileAdd,
}
