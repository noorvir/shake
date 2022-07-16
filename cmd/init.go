package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func InitCmd(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	if _, err := readShakefile(); err == nil {
		fmt.Println("Shakefile already exists")
		return nil
	}

	var shakefile Shakefile
	if err := writeShakefile(shakefile); err != nil {
		return fmt.Errorf("failed to create Shakefile, %v", err)
	}

	options := &promptui.Select{
		Label: "Would you like to create your first command?",
		Items: []string{"yes", "no"},
	}
	options.Run()

	return nil
}
