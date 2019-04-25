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
	Version       MessageType = "version"
	Update        MessageType = "update"
	SearchResult  MessageType = "searchResult"
	DirectoryList MessageType = "directoryList"
	PlaylistList  MessageType = "playlistList"
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

// VersionMsg creates a new Event including an error
func VersionMsg(str string) *Message {
	return NewMessage(Version, str)
}

// StatusData represents the data we get as MPC status
type StatusData struct {
	Duration int    `json:"duration"`
	Elapsed  int    `json:"elapsed"`
	State    string `json:"state"`
	Volume   int    `json:"volume"`
	Consume  bool   `json:"consume"`
	Random   bool   `json:"random"`
	Single   bool   `json:"single"`
	Repeat   bool   `json:"repeat"`
	Song     int    `json:"song"`
	NextSong int    `json:"nextsong"`
}

// StatusMsg creates a new Event including the status data mapped from mpd.Attrs
func statusData(attrs *mpd.Attrs) *StatusData {
	return &StatusData{
		Duration: conv.ToInt((*attrs)["duration"]),
		Elapsed:  conv.ToInt((*attrs)["elapsed"]),
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
		// this sorts recently add items first (actually it's last, but we reverse it)
		return q[i].Position < q[j].Position
	}
	// highest prio first (actually it's last, but we reverse it)
	return q[i].Prio < q[j].Prio
}

// UpdateData contains everything (status, active song, queue) the client needs
type UpdateData struct {
	Status     StatusData `json:"status"`
	ActiveSong SongData   `json:"activeSong"`
	Queue      []SongData `json:"queue"`
	Page       int        `json:"page"`
	LastPage   int        `json:"lastPage"`
}

// UpdateDataMsg creates a new Event including the current song data mapped from mpd.Attrs
func UpdateDataMsg(status *mpd.Attrs, song *mpd.Attrs, queue *[]mpd.Attrs,
	page int, perPage int) *Message {
	event := &UpdateData{}
	event.Status = *statusData(status)
	event.ActiveSong = *songData(song)
	event.Queue = make([]SongData, len(*queue))
	for i, attrs := range *queue {
		event.Queue[i] = *songData(&attrs)
	}
	if event.Status.Song != -1 {
		event.Queue[event.Status.Song].IsActive = true
		if event.Status.Duration == -1 { // Protocol Version < 0.20.0
			event.Status.Duration = event.Queue[event.Status.Song].Duration
		}
	}
	// FIXME: can Song and NextSong be the same?
	if event.Status.NextSong != -1 {
		event.Queue[event.Status.NextSong].IsNext = true
	}
	// order by song → nextsong → prio
	sort.Sort(sort.Reverse(queueData(event.Queue)))

	// pagination
	fmt.Printf("page: %d\n", page)
	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(event.Queue), page, perPage)
	//fmt.Printf("page: %d\n", queuePage)
	//fmt.Printf("size: %d; start: %d; end: %d\n", queueLength, start, end)
	event.Queue = event.Queue[start:end]

	return NewMessage(Update, event)
}

// SearchResultData converted from *mpd.attrs
type SearchResultData struct {
	SearchResult []SongData `json:"searchResult"`
	Page         int        `json:"page"`
	LastPage     int        `json:"lastPage"`
}

// SearchResultMsg from mpd.Attrs
func SearchResultMsg(attrArr *[]mpd.Attrs, page int, perPage int) *Message {
	event := &SearchResultData{}
	if attrArr == nil {
		return NewMessage(SearchResult, event)
	}
	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(*attrArr), page, perPage)
	//fmt.Printf("page: %d\n", queuePage)
	//fmt.Printf("size: %d; start: %d; end: %d\n", queueLength, start, end)
	iattrArr := (*attrArr)[start:end]
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

// PlaylistData name of the playlist
type PlaylistData struct {
	Playlist string `json:"playlist"`
}

// PlaylistListData holding playlists
type PlaylistListData struct {
	PlaylistList []PlaylistData `json:"playlistList"`
	Page         int            `json:"page"`
	LastPage     int            `json:"lastPage"`
}

type playlistData []PlaylistData

func (p playlistData) Len() int      { return len(p) }
func (p playlistData) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p playlistData) Less(i, j int) bool {
	return p[i].Playlist < p[j].Playlist
}

func paginate(length, page, perPage int) (currentPage int, lastPage int, start int, end int) {
	if length < perPage {
		currentPage = 1
		lastPage = 1
		start = 1
		end = length
	} else {
		lastPage = length / perPage
		currentPage = page
		if currentPage*perPage > length {
			currentPage = lastPage
		}
		start = (currentPage - 1) * perPage
		end = start + perPage
		if end > length {
			end = length
		}
	}
	return
}

// PlaylistListMsg from mpd.Attrs
func PlaylistListMsg(attrArr *[]mpd.Attrs, page int, perPage int) *Message {
	event := &PlaylistListData{}
	if attrArr == nil {
		return NewMessage(PlaylistList, event)
	}
	cnt := 0
	for _, attrs := range *attrArr {
		fmt.Printf("playlist: %+v\n", attrs)
		if attrs["playlist"] != "" {
			cnt++
		}
	}
	event.PlaylistList = make([]PlaylistData, cnt)
	i := 0
	for _, attrs := range *attrArr {
		log.Printf("%+v\n", attrs)
		if attrs["playlist"] != "" {
			event.PlaylistList[i] = PlaylistData{
				Playlist: attrs["playlist"],
			}
			i++
		}
	}
	sort.Sort(playlistData(event.PlaylistList))

	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(event.PlaylistList), page, perPage)
	fmt.Printf("page: %d\n", event.Page)
	fmt.Printf("size: %d; start: %d; end: %d\n", len(event.PlaylistList), start, end)
	event.PlaylistList = event.PlaylistList[start:end]
	return NewMessage(PlaylistList, event)
}

// PlaylistList payload of an event
func (msg *Message) PlaylistList() *PlaylistListData {
	if msg.Type == PlaylistList { // FIXME: how to inform the develeoper?
		return msg.Data.(*PlaylistListData)
	}
	return nil
}

// eof
