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
	AddPrio       CommandType = "addPrio"
	Clean         CommandType = "clean"
	Remove        CommandType = "remove"
	Search        CommandType = "search"
	SearchPage    CommandType = "searchPage"
	UpdateRequest CommandType = "updateRequest"
	Browse        CommandType = "browse"
	Playlists     CommandType = "playlists"
	PlaylistsPage CommandType = "playlistsPage"
	ModeRandom    CommandType = "random"
	ModeRepeat    CommandType = "repeat"
	ModeSingle    CommandType = "single"
	ModeConsume   CommandType = "consume"
	Prio          CommandType = "prio"
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
