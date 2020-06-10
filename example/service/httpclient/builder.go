// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"github.com/valyala/fasthttp"
)

const (
	httpMethodGetWarehouses                = "GET"
	uriPathClientGetWarehouses             = "/api/v1/getWarehouses"
	httpMethodGetDetails                   = "GET"
	uriPathClientGetDetails                = "/api/v1/namespaces/%s/details/%s"
	httpMethodGetDetailsEmbedStruct        = "GET"
	uriPathClientGetDetailsEmbedStruct     = "/api/v1/namespaces/%s/details-embed/%s"
	httpMethodGetDetailsListEmbedStruct    = "GET"
	uriPathClientGetDetailsListEmbedStruct = "/api/v1/namespaces/%s/details-embed-array/%s"
	httpMethodPutDetails                   = "PUT"
	uriPathClientPutDetails                = "/api/v1/namespaces/%s/details/%s"
	httpMethodGetSomeElseDataUtf8          = "GET"
	uriPathClientGetSomeElseDataUtf8       = "/api/v1/someelsedata"
)

type errorProcessor interface {
	Decode(r *fasthttp.Response) error
}

// New ...
func New(serverURL string, serverHost string, maxConns int, errorProcessor errorProcessor, options map[interface{}]Option) SomeService {
	transportGetWarehouses := NewGetWarehousesTransport(
		errorProcessor,
		serverURL+uriPathClientGetWarehouses,
		httpMethodGetWarehouses,
	)
	transportGetDetails := NewGetDetailsTransport(
		errorProcessor,
		serverURL+uriPathClientGetDetails,
		httpMethodGetDetails,
	)
	transportGetDetailsEmbedStruct := NewGetDetailsEmbedStructTransport(
		errorProcessor,
		serverURL+uriPathClientGetDetailsEmbedStruct,
		httpMethodGetDetailsEmbedStruct,
	)
	transportGetDetailsListEmbedStruct := NewGetDetailsListEmbedStructTransport(
		errorProcessor,
		serverURL+uriPathClientGetDetailsListEmbedStruct,
		httpMethodGetDetailsListEmbedStruct,
	)
	transportPutDetails := NewPutDetailsTransport(
		errorProcessor,
		serverURL+uriPathClientPutDetails,
		httpMethodPutDetails,
	)
	transportGetSomeElseDataUtf8 := NewGetSomeElseDataUtf8Transport(
		errorProcessor,
		serverURL+uriPathClientGetSomeElseDataUtf8,
		httpMethodGetSomeElseDataUtf8,
	)

	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},
		transportGetWarehouses,
		transportGetDetails,
		transportGetDetailsEmbedStruct,
		transportGetDetailsListEmbedStruct,
		transportPutDetails,
		transportGetSomeElseDataUtf8,
		options,
	)
}
