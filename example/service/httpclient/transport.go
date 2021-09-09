// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

type uploadDocumentRequest struct {
	CategoryID string
	ContractID *bool
	Data       *multipart.FileHeader
	Extension  string
	Name       []string
	SupplierID []int64
}

// UploadDocumentTransport transport interface
type UploadDocumentTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, token *string, name []string, extension string, categoryID string, supplierID []int64, contractID *bool, data *multipart.FileHeader) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (err error)
}

type uploadDocumentTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *uploadDocumentTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, token *string, name []string, extension string, categoryID string, supplierID []int64, contractID *bool, data *multipart.FileHeader) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)

	r.Header.Set("Authorization", *token)

	r.Header.Set("Content-Type", "multipart/form-data")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("categoryID", categoryID)

	if contractID != nil {

		writer.WriteField("contractID", fmt.Sprintf("%t", *contractID))

	}

	writer.WriteField("extension", extension)

	// todo generate then name is slice

	// todo generate then supplierID is slice

	writer.Close()
	r.Header.Set("Content-Type", writer.FormDataContentType())
	r.SetBody(body.Bytes())

	return
}

// DecodeResponse method for decoding response on server side
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

type getWarehousesResponse map[string]v1.Detail

// GetWarehousesTransport transport interface
type GetWarehousesTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, token *string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (pets map[string]v1.Detail, someCookie *string, err error)
}

type getWarehousesTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getWarehousesTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, token *string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)

	r.Header.Set("Authorization", *token)

	r.Header.Set("Content-Type", "application/json")

	return
}

// DecodeResponse method for decoding response on server side
func (t *getWarehousesTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (pets map[string]v1.Detail, someCookie *string, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	cookie := fasthttp.AcquireCookie()

	// cookies must be a *string type
	_someCookie := string(r.Header.PeekCookie("some_cookie"))
	someCookie = &_someCookie

	fasthttp.ReleaseCookie(cookie)

	var theResponse getWarehousesResponse
	if err = json.Unmarshal(r.Body(), &theResponse); err != nil {
		return
	}

	pets = theResponse

	return
}

// NewGetWarehousesTransport the transport creator for http requests
func NewGetWarehousesTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) GetWarehousesTransport {
	return &getWarehousesTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

type getDetailsResponse struct {
	Det v1.Detail
	Ns  v1.Namespace
}

// GetDetailsTransport transport interface
type GetDetailsTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string, fileID uint32, someID *uint64, token *string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (det v1.Detail, ns v1.Namespace, id *string, err error)
}

type getDetailsTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getDetailsTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string, fileID uint32, someID *uint64, token *string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, namespace, detail))

	r.URI().QueryArgs().Set("file", strconv.Itoa(int(fileID)))

	if someID != nil {
		r.URI().QueryArgs().Set("some", strconv.Itoa(int(*someID)))
	}

	r.Header.Set("X-Auth-Token", *token)

	r.Header.Set("Content-Type", "application/json")

	return
}

// DecodeResponse method for decoding response on server side
func (t *getDetailsTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (det v1.Detail, ns v1.Namespace, id *string, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse getDetailsResponse
	if err = json.Unmarshal(r.Body(), &theResponse); err != nil {
		return
	}

	det = theResponse.Det

	ns = theResponse.Ns

	id = ptr(r.Header.Peek("X-Auth-ID"))

	return
}

// NewGetDetailsTransport the transport creator for http requests
func NewGetDetailsTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) GetDetailsTransport {
	return &getDetailsTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

type getDetailsEmbedStructResponse struct {
	v1.GetDetailsEmbedStructResponse
}

// GetDetailsEmbedStructTransport transport interface
type GetDetailsEmbedStructTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string, token *string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.GetDetailsEmbedStructResponse, err error)
}

type getDetailsEmbedStructTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getDetailsEmbedStructTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string, token *string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, namespace, detail))

	r.Header.Set("Authorization", *token)

	r.Header.Set("Content-Type", "application/json")

	return
}

// DecodeResponse method for decoding response on server side
func (t *getDetailsEmbedStructTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.GetDetailsEmbedStructResponse, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse getDetailsEmbedStructResponse
	if err = json.Unmarshal(r.Body(), &theResponse); err != nil {
		return
	}

	response = theResponse.GetDetailsEmbedStructResponse

	return
}

