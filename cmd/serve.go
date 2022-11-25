package cmd

import (
	"github.com/mailpiggy/MailPiggy/storage"
	"github.com/spf13/cobra"
)

func serveArgs(cmd *cobra.Command, args []string) error {
	logManager().Warning("TODO: serveArgs Implement params validation")
	return nil
}

func serve(cmd *cobra.Command, args []string) {
	logManager().Debug("Start")
	storage.CreateDirectoryStorage("")
	storage.SetPerRoomLimit(2)
	logManager().Warning("TODO: serve Implement")
}
