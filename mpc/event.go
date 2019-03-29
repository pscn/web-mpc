package mpc

// EventType to identify the type of event
type EventType uint

const (
	// EventError an error happened
	EventError  EventType = 0
	EventString EventType = 1
)

// Event from the clients EventLoop
type Event struct {
	Type EventType
	Data interface{}
}

// NewEvent creates a new event with type and data
func NewEvent(eventType EventType, data interface{}) *Event {
	return &Event{eventType, data}
}

// ErrorEvent creates a new Event including an error
func ErrorEvent(err error) *Event {
	return NewEvent(EventError, err)
}

// StringEvent creates a new Event including an error
func StringEvent(str string) *Event {
	return NewEvent(EventString, str)
}

func (event *Event) Error() error {
	return event.Data.(error)
}

func (event *Event) String() string {
	return event.Data.(string)
}
