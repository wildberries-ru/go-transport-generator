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

	transportTestEasyJson := NewTestEasyJsonTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientTestEasyJson,
		httpMethodTestEasyJson,
	)

	cl := client{
		&fasthttp.HostClient{
			Addr:     parsedServerURL.Host,
			MaxConns: maxConns,
		},
		transportTestEasyJson,
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
