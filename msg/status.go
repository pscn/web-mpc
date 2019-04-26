package msg

import (
	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/conv"
)

// Status represents the data we get as MPC status
type Status struct {
	Consume  bool   `json:"consume"`
	Duration int    `json:"duration"`
	Elapsed  int    `json:"elapsed"`
	NextSong int    `json:"nextsong"`
	Random   bool   `json:"random"`
	Repeat   bool   `json:"repeat"`
	Single   bool   `json:"single"`
	Song     int    `json:"song"`
	State    string `json:"state"`
	Volume   int    `json:"volume"`
}

func mpd2Status(attrs *mpd.Attrs) *Status {
	return &Status{
		Consume:  conv.ToBool((*attrs)["consume"]),
		Duration: conv.ToInt((*attrs)["duration"]),
		Elapsed:  conv.ToInt((*attrs)["elapsed"]),
		NextSong: conv.ToInt((*attrs)["nextsong"]),
		Random:   conv.ToBool((*attrs)["random"]),
		Repeat:   conv.ToBool((*attrs)["repeat"]),
		Single:   conv.ToBool((*attrs)["single"]),
		Song:     conv.ToInt((*attrs)["song"]),
		State:    (*attrs)["state"],
		Volume:   conv.ToInt((*attrs)["volume"]),
	}
}

// eof
