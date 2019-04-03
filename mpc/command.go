package mpc

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
