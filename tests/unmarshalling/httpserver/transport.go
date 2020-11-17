// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"
)

var (
	emptyBytes = []byte("")
)

//easyjson:json
type testEasyJsonResponse struct {
	Field1 string `json:"field1"`

	Field2 string `json:"field2"`
}

// TestEasyJsonTransport transport interface
type TestEasyJsonTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (param1 int, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, field1 string, field2 string) (err error)
}

type testEasyJsonTransport struct {
	encodeJSONErrorCreator    errorCreator
	decodeTypeIntErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *testEasyJsonTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (param1 int, err error) {

	_param1 := ctx.QueryArgs().Peek("param1")
	if !bytes.Equal(_param1, emptyBytes) {
		var i int
		i, err = strconv.Atoi(string(_param1))
		if err != nil {
			err = t.decodeTypeIntErrorCreator(err)
			return
		}

		param1 = i

	}

	return
}

// EncodeResponse method for encoding response on server side
func (t *testEasyJsonTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, field1 string, field2 string) (err error) {

	r.Header.Set("Content-Type", "application/json")
	var theResponse testEasyJsonResponse

	theResponse.Field1 = field1

	theResponse.Field2 = field2

	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)

	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewTestEasyJsonTransport the transport creator for http requests
func NewTestEasyJsonTransport(

	encodeJSONErrorCreator errorCreator,

	decodeTypeIntErrorCreator errorCreator,
) TestEasyJsonTransport {
	return &testEasyJsonTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,

		decodeTypeIntErrorCreator: decodeTypeIntErrorCreator,
	}
}

func ptr(in []byte) *string {
	if bytes.Equal(in, emptyBytes) {
		return nil
	}
	i := string(in)
	return &i
}
