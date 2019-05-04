package msg

import (
	"github.com/fhs/gompd/mpd"
)

// MaxSearchResults to return when searching
const MaxSearchResults = 50

// TypeSearchResult message of type SearchResult
const TypeSearchResult MessageType = "search"

// searchResult converted from *mpd.attrs
type searchResult struct {
	Songs    songs `json:"songs"`
	Page     int   `json:"page"`
	LastPage int   `json:"lastPage"`
}

// SearchResult from mpd.Attrs
func SearchResult(attrArr *[]mpd.Attrs, queueArr *[]mpd.Attrs, page int, perPage int) *Message {
	event := &searchResult{}
	if attrArr == nil {
		return New(TypeSearchResult, event)
	}
	// read queue to provide for IsQueued
	var queue map[string]bool
	if len(*queueArr) > 0 {
		queue = make(map[string]bool, len(*queueArr))
		for _, attrs := range *queueArr {
			queue[attrs["file"]] = true
		}
	}
	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(*attrArr), page, perPage)
	//fmt.Printf("page: %d\n", queuePage)
	//fmt.Printf("size: %d; start: %d; end: %d\n", queueLength, start, end)
	iattrArr := (*attrArr)[start:end]
	event.Songs = make(songs, len(iattrArr))
	for i, attrs := range iattrArr {
		event.Songs[i] = *mpd2Song(&attrs)
		if queue != nil {
			if _, ok := queue[event.Songs[i].File]; ok {
				event.Songs[i].IsQueued = true
			}
		}
	}
	return New(TypeSearchResult, event)
}

// eof
