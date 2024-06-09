package svc

import (
	"dependency-injection/store"
	"fmt"
)

type Greeter struct {
	store.Message
}

// NewGreeter return new greeter of message
func NewGreeter(m store.Message) Greeter {
	fmt.Println("Creating greeter!!!", m)
	return Greeter{m}
}
