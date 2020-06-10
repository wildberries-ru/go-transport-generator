// Package httperrors ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httperrors

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"
)

type httpError struct {
	Code    int
	Message string
}

// Error returns a text message corresponding to the given error.
func (e *httpError) Error() string {
	return e.Message
}

// StatusCode returns an HTTP status code corresponding to the given error.
func (e *httpError) StatusCode() int {
	return e.Code
}

type responseWithError struct {
	Error            bool              `json:"error"`
	ErrorText        string            `json:"errorText"`
	Data             *bool             `json:"data"`
	AdditionalErrors map[string]string `json:"additionalErrors"`
}

// UIErrorProcessor ...
type UIErrorProcessor struct {
	defaultCode    int
	defaultMessage string
	errors         map[string]string
	errDefault     string
}

// Encode writes a svc error to the given http.ResponseWriter.
func (e *UIErrorProcessor) Encode(ctx context.Context, r *fasthttp.Response, err error) {
	numberOfError := err.Error()[:strings.Index(err.Error(), ":")]
	errorText, ok := e.errors[numberOfError]
	if !ok {
		errorText = e.errDefault
	}
	res := responseWithError{
		Error:     true,
		ErrorText: errorText,
	}
	r.SetStatusCode(http.StatusOK)
	r.Header.Set("Content-Type", "application/json")
	body, err := json.Marshal(res)
	if err != nil {
		return
	}
	r.SetBody(body)
	return
}

// Decode reads a Service error from the given *http.Response.
func (e *UIErrorProcessor) Decode(r *fasthttp.Response) error {
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

// NewUIErrorProcessor ...
func NewUIErrorProcessor(defaultCode int, defaultMessage string, errors map[string]string, errDefault string) *UIErrorProcessor {
	return &UIErrorProcessor{
		defaultCode:    defaultCode,
		defaultMessage: defaultMessage,
		errors:         errors,
		errDefault:     errDefault,
	}
}
