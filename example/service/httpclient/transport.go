// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

//easyjson:json
type getWarehousesResponse struct {
	Pets []v1.Detail
}

// GetWarehousesTransport transport interface
type GetWarehousesTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (pets []v1.Detail, err error)
}

type getWarehousesTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getWarehousesTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	return
}

// DecodeResponse method for encoding response on server side
func (t *getWarehousesTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (pets []v1.Detail, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	var theResponse getWarehousesResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
		return
	}

	pets = theResponse.Pets
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

//easyjson:json
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

// DecodeResponse method for encoding response on server side
func (t *getDetailsTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (det v1.Detail, ns v1.Namespace, id *string, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	var theResponse getDetailsResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
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

//easyjson:json
type getDetailsEmbedStructResponse struct {
	v1.GetDetailsEmbedStructResponse
}

// GetDetailsEmbedStructTransport transport interface
type GetDetailsEmbedStructTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.GetDetailsEmbedStructResponse, err error)
}

type getDetailsEmbedStructTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getDetailsEmbedStructTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, namespace, detail))
	r.Header.Set("Content-Type", "application/json")
	return
}

// DecodeResponse method for encoding response on server side
func (t *getDetailsEmbedStructTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.GetDetailsEmbedStructResponse, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	var theResponse getDetailsEmbedStructResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
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

//easyjson:json
type getDetailsListEmbedStructResponse []v1.Detail

// GetDetailsListEmbedStructTransport transport interface
type GetDetailsListEmbedStructTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (details []v1.Detail, err error)
}

type getDetailsListEmbedStructTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getDetailsListEmbedStructTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, namespace string, detail string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, namespace, detail))
	r.Header.Set("Content-Type", "application/json")
	return
}

// DecodeResponse method for encoding response on server side
func (t *getDetailsListEmbedStructTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (details []v1.Detail, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	var theResponse getDetailsListEmbedStructResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
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

//easyjson:json
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
	body, err := request.MarshalJSON()
	if err != nil {
		return
	}
	r.SetBody(body)
	return
}

// DecodeResponse method for encoding response on server side
func (t *putDetailsTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	var theResponse putDetailsResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
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

//easyjson:json
type getSomeElseDataUtf8Response struct {
	Cool    v1.Detail    `json:"cool"`
	Id      *string      `json:"id"`
	Nothing v1.Namespace `json:"TheNothing"`
}

// GetSomeElseDataUtf8Transport transport interface
type GetSomeElseDataUtf8Transport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
}

type getSomeElseDataUtf8Transport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getSomeElseDataUtf8Transport) EncodeRequest(ctx context.Context, r *fasthttp.Request) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	return
}

// DecodeResponse method for encoding response on server side
func (t *getSomeElseDataUtf8Transport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	var theResponse getSomeElseDataUtf8Response
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
		return
	}

	cool = theResponse.Cool
	id = theResponse.Id
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

func ptr(in []byte) *string {
	i := string(in)
	return &i
}
