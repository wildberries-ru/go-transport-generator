// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"
)

//easyjson:json
type testEasyJsonResponse struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

// TestEasyJsonTransport transport interface
type TestEasyJsonTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, param1 int) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (field1 string, field2 string, err error)
}

type testEasyJsonTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *testEasyJsonTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, param1 int) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)

	r.URI().QueryArgs().Set("param1", strconv.Itoa(param1))

	r.Header.Set("Content-Type", "application/json")

	return
}

// DecodeResponse method for encoding response on server side
func (t *testEasyJsonTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (field1 string, field2 string, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse testEasyJsonResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
		return
	}

	field1 = theResponse.Field1

	field2 = theResponse.Field2

	return
}

// NewTestEasyJsonTransport the transport creator for http requests
func NewTestEasyJsonTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) TestEasyJsonTransport {
	return &testEasyJsonTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

func ptr(in []byte) *string {
	i := string(in)
	return &i
}
