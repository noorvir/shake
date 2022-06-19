package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	return nil
}
