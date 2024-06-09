package controller

import (
	"dependency-injection/svc"
	"fmt"
)

type EventCtrl struct {
	svc.Greeter
}

// NewEvent return new event
func NewEvent(greeter svc.Greeter) EventCtrl {
	return EventCtrl{
		Greeter: greeter,
	}
}

// PushMessage out
func (even EventCtrl) PushMessage() {
	fmt.Println("Push this message. ", even.Greeter.Message)
}
