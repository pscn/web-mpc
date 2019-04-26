package msg

import (
	"fmt"
	"log"
	"sort"

	"github.com/fhs/gompd/mpd"
)

const PlaylistType MessageType = "playlistList"

// playlistEntry name of the playlist
type playlistEntry struct {
	Playlist string `json:"playlist"`
}

type playlistEntries []playlistEntry

// playlists holding playlists
type playlists struct {
	PlaylistList playlistEntries `json:"playlistList"`
	Page         int             `json:"page"`
	LastPage     int             `json:"lastPage"`
}

func (p playlistEntries) Len() int      { return len(p) }
func (p playlistEntries) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p playlistEntries) Less(i, j int) bool {
	return p[i].Playlist < p[j].Playlist
}

// PlaylistListMsg from mpd.Attrs
func PlaylistListMsg(attrArr *[]mpd.Attrs, page int, perPage int) *Message {
	event := &playlists{}
	if attrArr == nil {
		return NewMessage(PlaylistType, event)
	}
	cnt := 0
	for _, attrs := range *attrArr {
		fmt.Printf("playlist: %+v\n", attrs)
		if attrs["playlist"] != "" {
			cnt++
		}
	}
	event.PlaylistList = make([]playlistEntry, cnt)
	i := 0
	for _, attrs := range *attrArr {
		log.Printf("%+v\n", attrs)
		if attrs["playlist"] != "" {
			event.PlaylistList[i] = playlistEntry{
				Playlist: attrs["playlist"],
			}
			i++
		}
	}
	sort.Sort(playlistEntries(event.PlaylistList))

	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(event.PlaylistList), page, perPage)
	fmt.Printf("page: %d\n", event.Page)
	fmt.Printf("size: %d; start: %d; end: %d\n", len(event.PlaylistList), start, end)
	event.PlaylistList = event.PlaylistList[start:end]
	return NewMessage(PlaylistType, event)
}

// PlaylistList payload of an event
func (msg *Message) PlaylistList() *playlists {
	if msg.Type == PlaylistType { // FIXME: how to inform the develeoper?
		return msg.Data.(*playlists)
	}
	return nil
}

// eof
