package msg

import (
	"fmt"
	"sort"

	"github.com/fhs/gompd/mpd"
)

// DirectoryListType message
const DirectoryListType MessageType = "directoryList"

// directoryEntry an entry inside of a directory
// Type can be file or directory
// FIXME: naming?  file should be song?
type directoryEntry struct {
	Type string `json:"type"`
	Name string `json:"directory"`
	SongData
}

type directoryEntries []directoryEntry

// directory holds the contents of the directory Name
// Parent to give the UI something to link back to
// HasParent to distinguish between root & first level directory, both having
// a Parent == ""
type directory struct {
	Name      string           `json:"name"`
	Parent    string           `json:"parent"`
	HasParent bool             `json:"hasParent"`
	Entries   directoryEntries `json:"directoryList"`
	Page      int              `json:"page"`
	LastPage  int              `json:"lastPage"`
}

func (p directoryEntries) Len() int      { return len(p) }
func (p directoryEntries) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p directoryEntries) Less(i, j int) bool {
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

func mpd2DirectoryEntry(attrs *mpd.Attrs) *directoryEntry {
	if (*attrs)["directory"] != "" {
		return &directoryEntry{
			Type: "directory",
			Name: (*attrs)["directory"],
		}
	} else if (*attrs)["file"] != "" {
		return &directoryEntry{
			Type:     "file",
			SongData: *mpd2SongData(attrs),
		}
	}
	return nil
}

// DirectoryList from mpd.Attrs
func DirectoryList(current string, previous string, hasPrevious bool, attrsList *[]mpd.Attrs, page int, perPage int) *Message {
	event := &directory{
		Name:      current,
		Parent:    previous,
		HasParent: hasPrevious,
	}
	if attrsList == nil {
		return NewMessage(DirectoryListType, event)
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

	return NewMessage(DirectoryListType, event)
}

// eof
