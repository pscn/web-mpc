package mpc

import "fmt"

// Command simple command from the web

// CommandType to identify the type of command
type CommandType uint

// CommandTypes
const (
	Play          CommandType = 0
	Resume        CommandType = 1
	Pause         CommandType = 2
	Stop          CommandType = 3
	Next          CommandType = 4
	Previous      CommandType = 5
	Add           CommandType = 6
	Remove        CommandType = 7
	Search        CommandType = 8
	StatusRequest CommandType = 9
)

// Command from the web
type Command struct {
	Command CommandType `json:"command"`
	Data    string      `json:"data"`
}

func (cmd *Command) String() string {
	switch cmd.Command {
	case Play:
		return fmt.Sprintf("Play %s", cmd.Data)
	case Resume:
		return "Resume"
	case Pause:
		return "Pause"
	case Stop:
		return "Stop"
	case Next:
		return "Next"
	case Previous:
		return "Previous"
	case Add:
		return fmt.Sprintf("Add %s", cmd.Data)
	case Remove:
		return fmt.Sprintf("Remove %s", cmd.Data)
	case Search:
		return fmt.Sprintf("Search %s", cmd.Data)
	case StatusRequest:
		return "StatusRequest"
	default:
		return fmt.Sprintf("Unknown code: %d, data: %s", cmd.Command, cmd.Data)
	}
}

// eof
