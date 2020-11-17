// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"
)

// Options ...
var (
	TestEasyJson = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service ...
type Service interface {
	TestEasyJson(ctx context.Context, param1 int) (field1 string, field2 string, err error)
}

type client struct {
	cli                   *fasthttp.HostClient
	transportTestEasyJson TestEasyJsonTransport
	options               map[interface{}]Option
}

// TestEasyJson ...
func (s *client) TestEasyJson(ctx context.Context, param1 int) (field1 string, field2 string, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[TestEasyJson]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportTestEasyJson.EncodeRequest(ctx, req, param1); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportTestEasyJson.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,
	transportTestEasyJson TestEasyJsonTransport,
	options map[interface{}]Option,
) Service {
	return &client{
		cli:                   cli,
		transportTestEasyJson: transportTestEasyJson,
		options:               options,
	}
}
