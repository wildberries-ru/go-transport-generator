// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

var (
	emptyBytes = []byte("")
)

// UploadDocumentTransport transport interface
type UploadDocumentTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (token *string, name string, extension string, categoryID string, supplierID *int64, contractID *int64, data multipart.File, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error)
}

type uploadDocumentTransport struct {
	decodeTypeIntErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *uploadDocumentTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (token *string, name string, extension string, categoryID string, supplierID *int64, contractID *int64, data multipart.File, err error) {

	token = ptr(r.Header.Peek("Authorization"))

	var form *multipart.Form
	if form, err = ctx.MultipartForm(); err != nil {
		err = errors.Wrap(err, "failed to read MultipartForm")
		return
	}

	_categoryID := form.Value["categoryID"]

	if len(_categoryID) != 1 {
		err = errors.New("failed to read categoryID in MultipartForm")
		return
	}

	categoryID = _categoryID[0]

	_contractID := form.Value["contractID"]

	if len(_contractID) == 1 {

		_ContractID := _contractID[0]
		if _ContractID != "" {
			var i int
			i, err = strconv.Atoi(_ContractID)
			if err != nil {
				err = t.decodeTypeIntErrorCreator(err)
				return
			}

			ii := int64(i)
			contractID = &ii

		}

	}

	_extension := form.Value["extension"]

	if len(_extension) != 1 {
		err = errors.New("failed to read extension in MultipartForm")
		return
	}

	extension = _extension[0]

	_name := form.Value["name"]

	if len(_name) != 1 {
		err = errors.New("failed to read name in MultipartForm")
		return
	}

	name = _name[0]

	_supplierID := form.Value["supplierID"]

	if len(_supplierID) == 1 {

		_SupplierID := _supplierID[0]
		if _SupplierID != "" {
			var i int
			i, err = strconv.Atoi(_SupplierID)
			if err != nil {
				err = t.decodeTypeIntErrorCreator(err)
				return
			}

			ii := int64(i)
			supplierID = &ii

		}

	}

	_data := form.File["document"]
	if _data == nil {
		err = errors.New("failed to read document in MultipartForm")
		return
	}
	if len(_data) != 1 {
		err = errors.New("failed to read file in MultipartForm: too many files in document")
		return
	}
	data, err = _data[0].Open()
	if err != nil {
		err = errors.Wrap(err, "failed to open file document in MultipartForm")
		return
	}

	return
}

// EncodeResponse method for encoding response on server side
func (t *uploadDocumentTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error) {

	r.Header.Set("Content-Type", "application/json")

	r.Header.SetStatusCode(201)
	return
}

// NewUploadDocumentTransport the transport creator for http requests
func NewUploadDocumentTransport(

	decodeTypeIntErrorCreator errorCreator,
) UploadDocumentTransport {
	return &uploadDocumentTransport{

		decodeTypeIntErrorCreator: decodeTypeIntErrorCreator,
	}
}

//easyjson:json
type getWarehousesResponse map[string]v1.Detail

// GetWarehousesTransport transport interface
type GetWarehousesTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, pets map[string]v1.Detail) (err error)
}

type getWarehousesTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getWarehousesTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (err error) {

	return
}

// EncodeResponse method for encoding response on server side
func (t *getWarehousesTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, pets map[string]v1.Detail) (err error) {

	r.Header.Set("Content-Type", "application/json")
	var theResponse getWarehousesResponse

	theResponse = pets

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
func NewGetWarehousesTransport(

	encodeJSONErrorCreator errorCreator,

) GetWarehousesTransport {
	return &getWarehousesTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

//easyjson:json
type getDetailsResponse struct {
	Det v1.Detail

	Ns v1.Namespace
}

// GetDetailsTransport transport interface
type GetDetailsTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, fileID uint32, someID *uint64, token *string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, det v1.Detail, ns v1.Namespace, id *string) (err error)
}

type getDetailsTransport struct {
	encodeJSONErrorCreator    errorCreator
	decodeTypeIntErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getDetailsTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (namespace string, detail string, fileID uint32, someID *uint64, token *string, err error) {

	namespace = ctx.UserValue("namespace").(string)

	detail = ctx.UserValue("detail").(string)

	_fileID := ctx.QueryArgs().Peek("file")
	if !bytes.Equal(_fileID, emptyBytes) {
		var i int
		i, err = strconv.Atoi(string(_fileID))
		if err != nil {
			err = t.decodeTypeIntErrorCreator(err)
			return
		}

		fileID = uint32(i)

	}

	_someID := ctx.QueryArgs().Peek("some")
	if !bytes.Equal(_someID, emptyBytes) {
		var i int
		i, err = strconv.Atoi(string(_someID))
		if err != nil {
			err = t.decodeTypeIntErrorCreator(err)
			return
		}

		ii := uint64(i)
		someID = &ii

	}

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
func NewGetDetailsTransport(

	encodeJSONErrorCreator errorCreator,

	decodeTypeIntErrorCreator errorCreator,
) GetDetailsTransport {
	return &getDetailsTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,

		decodeTypeIntErrorCreator: decodeTypeIntErrorCreator,
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
func NewGetDetailsEmbedStructTransport(

	encodeJSONErrorCreator errorCreator,

) GetDetailsEmbedStructTransport {
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
func NewGetDetailsListEmbedStructTransport(

	encodeJSONErrorCreator errorCreator,

) GetDetailsListEmbedStructTransport {
	return &getDetailsListEmbedStructTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

type putDetailsRequest struct {
	Pretty v1.Detail `json:"ThePretty"`

	Yang v1.Namespace
}

//easyjson:json
type putDetailsResponse struct {
	Cool v1.Detail

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
func NewPutDetailsTransport(

	decodeJSONErrorCreator errorCreator,

	encodeJSONErrorCreator errorCreator,

) PutDetailsTransport {
	return &putDetailsTransport{

		decodeJSONErrorCreator: decodeJSONErrorCreator,

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

//easyjson:json
type getSomeElseDataUtf8Response struct {
	Cool v1.Detail `json:"cool"`

	ID *string `json:"id"`

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

	theResponse.ID = id

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
func NewGetSomeElseDataUtf8Transport(

	encodeJSONErrorCreator errorCreator,

) GetSomeElseDataUtf8Transport {
	return &getSomeElseDataUtf8Transport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

// GetFileTransport transport interface
type GetFileTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, data []byte, fileName string) (err error)
}

type getFileTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *getFileTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (err error) {

	return
}

// EncodeResponse method for encoding response on server side
func (t *getFileTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, data []byte, fileName string) (err error) {

	r.Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	r.SetBody(data)
	if len(fileName) > 0 {
		r.Header.Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	}
	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewGetFileTransport the transport creator for http requests
func NewGetFileTransport() GetFileTransport {
	return &getFileTransport{}
}

func ptr(in []byte) *string {
	if bytes.Equal(in, emptyBytes) {
		return nil
	}
	i := string(in)
	return &i
}
