package mpc

import "fmt"

// Command simple command from the web

// CommandType to identify the type of command
type CommandType string

// CommandTypes
const (
	Play          CommandType = "play"
	Resume        CommandType = "resume"
	Pause         CommandType = "pause"
	Stop          CommandType = "stop"
	Next          CommandType = "next"
	Previous      CommandType = "previous"
	Add           CommandType = "add"
	Remove        CommandType = "remove"
	Search        CommandType = "search"
	StatusRequest CommandType = "statusRequest"
)

// Command from the web
type Command struct {
	Command CommandType `json:"command"`
	Data    string      `json:"data"`
}

func (cmd *Command) String() string {
	return fmt.Sprintf("%s: %s", cmd.Command, cmd.Data)
}

// eof
