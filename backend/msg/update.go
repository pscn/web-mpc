package msg

import (
	"fmt"
	"sort"

	"github.com/fhs/gompd/mpd"
)

// TypeUpdate message containing status, queue and active song FIXME: naming
const TypeUpdate MessageType = "update"

type queue struct {
	Songs    queuedSongs `json:"songs"`
	Page     int         `json:"page"`
	LastPage int         `json:"lastPage"`
}

// update contains everything (status, active song, queue) the client needs
type update struct {
	Status     status     `json:"status"`
	ActiveSong queuedSong `json:"activeSong"`
	Queue      queue      `json:"queue"`
}

// Update creates a new Event including the current song data mapped from mpd.Attrs
func Update(status *mpd.Attrs, song *mpd.Attrs, songs *[]mpd.Attrs,
	page int, perPage int) *Message {
	event := &update{
		Status:     *mpd2Status(status),
		ActiveSong: *mpd2QueuedSong(song),
		Queue:      queue{},
	}
	event.Queue.Songs = make([]queuedSong, len(*songs))
	if len(*songs) > 0 {
		for i, attrs := range *songs {
			event.Queue.Songs[i] = *mpd2QueuedSong(&attrs)
		}
		if event.Status.Song != -1 {
			event.Queue.Songs[event.Status.Song].IsActive = true
			if event.Status.Duration == -1 { // Protocol Version < 0.20.0
				event.Status.Duration = event.Queue.Songs[event.Status.Song].Duration
			}
		}
		// FIXME: can Song and NextSong be the same?
		if event.Status.NextSong != -1 {
			event.Queue.Songs[event.Status.NextSong].IsNext = true
		}
		// order by song → nextsong → prio
		sort.Sort(sort.Reverse(queuedSongs(event.Queue.Songs)))

		// pagination
		fmt.Printf("page: %d\n", page)
		var start, end int
		event.Queue.Page, event.Queue.LastPage, start, end = paginate(len(event.Queue.Songs), page, perPage)
		//fmt.Printf("page: %d\n", queuePage)
		//fmt.Printf("size: %d; start: %d; end: %d\n", queueLength, start, end)
		event.Queue.Songs = event.Queue.Songs[start:end]
	}
	return New(TypeUpdate, event)
}

// eof
