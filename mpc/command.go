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
	Browse        CommandType = "browse"
	Random        CommandType = "random"
	Repeat        CommandType = "repeat"
	Single        CommandType = "single"
	Consume       CommandType = "consume"
)

// Command from the web
type Command struct {
	Command CommandType `json:"command"`
	Data    string      `json:"data"`
}

func (cmd *Command) String() string {
	if cmd.Data == "" {
		return fmt.Sprintf("%s", cmd.Command)
	}
	return fmt.Sprintf("%s(%s)", cmd.Command, cmd.Data)
}

// eof
