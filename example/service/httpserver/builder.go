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
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type errorCreator func(err error) error

// New ...
func New(router *fasthttprouter.Router, svc service, decodeJSONErrorCreator errorCreator, encodeJSONErrorCreator errorCreator, encodeQueryTypeIntErrorCreator errorCreator, errorProcessor errorProcessor) {

	getWarehousesTransport := NewGetWarehousesTransport(encodeJSONErrorCreator)
	router.Handle(httpMethodGetWarehouses, uriPathGetWarehouses, NewGetWarehousesSwaggerInfo(getWarehousesTransport, svc, errorProcessor))

	getDetailsTransport := NewGetDetailsTransport(encodeJSONErrorCreator, encodeQueryTypeIntErrorCreator)
	router.Handle(httpMethodGetDetails, uriPathGetDetails, NewGetDetailsSwaggerInfo(getDetailsTransport, svc, errorProcessor))

	getDetailsEmbedStructTransport := NewGetDetailsEmbedStructTransport(encodeJSONErrorCreator)
	router.Handle(httpMethodGetDetailsEmbedStruct, uriPathGetDetailsEmbedStruct, NewGetDetailsEmbedStructSwaggerInfo(getDetailsEmbedStructTransport, svc, errorProcessor))

	getDetailsListEmbedStructTransport := NewGetDetailsListEmbedStructTransport(encodeJSONErrorCreator)
	router.Handle(httpMethodGetDetailsListEmbedStruct, uriPathGetDetailsListEmbedStruct, NewGetDetailsListEmbedStructSwaggerInfo(getDetailsListEmbedStructTransport, svc, errorProcessor))

	putDetailsTransport := NewPutDetailsTransport(decodeJSONErrorCreator, encodeJSONErrorCreator)
	router.Handle(httpMethodPutDetails, uriPathPutDetails, NewPutDetailsSwaggerInfo(putDetailsTransport, svc, errorProcessor))

	getSomeElseDataUtf8Transport := NewGetSomeElseDataUtf8Transport(encodeJSONErrorCreator)
	router.Handle(httpMethodGetSomeElseDataUtf8, uriPathGetSomeElseDataUtf8, NewGetSomeElseDataUtf8SwaggerInfo(getSomeElseDataUtf8Transport, svc, errorProcessor))

	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
}
