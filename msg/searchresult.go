package msg

import (
	"github.com/fhs/gompd/mpd"
)

// MaxSearchResults to return when searching
const MaxSearchResults = 50

// TypeSearchResult message of type SearchResult
const TypeSearchResult MessageType = "searchResult"

// searchResult converted from *mpd.attrs
type searchResult struct {
	SearchResult songs `json:"searchResult"`
	Page         int   `json:"page"`
	LastPage     int   `json:"lastPage"`
}

// SearchResult from mpd.Attrs
func SearchResult(attrArr *[]mpd.Attrs, queueArr *[]mpd.Attrs, page int, perPage int) *Message {
	event := &searchResult{}
	if attrArr == nil {
		return New(TypeSearchResult, event)
	}
	// read queue to provide for IsQueued
	queue := make(map[string]bool, len(*queueArr))
	for _, attrs := range *queueArr {
		queue[attrs["file"]] = true
	}
	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(*attrArr), page, perPage)
	//fmt.Printf("page: %d\n", queuePage)
	//fmt.Printf("size: %d; start: %d; end: %d\n", queueLength, start, end)
	iattrArr := (*attrArr)[start:end]
	event.SearchResult = make(songs, len(iattrArr))
	for i, attrs := range iattrArr {
		event.SearchResult[i] = *mpd2Song(&attrs)
		if _, ok := queue[event.SearchResult[i].File]; ok {
			event.SearchResult[i].IsQueued = true
		}
	}
	return New(TypeSearchResult, event)
}

// eof
