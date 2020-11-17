// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/valyala/fasthttp"
)

func Test_client_TestEasyJson(t *testing.T) {

	var param1 int
	_ = faker.FakeData(&param1)

	var field1 string
	_ = faker.FakeData(&field1)

	var field2 string
	_ = faker.FakeData(&field2)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
			Field1 string `json:"field1"`

			Field2 string `json:"field2"`
		}{

			Field1: field1,

			Field2: field2,
		}

		b, _ := json.Marshal(result)
		w.Write(b)
	}))
	defer ts.Close()

	parsedServerURL, _ := url.Parse(ts.URL)

	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: maxConns,
	}

	transportTestEasyJson := NewTestEasyJsonTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientTestEasyJson,
		httpMethodTestEasyJson,
	)

	type fields struct {
		cli                   *fasthttp.HostClient
		transportTestEasyJson TestEasyJsonTransport
		options               map[interface{}]Option
	}
	type args struct {
		ctx    context.Context
		param1 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantField1 string

		wantField2 string

		wantErr bool
	}{
		{
			"test TestEasyJson",
			fields{hostClient, transportTestEasyJson, opts},
			args{context.Background(), param1},
			field1,
			field2,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                   tt.fields.cli,
				transportTestEasyJson: tt.fields.transportTestEasyJson,
				options:               tt.fields.options,
			}
			gotField1, gotField2, err := s.TestEasyJson(tt.args.ctx, tt.args.param1)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.TestEasyJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotField1, tt.wantField1) {
				t.Errorf("client.field1() = %v, want %v", gotField1, tt.wantField1)
			}

			if !reflect.DeepEqual(gotField2, tt.wantField2) {
				t.Errorf("client.field2() = %v, want %v", gotField2, tt.wantField2)
			}

		})
	}
}

func TestNewClient(t *testing.T) {
	serverURL := fmt.Sprintf("https://%v.com", time.Now().UnixNano())
	parsedServerURL, _ := url.Parse(serverURL)
	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: rand.Int(),
	}
	opts := map[interface{}]Option{}

	transportTestEasyJson := NewTestEasyJsonTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientTestEasyJson,
		httpMethodTestEasyJson,
	)

	cl := &client{
		hostClient,
		transportTestEasyJson,
		opts,
	}

	type args struct {
		cli *fasthttp.HostClient

		transportTestEasyJson TestEasyJsonTransport

		options map[interface{}]Option
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"test new client", args{hostClient, transportTestEasyJson, opts}, cl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.cli, tt.args.transportTestEasyJson, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
