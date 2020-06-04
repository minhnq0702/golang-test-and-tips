package store

type Message string

// NewMessage return default message
func NewMessage() Message {
	return Message("This is message")
}
