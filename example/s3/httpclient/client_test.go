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
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
			Data v1.CreateMultipartUploadData `json:"data"`

			ErrorFlag bool `json:"error"`

			ErrorText string `json:"errorText"`

			AdditionalErrors *v1.AdditionalErrors `json:"additionalErrors"`
		}{

			Data: data,

			ErrorFlag: errorFlag,

			ErrorText: errorText,

			AdditionalErrors: additionalErrors,
		}

		b, _ := json.Marshal(result)
		w.WriteHeader(201)
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
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientCreateMultipartUpload,
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

			if !cmp.Equal(gotData, tt.wantData, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.data() = %s", cmp.Diff(gotData, tt.wantData))
			}

			if !cmp.Equal(gotErrorFlag, tt.wantErrorFlag, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.errorFlag() = %s", cmp.Diff(gotErrorFlag, tt.wantErrorFlag))
			}

			if !cmp.Equal(gotErrorText, tt.wantErrorText, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.errorText() = %s", cmp.Diff(gotErrorText, tt.wantErrorText))
			}

			if !cmp.Equal(gotAdditionalErrors, tt.wantAdditionalErrors, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.additionalErrors() = %s", cmp.Diff(gotAdditionalErrors, tt.wantAdditionalErrors))
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
		w.WriteHeader(200)
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
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientUploadPartDocument,
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
		w.WriteHeader(200)
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
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientCompleteUpload,
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
		w.WriteHeader(201)
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
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientUploadDocument,
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
		w.WriteHeader(200)
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
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientDownloadDocument,
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

			if !cmp.Equal(gotDocument, tt.wantDocument, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.document() = %s", cmp.Diff(gotDocument, tt.wantDocument))
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
			Token string `json:"token"`

			ExpiresIn int `json:"expiresIn"`
		}{

			Token: token,

			ExpiresIn: expiresIn,
		}

		b, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
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
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetToken,
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

			if !cmp.Equal(gotToken, tt.wantToken, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.token() = %s", cmp.Diff(gotToken, tt.wantToken))
			}

			if !cmp.Equal(gotExpiresIn, tt.wantExpiresIn, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.expiresIn() = %s", cmp.Diff(gotExpiresIn, tt.wantExpiresIn))
			}

		})
	}
}

func Test_client_GetBranches(t *testing.T) {

	var authToken *string
	_ = faker.FakeData(&authToken)

	var supplierID *string
	_ = faker.FakeData(&supplierID)

	var branches []int
	_ = faker.FakeData(&branches)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := branches

		b, _ := json.Marshal(result)
		w.WriteHeader(200)
		w.Write(b)
	}))
	defer ts.Close()

	parsedServerURL, _ := url.Parse(ts.URL)

	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: maxConns,
	}

	transportGetBranches := NewGetBranchesTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetBranches,
		httpMethodGetBranches,
	)

	type fields struct {
		cli                  *fasthttp.HostClient
		transportGetBranches GetBranchesTransport
		options              map[interface{}]Option
	}
	type args struct {
		ctx        context.Context
		authToken  *string
		supplierID *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantBranches []int

		wantErr bool
	}{
		{
			"test GetBranches",
			fields{hostClient, transportGetBranches, opts},
			args{context.Background(), authToken, supplierID},
			branches,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                  tt.fields.cli,
				transportGetBranches: tt.fields.transportGetBranches,
				options:              tt.fields.options,
			}
			gotBranches, err := s.GetBranches(tt.args.ctx, tt.args.authToken, tt.args.supplierID)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetBranches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(gotBranches, tt.wantBranches, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.branches() = %s", cmp.Diff(gotBranches, tt.wantBranches))
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

	cl := &client{
		hostClient,
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
		cli *fasthttp.HostClient

		transportCreateMultipartUpload CreateMultipartUploadTransport

		transportUploadPartDocument UploadPartDocumentTransport

		transportCompleteUpload CompleteUploadTransport

		transportUploadDocument UploadDocumentTransport

		transportDownloadDocument DownloadDocumentTransport

		transportGetToken GetTokenTransport

		transportGetBranches GetBranchesTransport

		options map[interface{}]Option
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"test new client", args{hostClient, transportCreateMultipartUpload, transportUploadPartDocument, transportCompleteUpload, transportUploadDocument, transportDownloadDocument, transportGetToken, transportGetBranches, opts}, cl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.cli, tt.args.transportCreateMultipartUpload, tt.args.transportUploadPartDocument, tt.args.transportCompleteUpload, tt.args.transportUploadDocument, tt.args.transportDownloadDocument, tt.args.transportGetToken, tt.args.transportGetBranches, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
