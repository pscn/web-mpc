package mpc

import (
	"fmt"
	"log"

	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/helpers"
)

// MessageType to identify the type of message
type MessageType string

// MessageTypes
const (
	Error          MessageType = "error"
	Info           MessageType = "info"
	Status         MessageType = "status"
	ActiveSong     MessageType = "activeSong"
	ActivePlaylist MessageType = "activePlaylist"
	SearchResult   MessageType = "searchResult"
	DirectoryList  MessageType = "directoryList"
)

// Message from the clients EventLoop
type Message struct {
	Type MessageType `json:"type"`
	Data interface{} `json:"data"`
}

// NewMessage creates a new event with type and data
func NewMessage(msgType MessageType, data interface{}) *Message {
	return &Message{msgType, data}
}

func (msg *Message) String() string {
	switch msg.Type {
	case Error, Info:
		return fmt.Sprintf("%s: %s", msg.Type, msg.Data.(string))
	case Status:
		return fmt.Sprintf("%s: %+v", msg.Type, msg.Data.(*StatusData))
	case ActiveSong:
		return fmt.Sprintf("%s: %+v", msg.Type, msg.Data.(*SongData))
	case ActivePlaylist:
		return fmt.Sprintf("%s: %+v", msg.Type, msg.Data.(*PlaylistData))
	case SearchResult:
		return fmt.Sprintf("%s: %+v", msg.Type, msg.Data.(*SearchResultData))
	}
	return fmt.Sprintf("Unknown type: %s", msg.Type)
}

// ErrorMsg creates a new Event including an error
func ErrorMsg(str string) *Message {
	return NewMessage(Error, str)
}

func (msg *Message) Error() string {
	if msg.Type == Error { // FIXME: how to inform the develeoper?
		return msg.Data.(string)
	}
	return ""
}

// InfoMsg creates a new Event including an error
func InfoMsg(str string) *Message {
	return NewMessage(Info, str)
}

// StatusData represents the data we get as MPC status
type StatusData struct {
	Duration float32 `json:"duration"`
	Elapsed  float32 `json:"elapsed"`
	State    string  `json:"state"`
	Volume   int     `json:"volume"`
	Consume  bool    `json:"consume"`
	Random   bool    `json:"random"`
	Single   bool    `json:"single"`
	Repeat   bool    `json:"repeat"`
}

// StatusMsg creates a new Event including the status data mapped from mpd.Attrs
func StatusMsg(attrs *mpd.Attrs) *Message {
	return NewMessage(Status, &StatusData{
		Duration: helpers.ToFloat((*attrs)["duration"]),
		Elapsed:  helpers.ToFloat((*attrs)["elapsed"]),
		Consume:  helpers.ToBool((*attrs)["consume"]),
		Random:   helpers.ToBool((*attrs)["random"]),
		Single:   helpers.ToBool((*attrs)["single"]),
		Repeat:   helpers.ToBool((*attrs)["repeat"]),
		Volume:   helpers.ToInt((*attrs)["volume"]),
		State:    (*attrs)["state"],
	})
}

// Status payload of an event
func (msg *Message) Status() *StatusData {
	if msg.Type == Status { // FIXME: how to inform the develeoper?
		return msg.Data.(*StatusData)
	}
	return nil
}

