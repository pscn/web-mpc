package msg

// TypeVersion message of type Version
const TypeVersion MessageType = "version"

// Version creates a new Message including the version
func Version(str string) *Message {
	return New(TypeVersion, str)
}

// eof
