package msg

import (
	"fmt"
	"log"
	"sort"

	"github.com/fhs/gompd/mpd"
)

const PlaylistListType MessageType = "playlistList"

// playlistListEntry name of the playlist
// FIXME: naming is so wrong
type playlistListEntry struct {
	Name string `json:"playlist"`
}

type playlistListEntries []playlistListEntry

// playlists holding playlists
type playlists struct {
	Playlists playlistListEntries `json:"playlistList"`
	Page      int                 `json:"page"`
	LastPage  int                 `json:"lastPage"`
}

func (p playlistListEntries) Len() int      { return len(p) }
func (p playlistListEntries) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p playlistListEntries) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

// FIXME: naming
func mpd2PlaylistListEntry(attrs *mpd.Attrs) *playlistListEntry {
	return &playlistListEntry{
		Name: (*attrs)["playlist"],
	}
}

// PlaylistList from mpd.Attrs
func PlaylistList(attrsList *[]mpd.Attrs, page int, perPage int) *Message {
	event := &playlists{}
	if attrsList == nil {
		return NewMessage(PlaylistListType, event)
	}
	for _, attrs := range *attrsList {
		log.Printf("%+v\n", attrs)
		event.Playlists = append(event.Playlists, *mpd2PlaylistListEntry(&attrs))
	}
	if event.Playlists == nil || event.Playlists.Len() == 0 {
		return NewMessage(PlaylistListType, event)
	}
	// sort alphabeticaly
	sort.Sort(event.Playlists)

	// paginate
	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(event.Playlists), page, perPage)
	fmt.Printf("page: %d\n", event.Page)
	fmt.Printf("size: %d; start: %d; end: %d\n", len(event.Playlists), start, end)
	event.Playlists = event.Playlists[start:end]
	return NewMessage(PlaylistListType, event)
}

// PlaylistList payload of an event
func (msg *Message) PlaylistList() *playlists {
	if msg.Type == PlaylistListType { // FIXME: how to inform the develeoper?
		return msg.Data.(*playlists)
	}
	return nil
}

// eof
