package mpc

// Command simple command from the web

// CommandType to identify the type of command
type CommandType uint

const (
	Play   CommandType = 0
	Add    CommandType = 1
	Search CommandType = 2
)

// Command from the web
type Command struct {
	Command string `json:"command"`
	Data    string `json:"data"`
}
