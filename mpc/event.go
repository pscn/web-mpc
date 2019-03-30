package mpc

import (
	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/helpers"
)

// EventType to identify the type of event
type EventType uint

const (
	// EventTypeError the Event contains an error as payload
	EventTypeError EventType = 0
	// EventTypeString the Event contains a string as payload
	EventTypeString EventType = 1
	// EventTypeStatus the Event contains a status as payload
	EventTypeStatus EventType = 2
	// EventTypeCurrentSong the Event contains the currently active song
	EventTypeCurrentSong EventType = 3
)

// Event from the clients EventLoop
type Event struct {
	Type EventType   `json:"type"`
	Data interface{} `json:"data"`
}

// NewEvent creates a new event with type and data
func NewEvent(eventType EventType, data interface{}) *Event {
	return &Event{eventType, data}
}

// NewErrorEvent creates a new Event including an error
func NewErrorEvent(err error) *Event {
	return NewEvent(EventTypeError, err)
}

func (event *Event) Error() error {
	if event.Type == EventTypeError { // FIXME: how to inform the develeoper?
		return event.Data.(error)
	}
	return nil
}

// NewStringEvent creates a new Event including an error
func NewStringEvent(str string) *Event {
	return NewEvent(EventTypeString, str)
}

func (event *Event) String() string {
	if event.Type == EventTypeString { // FIXME: how to inform the develeoper?
		return event.Data.(string)
	}
	return ""
}

// EventDataStatus represents the data we get as MPC status
type EventDataStatus struct {
	Duration float64 `json:"duration"`
	Elapsed  float64 `json:"elapsed"`
	State    string  `json:"state"`
	Volume   int64   `json:"volume"`
	Consume  bool    `json:"consume"`
	Random   bool    `json:"random"`
	Single   bool    `json:"single"`
	Repeat   bool    `json:"repeat"`
}

// NewStatusEvent creates a new Event including the status data mapped from mpd.Attrs
func NewStatusEvent(attrs *mpd.Attrs) *Event {
	return NewEvent(EventTypeStatus, &EventDataStatus{
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
func (event *Event) Status() *EventDataStatus {
	if event.Type == EventTypeStatus { // FIXME: how to inform the develeoper?
		return event.Data.(*EventDataStatus)
	}
	return nil
}

// EventDataCurrentSong converted from *mpd.attrs
type EventDataCurrentSong struct {
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	AlbumArtist string `json:"album_artist"`
	Title       string `json:"title"`
	Length      int64  `json:"length"`
	File        string `json:"file"`
	Genre       string `json:"genre"`
	Released    string `json:"released"`
}

// NewCurrentSongEvent creates a new Event including the current song data mapped from mpd.Attrs
func NewCurrentSongEvent(attrs *mpd.Attrs) *Event {
	if (*attrs)["AlbumArtist"] == "" {
		(*attrs)["AlbumArtist"] = (*attrs)["Artist"]
	}
	return NewEvent(EventTypeCurrentSong, &EventDataCurrentSong{
		Artist:      (*attrs)["Artist"],
		Album:       (*attrs)["Album"],
		AlbumArtist: (*attrs)["AlbumArtist"],
		Title:       (*attrs)["Title"],
		Length:      helpers.ToInt64((*attrs)["Time"]),
		File:        (*attrs)["file"],
		Genre:       (*attrs)["Genre"],
		Released:    (*attrs)["Released"],
	})
}

// CurrentSong payload of an event
func (event *Event) CurrentSong() *EventDataCurrentSong {
	if event.Type == EventTypeCurrentSong { // FIXME: how to inform the develeoper?
		return event.Data.(*EventDataCurrentSong)
	}
	return nil
}
