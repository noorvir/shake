package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RootCmd(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	// parse the args and validate them against the Shakefile
	shakefile, err := readShakefile()
	if err != nil {
		return fmt.Errorf("failed to load Shakefile, %v", err)
	}
	fmt.Printf("%s\n", shakefile)

	return nil
}
