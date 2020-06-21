// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"fmt"
	"net/url"

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
func New(
	serverURL string,
	maxConns int,
	errorProcessor errorProcessor,
	options map[interface{}]Option,
) (client Service, err error) {
	parsedServerURL, err := url.Parse(serverURL)
	if err != nil {
		err = fmt.Errorf("failed to parse apiserver url", err)
		return
	}
	transportCreateMultipartUpload := NewCreateMultipartUploadTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientCreateMultipartUpload,
		httpMethodCreateMultipartUpload,
	)
	transportUploadPartDocument := NewUploadPartDocumentTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientUploadPartDocument,
		httpMethodUploadPartDocument,
	)
	transportCompleteUpload := NewCompleteUploadTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientCompleteUpload,
		httpMethodCompleteUpload,
	)
	transportUploadDocument := NewUploadDocumentTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientUploadDocument,
		httpMethodUploadDocument,
	)
	transportDownloadDocument := NewDownloadDocumentTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientDownloadDocument,
		httpMethodDownloadDocument,
	)

	client = NewClient(
		&fasthttp.HostClient{
			Addr:     parsedServerURL.Host,
			MaxConns: maxConns,
		},
		transportCreateMultipartUpload,
		transportUploadPartDocument,
		transportCompleteUpload,
		transportUploadDocument,
		transportDownloadDocument,
		options,
	)
	return
}
