// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"mime/multipart"
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

func Test_client_UploadDocument(t *testing.T) {

	var token *string
	_ = faker.FakeData(&token)

	var name []string
	_ = faker.FakeData(&name)

	var extension string
	_ = faker.FakeData(&extension)

	var categoryID string
	_ = faker.FakeData(&categoryID)

	var supplierID []int64
	_ = faker.FakeData(&supplierID)

	var contractID *bool
	_ = faker.FakeData(&contractID)

	var data *multipart.FileHeader

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
		ctx        context.Context
		token      *string
		name       []string
		extension  string
		categoryID string
		supplierID []int64
		contractID *bool
		data       *multipart.FileHeader
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
			args{context.Background(), token, name, extension, categoryID, supplierID, contractID, data},

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                     tt.fields.cli,
				transportUploadDocument: tt.fields.transportUploadDocument,
				options:                 tt.fields.options,
			}
			err := s.UploadDocument(tt.args.ctx, tt.args.token, tt.args.name, tt.args.extension, tt.args.categoryID, tt.args.supplierID, tt.args.contractID, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UploadDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_client_GetWarehouses(t *testing.T) {

	var token *string
	_ = faker.FakeData(&token)

	var pets map[string]v1.Detail
	_ = faker.FakeData(&pets)

	var someCookie *string
	_ = faker.FakeData(&someCookie)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
			Pets map[string]v1.Detail `json:"Pets"`

			SomeCookie *string `json:"SomeCookie"`
		}{

			Pets: pets,

			SomeCookie: someCookie,
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

	transportGetWarehouses := NewGetWarehousesTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetWarehouses,
		httpMethodGetWarehouses,
	)

	type fields struct {
		cli                    *fasthttp.HostClient
		transportGetWarehouses GetWarehousesTransport
		options                map[interface{}]Option
	}
	type args struct {
		ctx   context.Context
		token *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantPets map[string]v1.Detail

		wantSomeCookie *string

		wantErr bool
	}{
		{
			"test GetWarehouses",
			fields{hostClient, transportGetWarehouses, opts},
			args{context.Background(), token},
			pets,
			someCookie,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                    tt.fields.cli,
				transportGetWarehouses: tt.fields.transportGetWarehouses,
				options:                tt.fields.options,
			}
			gotPets, gotSomeCookie, err := s.GetWarehouses(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetWarehouses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(gotPets, tt.wantPets, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.pets() = %s", cmp.Diff(gotPets, tt.wantPets))
			}

			if !cmp.Equal(gotSomeCookie, tt.wantSomeCookie, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.someCookie() = %s", cmp.Diff(gotSomeCookie, tt.wantSomeCookie))
			}

		})
	}
}

func Test_client_GetDetails(t *testing.T) {

	var namespace string
	_ = faker.FakeData(&namespace)

	var detail string
	_ = faker.FakeData(&detail)

	var fileID uint32
	_ = faker.FakeData(&fileID)

	var someID *uint64
	_ = faker.FakeData(&someID)

	var token *string
	_ = faker.FakeData(&token)

	var det v1.Detail
	_ = faker.FakeData(&det)

	var ns v1.Namespace
	_ = faker.FakeData(&ns)

	var id *string
	_ = faker.FakeData(&id)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
			Det v1.Detail `json:"Det"`

			Ns v1.Namespace `json:"Ns"`

			ID *string `json:"ID"`
		}{

			Det: det,

			Ns: ns,

			ID: id,
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

	transportGetDetails := NewGetDetailsTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetDetails,
		httpMethodGetDetails,
	)

	type fields struct {
		cli                 *fasthttp.HostClient
		transportGetDetails GetDetailsTransport
		options             map[interface{}]Option
	}
	type args struct {
		ctx       context.Context
		namespace string
		detail    string
		fileID    uint32
		someID    *uint64
		token     *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantDet v1.Detail

		wantNs v1.Namespace

		wantID *string

		wantErr bool
	}{
		{
			"test GetDetails",
			fields{hostClient, transportGetDetails, opts},
			args{context.Background(), namespace, detail, fileID, someID, token},
			det,
			ns,
			id,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                 tt.fields.cli,
				transportGetDetails: tt.fields.transportGetDetails,
				options:             tt.fields.options,
			}
			gotDet, gotNs, gotID, err := s.GetDetails(tt.args.ctx, tt.args.namespace, tt.args.detail, tt.args.fileID, tt.args.someID, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(gotDet, tt.wantDet, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.det() = %s", cmp.Diff(gotDet, tt.wantDet))
			}

			if !cmp.Equal(gotNs, tt.wantNs, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.ns() = %s", cmp.Diff(gotNs, tt.wantNs))
			}

			if !cmp.Equal(gotID, tt.wantID, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.id() = %s", cmp.Diff(gotID, tt.wantID))
			}

		})
	}
}

