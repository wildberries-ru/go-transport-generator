// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"github.com/valyala/fasthttp"
)

const (
	httpMethodCreateMultipartUpload    = "POST"
	uriPathClientCreateMultipartUpload = "/api/v1/multipart/%s/%s"
	httpMethodUploadPartDocument       = "PATCH"
	uriPathClientUploadPartDocument    = "/api/v1/multipart/%s/%s"
	httpMethodCompleteUpload           = "PUT"
	uriPathClientCompleteUpload        = "/api/v1/multipart/%s/%s"
	httpMethodUploadDocument           = "POST"
	uriPathClientUploadDocument        = "/api/v1/doc/%s/%s"
	httpMethodDownloadDocument         = "GET"
	uriPathClientDownloadDocument      = "/api/v1/doc/%s/%s"
)

type errorProcessor interface {
	Decode(r *fasthttp.Response) error
}

// New ...
func New(serverURL string, serverHost string, maxConns int, errorProcessor errorProcessor, options map[interface{}]Option) Service {
	transportCreateMultipartUpload := NewCreateMultipartUploadTransport(
		errorProcessor,
		serverURL+uriPathClientCreateMultipartUpload,
		httpMethodCreateMultipartUpload,
	)
	transportUploadPartDocument := NewUploadPartDocumentTransport(
		errorProcessor,
		serverURL+uriPathClientUploadPartDocument,
		httpMethodUploadPartDocument,
	)
	transportCompleteUpload := NewCompleteUploadTransport(
		errorProcessor,
		serverURL+uriPathClientCompleteUpload,
		httpMethodCompleteUpload,
	)
	transportUploadDocument := NewUploadDocumentTransport(
		errorProcessor,
		serverURL+uriPathClientUploadDocument,
		httpMethodUploadDocument,
	)
	transportDownloadDocument := NewDownloadDocumentTransport(
		errorProcessor,
		serverURL+uriPathClientDownloadDocument,
		httpMethodDownloadDocument,
	)

	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},
		transportCreateMultipartUpload,
		transportUploadPartDocument,
		transportCompleteUpload,
		transportUploadDocument,
		transportDownloadDocument,
		options,
	)
}
