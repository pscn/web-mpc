package msg

import (
	"fmt"
	"sort"

	"github.com/fhs/gompd/mpd"
)

// TypePlaylistList message of tpye PlaylistList
const TypePlaylistList MessageType = "playlistList"

// playlistListEntry name of the playlist
// FIXME: naming is so wrong
type playlistListEntry struct {
	Name string `json:"playlist"`
}

type playlistListEntries []playlistListEntry

// playlistList holding playlistList
type playlistList struct {
	Playlists playlistListEntries `json:"playlistList"`
	Page      int                 `json:"page"`
	LastPage  int                 `json:"lastPage"`
}

func (p playlistListEntries) Len() int           { return len(p) }
func (p playlistListEntries) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p playlistListEntries) Less(i, j int) bool { return p[i].Name < p[j].Name }

// FIXME: naming
func mpd2PlaylistListEntry(attrs *mpd.Attrs) *playlistListEntry {
	return &playlistListEntry{
		Name: (*attrs)["playlist"],
	}
}

// PlaylistList from mpd.Attrs
func PlaylistList(attrsList *[]mpd.Attrs, page int, perPage int) *Message {
	event := &playlistList{}
	if attrsList == nil {
		return New(TypePlaylistList, event)
	}
	event.Playlists = make(playlistListEntries, len(*attrsList))
	for i, attrs := range *attrsList {
		fmt.Printf("%+v\n", attrs)
		event.Playlists[i] = *mpd2PlaylistListEntry(&attrs)
	}

	// sort alphabeticaly
	sort.Sort(event.Playlists)

	// paginate
	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(event.Playlists), page, perPage)
	fmt.Printf("page: %d\n", event.Page)
	fmt.Printf("size: %d; start: %d; end: %d\n", len(event.Playlists), start, end)
	event.Playlists = event.Playlists[start:end]
	return New(TypePlaylistList, event)
}

// eof