func Test_client_GetDetailsEmbedStruct(t *testing.T) {

	var namespace string
	_ = faker.FakeData(&namespace)

	var detail string
	_ = faker.FakeData(&detail)

	var token *string
	_ = faker.FakeData(&token)

	var response v1.GetDetailsEmbedStructResponse
	_ = faker.FakeData(&response)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := response

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

	transportGetDetailsEmbedStruct := NewGetDetailsEmbedStructTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetDetailsEmbedStruct,
		httpMethodGetDetailsEmbedStruct,
	)

	type fields struct {
		cli                            *fasthttp.HostClient
		transportGetDetailsEmbedStruct GetDetailsEmbedStructTransport
		options                        map[interface{}]Option
	}
	type args struct {
		ctx       context.Context
		namespace string
		detail    string
		token     *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantResponse v1.GetDetailsEmbedStructResponse

		wantErr bool
	}{
		{
			"test GetDetailsEmbedStruct",
			fields{hostClient, transportGetDetailsEmbedStruct, opts},
			args{context.Background(), namespace, detail, token},
			response,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                            tt.fields.cli,
				transportGetDetailsEmbedStruct: tt.fields.transportGetDetailsEmbedStruct,
				options:                        tt.fields.options,
			}
			gotResponse, err := s.GetDetailsEmbedStruct(tt.args.ctx, tt.args.namespace, tt.args.detail, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetDetailsEmbedStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(gotResponse, tt.wantResponse, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.response() = %s", cmp.Diff(gotResponse, tt.wantResponse))
			}

		})
	}
}

func Test_client_GetDetailsListEmbedStruct(t *testing.T) {

	var namespace string
	_ = faker.FakeData(&namespace)

	var detail string
	_ = faker.FakeData(&detail)

	var token *string
	_ = faker.FakeData(&token)

	var details []v1.Detail
	_ = faker.FakeData(&details)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := details

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

	transportGetDetailsListEmbedStruct := NewGetDetailsListEmbedStructTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetDetailsListEmbedStruct,
		httpMethodGetDetailsListEmbedStruct,
	)

	type fields struct {
		cli                                *fasthttp.HostClient
		transportGetDetailsListEmbedStruct GetDetailsListEmbedStructTransport
		options                            map[interface{}]Option
	}
	type args struct {
		ctx       context.Context
		namespace string
		detail    string
		token     *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantDetails []v1.Detail

		wantErr bool
	}{
		{
			"test GetDetailsListEmbedStruct",
			fields{hostClient, transportGetDetailsListEmbedStruct, opts},
			args{context.Background(), namespace, detail, token},
			details,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                                tt.fields.cli,
				transportGetDetailsListEmbedStruct: tt.fields.transportGetDetailsListEmbedStruct,
				options:                            tt.fields.options,
			}
			gotDetails, err := s.GetDetailsListEmbedStruct(tt.args.ctx, tt.args.namespace, tt.args.detail, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetDetailsListEmbedStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(gotDetails, tt.wantDetails, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.details() = %s", cmp.Diff(gotDetails, tt.wantDetails))
			}

		})
	}
}

