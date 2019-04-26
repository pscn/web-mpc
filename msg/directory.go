package msg

import (
	"fmt"
	"sort"

	"github.com/fhs/gompd/mpd"
)

// TypeDirectoryList message
const TypeDirectoryList MessageType = "directoryList"

// dirEntry an entry inside of a directory
// Type can be file or directory
// FIXME: naming?  file should be song?
type dirEntry struct {
	Type string `json:"type"`
	Name string `json:"directory"`
	song
}

type dirEntries []dirEntry

// dir holds the contents of the dir Name
// Parent to give the UI something to link back to
// HasParent to distinguish between root & first level dir, both having
// a Parent == ""
type dir struct {
	Name      string     `json:"name"`
	Parent    string     `json:"parent"`
	HasParent bool       `json:"hasParent"`
	Entries   dirEntries `json:"directoryList"`
	Page      int        `json:"page"`
	LastPage  int        `json:"lastPage"`
}

func (p dirEntries) Len() int      { return len(p) }
func (p dirEntries) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p dirEntries) Less(i, j int) bool {
	if p[i].Type != p[j].Type {
		if p[i].Type == "file" {
			return true
		}
		return false
	}
	if p[i].Disc != p[j].Disc {
		return p[i].Disc < p[j].Disc
	}
	if p[i].Track != p[j].Track {
		return p[i].Track < p[j].Track
	}
	return p[i].Name < p[j].Name
}

func mpd2DirectoryEntry(attrs *mpd.Attrs) *dirEntry {
	if (*attrs)["directory"] != "" {
		return &dirEntry{
			Type: "directory",
			Name: (*attrs)["directory"],
		}
	} else if (*attrs)["file"] != "" {
		return &dirEntry{
			Type: "file",
			song: *mpd2Song(attrs),
		}
	}
	return nil
}

// Directory from mpd.Attrs
func Directory(cur string, prev string, hasPrev bool, attrsList *[]mpd.Attrs, page int, perPage int) *Message {
	event := &dir{
		Name:      cur,
		Parent:    prev,
		HasParent: hasPrev,
	}
	if attrsList == nil {
		return New(TypeDirectoryList, event)
	}
	for _, attrs := range *attrsList {
		entry := mpd2DirectoryEntry(&attrs)
		if entry != nil {
			event.Entries = append(event.Entries, *entry)
		}
	}

	// sort alphabeticaly
	sort.Sort(event.Entries)

	// paginate
	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(event.Entries), page, perPage)
	fmt.Printf("page: %d\n", event.Page)
	fmt.Printf("size: %d; start: %d; end: %d\n", len(event.Entries), start, end)
	event.Entries = event.Entries[start:end]

	return New(TypeDirectoryList, event)
}

// eof
