// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
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
	httpMethodGetFile                      = "GET"
	uriPathClientGetFile                   = "/api/v1/file"
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
	defaultTimeOut time.Duration,
) (client SomeService, err error) {
	parsedServerURL, err := url.Parse(serverURL)
	if err != nil {
		err = errors.Wrap(err, "failed to parse server url")
		return
	}
	transportGetWarehouses := NewGetWarehousesTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetWarehouses,
		httpMethodGetWarehouses,
	)
	transportGetDetails := NewGetDetailsTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetDetails,
		httpMethodGetDetails,
	)
	transportGetDetailsEmbedStruct := NewGetDetailsEmbedStructTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetDetailsEmbedStruct,
		httpMethodGetDetailsEmbedStruct,
	)
	transportGetDetailsListEmbedStruct := NewGetDetailsListEmbedStructTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetDetailsListEmbedStruct,
		httpMethodGetDetailsListEmbedStruct,
	)
	transportPutDetails := NewPutDetailsTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientPutDetails,
		httpMethodPutDetails,
	)
	transportGetSomeElseDataUtf8 := NewGetSomeElseDataUtf8Transport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetSomeElseDataUtf8,
		httpMethodGetSomeElseDataUtf8,
	)
	transportGetFile := NewGetFileTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetFile,
		httpMethodGetFile,
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
		transportGetFile,
		options,
		defaultTimeOut,
	)
	return
}