// NewGetDetailsEmbedStructTransport the transport creator for http requests
func NewGetDetailsEmbedStructTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) GetDetailsEmbedStructTransport {
	return &getDetailsEmbedStructTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

type getDetailsListEmbedStructResponse []v1.Detail

// GetDetailsListEmbedStructTransport transport interface
type GetDetailsListEmbedStructTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string, token *string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (details []v1.Detail, err error)
}

type getDetailsListEmbedStructTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getDetailsListEmbedStructTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string, token *string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, namespace, detail))

	r.Header.Set("Authorization", *token)

	r.Header.Set("Content-Type", "application/json")

	return
}

// DecodeResponse method for decoding response on server side
func (t *getDetailsListEmbedStructTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (details []v1.Detail, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse getDetailsListEmbedStructResponse
	if err = json.Unmarshal(r.Body(), &theResponse); err != nil {
		return
	}

	details = theResponse

	return
}

// NewGetDetailsListEmbedStructTransport the transport creator for http requests
func NewGetDetailsListEmbedStructTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) GetDetailsListEmbedStructTransport {
	return &getDetailsListEmbedStructTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

type putDetailsRequest struct {
	Pretty v1.Detail `json:"ThePretty"`
	Yang   v1.Namespace
}

type putDetailsResponse struct {
	Cool    v1.Detail
	Nothing v1.Namespace
}

// PutDetailsTransport transport interface
type PutDetailsTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
}

type putDetailsTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *putDetailsTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, namespace, detail))

	if blaID != nil {
		r.URI().QueryArgs().Set("bla", *blaID)
	}

	r.URI().QueryArgs().Set("test", testID)

	r.Header.Set("X-Auth-Token", *token)

	r.Header.Set("Content-Type", "application/json")

	var request putDetailsRequest

	request.Pretty = pretty

	request.Yang = yang

	body, err := json.Marshal(request)
	if err != nil {
		return
	}
	r.SetBody(body)

	return
}

// DecodeResponse method for decoding response on server side
func (t *putDetailsTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse putDetailsResponse
	if err = json.Unmarshal(r.Body(), &theResponse); err != nil {
		return
	}

	cool = theResponse.Cool

	nothing = theResponse.Nothing

	id = ptr(r.Header.Peek("X-Auth-ID"))

	return
}

// NewPutDetailsTransport the transport creator for http requests
func NewPutDetailsTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) PutDetailsTransport {
	return &putDetailsTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

type getSomeElseDataUtf8Response struct {
	Cool    v1.Detail    `json:"cool"`
	ID      *string      `json:"id"`
	Nothing v1.Namespace `json:"TheNothing"`
}

// GetSomeElseDataUtf8Transport transport interface
type GetSomeElseDataUtf8Transport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, token *string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
}

type getSomeElseDataUtf8Transport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getSomeElseDataUtf8Transport) EncodeRequest(ctx context.Context, r *fasthttp.Request, token *string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)

	r.Header.Set("Authorization", *token)

	r.Header.Set("Content-Type", "application/json")

	return
}

// DecodeResponse method for decoding response on server side
func (t *getSomeElseDataUtf8Transport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse getSomeElseDataUtf8Response
	if err = json.Unmarshal(r.Body(), &theResponse); err != nil {
		return
	}

	cool = theResponse.Cool

	id = theResponse.ID

	nothing = theResponse.Nothing

	return
}

// NewGetSomeElseDataUtf8Transport the transport creator for http requests
func NewGetSomeElseDataUtf8Transport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) GetSomeElseDataUtf8Transport {
	return &getSomeElseDataUtf8Transport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

type getFileResponse struct {
	Data     []byte
	FileName string
}

// GetFileTransport transport interface
type GetFileTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, token *string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (data []byte, fileName string, err error)
}

type getFileTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getFileTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, token *string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)

	r.Header.Set("Authorization", *token)

	return
}

// DecodeResponse method for decoding response on server side
func (t *getFileTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (data []byte, fileName string, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	data = r.Body()
	return
}

// NewGetFileTransport the transport creator for http requests
func NewGetFileTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) GetFileTransport {
	return &getFileTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

func ptr(in []byte) *string {
	i := string(in)
	return &i
}
