package msg

import (
	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/conv"
)

// QueuedSongData enhances Song with Queue data
type QueuedSongData struct {
	SongData
	IsActive bool `json:"isActive"`
	IsNext   bool `json:"isNext"`
	Position int  `json:"position"`
	Prio     int  `json:"prio"`
}

func mpd2QueuedSongData(attrs *mpd.Attrs) *QueuedSongData {
	if (*attrs)["Prio"] == "" {
		(*attrs)["Prio"] = "0"
	}
	// fmt.Printf("attrs: %+v\n", attrs)
	return &QueuedSongData{
		SongData: *mpd2SongData(attrs),
		IsActive: false,
		IsNext:   false,
		Position: conv.ToInt((*attrs)["Pos"]),
		Prio:     conv.ToInt((*attrs)["Prio"]),
	}
}

// eof
