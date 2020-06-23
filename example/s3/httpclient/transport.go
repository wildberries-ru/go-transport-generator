// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"
	"fmt"

	"github.com/valyala/fasthttp"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

//easyjson:json
type createMultipartUploadResponse struct {
	AdditionalErrors *v1.AdditionalErrors         `json:"additionalErrors"`
	Data             v1.CreateMultipartUploadData `json:"data"`
	ErrorFlag        bool                         `json:"error"`
	ErrorText        string                       `json:"errorText"`
}

// CreateMultipartUploadTransport transport interface
type CreateMultipartUploadTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error)
}

type createMultipartUploadTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *createMultipartUploadTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, bucket, key))

	return
}

// DecodeResponse method for encoding response on server side
func (t *createMultipartUploadTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error) {
	if r.StatusCode() != 201 {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse createMultipartUploadResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
		return
	}

	additionalErrors = theResponse.AdditionalErrors

	data = theResponse.Data

	errorFlag = theResponse.ErrorFlag

	errorText = theResponse.ErrorText

	return
}

// NewCreateMultipartUploadTransport the transport creator for http requests
func NewCreateMultipartUploadTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) CreateMultipartUploadTransport {
	return &createMultipartUploadTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

type uploadPartDocumentRequest struct {
	Document   []byte
	PartNumber int64
	UploadID   string
}

// UploadPartDocumentTransport transport interface
type UploadPartDocumentTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (err error)
}

type uploadPartDocumentTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *uploadPartDocumentTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, bucket, key))

	return
}

// DecodeResponse method for encoding response on server side
func (t *uploadPartDocumentTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (err error) {
	if r.StatusCode() != 200 {
		err = t.errorProcessor.Decode(r)
		return
	}

	return
}

// NewUploadPartDocumentTransport the transport creator for http requests
func NewUploadPartDocumentTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) UploadPartDocumentTransport {
	return &uploadPartDocumentTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

type completeUploadRequest struct {
	UploadID string
}

// CompleteUploadTransport transport interface
type CompleteUploadTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string, uploadID string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (err error)
}

type completeUploadTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *completeUploadTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string, uploadID string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, bucket, key))

	return
}

// DecodeResponse method for encoding response on server side
func (t *completeUploadTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (err error) {
	if r.StatusCode() != 200 {
		err = t.errorProcessor.Decode(r)
		return
	}

	return
}

// NewCompleteUploadTransport the transport creator for http requests
func NewCompleteUploadTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) CompleteUploadTransport {
	return &completeUploadTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

type uploadDocumentRequest struct {
	Document []byte
}

// UploadDocumentTransport transport interface
type UploadDocumentTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string, document []byte) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (err error)
}

type uploadDocumentTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *uploadDocumentTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string, document []byte) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, bucket, key))

	return
}

// DecodeResponse method for encoding response on server side
func (t *uploadDocumentTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (err error) {
	if r.StatusCode() != 201 {
		err = t.errorProcessor.Decode(r)
		return
	}

	return
}

// NewUploadDocumentTransport the transport creator for http requests
func NewUploadDocumentTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) UploadDocumentTransport {
	return &uploadDocumentTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

//easyjson:json
type downloadDocumentResponse struct {
	Document []byte
}

// DownloadDocumentTransport transport interface
type DownloadDocumentTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (document []byte, err error)
}

type downloadDocumentTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *downloadDocumentTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, bucket string, key string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, bucket, key))

	return
}

// DecodeResponse method for encoding response on server side
func (t *downloadDocumentTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (document []byte, err error) {
	if r.StatusCode() != 200 {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse downloadDocumentResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
		return
	}

	document = theResponse.Document

	return
}

// NewDownloadDocumentTransport the transport creator for http requests
func NewDownloadDocumentTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) DownloadDocumentTransport {
	return &downloadDocumentTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

func ptr(in []byte) *string {
	i := string(in)
	return &i
}
