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
	Track       int    `json:"track"`
	Disc        int    `json:"disc"`
}

func mpd2SongData(attrs *mpd.Attrs) *SongData {
	// for some requests MPD returns it camelcase, for others all lowercase
	if (*attrs)["Artist"] != "" { // camelcase
		if (*attrs)["AlbumArtist"] == "" {
			(*attrs)["AlbumArtist"] = (*attrs)["Artist"]
		}
		if (*attrs)["Duration"] == "" {
			(*attrs)["Duration"] = (*attrs)["Time"]
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
			Track:       conv.ToInt((*attrs)["Track"]),
			Disc:        conv.ToInt((*attrs)["Disc"]),
		}
	}
	// lowercase

	if (*attrs)["albumartist"] == "" {
		(*attrs)["albumartist"] = (*attrs)["artist"]
	}
	if (*attrs)["duration"] == "" {
		(*attrs)["duration"] = (*attrs)["time"]
	}
	// fmt.Printf("attrs: %+v\n", attrs)
	return &SongData{
		Album:       (*attrs)["album"],
		AlbumArtist: (*attrs)["albumartist"],
		Artist:      (*attrs)["artist"],
		Duration:    conv.ToInt((*attrs)["duration"]),
		File:        (*attrs)["file"],
		Genre:       (*attrs)["genre"],
		Released:    (*attrs)["date"],
		Title:       (*attrs)["title"],
		Track:       conv.ToInt((*attrs)["track"]),
		Disc:        conv.ToInt((*attrs)["disc"]),
	}
}

// eof
