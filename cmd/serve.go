package cmd

import (
	"github.com/mailhedgehog/MailHedgehog/config"
	"github.com/mailhedgehog/MailHedgehog/server"
	"github.com/spf13/cobra"
)

func serveArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.RangeArgs(0, 1)(cmd, args); err != nil {
		return err
	}

	return nil
}

func serve(cmd *cobra.Command, args []string) {
	filePath := ""
	if len(args) > 0 {
		filePath = args[0]
	}

	server.Start(server.Configure(config.ParseConfig(filePath)))
}