func Test_client_PutDetails(t *testing.T) {

	var namespace string
	_ = faker.FakeData(&namespace)

	var detail string
	_ = faker.FakeData(&detail)

	var testID string
	_ = faker.FakeData(&testID)

	var blaID *string
	_ = faker.FakeData(&blaID)

	var token *string
	_ = faker.FakeData(&token)

	var pretty v1.Detail
	_ = faker.FakeData(&pretty)

	var yang v1.Namespace
	_ = faker.FakeData(&yang)

	var cool v1.Detail
	_ = faker.FakeData(&cool)

	var nothing v1.Namespace
	_ = faker.FakeData(&nothing)

	var id *string
	_ = faker.FakeData(&id)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
			Cool v1.Detail `json:"Cool"`

			Nothing v1.Namespace `json:"Nothing"`

			ID *string `json:"ID"`
		}{

			Cool: cool,

			Nothing: nothing,

			ID: id,
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

	transportPutDetails := NewPutDetailsTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientPutDetails,
		httpMethodPutDetails,
	)

	type fields struct {
		cli                 *fasthttp.HostClient
		transportPutDetails PutDetailsTransport
		options             map[interface{}]Option
	}
	type args struct {
		ctx       context.Context
		namespace string
		detail    string
		testID    string
		blaID     *string
		token     *string
		pretty    v1.Detail
		yang      v1.Namespace
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantCool v1.Detail

		wantNothing v1.Namespace

		wantID *string

		wantErr bool
	}{
		{
			"test PutDetails",
			fields{hostClient, transportPutDetails, opts},
			args{context.Background(), namespace, detail, testID, blaID, token, pretty, yang},
			cool,
			nothing,
			id,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                 tt.fields.cli,
				transportPutDetails: tt.fields.transportPutDetails,
				options:             tt.fields.options,
			}
			gotCool, gotNothing, gotID, err := s.PutDetails(tt.args.ctx, tt.args.namespace, tt.args.detail, tt.args.testID, tt.args.blaID, tt.args.token, tt.args.pretty, tt.args.yang)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.PutDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(gotCool, tt.wantCool, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.cool() = %s", cmp.Diff(gotCool, tt.wantCool))
			}

			if !cmp.Equal(gotNothing, tt.wantNothing, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.nothing() = %s", cmp.Diff(gotNothing, tt.wantNothing))
			}

			if !cmp.Equal(gotID, tt.wantID, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.id() = %s", cmp.Diff(gotID, tt.wantID))
			}

		})
	}
}

func Test_client_GetSomeElseDataUtf8(t *testing.T) {

	var token *string
	_ = faker.FakeData(&token)

	var cool v1.Detail
	_ = faker.FakeData(&cool)

	var nothing v1.Namespace
	_ = faker.FakeData(&nothing)

	var id *string
	_ = faker.FakeData(&id)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
			Cool v1.Detail `json:"cool"`

			Nothing v1.Namespace `json:"TheNothing"`

			ID *string `json:"id"`
		}{

			Cool: cool,

			Nothing: nothing,

			ID: id,
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

	transportGetSomeElseDataUtf8 := NewGetSomeElseDataUtf8Transport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetSomeElseDataUtf8,
		httpMethodGetSomeElseDataUtf8,
	)

	type fields struct {
		cli                          *fasthttp.HostClient
		transportGetSomeElseDataUtf8 GetSomeElseDataUtf8Transport
		options                      map[interface{}]Option
	}
	type args struct {
		ctx   context.Context
		token *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantCool v1.Detail

		wantNothing v1.Namespace

		wantID *string

		wantErr bool
	}{
		{
			"test GetSomeElseDataUtf8",
			fields{hostClient, transportGetSomeElseDataUtf8, opts},
			args{context.Background(), token},
			cool,
			nothing,
			id,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                          tt.fields.cli,
				transportGetSomeElseDataUtf8: tt.fields.transportGetSomeElseDataUtf8,
				options:                      tt.fields.options,
			}
			gotCool, gotNothing, gotID, err := s.GetSomeElseDataUtf8(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetSomeElseDataUtf8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(gotCool, tt.wantCool, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.cool() = %s", cmp.Diff(gotCool, tt.wantCool))
			}

			if !cmp.Equal(gotNothing, tt.wantNothing, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.nothing() = %s", cmp.Diff(gotNothing, tt.wantNothing))
			}

			if !cmp.Equal(gotID, tt.wantID, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.id() = %s", cmp.Diff(gotID, tt.wantID))
			}

		})
	}
}

