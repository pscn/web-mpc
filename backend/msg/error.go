package msg

// TypeError Message
const TypeError MessageType = "error"

// Error creates a new Message including an error
func Error(str string) *Message {
	return New(TypeError, str)
}

// eof
