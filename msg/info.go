package msg

const Info MessageType = "info"

// InfoMsg creates a new Message including an information
func InfoMsg(str string) *Message {
	return NewMessage(Info, str)
}

// eof
