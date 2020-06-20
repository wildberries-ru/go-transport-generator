// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"fmt"
	"net/url"

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
func New(
	serverURL string,
	maxConns int,
	errorProcessor errorProcessor,
	options map[interface{}]Option,
) (client SomeService, err error) {
	parsedServerURL, err := url.Parse(serverURL)
	if err != nil {
		err = fmt.Errorf("failed to parse apiserver url", err)
		return
	}
	transportGetWarehouses := NewGetWarehousesTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetWarehouses,
		httpMethodGetWarehouses,
	)
	transportGetDetails := NewGetDetailsTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetDetails,
		httpMethodGetDetails,
	)
	transportGetDetailsEmbedStruct := NewGetDetailsEmbedStructTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetDetailsEmbedStruct,
		httpMethodGetDetailsEmbedStruct,
	)
	transportGetDetailsListEmbedStruct := NewGetDetailsListEmbedStructTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetDetailsListEmbedStruct,
		httpMethodGetDetailsListEmbedStruct,
	)
	transportPutDetails := NewPutDetailsTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientPutDetails,
		httpMethodPutDetails,
	)
	transportGetSomeElseDataUtf8 := NewGetSomeElseDataUtf8Transport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetSomeElseDataUtf8,
		httpMethodGetSomeElseDataUtf8,
	)

	client = NewClient(
		&fasthttp.HostClient{
			Addr:     parsedServerURL.Host,
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
	return
}
