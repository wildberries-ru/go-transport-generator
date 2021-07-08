// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"net/url"

	"github.com/pkg/errors"
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
	httpMethodGetToken                 = "POST"
	uriPathClientGetToken              = "/token"
	httpMethodGetBranches              = "GET"
	uriPathClientGetBranches           = "/api/v1/branches"
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
		err = errors.Wrap(err, "failed to parse server url")
		return
	}
	transportCreateMultipartUpload := NewCreateMultipartUploadTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientCreateMultipartUpload,
		httpMethodCreateMultipartUpload,
	)
	transportUploadPartDocument := NewUploadPartDocumentTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientUploadPartDocument,
		httpMethodUploadPartDocument,
	)
	transportCompleteUpload := NewCompleteUploadTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientCompleteUpload,
		httpMethodCompleteUpload,
	)
	transportUploadDocument := NewUploadDocumentTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientUploadDocument,
		httpMethodUploadDocument,
	)
	transportDownloadDocument := NewDownloadDocumentTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientDownloadDocument,
		httpMethodDownloadDocument,
	)
	transportGetToken := NewGetTokenTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetToken,
		httpMethodGetToken,
	)
	transportGetBranches := NewGetBranchesTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetBranches,
		httpMethodGetBranches,
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
		transportGetToken,
		transportGetBranches,
		options,
	)
	return
}
