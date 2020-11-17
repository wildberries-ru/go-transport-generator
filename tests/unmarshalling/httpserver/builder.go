// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"
	"net/http/pprof"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const (
	httpMethodTestEasyJson = "GET"
	uriPathTestEasyJson    = "/api/testeasyjson"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type errorCreator func(err error) error

// New ...
func New(router *fasthttprouter.Router, svc service, decodeJSONErrorCreator errorCreator, encodeJSONErrorCreator errorCreator, decodeTypeIntErrorCreator errorCreator, errorProcessor errorProcessor) {

	testEasyJsonTransport := NewTestEasyJsonTransport(

		encodeJSONErrorCreator,
		decodeTypeIntErrorCreator,
	)
	router.Handle(httpMethodTestEasyJson, uriPathTestEasyJson, NewTestEasyJson(testEasyJsonTransport, svc, errorProcessor))

	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
}
