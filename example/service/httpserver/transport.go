// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"bytes"
	"net/http"
	"strconv"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
	"github.com/valyala/fasthttp"
)

var (
	emptyBytes = []byte("")
)

//easyjson:json
type getWarehousesResponse struct {
	Pets []v1.Detail
}

// GetWarehousesTransport transport interface
type GetWarehousesTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, pets []v1.Detail) (err error)
}

type getWarehousesTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getWarehousesTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (err error) {
	return
}

// EncodeResponse method for encoding response on server side
func (t *getWarehousesTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, pets []v1.Detail) (err error) {
	r.Header.Set("Content-Type", "application/json")
	var theResponse getWarehousesResponse
	theResponse.Pets = pets
	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)
	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewGetWarehousesTransport the transport creator for http requests
func NewGetWarehousesTransport(encodeJSONErrorCreator errorCreator) GetWarehousesTransport {
	return &getWarehousesTransport{
		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

//easyjson:json
type getDetailsResponse struct {
	Det v1.Detail
	Ns  v1.Namespace
}

// GetDetailsTransport transport interface
type GetDetailsTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, fileID uint32, someID *uint64, token *string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, det v1.Detail, ns v1.Namespace, id *string) (err error)
}

type getDetailsTransport struct {
	encodeJSONErrorCreator         errorCreator
	encodeQueryTypeIntErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getDetailsTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, fileID uint32, someID *uint64, token *string, err error) {
	namespace = ctx.UserValue("namespace").(string)
	detail = ctx.UserValue("detail").(string)
	_fileID, err := atoi(ctx.QueryArgs().Peek("file"))
	if err != nil {
		err = t.encodeQueryTypeIntErrorCreator(err)
		return
	}
	fileID = uint32(_fileID)
	_someID, err := atoi(ctx.QueryArgs().Peek("some"))
	if err != nil {
		err = t.encodeQueryTypeIntErrorCreator(err)
		return
	}
	__someID := uint64(_someID)
	someID = &__someID

	token = ptr(r.Header.Peek("X-Auth-Token"))
	return
}

// EncodeResponse method for encoding response on server side
func (t *getDetailsTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, det v1.Detail, ns v1.Namespace, id *string) (err error) {
	r.Header.Set("Content-Type", "application/json")
	var theResponse getDetailsResponse
	theResponse.Det = det
	theResponse.Ns = ns
	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)
	r.Header.Set("X-Auth-ID", "id")
	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewGetDetailsTransport the transport creator for http requests
func NewGetDetailsTransport(encodeJSONErrorCreator errorCreator, encodeQueryTypeIntErrorCreator errorCreator) GetDetailsTransport {
	return &getDetailsTransport{
		encodeJSONErrorCreator:         encodeJSONErrorCreator,
		encodeQueryTypeIntErrorCreator: encodeQueryTypeIntErrorCreator,
	}
}

//easyjson:json
type getDetailsEmbedStructResponse struct {
	v1.GetDetailsEmbedStructResponse
}

// GetDetailsEmbedStructTransport transport interface
type GetDetailsEmbedStructTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.GetDetailsEmbedStructResponse) (err error)
}

type getDetailsEmbedStructTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getDetailsEmbedStructTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, err error) {
	namespace = ctx.UserValue("namespace").(string)
	detail = ctx.UserValue("detail").(string)
	return
}

// EncodeResponse method for encoding response on server side
func (t *getDetailsEmbedStructTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.GetDetailsEmbedStructResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	var theResponse getDetailsEmbedStructResponse
	theResponse.GetDetailsEmbedStructResponse = response
	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)
	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewGetDetailsEmbedStructTransport the transport creator for http requests
func NewGetDetailsEmbedStructTransport(encodeJSONErrorCreator errorCreator) GetDetailsEmbedStructTransport {
	return &getDetailsEmbedStructTransport{
		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

//easyjson:json
type getDetailsListEmbedStructResponse []v1.Detail

// GetDetailsListEmbedStructTransport transport interface
type GetDetailsListEmbedStructTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, details []v1.Detail) (err error)
}

type getDetailsListEmbedStructTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getDetailsListEmbedStructTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, err error) {
	namespace = ctx.UserValue("namespace").(string)
	detail = ctx.UserValue("detail").(string)
	return
}

// EncodeResponse method for encoding response on server side
func (t *getDetailsListEmbedStructTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, details []v1.Detail) (err error) {
	r.Header.Set("Content-Type", "application/json")
	var theResponse getDetailsListEmbedStructResponse
	theResponse = details
	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)
	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewGetDetailsListEmbedStructTransport the transport creator for http requests
func NewGetDetailsListEmbedStructTransport(encodeJSONErrorCreator errorCreator) GetDetailsListEmbedStructTransport {
	return &getDetailsListEmbedStructTransport{
		encodeJSONErrorCreator: encodeJSONErrorCreator,
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
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, cool v1.Detail, nothing v1.Namespace, id *string) (err error)
}

type putDetailsTransport struct {
	decodeJSONErrorCreator errorCreator
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *putDetailsTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace, err error) {
	namespace = ctx.UserValue("namespace").(string)
	detail = ctx.UserValue("detail").(string)
	blaID = ptr(ctx.QueryArgs().Peek("bla"))
	testID = string(ctx.QueryArgs().Peek("test"))
	token = ptr(r.Header.Peek("X-Auth-Token"))
	var request putDetailsRequest
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		err = t.decodeJSONErrorCreator(err)
		return
	}
	pretty = request.Pretty
	yang = request.Yang
	return
}

// EncodeResponse method for encoding response on server side
func (t *putDetailsTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, cool v1.Detail, nothing v1.Namespace, id *string) (err error) {
	r.Header.Set("Content-Type", "application/json")
	var theResponse putDetailsResponse
	theResponse.Cool = cool
	theResponse.Nothing = nothing
	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)
	r.Header.Set("X-Auth-ID", "id")
	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewPutDetailsTransport the transport creator for http requests
func NewPutDetailsTransport(decodeJSONErrorCreator errorCreator, encodeJSONErrorCreator errorCreator) PutDetailsTransport {
	return &putDetailsTransport{
		decodeJSONErrorCreator: decodeJSONErrorCreator,
		encodeJSONErrorCreator: encodeJSONErrorCreator,
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
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, cool v1.Detail, nothing v1.Namespace, id *string) (err error)
}

type getSomeElseDataUtf8Transport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getSomeElseDataUtf8Transport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (err error) {
	return
}

// EncodeResponse method for encoding response on server side
func (t *getSomeElseDataUtf8Transport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, cool v1.Detail, nothing v1.Namespace, id *string) (err error) {
	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	var theResponse getSomeElseDataUtf8Response
	theResponse.Cool = cool
	theResponse.Id = id
	theResponse.Nothing = nothing
	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)
	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewGetSomeElseDataUtf8Transport the transport creator for http requests
func NewGetSomeElseDataUtf8Transport(encodeJSONErrorCreator errorCreator) GetSomeElseDataUtf8Transport {
	return &getSomeElseDataUtf8Transport{
		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

func ptr(in []byte) *string {
	i := string(in)
	return &i
}

func atoiptr(in []byte) (out *int, err error) {
	var (
		o int
		i = string(in)
	)
	if i != "" {
		if o, err = strconv.Atoi(i); err == nil {
			out = &o
		}
	}
	return
}

func atoi(in []byte) (out int, err error) {
	if bytes.Equal(in, emptyBytes) {
		return
	}
	return strconv.Atoi(string(in))
}
