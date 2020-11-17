// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

const (
	httpMethodTestEasyJson    = "GET"
	uriPathClientTestEasyJson = "/api/testeasyjson"
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
) (client Service, err error) {
	parsedServerURL, err := url.Parse(serverURL)
	if err != nil {
		err = errors.Wrap(err, "failed to parse server url")
		return
	}
	transportTestEasyJson := NewTestEasyJsonTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientTestEasyJson,
		httpMethodTestEasyJson,
	)

	client = NewClient(
		&fasthttp.HostClient{
			Addr:     parsedServerURL.Host,
			MaxConns: maxConns,
		},
		transportTestEasyJson,
		options,
	)
	return
}
