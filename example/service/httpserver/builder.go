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
	httpMethodGetWarehouses             = "GET"
	uriPathGetWarehouses                = "/api/v1/getWarehouses"
	httpMethodGetDetails                = "GET"
	uriPathGetDetails                   = "/api/v1/namespaces/:namespace/details/:detail"
	httpMethodGetDetailsEmbedStruct     = "GET"
	uriPathGetDetailsEmbedStruct        = "/api/v1/namespaces/:namespace/details-embed/:detail"
	httpMethodGetDetailsListEmbedStruct = "GET"
	uriPathGetDetailsListEmbedStruct    = "/api/v1/namespaces/:namespace/details-embed-array/:detail"
	httpMethodPutDetails                = "PUT"
	uriPathPutDetails                   = "/api/v1/namespaces/:namespace/details/:detail"
	httpMethodGetSomeElseDataUtf8       = "GET"
	uriPathGetSomeElseDataUtf8          = "/api/v1/someelsedata"
	httpMethodGetFile                   = "GET"
	uriPathGetFile                      = "/api/v1/file"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type errorCreator func(err error) error

// New ...
func New(router *fasthttprouter.Router, svc service, decodeJSONErrorCreator errorCreator, encodeJSONErrorCreator errorCreator, decodeTypeIntErrorCreator errorCreator, errorProcessor errorProcessor) {

	getWarehousesTransport := NewGetWarehousesTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodGetWarehouses, uriPathGetWarehouses, NewGetWarehouses(getWarehousesTransport, svc, errorProcessor))

	getDetailsTransport := NewGetDetailsTransport(

		encodeJSONErrorCreator,
		decodeTypeIntErrorCreator,
	)
	router.Handle(httpMethodGetDetails, uriPathGetDetails, NewGetDetails(getDetailsTransport, svc, errorProcessor))

	getDetailsEmbedStructTransport := NewGetDetailsEmbedStructTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodGetDetailsEmbedStruct, uriPathGetDetailsEmbedStruct, NewGetDetailsEmbedStruct(getDetailsEmbedStructTransport, svc, errorProcessor))

	getDetailsListEmbedStructTransport := NewGetDetailsListEmbedStructTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodGetDetailsListEmbedStruct, uriPathGetDetailsListEmbedStruct, NewGetDetailsListEmbedStruct(getDetailsListEmbedStructTransport, svc, errorProcessor))

	putDetailsTransport := NewPutDetailsTransport(
		decodeJSONErrorCreator,
		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodPutDetails, uriPathPutDetails, NewPutDetails(putDetailsTransport, svc, errorProcessor))

	getSomeElseDataUtf8Transport := NewGetSomeElseDataUtf8Transport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodGetSomeElseDataUtf8, uriPathGetSomeElseDataUtf8, NewGetSomeElseDataUtf8(getSomeElseDataUtf8Transport, svc, errorProcessor))

	getFileTransport := NewGetFileTransport()
	router.Handle(httpMethodGetFile, uriPathGetFile, NewGetFile(getFileTransport, svc, errorProcessor))

}

// NewPprofWrapper wraps router in pprof
func NewPprofWrapper(router *fasthttprouter.Router) {
	router.Handle("GET", "/debug/pprof", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
	router.Handle("GET", "/debug/pprof/cmdline", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline))
	router.Handle("GET", "/debug/pprof/symbol", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol))
	router.Handle("GET", "/debug/pprof/trace", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace))
	router.Handle("GET", "/debug/pprof/goroutine", fasthttpadaptor.NewFastHTTPHandler(pprof.Handler("goroutine")))
	router.Handle("GET", "/debug/pprof/heap", fasthttpadaptor.NewFastHTTPHandler(pprof.Handler("heap")))
	router.Handle("GET", "/debug/pprof/threadcreate", fasthttpadaptor.NewFastHTTPHandler(pprof.Handler("threadcreate")))
	router.Handle("GET", "/debug/pprof/block", fasthttpadaptor.NewFastHTTPHandler(pprof.Handler("block")))
}
