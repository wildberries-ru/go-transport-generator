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
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientCreateMultipartUpload,
		httpMethodCreateMultipartUpload,
	)
	transportUploadPartDocument := NewUploadPartDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientUploadPartDocument,
		httpMethodUploadPartDocument,
	)
	transportCompleteUpload := NewCompleteUploadTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientCompleteUpload,
		httpMethodCompleteUpload,
	)
	transportUploadDocument := NewUploadDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientUploadDocument,
		httpMethodUploadDocument,
	)
	transportDownloadDocument := NewDownloadDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientDownloadDocument,
		httpMethodDownloadDocument,
	)
	transportGetToken := NewGetTokenTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetToken,
		httpMethodGetToken,
	)
	transportGetBranches := NewGetBranchesTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetBranches,
		httpMethodGetBranches,
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
		transportGetBranches,
		opts,
	}

	type args struct {
		config         Config
		errorProcessor errorProcessor
		options        map[interface{}]Option
	}
	tests := []struct {
		name       string
		args       args
		wantClient Service
		wantErr    bool
	}{
		{"test new builder", args{Config{ServerURL: serverURL, MaxConns: &maxConns}, &testErrorProcessor{}, opts}, &cl, false},
		{"test new builder incorrect URL", args{Config{ServerURL: " http:example%20.com", MaxConns: &maxConns}, &testErrorProcessor{}, opts}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient, err := New(tt.args.config, tt.args.errorProcessor, tt.args.options)
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
