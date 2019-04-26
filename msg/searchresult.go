package msg

import (
	"github.com/fhs/gompd/mpd"
)

const SearchResult MessageType = "searchResult"

// SearchResultData converted from *mpd.attrs
type SearchResultData struct {
	SearchResult []QueuedSongData `json:"searchResult"`
	Page         int              `json:"page"`
	LastPage     int              `json:"lastPage"`
}

// SearchResultMsg from mpd.Attrs
func SearchResultMsg(attrArr *[]mpd.Attrs, page int, perPage int) *Message {
	event := &SearchResultData{}
	if attrArr == nil {
		return NewMessage(SearchResult, event)
	}
	var start, end int
	event.Page, event.LastPage, start, end = paginate(len(*attrArr), page, perPage)
	//fmt.Printf("page: %d\n", queuePage)
	//fmt.Printf("size: %d; start: %d; end: %d\n", queueLength, start, end)
	iattrArr := (*attrArr)[start:end]
	event.SearchResult = make([]QueuedSongData, len(iattrArr))
	for i, attrs := range iattrArr {
		event.SearchResult[i] = *mpd2QueuedSongData(&attrs)
	}
	return NewMessage(SearchResult, event)
}

// SearchResult payload of an event
func (msg *Message) SearchResult() *SearchResultData {
	if msg.Type == SearchResult { // FIXME: how to inform the develeoper?
		return msg.Data.(*SearchResultData)
	}
	return nil
}

// eof
