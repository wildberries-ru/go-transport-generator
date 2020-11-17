// Package httperrors ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httperrors

import (
	"context"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"
)

// ClientErrorProcessor ...
type ClientErrorProcessor struct {
	defaultCode    int
	defaultMessage string
}

// Encode writes a svc error to the given http.ResponseWriter.
func (e *ClientErrorProcessor) Encode(ctx context.Context, r *fasthttp.Response, err error) {
	code := e.defaultCode
	message := e.defaultMessage
	if err, ok := err.(*httpError); ok {
		if err.Code != e.defaultCode {
			code = err.Code
			message = err.Message
		}
	}
	r.SetStatusCode(code)
	r.SetBodyString(message)
	return
}

// Decode reads a Service error from the given *http.Response.
func (e *ClientErrorProcessor) Decode(r *fasthttp.Response) error {
	msgBytes := r.Body()
	msg := strings.TrimSpace(string(msgBytes))
	if msg == "" {
		msg = http.StatusText(r.StatusCode())
	}
	return &httpError{
		Code:    r.StatusCode(),
		Message: msg,
	}
}

// NewClientErrorProcessor ...
func NewClientErrorProcessor(defaultCode int, defaultMessage string) *ClientErrorProcessor {
	return &ClientErrorProcessor{
		defaultCode:    defaultCode,
		defaultMessage: defaultMessage,
	}
}
