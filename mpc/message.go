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
	Error         MessageType = "error"
	Info          MessageType = "info"
	Status        MessageType = "status"
	CurrentSong   MessageType = "currentSong"
	Playlist      MessageType = "playlist"
	SearchResult  MessageType = "searchResult"
	DirectoryList MessageType = "directoryList"
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
	case CurrentSong:
		return fmt.Sprintf("%s: %+v", msg.Type, msg.Data.(*SongData))
	case Playlist:
		return fmt.Sprintf("%s: %+v", msg.Type, msg.Data.(*PlaylistData))
	case SearchResult:
		return fmt.Sprintf("%s: %+v", msg.Type, msg.Data.(*SearchResultData))
	}
	return fmt.Sprintf("Unknown type: %s", msg.Type)
}

// NewError creates a new Event including an error
func NewError(str string) *Message {
	return NewMessage(Error, str)
}

func (msg *Message) Error() string {
	if msg.Type == Error { // FIXME: how to inform the develeoper?
		return msg.Data.(string)
	}
	return ""
}

// NewInfo creates a new Event including an error
func NewInfo(str string) *Message {
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

// NewStatus creates a new Event including the status data mapped from mpd.Attrs
func NewStatus(attrs *mpd.Attrs) *Message {
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

// NewCurrentSong creates a new Event including the current song data mapped from mpd.Attrs
func NewCurrentSong(attrs *mpd.Attrs) *Message {
	if (*attrs)["AlbumArtist"] == "" {
		(*attrs)["AlbumArtist"] = (*attrs)["Artist"]
	}
	return NewMessage(CurrentSong, &SongData{
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
	if msg.Type == CurrentSong { // FIXME: how to inform the develeoper?
		return msg.Data.(*SongData)
	}
	return nil
}

// PlaylistData converted from *mpd.attrs
type PlaylistData struct {
	Playlist []SongData
}

// NewCurrentPlaylist creates a new Event including the current song data mapped from mpd.Attrs
func NewCurrentPlaylist(attrArr *[]mpd.Attrs) *Message {
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
	return NewMessage(Playlist, event)
}

// CurrentPlaylist payload of an event
func (msg *Message) CurrentPlaylist() *PlaylistData {
	if msg.Type == Playlist { // FIXME: how to inform the develeoper?
		return msg.Data.(*PlaylistData)
	}
	return nil
}

// SearchResultData converted from *mpd.attrs
type SearchResultData struct {
	SearchResult []SongData
}

// NewSearchResult from mpd.Attrs
func NewSearchResult(attrArr *[]mpd.Attrs) *Message {
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

// NewDirectoryList from mpd.Attrs
func NewDirectoryList(currentDirectory string, attrArr *[]mpd.Attrs) *Message {
	event := &DirectoryListData{
		Parent: currentDirectory,
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
