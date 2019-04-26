package msg

import (
	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/conv"
)

// SongData converted from *mpd.attrs
type SongData struct {
	Album       string `json:"album"`
	AlbumArtist string `json:"album_artist"`
	Artist      string `json:"artist"`
	Duration    int    `json:"duration"`
	File        string `json:"file"`
	Genre       string `json:"genre"`
	Released    string `json:"released"`
	Title       string `json:"title"`
}

func mpd2SongData(attrs *mpd.Attrs) *SongData {
	if (*attrs)["AlbumArtist"] == "" {
		(*attrs)["AlbumArtist"] = (*attrs)["Artist"]
	}
	// fmt.Printf("attrs: %+v\n", attrs)
	return &SongData{
		Album:       (*attrs)["Album"],
		AlbumArtist: (*attrs)["AlbumArtist"],
		Artist:      (*attrs)["Artist"],
		Duration:    conv.ToInt((*attrs)["Time"]),
		File:        (*attrs)["file"],
		Genre:       (*attrs)["Genre"],
		Released:    (*attrs)["Date"],
		Title:       (*attrs)["Title"],
	}
}

// eof
