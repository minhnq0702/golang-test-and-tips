// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "RESTAPI": Application Controllers
//
// Command:
// $ goagen
// --design=MYGO/studying/goa-design/design
// --out=$(GOPATH)/src/MYGO/studying/goa-design/infra
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// PartnerController is the controller interface for the Partner actions.
type PartnerController interface {
	goa.Muxer
	List(*ListPartnerContext) error
}

// MountPartnerController "mounts" a Partner resource controller on the given service.
func MountPartnerController(service *goa.Service, ctrl PartnerController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListPartnerContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/partners/", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Partner", "action", "List", "route", "GET /api/partners/")
}
