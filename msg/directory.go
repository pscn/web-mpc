package msg

import (
	"github.com/fhs/gompd/mpd"
)

const DirectoryListType MessageType = "directoryList"

// directoryEntry an entry inside of a directory
// Type can be file or directory
// FIXME: naming?  file should be song?
type directoryEntry struct {
	Type      string `json:"type"`
	Directory string `json:"directory"`
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
}

func mpd2Directory(attrs *mpd.Attrs) *directoryEntry {
	if (*attrs)["directory"] != "" {
		return &directoryEntry{
			Type:      "directory",
			Directory: (*attrs)["directory"],
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
func DirectoryList(current string, previous string, hasPrevious bool, attrArr *[]mpd.Attrs) *Message {
	event := &directory{
		Name:      current,
		Parent:    previous,
		HasParent: hasPrevious,
	}
	if attrArr == nil {
		return NewMessage(DirectoryListType, event)
	}
	for _, attrs := range *attrArr {
		entry := mpd2Directory(&attrs)
		if entry != nil {
			event.Entries = append(event.Entries, *entry)
		}
	}
	return NewMessage(DirectoryListType, event)
}

// eof