func Test_client_GetFile(t *testing.T) {

	var token *string
	_ = faker.FakeData(&token)

	var data []byte
	_ = faker.FakeData(&data)

	var fileName string
	_ = faker.FakeData(&fileName)

	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := struct {
			Data []byte `json:"Data"`

			FileName string `json:"FileName"`
		}{

			Data: data,

			FileName: fileName,
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

	transportGetFile := NewGetFileTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetFile,
		httpMethodGetFile,
	)

	type fields struct {
		cli              *fasthttp.HostClient
		transportGetFile GetFileTransport
		options          map[interface{}]Option
	}
	type args struct {
		ctx   context.Context
		token *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantData []byte

		wantFileName string

		wantErr bool
	}{
		{
			"test GetFile",
			fields{hostClient, transportGetFile, opts},
			args{context.Background(), token},
			data,
			fileName,

			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:              tt.fields.cli,
				transportGetFile: tt.fields.transportGetFile,
				options:          tt.fields.options,
			}
			gotData, gotFileName, err := s.GetFile(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(gotData, tt.wantData, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.data() = %s", cmp.Diff(gotData, tt.wantData))
			}

			if !cmp.Equal(gotFileName, tt.wantFileName, cmpopts.EquateApproxTime(1*time.Second)) {
				t.Errorf("client.fileName() = %s", cmp.Diff(gotFileName, tt.wantFileName))
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

	transportUploadDocument := NewUploadDocumentTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientUploadDocument,
		httpMethodUploadDocument,
	)

	transportGetWarehouses := NewGetWarehousesTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetWarehouses,
		httpMethodGetWarehouses,
	)

	transportGetDetails := NewGetDetailsTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetDetails,
		httpMethodGetDetails,
	)

	transportGetDetailsEmbedStruct := NewGetDetailsEmbedStructTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetDetailsEmbedStruct,
		httpMethodGetDetailsEmbedStruct,
	)

	transportGetDetailsListEmbedStruct := NewGetDetailsListEmbedStructTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetDetailsListEmbedStruct,
		httpMethodGetDetailsListEmbedStruct,
	)

	transportPutDetails := NewPutDetailsTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientPutDetails,
		httpMethodPutDetails,
	)

	transportGetSomeElseDataUtf8 := NewGetSomeElseDataUtf8Transport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetSomeElseDataUtf8,
		httpMethodGetSomeElseDataUtf8,
	)

	transportGetFile := NewGetFileTransport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClientGetFile,
		httpMethodGetFile,
	)

	cl := &client{
		hostClient,
		transportUploadDocument,
		transportGetWarehouses,
		transportGetDetails,
		transportGetDetailsEmbedStruct,
		transportGetDetailsListEmbedStruct,
		transportPutDetails,
		transportGetSomeElseDataUtf8,
		transportGetFile,
		opts,
	}

	type args struct {
		cli *fasthttp.HostClient

		transportUploadDocument UploadDocumentTransport

		transportGetWarehouses GetWarehousesTransport

		transportGetDetails GetDetailsTransport

		transportGetDetailsEmbedStruct GetDetailsEmbedStructTransport

		transportGetDetailsListEmbedStruct GetDetailsListEmbedStructTransport

		transportPutDetails PutDetailsTransport

		transportGetSomeElseDataUtf8 GetSomeElseDataUtf8Transport

		transportGetFile GetFileTransport

		options map[interface{}]Option
	}
	tests := []struct {
		name string
		args args
		want SomeService
	}{
		{"test new client", args{hostClient, transportUploadDocument, transportGetWarehouses, transportGetDetails, transportGetDetailsEmbedStruct, transportGetDetailsListEmbedStruct, transportPutDetails, transportGetSomeElseDataUtf8, transportGetFile, opts}, cl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.cli, tt.args.transportUploadDocument, tt.args.transportGetWarehouses, tt.args.transportGetDetails, tt.args.transportGetDetailsEmbedStruct, tt.args.transportGetDetailsListEmbedStruct, tt.args.transportPutDetails, tt.args.transportGetSomeElseDataUtf8, tt.args.transportGetFile, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
