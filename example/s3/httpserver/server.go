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
	CreateMultipartUpload(ctx context.Context, bucket string, key string) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error)
	UploadPartDocument(ctx context.Context, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error)
	CompleteUpload(ctx context.Context, bucket string, key string, uploadID string) (err error)
	UploadDocument(ctx context.Context, bucket string, key string, document []byte) (err error)
	DownloadDocument(ctx context.Context, bucket string, key string) (document []byte, err error)
	GetToken(ctx context.Context, authToken *string, scope string, grantType string) (token string, expiresIn int, err error)
	GetBranches(ctx context.Context, authToken *string, supplierID *string) (branches []int, err error)
}

type createMultipartUpload struct {
	transport      CreateMultipartUploadTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *createMultipartUpload) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewCreateMultipartUpload the server creator
func NewCreateMultipartUpload(transport CreateMultipartUploadTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := createMultipartUpload{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type uploadPartDocument struct {
	transport      UploadPartDocumentTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *uploadPartDocument) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewUploadPartDocument the server creator
func NewUploadPartDocument(transport UploadPartDocumentTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := uploadPartDocument{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type completeUpload struct {
	transport      CompleteUploadTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *completeUpload) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewCompleteUpload the server creator
func NewCompleteUpload(transport CompleteUploadTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := completeUpload{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type uploadDocument struct {
	transport      UploadDocumentTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *uploadDocument) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewUploadDocument the server creator
func NewUploadDocument(transport UploadDocumentTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := uploadDocument{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type downloadDocument struct {
	transport      DownloadDocumentTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *downloadDocument) ServeHTTP(ctx *fasthttp.RequestCtx) {
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

// NewDownloadDocument the server creator
func NewDownloadDocument(transport DownloadDocumentTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := downloadDocument{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getToken struct {
	transport      GetTokenTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getToken) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		authToken *string
		scope     string
		grantType string
		token     string
		expiresIn int
		err       error
	)
	authToken, scope, grantType, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	token, expiresIn, err = s.service.GetToken(ctx, authToken, scope, grantType)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, token, expiresIn); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetToken the server creator
func NewGetToken(transport GetTokenTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getToken{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getBranches struct {
	transport      GetBranchesTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getBranches) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		authToken  *string
		supplierID *string
		branches   []int
		err        error
	)
	authToken, supplierID, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	branches, err = s.service.GetBranches(ctx, authToken, supplierID)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, branches); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetBranches the server creator
func NewGetBranches(transport GetBranchesTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getBranches{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
