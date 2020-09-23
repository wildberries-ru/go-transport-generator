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
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

func Test_client_CreateMultipartUpload(t *testing.T) {

	var bucket string
	_ = faker.FakeData(&bucket)

	var key string
	_ = faker.FakeData(&key)

	var data v1.CreateMultipartUploadData
	_ = faker.FakeData(&data)

	var errorFlag bool
	_ = faker.FakeData(&errorFlag)

	var errorText string
	_ = faker.FakeData(&errorText)

	var additionalErrors *v1.AdditionalErrors
	_ = faker.FakeData(&additionalErrors)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
			Data v1.CreateMultipartUploadData `json:"Data"`

			ErrorFlag bool `json:"ErrorFlag"`

			ErrorText string `json:"ErrorText"`

			AdditionalErrors *v1.AdditionalErrors `json:"AdditionalErrors"`
		}{

			Data: data,

			ErrorFlag: errorFlag,

			ErrorText: errorText,

			AdditionalErrors: additionalErrors,
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

	transportCreateMultipartUpload := NewCreateMultipartUploadTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientCreateMultipartUpload,
		httpMethodCreateMultipartUpload,
	)

	type fields struct {
		cli                            *fasthttp.HostClient
		transportCreateMultipartUpload CreateMultipartUploadTransport
		options                        map[interface{}]Option
	}
	type args struct {
		ctx    context.Context
		bucket string
		key    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantData v1.CreateMultipartUploadData

		wantErrorFlag bool

		wantErrorText string

		wantAdditionalErrors *v1.AdditionalErrors

		wantErr bool
	}{
		{
			"test CreateMultipartUpload",
			fields{hostClient, transportCreateMultipartUpload, opts},
			args{context.Background(), bucket, key},
			data,
			errorFlag,
			errorText,
			additionalErrors,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                            tt.fields.cli,
				transportCreateMultipartUpload: tt.fields.transportCreateMultipartUpload,
				options:                        tt.fields.options,
			}
			gotData, gotErrorFlag, gotErrorText, gotAdditionalErrors, err := s.CreateMultipartUpload(tt.args.ctx, tt.args.bucket, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CreateMultipartUpload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("client.data() = %v, want %v", gotData, tt.wantData)
			}

			if !reflect.DeepEqual(gotErrorFlag, tt.wantErrorFlag) {
				t.Errorf("client.errorFlag() = %v, want %v", gotErrorFlag, tt.wantErrorFlag)
			}

			if !reflect.DeepEqual(gotErrorText, tt.wantErrorText) {
				t.Errorf("client.errorText() = %v, want %v", gotErrorText, tt.wantErrorText)
			}

			if !reflect.DeepEqual(gotAdditionalErrors, tt.wantAdditionalErrors) {
				t.Errorf("client.additionalErrors() = %v, want %v", gotAdditionalErrors, tt.wantAdditionalErrors)
			}

		})
	}
}

