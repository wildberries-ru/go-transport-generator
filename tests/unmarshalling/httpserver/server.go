// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"

	"github.com/valyala/fasthttp"
)

type service interface {
	TestEasyJson(ctx context.Context, param1 int) (field1 string, field2 string, err error)
}

type testEasyJson struct {
	transport      TestEasyJsonTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *testEasyJson) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		param1 int
		field1 string
		field2 string
		err    error
	)
	param1, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	field1, field2, err = s.service.TestEasyJson(ctx, param1)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, field1, field2); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewTestEasyJson the server creator
func NewTestEasyJson(transport TestEasyJsonTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := testEasyJson{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
