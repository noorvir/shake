package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ShakeCommand struct {
	Use  string
	Args []string
}

type Shakefile struct {
	Backend  string         `json:"backend"`
	Commands []ShakeCommand `json:"commands"`
}

func (s Shakefile) String() string {
	res, _ := json.MarshalIndent(s, "", "  ")
	return string(res)
}

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

func writeShakefile(f Shakefile) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := filepath.Join(pwd, "Shakefile")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	res, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(res)
	if err != nil {
		return err
	}

	return nil
}
