package mpc

import (
	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/helpers"
)

// Type to identify the type of event
type Type uint

const (
	// Error the Event contains an error as payload
	Error Type = 0
	// Info the Event contains a string as payload
	Info Type = 1
	// Status the Event contains a status as payload
	Status Type = 2
	// CurrentSong the Event contains the currently active song
	CurrentSong Type = 3
	// Playlist the Event contains the currently active song
	Playlist Type = 4
	// SearchResult the result of a search
	SearchResult Type = 5
)

// Message from the clients EventLoop
type Message struct {
	Type Type        `json:"type"`
	Data interface{} `json:"data"`
}

// NewMessage creates a new event with type and data
func NewMessage(msgType Type, data interface{}) *Message {
	return &Message{msgType, data}
}

// NewErrorEvent creates a new Event including an error
func NewErrorEvent(err error) *Message {
	return NewMessage(Error, err)
}

func (msg *Message) Error() error {
	if msg.Type == Error { // FIXME: how to inform the develeoper?
		return msg.Data.(error)
	}
	return nil
}

// NewStringEvent creates a new Event including an error
func NewStringEvent(str string) *Message {
	return NewMessage(Info, str)
}

func (msg *Message) String() string {
	if msg.Type == Info { // FIXME: how to inform the develeoper?
		return msg.Data.(string)
	}
	return ""
}

// StatusData represents the data we get as MPC status
type StatusData struct {
	Duration float64 `json:"duration"`
	Elapsed  float64 `json:"elapsed"`
	State    string  `json:"state"`
	Volume   int64   `json:"volume"`
	Consume  bool    `json:"consume"`
	Random   bool    `json:"random"`
	Single   bool    `json:"single"`
	Repeat   bool    `json:"repeat"`
}

// NewStatus creates a new Event including the status data mapped from mpd.Attrs
func NewStatus(attrs *mpd.Attrs) *Message {
	return NewMessage(Status, &StatusData{
		Duration: helpers.ToFloat64((*attrs)["duration"]),
		Elapsed:  helpers.ToFloat64((*attrs)["elapsed"]),
		Consume:  helpers.ToBool((*attrs)["consume"]),
		Random:   helpers.ToBool((*attrs)["random"]),
		Single:   helpers.ToBool((*attrs)["single"]),
		Repeat:   helpers.ToBool((*attrs)["repeat"]),
		Volume:   helpers.ToInt64((*attrs)["volume"]),
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
	Duration    int64  `json:"duration"`
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
		Duration:    helpers.ToInt64((*attrs)["Time"]),
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
			Duration:    helpers.ToInt64(attrs["Time"]),
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
			Duration:    helpers.ToInt64(attrs["Time"]),
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
