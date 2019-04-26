package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pscn/web-mpc/conv"
	"github.com/pscn/web-mpc/mpc"
	"github.com/pscn/web-mpc/msg"
)

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
	Command      CommandType `json:"command"`
	Data         string      `json:"data"`
	Page         int
	SearchPage   int
	PlaylistPage int
	LastSearch   string
}

// FromJSON the JSON from data to Command
func FromJSON(data []byte) (*Command, error) {
	var c Command
	err := json.Unmarshal(data, &c)
	return &c, err
}

func (c *Command) String() string {
	if c.Data == "" {
		return fmt.Sprintf("%s", c.Command)
	}
	return fmt.Sprintf("%s(%s)", c.Command, c.Data)
}

// Exec the command and return a Message if a reply is required
func (c *Command) Exec(client *mpc.Client) (msg *msg.Message, err error) {
	switch c.Command {
	case TypePlay:
		if c.Data != "" {
			err = client.Play(conv.ToInt(c.Data))
		} else {
			err = client.Play(-1)
		}

	case TypeResume:
		err = client.Pause(false)

	case TypePause:
		err = client.Pause(true)

	case TypeStop:
		err = client.Stop()

	case TypeNext:
		err = client.Next()

	case TypePrevious:
		err = client.Previous()

	case TypeAdd:
		err = client.Add(c.Data)

	case TypeRemove:
		err = client.Delete(conv.ToInt(c.Data), -1)

	case TypeClean:
		err = client.Clean()

	case TypePrio:
		args := strings.Split(c.Data, ":")
		err = client.SetPriority(conv.ToInt(args[0]), conv.ToInt(args[1]), -1)

	case TypeAddPrio:
		args := strings.Split(c.Data, ":")
		err = client.AddPrio(conv.ToInt(args[0]), args[1])

	case TypeModeConsume, TypeModeRepeat, TypeModeSingle, TypeModeRandom:
		target := true
		if c.Data == "disable" {
			target = false
		}
		switch c.Command {
		case TypeModeConsume:
			client.Consume(target)

		case TypeModeRepeat:
			client.Repeat(target)

		case TypeModeSingle:
			client.Single(target)

		case TypeModeRandom:
			client.Random(target)
		}

	case TypeUpdateRequest:
		msg = client.Update(conv.ToInt(c.Data))

	case TypeSearch:
		c.SearchPage = 1
		c.LastSearch = c.Data
		msg = client.Search(c.Data, c.SearchPage)

	case TypeSearchPage:
		c.SearchPage = conv.ToInt(c.Data)
		msg = client.Search(c.LastSearch, c.SearchPage)

	case TypeBrowse:
		msg = client.ListDirectory(c.Data)

	case TypePlaylists:
		c.PlaylistPage = 1
		msg = client.ListPlaylists(c.PlaylistPage)

	case TypePlaylistsPage:
		c.PlaylistPage = conv.ToInt(c.Data)
		msg = client.ListPlaylists(c.PlaylistPage)
	}
	return
}

// eof
