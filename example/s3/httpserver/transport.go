// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"bytes"
	"net/http"

	"github.com/valyala/fasthttp"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

var (
	emptyBytes = []byte("")
)

//easyjson:json
type createMultipartUploadResponse struct {
	AdditionalErrors *v1.AdditionalErrors `json:"additionalErrors"`

	Data v1.CreateMultipartUploadData `json:"data"`

	ErrorFlag bool `json:"error"`

	ErrorText string `json:"errorText"`
}

// CreateMultipartUploadTransport transport interface
type CreateMultipartUploadTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors) (err error)
}

type createMultipartUploadTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *createMultipartUploadTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, err error) {

	bucket = ctx.UserValue("bucket").(string)

	key = ctx.UserValue("key").(string)

	return
}

// EncodeResponse method for encoding response on server side
func (t *createMultipartUploadTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors) (err error) {

	r.Header.Set("Content-Type", "application/json")
	var theResponse createMultipartUploadResponse

	theResponse.AdditionalErrors = additionalErrors

	theResponse.Data = data

	theResponse.ErrorFlag = errorFlag

	theResponse.ErrorText = errorText

	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)

	r.Header.SetStatusCode(201)
	return
}

// NewCreateMultipartUploadTransport the transport creator for http requests
func NewCreateMultipartUploadTransport(

	encodeJSONErrorCreator errorCreator,

) CreateMultipartUploadTransport {
	return &createMultipartUploadTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

// UploadPartDocumentTransport transport interface
type UploadPartDocumentTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, uploadID string, partNumber int64, document []byte, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error)
}

type uploadPartDocumentTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *uploadPartDocumentTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, uploadID string, partNumber int64, document []byte, err error) {

	bucket = ctx.UserValue("bucket").(string)

	key = ctx.UserValue("key").(string)

	return
}

// EncodeResponse method for encoding response on server side
func (t *uploadPartDocumentTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error) {

	r.Header.Set("Content-Type", "application/json")

	r.Header.SetStatusCode(200)
	return
}

// NewUploadPartDocumentTransport the transport creator for http requests
func NewUploadPartDocumentTransport() UploadPartDocumentTransport {
	return &uploadPartDocumentTransport{}
}

// CompleteUploadTransport transport interface
type CompleteUploadTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, uploadID string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error)
}

type completeUploadTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *completeUploadTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, uploadID string, err error) {

	bucket = ctx.UserValue("bucket").(string)

	key = ctx.UserValue("key").(string)

	return
}

// EncodeResponse method for encoding response on server side
func (t *completeUploadTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error) {

	r.Header.Set("Content-Type", "application/json")

	r.Header.SetStatusCode(200)
	return
}

// NewCompleteUploadTransport the transport creator for http requests
func NewCompleteUploadTransport() CompleteUploadTransport {
	return &completeUploadTransport{}
}

// UploadDocumentTransport transport interface
type UploadDocumentTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, document []byte, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error)
}

type uploadDocumentTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *uploadDocumentTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, document []byte, err error) {

	bucket = ctx.UserValue("bucket").(string)

	key = ctx.UserValue("key").(string)

	return
}

// EncodeResponse method for encoding response on server side
func (t *uploadDocumentTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error) {

	r.Header.Set("Content-Type", "application/json")

	r.Header.SetStatusCode(201)
	return
}

// NewUploadDocumentTransport the transport creator for http requests
func NewUploadDocumentTransport() UploadDocumentTransport {
	return &uploadDocumentTransport{}
}

//easyjson:json
type downloadDocumentResponse struct {
	Document []byte `json:"document"`
}

// DownloadDocumentTransport transport interface
type DownloadDocumentTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, document []byte) (err error)
}

type downloadDocumentTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *downloadDocumentTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (bucket string, key string, err error) {

	bucket = ctx.UserValue("bucket").(string)

	key = ctx.UserValue("key").(string)

	return
}

// EncodeResponse method for encoding response on server side
func (t *downloadDocumentTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, document []byte) (err error) {

	r.Header.Set("Content-Type", "application/json")
	var theResponse downloadDocumentResponse

	theResponse.Document = document

	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)

	r.Header.SetStatusCode(200)
	return
}

// NewDownloadDocumentTransport the transport creator for http requests
func NewDownloadDocumentTransport(

	encodeJSONErrorCreator errorCreator,

) DownloadDocumentTransport {
	return &downloadDocumentTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

//easyjson:json
type getTokenResponse struct {
	ExpiresIn int `json:"expiresIn"`

	Token string `json:"token"`
}

// GetTokenTransport transport interface
type GetTokenTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (authToken *string, scope string, grantType string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, token string, expiresIn int) (err error)
}

type getTokenTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getTokenTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (authToken *string, scope string, grantType string, err error) {

	authToken = ptr(r.Header.Peek("Authorization"))

	return
}

// EncodeResponse method for encoding response on server side
func (t *getTokenTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, token string, expiresIn int) (err error) {

	r.Header.Set("Content-Type", "application/json")
	var theResponse getTokenResponse

	theResponse.ExpiresIn = expiresIn

	theResponse.Token = token

	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)

	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewGetTokenTransport the transport creator for http requests
func NewGetTokenTransport(

	encodeJSONErrorCreator errorCreator,

) GetTokenTransport {
	return &getTokenTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

func ptr(in []byte) *string {
	if bytes.Equal(in, emptyBytes) {
		return nil
	}
	i := string(in)
	return &i
}
