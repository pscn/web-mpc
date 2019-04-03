package mpc

import (
	"fmt"

	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/helpers"
)

// MessageType to identify the type of message
type MessageType uint

// MessageTypes
const (
	Error        MessageType = 0
	Info         MessageType = 1
	Status       MessageType = 2
	CurrentSong  MessageType = 3
	Playlist     MessageType = 4
	SearchResult MessageType = 5
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
	case Error:
		return fmt.Sprintf("Error: %s", msg.Data.(string))
	case Info:
		return fmt.Sprintf("Info: %s", msg.Data.(string))
	case Status:
		return fmt.Sprintf("Status: %+v", msg.Data.(*StatusData))
	case CurrentSong:
		return fmt.Sprintf("CurrentSong: %+v", msg.Data.(*SongData))
	case Playlist:
		return fmt.Sprintf("Playlist: %+v", msg.Data.(*PlaylistData))
	case SearchResult:
		return fmt.Sprintf("SearchResult: %+v", msg.Data.(*SearchResultData))
	}
	return fmt.Sprintf("Unknown type: %d", msg.Type)
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
	Playlist []SongData
}

// NewSearchResult from mpd.Attrs
func NewSearchResult(attrArr *[]mpd.Attrs) *Message {
	event := &SearchResultData{}
	if attrArr == nil {
		return NewMessage(SearchResult, event)
	}
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
	return NewMessage(SearchResult, event)
}

// SearchResult payload of an event
func (msg *Message) SearchResult() *SearchResultData {
	if msg.Type == SearchResult { // FIXME: how to inform the develeoper?
		return msg.Data.(*SearchResultData)
	}
	return nil
}
