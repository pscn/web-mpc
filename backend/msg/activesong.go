package msg

import (
	"github.com/fhs/gompd/mpd"
)

// queuedSong enhances Song with Queue data
type activeSong struct {
	song
	Elapsed int  `json:"elapsed"`
	Playing bool `json:"playing"`
}

func mpd2ActiveSong(attrs *mpd.Attrs) *activeSong {
	if (*attrs)["Prio"] == "" {
		(*attrs)["Prio"] = "0"
	}
	// fmt.Printf("attrs: %+v\n", attrs)
	return &activeSong{
		song:    *mpd2Song(attrs),
		Elapsed: -1, // needs to be filled in from the status
	}
}

// eof
