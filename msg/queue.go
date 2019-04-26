package msg

// QueuedSongsData helps in sorting []SongData FIXME: is there a better way to do this?
type QueuedSongsData []QueuedSongData

func (q QueuedSongsData) Len() int      { return len(q) }
func (q QueuedSongsData) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q QueuedSongsData) Less(i, j int) bool {
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
