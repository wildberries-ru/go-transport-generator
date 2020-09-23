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
	CreateMultipartUpload = option{}
	UploadPartDocument    = option{}
	CompleteUpload        = option{}
	UploadDocument        = option{}
	DownloadDocument      = option{}
	GetToken              = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service ...
type Service interface {
	CreateMultipartUpload(ctx context.Context, bucket string, key string) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error)
	UploadPartDocument(ctx context.Context, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error)
	CompleteUpload(ctx context.Context, bucket string, key string, uploadID string) (err error)
	UploadDocument(ctx context.Context, bucket string, key string, document []byte) (err error)
	DownloadDocument(ctx context.Context, bucket string, key string) (document []byte, err error)
	GetToken(ctx context.Context, authToken *string, scope string, grantType string) (token string, expiresIn int, err error)
}

type client struct {
	cli                            *fasthttp.HostClient
	transportCreateMultipartUpload CreateMultipartUploadTransport
	transportUploadPartDocument    UploadPartDocumentTransport
	transportCompleteUpload        CompleteUploadTransport
	transportUploadDocument        UploadDocumentTransport
	transportDownloadDocument      DownloadDocumentTransport
	transportGetToken              GetTokenTransport
	options                        map[interface{}]Option
}

// CreateMultipartUpload ...
func (s *client) CreateMultipartUpload(ctx context.Context, bucket string, key string) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[CreateMultipartUpload]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportCreateMultipartUpload.EncodeRequest(ctx, req, bucket, key); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportCreateMultipartUpload.DecodeResponse(ctx, res)
}

// UploadPartDocument ...
func (s *client) UploadPartDocument(ctx context.Context, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[UploadPartDocument]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportUploadPartDocument.EncodeRequest(ctx, req, bucket, key, uploadID, partNumber, document); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportUploadPartDocument.DecodeResponse(ctx, res)
}

// CompleteUpload ...
func (s *client) CompleteUpload(ctx context.Context, bucket string, key string, uploadID string) (err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[CompleteUpload]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportCompleteUpload.EncodeRequest(ctx, req, bucket, key, uploadID); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportCompleteUpload.DecodeResponse(ctx, res)
}

// UploadDocument ...
func (s *client) UploadDocument(ctx context.Context, bucket string, key string, document []byte) (err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[UploadDocument]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportUploadDocument.EncodeRequest(ctx, req, bucket, key, document); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportUploadDocument.DecodeResponse(ctx, res)
}

// DownloadDocument ...
func (s *client) DownloadDocument(ctx context.Context, bucket string, key string) (document []byte, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[DownloadDocument]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportDownloadDocument.EncodeRequest(ctx, req, bucket, key); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportDownloadDocument.DecodeResponse(ctx, res)
}

// GetToken ...
func (s *client) GetToken(ctx context.Context, authToken *string, scope string, grantType string) (token string, expiresIn int, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetToken]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetToken.EncodeRequest(ctx, req, authToken, scope, grantType); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetToken.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,
	transportCreateMultipartUpload CreateMultipartUploadTransport,
	transportUploadPartDocument UploadPartDocumentTransport,
	transportCompleteUpload CompleteUploadTransport,
	transportUploadDocument UploadDocumentTransport,
	transportDownloadDocument DownloadDocumentTransport,
	transportGetToken GetTokenTransport,
	options map[interface{}]Option,
) Service {
	return &client{
		cli:                            cli,
		transportCreateMultipartUpload: transportCreateMultipartUpload,
		transportUploadPartDocument:    transportUploadPartDocument,
		transportCompleteUpload:        transportCompleteUpload,
		transportUploadDocument:        transportUploadDocument,
		transportDownloadDocument:      transportDownloadDocument,
		transportGetToken:              transportGetToken,
		options:                        options,
	}
}
