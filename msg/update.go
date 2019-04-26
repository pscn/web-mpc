package msg

import (
	"fmt"
	"sort"

	"github.com/fhs/gompd/mpd"
)

// Update message containing status, queue and active song FIXME: naming
const Update MessageType = "update"

// UpdateData contains everything (status, active song, queue) the client needs
type UpdateData struct {
	Status     Status          `json:"status"`
	ActiveSong QueuedSongData  `json:"activeSong"`
	Queue      QueuedSongsData `json:"queue"`
	Page       int             `json:"page"`
	LastPage   int             `json:"lastPage"`
}

// NewUpdate creates a new Event including the current song data mapped from mpd.Attrs
func NewUpdate(status *mpd.Attrs, song *mpd.Attrs, queue *[]mpd.Attrs,
	page int, perPage int) *Message {
	event := &UpdateData{}
	event.Status = *mpd2Status(status)
	event.ActiveSong = *mpd2QueuedSongData(song)
	event.Queue = make([]QueuedSongData, len(*queue))
	for i, attrs := range *queue {
		event.Queue[i] = *mpd2QueuedSongData(&attrs)
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
	sort.Sort(sort.Reverse(QueuedSongsData(event.Queue)))

	// pagination
	fmt.Printf("page: %d\n", page)
	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(event.Queue), page, perPage)
	//fmt.Printf("page: %d\n", queuePage)
	//fmt.Printf("size: %d; start: %d; end: %d\n", queueLength, start, end)
	event.Queue = event.Queue[start:end]

	return NewMessage(Update, event)
}

// eof
