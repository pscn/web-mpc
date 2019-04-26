package msg

// Error Message
const Error MessageType = "error"

// NewError creates a new Message including an error
func NewError(str string) *Message {
	return NewMessage(Error, str)
}

// eof
