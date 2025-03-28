// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"net/url"
	"time"

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

// Config ...
type Config struct {
	ServerURL           string
	MaxConns            *int
	MaxConnDuration     *time.Duration
	MaxIdleConnDuration *time.Duration
	ReadBufferSize      *int
	WriteBufferSize     *int
	ReadTimeout         *time.Duration
	WriteTimeout        *time.Duration
	MaxResponseBodySize *int
}

// New ...
func New(
	config Config,
	errorProcessor errorProcessor,
	options map[interface{}]Option,
) (client Service, err error) {
	parsedServerURL, err := url.Parse(config.ServerURL)
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

	cli := fasthttp.HostClient{
		Addr: parsedServerURL.Host,
	}
	if config.MaxConns != nil {
		cli.MaxConns = *config.MaxConns
	}
	if config.MaxConnDuration != nil {
		cli.MaxConnDuration = *config.MaxConnDuration
	}
	if config.MaxIdleConnDuration != nil {
		cli.MaxIdleConnDuration = *config.MaxIdleConnDuration
	}
	if config.ReadBufferSize != nil {
		cli.ReadBufferSize = *config.ReadBufferSize
	}
	if config.WriteBufferSize != nil {
		cli.WriteBufferSize = *config.WriteBufferSize
	}
	if config.ReadTimeout != nil {
		cli.ReadTimeout = *config.ReadTimeout
	}
	if config.WriteTimeout != nil {
		cli.WriteTimeout = *config.WriteTimeout
	}
	if config.MaxResponseBodySize != nil {
		cli.MaxResponseBodySize = *config.MaxResponseBodySize
	}

	client = NewClient(
		&cli,
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
