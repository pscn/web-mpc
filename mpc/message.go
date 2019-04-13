package mpc

import (
	"fmt"
	"log"
	"sort"

	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/conv"
)

// MessageType to identify the type of message
type MessageType string

// MessageTypes
const (
	Error         MessageType = "error"
	Info          MessageType = "info"
	Update        MessageType = "update"
	SearchResult  MessageType = "searchResult"
	DirectoryList MessageType = "directoryList"
)

// MaxSearchResults to return when searching
const MaxSearchResults = 50

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
	return fmt.Sprintf("%s: %+v", msg.Type, msg.Data)
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
	Song     int     `json:"song"`
	NextSong int     `json:"nextsong"`
}

// StatusMsg creates a new Event including the status data mapped from mpd.Attrs
func statusData(attrs *mpd.Attrs) *StatusData {
	return &StatusData{
		Duration: conv.ToFloat((*attrs)["duration"]),
		Elapsed:  conv.ToFloat((*attrs)["elapsed"]),
		Consume:  conv.ToBool((*attrs)["consume"]),
		Random:   conv.ToBool((*attrs)["random"]),
		Single:   conv.ToBool((*attrs)["single"]),
		Repeat:   conv.ToBool((*attrs)["repeat"]),
		Volume:   conv.ToInt((*attrs)["volume"]),
		State:    (*attrs)["state"],
		Song:     conv.ToInt((*attrs)["song"]),
		NextSong: conv.ToInt((*attrs)["nextsong"]),
	}
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
	Prio        int    `json:"prio"`
	PrioUp      int    `json:"prioUp"`   // the required prio to move 1 up
	PrioDown    int    `json:"prioDown"` // the required prio to move 1 down
	IsActive    bool   `json:"isActive"`
	IsNext      bool   `json:"isNext"`
	Position    int    `json:"position"`
}

// ActiveSongMsg creates a new Event including the current song data mapped from mpd.Attrs
func songData(attrs *mpd.Attrs) *SongData {
	if (*attrs)["AlbumArtist"] == "" {
		(*attrs)["AlbumArtist"] = (*attrs)["Artist"]
	}
	if (*attrs)["Prio"] == "" {
		(*attrs)["Prio"] = "0"
	}
	// fmt.Printf("attrs: %+v\n", attrs)
	return &SongData{
		Artist:      (*attrs)["Artist"],
		Album:       (*attrs)["Album"],
		AlbumArtist: (*attrs)["AlbumArtist"],
		Title:       (*attrs)["Title"],
		Duration:    conv.ToInt((*attrs)["Time"]),
		File:        (*attrs)["file"],
		Genre:       (*attrs)["Genre"],
		Released:    (*attrs)["Date"],
		Prio:        conv.ToInt((*attrs)["Prio"]),
		IsActive:    false,
		IsNext:      false,
		Position:    conv.ToInt((*attrs)["Pos"]),
	}
}

// queueData helps in sorting []SongData FIXME: is there a better way to do this?
type queueData []SongData

func (q queueData) Len() int      { return len(q) }
func (q queueData) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q queueData) Less(i, j int) bool {
	// We
	if q[i].IsActive {
		return false
	}
	if q[j].IsActive {
		return true
	}
	if q[i].IsNext {
		return false
	}
	if q[j].IsNext {
		return true
	}
	if q[i].Prio == q[j].Prio {
		return q[i].Position < q[j].Position
	}
	return q[i].Prio < q[j].Prio
}

// UpdateData contains everything (status, active song, queue) the client needs
type UpdateData struct {
	Status     StatusData `json:"status"`
	ActiveSong SongData   `json:"activeSong"`
	Queue      []SongData `json:"queue"`
}

// UpdateDataMsg creates a new Event including the current song data mapped from mpd.Attrs
func UpdateDataMsg(status *mpd.Attrs, song *mpd.Attrs, queue *[]mpd.Attrs) *Message {
	event := &UpdateData{}
	event.Status = *statusData(status)
	event.ActiveSong = *songData(song)
	event.Queue = make([]SongData, len(*queue))
	for i, attrs := range *queue {
		event.Queue[i] = *songData(&attrs)
	}
	if event.Status.Song != -1 {
		event.Queue[event.Status.Song].IsActive = true
	}
	// FIXME: can Song and NextSong be the same?
	if event.Status.NextSong != -1 {
		event.Queue[event.Status.NextSong].IsNext = true
	}
	// order by song → nextsong → prio
	sort.Sort(sort.Reverse(queueData(event.Queue)))
	return NewMessage(Update, event)
}

// SearchResultData converted from *mpd.attrs
type SearchResultData struct {
	SearchResult []SongData `json:"searchResult"`
	Truncated    bool       `json:"truncated"` // have we omited some results?
	MaxResults   int        `json:"maxResults"`
}

// SearchResultMsg from mpd.Attrs
func SearchResultMsg(attrArr *[]mpd.Attrs) *Message {
	event := &SearchResultData{
		Truncated:  false,
		MaxResults: MaxSearchResults,
	}
	if attrArr == nil {
		return NewMessage(SearchResult, event)
	}
	iattrArr := *attrArr
	if len(iattrArr) > 50 {
		event.Truncated = true
		iattrArr = iattrArr[:50]
	}
	event.SearchResult = make([]SongData, len(iattrArr))
	for i, attrs := range iattrArr {
		event.SearchResult[i] = *songData(&attrs)
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

// DirectoryListEntry an entry for a directory list
type DirectoryListEntry struct {
	Type        string `json:"type"`
	Directory   string `json:"directory"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	AlbumArtist string `json:"album_artist"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	File        string `json:"file"`
	Genre       string `json:"genre"`
	Released    string `json:"released"`
}

// DirectoryListData helper to remember the parent of the current directory
// to give the UI something to link back to
type DirectoryListData struct {
	Parent        string               `json:"parent"`
	HasParent     bool                 `json:"hasParent"`
	DirectoryList []DirectoryListEntry `json:"directoryList"`
}

// DirectoryListMsg from mpd.Attrs
func DirectoryListMsg(previousDirectory string, hasPreviousDirectory bool, attrArr *[]mpd.Attrs) *Message {
	event := &DirectoryListData{
		Parent:    previousDirectory,
		HasParent: hasPreviousDirectory,
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
				Duration:    conv.ToInt(attrs["time"]),
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
