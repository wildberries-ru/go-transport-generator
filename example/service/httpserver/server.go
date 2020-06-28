// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"
	"mime/multipart"

	"github.com/valyala/fasthttp"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

type service interface {
	UploadDocument(ctx context.Context, token *string, name string, extension string, categoryID string, supplierID *int64, contractID *int64, data multipart.File) (err error)
	GetWarehouses(ctx context.Context) (pets map[string]v1.Detail, err error)
	GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error)
	GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error)
	GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error)
	PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
	GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
	GetFile(ctx context.Context) (data []byte, fileName string, err error)
}

type uploadDocument struct {
	transport      UploadDocumentTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *uploadDocument) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		token      *string
		name       string
		extension  string
		categoryID string
		supplierID *int64
		contractID *int64
		data       multipart.File
		err        error
	)
	token, name, extension, categoryID, supplierID, contractID, data, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	err = s.service.UploadDocument(ctx, token, name, extension, categoryID, supplierID, contractID, data)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewUploadDocument the server creator
func NewUploadDocument(transport UploadDocumentTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := uploadDocument{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getWarehouses struct {
	transport      GetWarehousesTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getWarehouses) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		pets map[string]v1.Detail
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

// NewGetWarehouses the server creator
func NewGetWarehouses(transport GetWarehousesTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getWarehouses{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getDetails struct {
	transport      GetDetailsTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getDetails) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewGetDetails the server creator
func NewGetDetails(transport GetDetailsTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getDetails{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getDetailsEmbedStruct struct {
	transport      GetDetailsEmbedStructTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getDetailsEmbedStruct) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewGetDetailsEmbedStruct the server creator
func NewGetDetailsEmbedStruct(transport GetDetailsEmbedStructTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getDetailsEmbedStruct{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getDetailsListEmbedStruct struct {
	transport      GetDetailsListEmbedStructTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getDetailsListEmbedStruct) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewGetDetailsListEmbedStruct the server creator
func NewGetDetailsListEmbedStruct(transport GetDetailsListEmbedStructTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getDetailsListEmbedStruct{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type putDetails struct {
	transport      PutDetailsTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *putDetails) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewPutDetails the server creator
func NewPutDetails(transport PutDetailsTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := putDetails{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getSomeElseDataUtf8 struct {
	transport      GetSomeElseDataUtf8Transport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getSomeElseDataUtf8) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewGetSomeElseDataUtf8 the server creator
func NewGetSomeElseDataUtf8(transport GetSomeElseDataUtf8Transport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getSomeElseDataUtf8{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getFile struct {
	transport      GetFileTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getFile) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		data     []byte
		fileName string
		err      error
	)
	err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	data, fileName, err = s.service.GetFile(ctx)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, data, fileName); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetFile the server creator
func NewGetFile(transport GetFileTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getFile{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
