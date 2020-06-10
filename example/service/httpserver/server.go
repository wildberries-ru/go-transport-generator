// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"

	"github.com/valyala/fasthttp"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

type service interface {
	GetWarehouses(ctx context.Context) (pets []v1.Detail, err error)
	GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error)
	GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error)
	GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error)
	PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
	GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
}

type getWarehousesSwaggerInfo struct {
	transport      GetWarehousesTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getWarehousesSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		pets []v1.Detail
		err  error
	)
	err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	pets, err = s.service.GetWarehouses(ctx)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, pets); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetWarehousesSwaggerInfo the server creator
func NewGetWarehousesSwaggerInfo(transport GetWarehousesTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getWarehousesSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getDetailsSwaggerInfo struct {
	transport      GetDetailsTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getDetailsSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		namespace string
		detail    string
		fileID    uint32
		someID    *uint64
		token     *string
		det       v1.Detail
		ns        v1.Namespace
		id        *string
		err       error
	)
	namespace, detail, fileID, someID, token, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	det, ns, id, err = s.service.GetDetails(ctx, namespace, detail, fileID, someID, token)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, det, ns, id); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetDetailsSwaggerInfo the server creator
func NewGetDetailsSwaggerInfo(transport GetDetailsTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getDetailsSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getDetailsEmbedStructSwaggerInfo struct {
	transport      GetDetailsEmbedStructTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getDetailsEmbedStructSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		namespace string
		detail    string
		response  v1.GetDetailsEmbedStructResponse
		err       error
	)
	namespace, detail, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err = s.service.GetDetailsEmbedStruct(ctx, namespace, detail)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetDetailsEmbedStructSwaggerInfo the server creator
func NewGetDetailsEmbedStructSwaggerInfo(transport GetDetailsEmbedStructTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getDetailsEmbedStructSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getDetailsListEmbedStructSwaggerInfo struct {
	transport      GetDetailsListEmbedStructTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getDetailsListEmbedStructSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		namespace string
		detail    string
		details   []v1.Detail
		err       error
	)
	namespace, detail, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	details, err = s.service.GetDetailsListEmbedStruct(ctx, namespace, detail)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, details); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetDetailsListEmbedStructSwaggerInfo the server creator
func NewGetDetailsListEmbedStructSwaggerInfo(transport GetDetailsListEmbedStructTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getDetailsListEmbedStructSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type putDetailsSwaggerInfo struct {
	transport      PutDetailsTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *putDetailsSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		namespace string
		detail    string
		testID    string
		blaID     *string
		token     *string
		pretty    v1.Detail
		yang      v1.Namespace
		cool      v1.Detail
		nothing   v1.Namespace
		id        *string
		err       error
	)
	namespace, detail, testID, blaID, token, pretty, yang, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	cool, nothing, id, err = s.service.PutDetails(ctx, namespace, detail, testID, blaID, token, pretty, yang)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, cool, nothing, id); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewPutDetailsSwaggerInfo the server creator
func NewPutDetailsSwaggerInfo(transport PutDetailsTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := putDetailsSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getSomeElseDataUtf8SwaggerInfo struct {
	transport      GetSomeElseDataUtf8Transport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getSomeElseDataUtf8SwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		cool    v1.Detail
		nothing v1.Namespace
		id      *string
		err     error
	)
	err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	cool, nothing, id, err = s.service.GetSomeElseDataUtf8(ctx)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, cool, nothing, id); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetSomeElseDataUtf8SwaggerInfo the server creator
func NewGetSomeElseDataUtf8SwaggerInfo(transport GetSomeElseDataUtf8Transport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getSomeElseDataUtf8SwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
