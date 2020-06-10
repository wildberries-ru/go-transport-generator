// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// Options ...
var (
	GetWarehouses             = option{}
	GetDetails                = option{}
	GetDetailsEmbedStruct     = option{}
	GetDetailsListEmbedStruct = option{}
	PutDetails                = option{}
	GetSomeElseDataUtf8       = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// SomeService ...
type SomeService interface {
	GetWarehouses(ctx context.Context) (pets []v1.Detail, err error)
	GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error)
	GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error)
	GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error)
	PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
	GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
}

type client struct {
	cli                                *fasthttp.HostClient
	transportGetWarehouses             GetWarehousesTransport
	transportGetDetails                GetDetailsTransport
	transportGetDetailsEmbedStruct     GetDetailsEmbedStructTransport
	transportGetDetailsListEmbedStruct GetDetailsListEmbedStructTransport
	transportPutDetails                PutDetailsTransport
	transportGetSomeElseDataUtf8       GetSomeElseDataUtf8Transport
	options                            map[interface{}]Option
}

// GetWarehouses ...
func (s *client) GetWarehouses(ctx context.Context) (pets []v1.Detail, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetWarehouses]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetWarehouses.EncodeRequest(ctx, req); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetWarehouses.DecodeResponse(ctx, res)
}

// GetDetails ...
func (s *client) GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetDetails]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetDetails.EncodeRequest(ctx, req, namespace, detail, fileID, someID, token); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetDetails.DecodeResponse(ctx, res)
}

// GetDetailsEmbedStruct ...
func (s *client) GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetDetailsEmbedStruct]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetDetailsEmbedStruct.EncodeRequest(ctx, req, namespace, detail); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetDetailsEmbedStruct.DecodeResponse(ctx, res)
}

// GetDetailsListEmbedStruct ...
func (s *client) GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetDetailsListEmbedStruct]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetDetailsListEmbedStruct.EncodeRequest(ctx, req, namespace, detail); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetDetailsListEmbedStruct.DecodeResponse(ctx, res)
}

// PutDetails ...
func (s *client) PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[PutDetails]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportPutDetails.EncodeRequest(ctx, req, namespace, detail, testID, blaID, token, pretty, yang); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportPutDetails.DecodeResponse(ctx, res)
}

// GetSomeElseDataUtf8 ...
func (s *client) GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetSomeElseDataUtf8]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetSomeElseDataUtf8.EncodeRequest(ctx, req); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetSomeElseDataUtf8.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,
	transportGetWarehouses GetWarehousesTransport,
	transportGetDetails GetDetailsTransport,
	transportGetDetailsEmbedStruct GetDetailsEmbedStructTransport,
	transportGetDetailsListEmbedStruct GetDetailsListEmbedStructTransport,
	transportPutDetails PutDetailsTransport,
	transportGetSomeElseDataUtf8 GetSomeElseDataUtf8Transport,
	options map[interface{}]Option,
) SomeService {
	return &client{
		cli:                                cli,
		transportGetWarehouses:             transportGetWarehouses,
		transportGetDetails:                transportGetDetails,
		transportGetDetailsEmbedStruct:     transportGetDetailsEmbedStruct,
		transportGetDetailsListEmbedStruct: transportGetDetailsListEmbedStruct,
		transportPutDetails:                transportPutDetails,
		transportGetSomeElseDataUtf8:       transportGetSomeElseDataUtf8,
		options:                            options,
	}
}
