package cmd

import "github.com/spf13/cobra"

func InitCmd(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	return nil
}