func Test_client_UploadPartDocument(t *testing.T) {

	var bucket string
	_ = faker.FakeData(&bucket)

	var key string
	_ = faker.FakeData(&key)

	var uploadID string
	_ = faker.FakeData(&uploadID)

	var partNumber int64
	_ = faker.FakeData(&partNumber)

	var document []byte
	_ = faker.FakeData(&document)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
		}{}

		b, _ := json.Marshal(result)
		w.Write(b)
	}))
	defer ts.Close()

	parsedServerURL, _ := url.Parse(ts.URL)

	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: maxConns,
	}

	transportUploadPartDocument := NewUploadPartDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientUploadPartDocument,
		httpMethodUploadPartDocument,
	)

	type fields struct {
		cli                         *fasthttp.HostClient
		transportUploadPartDocument UploadPartDocumentTransport
		options                     map[interface{}]Option
	}
	type args struct {
		ctx        context.Context
		bucket     string
		key        string
		uploadID   string
		partNumber int64
		document   []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantErr bool
	}{
		{
			"test UploadPartDocument",
			fields{hostClient, transportUploadPartDocument, opts},
			args{context.Background(), bucket, key, uploadID, partNumber, document},

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                         tt.fields.cli,
				transportUploadPartDocument: tt.fields.transportUploadPartDocument,
				options:                     tt.fields.options,
			}
			err := s.UploadPartDocument(tt.args.ctx, tt.args.bucket, tt.args.key, tt.args.uploadID, tt.args.partNumber, tt.args.document)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UploadPartDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_client_CompleteUpload(t *testing.T) {

	var bucket string
	_ = faker.FakeData(&bucket)

	var key string
	_ = faker.FakeData(&key)

	var uploadID string
	_ = faker.FakeData(&uploadID)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
		}{}

		b, _ := json.Marshal(result)
		w.Write(b)
	}))
	defer ts.Close()

	parsedServerURL, _ := url.Parse(ts.URL)

	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: maxConns,
	}

	transportCompleteUpload := NewCompleteUploadTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientCompleteUpload,
		httpMethodCompleteUpload,
	)

	type fields struct {
		cli                     *fasthttp.HostClient
		transportCompleteUpload CompleteUploadTransport
		options                 map[interface{}]Option
	}
	type args struct {
		ctx      context.Context
		bucket   string
		key      string
		uploadID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantErr bool
	}{
		{
			"test CompleteUpload",
			fields{hostClient, transportCompleteUpload, opts},
			args{context.Background(), bucket, key, uploadID},

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                     tt.fields.cli,
				transportCompleteUpload: tt.fields.transportCompleteUpload,
				options:                 tt.fields.options,
			}
			err := s.CompleteUpload(tt.args.ctx, tt.args.bucket, tt.args.key, tt.args.uploadID)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CompleteUpload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_client_UploadDocument(t *testing.T) {

	var bucket string
	_ = faker.FakeData(&bucket)

	var key string
	_ = faker.FakeData(&key)

	var document []byte
	_ = faker.FakeData(&document)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
		}{}

		b, _ := json.Marshal(result)
		w.Write(b)
	}))
	defer ts.Close()

	parsedServerURL, _ := url.Parse(ts.URL)

	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: maxConns,
	}

	transportUploadDocument := NewUploadDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientUploadDocument,
		httpMethodUploadDocument,
	)

	type fields struct {
		cli                     *fasthttp.HostClient
		transportUploadDocument UploadDocumentTransport
		options                 map[interface{}]Option
	}
	type args struct {
		ctx      context.Context
		bucket   string
		key      string
		document []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantErr bool
	}{
		{
			"test UploadDocument",
			fields{hostClient, transportUploadDocument, opts},
			args{context.Background(), bucket, key, document},

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                     tt.fields.cli,
				transportUploadDocument: tt.fields.transportUploadDocument,
				options:                 tt.fields.options,
			}
			err := s.UploadDocument(tt.args.ctx, tt.args.bucket, tt.args.key, tt.args.document)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UploadDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_client_DownloadDocument(t *testing.T) {

	var bucket string
	_ = faker.FakeData(&bucket)

	var key string
	_ = faker.FakeData(&key)

	var document []byte
	_ = faker.FakeData(&document)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := document

		b, _ := json.Marshal(result)
		w.Write(b)
	}))
	defer ts.Close()

	parsedServerURL, _ := url.Parse(ts.URL)

	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: maxConns,
	}

	transportDownloadDocument := NewDownloadDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientDownloadDocument,
		httpMethodDownloadDocument,
	)

	type fields struct {
		cli                       *fasthttp.HostClient
		transportDownloadDocument DownloadDocumentTransport
		options                   map[interface{}]Option
	}
	type args struct {
		ctx    context.Context
		bucket string
		key    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantDocument []byte

		wantErr bool
	}{
		{
			"test DownloadDocument",
			fields{hostClient, transportDownloadDocument, opts},
			args{context.Background(), bucket, key},
			document,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                       tt.fields.cli,
				transportDownloadDocument: tt.fields.transportDownloadDocument,
				options:                   tt.fields.options,
			}
			gotDocument, err := s.DownloadDocument(tt.args.ctx, tt.args.bucket, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.DownloadDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotDocument, tt.wantDocument) {
				t.Errorf("client.document() = %v, want %v", gotDocument, tt.wantDocument)
			}

		})
	}
}

func Test_client_GetToken(t *testing.T) {

	var authToken *string
	_ = faker.FakeData(&authToken)

	var scope string
	_ = faker.FakeData(&scope)

	var grantType string
	_ = faker.FakeData(&grantType)

	var token string
	_ = faker.FakeData(&token)

	var expiresIn int
	_ = faker.FakeData(&expiresIn)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
			Token string `json:"Token"`

			ExpiresIn int `json:"ExpiresIn"`
		}{

			Token: token,

			ExpiresIn: expiresIn,
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

	transportGetToken := NewGetTokenTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetToken,
		httpMethodGetToken,
	)

	type fields struct {
		cli               *fasthttp.HostClient
		transportGetToken GetTokenTransport
		options           map[interface{}]Option
	}
	type args struct {
		ctx       context.Context
		authToken *string
		scope     string
		grantType string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantToken string

		wantExpiresIn int

		wantErr bool
	}{
		{
			"test GetToken",
			fields{hostClient, transportGetToken, opts},
			args{context.Background(), authToken, scope, grantType},
			token,
			expiresIn,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:               tt.fields.cli,
				transportGetToken: tt.fields.transportGetToken,
				options:           tt.fields.options,
			}
			gotToken, gotExpiresIn, err := s.GetToken(tt.args.ctx, tt.args.authToken, tt.args.scope, tt.args.grantType)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotToken, tt.wantToken) {
				t.Errorf("client.token() = %v, want %v", gotToken, tt.wantToken)
			}

			if !reflect.DeepEqual(gotExpiresIn, tt.wantExpiresIn) {
				t.Errorf("client.expiresIn() = %v, want %v", gotExpiresIn, tt.wantExpiresIn)
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

	cl := &client{
		hostClient,
		transportCreateMultipartUpload,
		transportUploadPartDocument,
		transportCompleteUpload,
		transportUploadDocument,
		transportDownloadDocument,
		transportGetToken,
		opts,
	}

	type args struct {
		cli *fasthttp.HostClient

		transportCreateMultipartUpload CreateMultipartUploadTransport

		transportUploadPartDocument UploadPartDocumentTransport

		transportCompleteUpload CompleteUploadTransport

		transportUploadDocument UploadDocumentTransport

		transportDownloadDocument DownloadDocumentTransport

		transportGetToken GetTokenTransport

		options map[interface{}]Option
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"test new client", args{hostClient, transportCreateMultipartUpload, transportUploadPartDocument, transportCompleteUpload, transportUploadDocument, transportDownloadDocument, transportGetToken, opts}, cl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.cli, tt.args.transportCreateMultipartUpload, tt.args.transportUploadPartDocument, tt.args.transportCompleteUpload, tt.args.transportUploadDocument, tt.args.transportDownloadDocument, tt.args.transportGetToken, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
