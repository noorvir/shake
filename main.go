package shake

import (
	"os"

	"github.com/noorvir/shake/cmd"
	"github.com/noorvir/shake/constants"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "shake",
	Short:         "Manage you internal scripts like a pro",
	RunE:          cmd.RootCmd,
	SilenceUsage:  false,
	SilenceErrors: false,
}

func init() {
	rootCmd.Version = constants.Version
	rootCmd.Flags().BoolP("version", "v", false, "Get the version of current Shake CLI")

	// Init
	rootCmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "Init a new shake project",
		RunE:  cmd.InitCmd,
	})
}

func main() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
