package msg

// TypeInfo message of type Info
const TypeInfo MessageType = "info"

// Info creates a new Message including an information
func Info(str string) *Message {
	return New(TypeInfo, str)
}

// eof
