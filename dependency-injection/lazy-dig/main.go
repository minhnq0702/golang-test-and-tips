package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.uber.org/dig"
)

// RequestContext struct
type RequestContext struct {
	RequestID int
}

// RequestHandler struct
type RequestHandler struct {
	requestContextProvider func() *RequestContext
}

// NewRequestContext make new request
func NewRequestContext() *RequestContext {
	return &RequestContext{
		RequestID: rand.Intn(1000),
	}
}

// NewRequesthandler make new handler
func NewRequesthandler(rcp func() *RequestContext) *RequestHandler {
	return &RequestHandler{
		requestContextProvider: rcp,
	}
}

// HandleRequest handles request
func (handler *RequestHandler) HandleRequest() {
	request := handler.requestContextProvider()
	fmt.Println("handling request", request.RequestID)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := dig.New()

	// Wrap function make request by another function, so request instances will be provided multiple
	err := c.Provide(func() func() *RequestContext {
		return NewRequestContext
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(NewRequesthandler)
	if err != nil {
		panic(err)
	}

	err = c.Invoke(func(handler *RequestHandler) {
		handler.HandleRequest()
	})
	if err != nil {
		panic(err)
	}

	err = c.Invoke(func(handler *RequestHandler) {
		handler.HandleRequest()
	})
	if err != nil {
		panic(err)
	}

	graph, err := os.Create("./lazy-dig/dig.gv")
	if err != nil {
		panic(err)
	}
	err = dig.Visualize(c, graph)
	if err != nil {
		panic(err)
	}

}
