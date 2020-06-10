// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
	"github.com/valyala/fasthttp"
)

type service interface {
	CreateMultipartUpload(ctx context.Context, bucket string, key string) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error)
	UploadPartDocument(ctx context.Context, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error)
	CompleteUpload(ctx context.Context, bucket string, key string, uploadID string) (err error)
	UploadDocument(ctx context.Context, bucket string, key string, document []byte) (err error)
	DownloadDocument(ctx context.Context, bucket string, key string) (document []byte, err error)
}

type createMultipartUploadSwaggerInfo struct {
	transport      CreateMultipartUploadTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *createMultipartUploadSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		bucket           string
		key              string
		data             v1.CreateMultipartUploadData
		errorFlag        bool
		errorText        string
		additionalErrors *v1.AdditionalErrors
		err              error
	)
	bucket, key, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	data, errorFlag, errorText, additionalErrors, err = s.service.CreateMultipartUpload(ctx, bucket, key)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, data, errorFlag, errorText, additionalErrors); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewCreateMultipartUploadSwaggerInfo the server creator
func NewCreateMultipartUploadSwaggerInfo(transport CreateMultipartUploadTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := createMultipartUploadSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type uploadPartDocumentSwaggerInfo struct {
	transport      UploadPartDocumentTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *uploadPartDocumentSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		bucket     string
		key        string
		uploadID   string
		partNumber int64
		document   []byte
		err        error
	)
	bucket, key, uploadID, partNumber, document, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	err = s.service.UploadPartDocument(ctx, bucket, key, uploadID, partNumber, document)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewUploadPartDocumentSwaggerInfo the server creator
func NewUploadPartDocumentSwaggerInfo(transport UploadPartDocumentTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := uploadPartDocumentSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type completeUploadSwaggerInfo struct {
	transport      CompleteUploadTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *completeUploadSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		bucket   string
		key      string
		uploadID string
		err      error
	)
	bucket, key, uploadID, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	err = s.service.CompleteUpload(ctx, bucket, key, uploadID)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewCompleteUploadSwaggerInfo the server creator
func NewCompleteUploadSwaggerInfo(transport CompleteUploadTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := completeUploadSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type uploadDocumentSwaggerInfo struct {
	transport      UploadDocumentTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *uploadDocumentSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		bucket   string
		key      string
		document []byte
		err      error
	)
	bucket, key, document, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	err = s.service.UploadDocument(ctx, bucket, key, document)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewUploadDocumentSwaggerInfo the server creator
func NewUploadDocumentSwaggerInfo(transport UploadDocumentTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := uploadDocumentSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type downloadDocumentSwaggerInfo struct {
	transport      DownloadDocumentTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *downloadDocumentSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		bucket   string
		key      string
		document []byte
		err      error
	)
	bucket, key, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	document, err = s.service.DownloadDocument(ctx, bucket, key)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, document); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewDownloadDocumentSwaggerInfo the server creator
func NewDownloadDocumentSwaggerInfo(transport DownloadDocumentTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := downloadDocumentSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
