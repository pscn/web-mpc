package mpc

import "fmt"

// Command simple command from the web

// Command from the web
type Command struct {
	Command string `json:"command"`
	Data    string `json:"data"`
}

func (cmd *Command) String() string {
	return fmt.Sprintf("cmd: %s, data:%s", cmd.Command, cmd.Data)
}

// eof
