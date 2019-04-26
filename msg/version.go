package msg

const Version MessageType = "version"

// VersionMsg creates a new Message including the version
func VersionMsg(str string) *Message {
	return NewMessage(Version, str)
}

// eof
