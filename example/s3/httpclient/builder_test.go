// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

type testErrorProcessor struct{}

func TestNew(t *testing.T) {
	serverURL := fmt.Sprintf("https://%v.com", time.Now().UnixNano())
	parsedServerURL, _ := url.Parse(serverURL)
	maxConns := rand.Int()
	opts := map[interface{}]Option{}

	transportCreateMultipartUpload := NewCreateMultipartUploadTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientCreateMultipartUpload,
		httpMethodCreateMultipartUpload,
	)
	transportUploadPartDocument := NewUploadPartDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientUploadPartDocument,
		httpMethodUploadPartDocument,
	)
	transportCompleteUpload := NewCompleteUploadTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientCompleteUpload,
		httpMethodCompleteUpload,
	)
	transportUploadDocument := NewUploadDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientUploadDocument,
		httpMethodUploadDocument,
	)
	transportDownloadDocument := NewDownloadDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientDownloadDocument,
		httpMethodDownloadDocument,
	)
	transportGetToken := NewGetTokenTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetToken,
		httpMethodGetToken,
	)

	cl := client{
		&fasthttp.HostClient{
			Addr:     parsedServerURL.Host,
			MaxConns: maxConns,
		},
		transportCreateMultipartUpload,
		transportUploadPartDocument,
		transportCompleteUpload,
		transportUploadDocument,
		transportDownloadDocument,
		transportGetToken,
		opts,
	}

	type args struct {
		serverURL      string
		maxConns       int
		errorProcessor errorProcessor
		options        map[interface{}]Option
	}
	tests := []struct {
		name       string
		args       args
		wantClient Service
		wantErr    bool
	}{
		{"test new builder", args{serverURL, maxConns, &testErrorProcessor{}, opts}, &cl, false},
		{"test new builder incorrect URL", args{" http:example%20.com", maxConns, &testErrorProcessor{}, opts}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient, err := New(tt.args.serverURL, tt.args.maxConns, tt.args.errorProcessor, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotClient, tt.wantClient) {
				t.Errorf("New() = %v, want %v", gotClient, tt.wantClient)
			}
		})
	}
}

func (ep *testErrorProcessor) Decode(r *fasthttp.Response) error {
	return nil
}
