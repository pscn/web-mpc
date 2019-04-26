package msg

import (
	"encoding/json"
	"fmt"
)

// MessageType to identify the type of message
type MessageType string

// Message from the clients EventLoop
type Message struct {
	Type MessageType `json:"type"`
	Data interface{} `json:"data"`
}

// New creates a new Message with type and data
func New(msgType MessageType, data interface{}) *Message {
	return &Message{msgType, data}
}

func (msg *Message) String() string {
	return fmt.Sprintf("%s: %+v", msg.Type, msg.Data)
}

// JSON the message
func (msg *Message) JSON() (data []byte, err error) {
	data, err = json.Marshal(msg)
	if err != nil {
		err = fmt.Errorf("marshal: %v", err)
	}
	return
}

func paginate(length, page, perPage int) (currentPage int, lastPage int, start int, end int) {
	if length < perPage {
		currentPage = 1
		lastPage = 1
		start = 1
		end = length
	} else {
		lastPage = length / perPage
		currentPage = page
		if currentPage*perPage > length {
			currentPage = lastPage
		}
		start = (currentPage - 1) * perPage
		end = start + perPage
		if end > length {
			end = length
		}
	}
	return
}

// eof
