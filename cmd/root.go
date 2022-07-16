package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func readShakefile() (Shakefile, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return Shakefile{}, err
	}

	path := filepath.Join(pwd, "Shakefile")
	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return Shakefile{}, NotInitializedError
		}
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return Shakefile{}, err
	}

	var shakefile Shakefile
	if err = json.Unmarshal(file, &shakefile); err != nil {
		return Shakefile{}, InvalidShakefileError
	}
	return shakefile, nil
}

func writeShakefile() error {
	return nil
}

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
