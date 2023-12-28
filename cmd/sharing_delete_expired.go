package cmd

import (
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/config"
	"github.com/mailhedgehog/messageSharingStorageFileCsv"
	"github.com/spf13/cobra"
	"os"
)

func sharingDeleteExpiredArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.RangeArgs(0, 1)(cmd, args); err != nil {
		return err
	}

	return nil
}

func sharingDeleteExpired(cmd *cobra.Command, args []string) {
	filePath := ""
	if len(args) > 0 {
		filePath = args[0]
	}
	configuration := config.ParseConfig(filePath)
	switch configuration.Sharing.Use {
	case "csv":
		emailSharingDeleteExpired(messageSharingStorageFileCsv.CreateSharingEmailUsingCSV(&configuration.Sharing.CSV))
	case "mongodb":
		// TODO: implement
	default:
		logManager().Error(fmt.Sprintf("Unsupported sharing type [%s]", configuration.Authentication.Use))
	}
}

func emailSharingDeleteExpired(emailSharing *messageSharingStorageFileCsv.MessageSharingStorageFileCsv) {
	result, err := emailSharing.DeleteExpired()
	if err != nil {
		logManager().Critical(err.Error())
		os.Exit(0)
	}

	logManager().Info(fmt.Sprintf("Deleted expired links [%t].", result))
}