// SongData converted from *mpd.attrs
type SongData struct {
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	AlbumArtist string `json:"album_artist"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	File        string `json:"file"`
	Genre       string `json:"genre"`
	Released    string `json:"released"`
}

// ActiveSongMsg creates a new Event including the current song data mapped from mpd.Attrs
func ActiveSongMsg(attrs *mpd.Attrs) *Message {
	if (*attrs)["AlbumArtist"] == "" {
		(*attrs)["AlbumArtist"] = (*attrs)["Artist"]
	}
	return NewMessage(ActiveSong, &SongData{
		Artist:      (*attrs)["Artist"],
		Album:       (*attrs)["Album"],
		AlbumArtist: (*attrs)["AlbumArtist"],
		Title:       (*attrs)["Title"],
		Duration:    helpers.ToInt((*attrs)["Time"]),
		File:        (*attrs)["file"],
		Genre:       (*attrs)["Genre"],
		Released:    (*attrs)["Date"],
	})
}

// CurrentSong payload of an event
func (msg *Message) CurrentSong() *SongData {
	if msg.Type == ActiveSong { // FIXME: how to inform the develeoper?
		return msg.Data.(*SongData)
	}
	return nil
}

// PlaylistData converted from *mpd.attrs
type PlaylistData struct {
	Playlist []SongData
}

// ActivePlaylistMsg creates a new Event including the current song data mapped from mpd.Attrs
func ActivePlaylistMsg(attrArr *[]mpd.Attrs) *Message {
	event := &PlaylistData{}
	event.Playlist = make([]SongData, len(*attrArr))
	for i, attrs := range *attrArr {
		if attrs["AlbumArtist"] == "" {
			attrs["AlbumArtist"] = attrs["Artist"]
		}
		event.Playlist[i] = SongData{
			Artist:      attrs["Artist"],
			Album:       attrs["Album"],
			AlbumArtist: attrs["AlbumArtist"],
			Title:       attrs["Title"],
			Duration:    helpers.ToInt(attrs["Time"]),
			File:        attrs["file"],
			Genre:       attrs["Genre"],
			Released:    attrs["Date"],
		}
	}
	return NewMessage(ActivePlaylist, event)
}

// CurrentPlaylist payload of an event
func (msg *Message) CurrentPlaylist() *PlaylistData {
	if msg.Type == ActivePlaylist { // FIXME: how to inform the develeoper?
		return msg.Data.(*PlaylistData)
	}
	return nil
}

// SearchResultData converted from *mpd.attrs
type SearchResultData struct {
	SearchResult []SongData
}

// SearchResultMsg from mpd.Attrs
func SearchResultMsg(attrArr *[]mpd.Attrs) *Message {
	event := &SearchResultData{}
	if attrArr == nil {
		return NewMessage(SearchResult, event)
	}
	event.SearchResult = make([]SongData, len(*attrArr))
	for i, attrs := range *attrArr {
		if attrs["AlbumArtist"] == "" {
			attrs["AlbumArtist"] = attrs["Artist"]
		}
		event.SearchResult[i] = SongData{
			Artist:      attrs["Artist"],
			Album:       attrs["Album"],
			AlbumArtist: attrs["AlbumArtist"],
			Title:       attrs["Title"],
			Duration:    helpers.ToInt(attrs["Time"]),
			File:        attrs["file"],
			Genre:       attrs["Genre"],
			Released:    attrs["Date"],
		}
	}
	return NewMessage(SearchResult, event)
}

// SearchResult payload of an event
func (msg *Message) SearchResult() *SearchResultData {
	if msg.Type == SearchResult { // FIXME: how to inform the develeoper?
		return msg.Data.(*SearchResultData)
	}
	return nil
}

type DirectoryListEntry struct {
	Type      string `json:"type"`
	Directory string `json:"directory"`

	Artist      string `json:"artist"`
	Album       string `json:"album"`
	AlbumArtist string `json:"album_artist"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	File        string `json:"file"`
	Genre       string `json:"genre"`
	Released    string `json:"released"`
}
type DirectoryListData struct {
	Parent        string               `json:"parent"`
	DirectoryList []DirectoryListEntry `json:"directoryList"`
}

// DirectoryListMsg from mpd.Attrs
func DirectoryListMsg(previousDirectory string, attrArr *[]mpd.Attrs) *Message {
	event := &DirectoryListData{
		Parent: previousDirectory,
	}
	if attrArr == nil {
		return NewMessage(DirectoryList, event)
	}
	cnt := 0
	for _, attrs := range *attrArr {
		if attrs["directory"] != "" {
			cnt++
		} else if attrs["file"] != "" {
			cnt++
		}
	}
	event.DirectoryList = make([]DirectoryListEntry, cnt)
	i := 0
	for _, attrs := range *attrArr {
		log.Printf("%+v\n", attrs)
		if attrs["directory"] != "" {
			event.DirectoryList[i] = DirectoryListEntry{
				Type:      "directory",
				Directory: attrs["directory"],
			}
			i++
		} else if attrs["file"] != "" {
			if attrs["albumartist"] == "" {
				attrs["albumartist"] = attrs["artist"]
			}
			event.DirectoryList[i] = DirectoryListEntry{
				Type:        "file",
				Artist:      attrs["artist"],
				Album:       attrs["album"],
				AlbumArtist: attrs["albumartist"],
				Title:       attrs["title"],
				Duration:    helpers.ToInt(attrs["time"]),
				File:        attrs["file"],
				Genre:       attrs["genre"],
				Released:    attrs["date"],
			}
			i++
		}
	}
	return NewMessage(DirectoryList, event)
}

// DirectoryList payload of an event
func (msg *Message) DirectoryList() *DirectoryListData {
	if msg.Type == DirectoryList { // FIXME: how to inform the develeoper?
		return msg.Data.(*DirectoryListData)
	}
	return nil
}

// eof
