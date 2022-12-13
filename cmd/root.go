package cmd

import (
    "github.com/mailhedgehog/MailHedgehog/logger"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "MailHedgehog",
    Short: "Mail storage system",
    Long:  ``,
}

func Execute() error {
    return rootCmd.Execute()
}

var flagForce bool

func init() {
    initCmd.Flags().BoolVarP(&flagForce, "force", "F", false, "Force override files")

    rootCmd.AddCommand(initCmd)
    rootCmd.AddCommand(serveCmd)
    rootCmd.AddCommand(authFileAddCmd)
}

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
    if configuredLogger == nil {
        configuredLogger = logger.CreateLogger("MailHedgehog")
    }
    return configuredLogger
}

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Initialise application",
    Long:  `Create configuration files`,
    Args:  initApplicationArgs,
    Run:   initApplication,
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
