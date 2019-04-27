package msg

import (
	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/conv"
)

// queuedSong enhances Song with Queue data
type queuedSong struct {
	song
	IsActive bool `json:"isActive"`
	IsNext   bool `json:"isNext"`
	Position int  `json:"position"`
	Prio     int  `json:"prio"`
}

func mpd2QueuedSong(attrs *mpd.Attrs) *queuedSong {
	if (*attrs)["Prio"] == "" {
		(*attrs)["Prio"] = "0"
	}
	// fmt.Printf("attrs: %+v\n", attrs)
	return &queuedSong{
		song:     *mpd2Song(attrs),
		IsActive: false,
		IsNext:   false,
		Position: conv.ToInt((*attrs)["Pos"]),
		Prio:     conv.ToInt((*attrs)["Prio"]),
	}
}

// queuedSongs helps in sorting []SongData FIXME: is there a better way to do this?
type queuedSongs []queuedSong

func (q queuedSongs) Len() int      { return len(q) }
func (q queuedSongs) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q queuedSongs) Less(i, j int) bool {
	// We
	if q[i].IsActive {
		return false
	}
	if q[j].IsActive {
		return true
	}
	if q[i].IsNext {
		return false
	}
	if q[j].IsNext {
		return true
	}
	if q[i].Prio == q[j].Prio {
		// this sorts recently add items first (actually it's last, but we reverse it)
		return q[i].Position < q[j].Position
	}
	// highest prio first (actually it's last, but we reverse it)
	return q[i].Prio < q[j].Prio
}

// eof
