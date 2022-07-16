package cmd

import (
	"fmt"
)

type ShakeError error

var (
	InvalidShakefileError ShakeError = fmt.Errorf("invalid Shakefile")
	NotInitializedError   ShakeError = fmt.Errorf("%s\n", "No Shakefile found in current directory. Run `shake init` to create one.")
)
