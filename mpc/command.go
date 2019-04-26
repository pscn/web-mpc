package mpc

import "fmt"

// Command simple command from the web

// CommandType to identify the type of command
type CommandType string

// CommandTypes
const (
	TypePlay          CommandType = "play"
	TypeResume        CommandType = "resume"
	TypePause         CommandType = "pause"
	TypeStop          CommandType = "stop"
	TypeNext          CommandType = "next"
	TypePrevious      CommandType = "previous"
	TypeAdd           CommandType = "add"
	TypeAddPrio       CommandType = "addPrio"
	TypeClean         CommandType = "clean"
	TypeRemove        CommandType = "remove"
	TypeSearch        CommandType = "search"
	TypeSearchPage    CommandType = "searchPage"
	TypeUpdateRequest CommandType = "updateRequest"
	TypeBrowse        CommandType = "browse"
	TypePlaylists     CommandType = "playlists"
	TypePlaylistsPage CommandType = "playlistsPage"
	TypeModeRandom    CommandType = "random"
	TypeModeRepeat    CommandType = "repeat"
	TypeModeSingle    CommandType = "single"
	TypeModeConsume   CommandType = "consume"
	TypePrio          CommandType = "prio"
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
