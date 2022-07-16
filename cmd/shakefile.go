package cmd

import "encoding/json"

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
