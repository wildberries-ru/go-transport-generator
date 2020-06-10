package service

import (
	"context"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// SomeService ...
// @gtg http-server log metrics http-client mock http-errors
type SomeService interface {
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/getWarehouses
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	GetWarehouses(ctx context.Context) (pets []v1.Detail, err error)
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/namespaces/{namespace}/details/{detail}
	// @gtg http-server-query file={fileID}&some={someID}
	// @gtg http-server-header X-Auth-Token {token}
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-header X-Auth-ID {id}
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error)
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/namespaces/{namespace}/details-embed/{detail}
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	// @gtg http-server-response-body response
	GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error)
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/namespaces/{namespace}/details-embed-array/{detail}
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	// @gtg http-server-response-body details
	GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error)
	// @gtg http-server-method PUT
	// @gtg http-server-uri-path /api/v1/namespaces/{namespace}/details/{detail}
	// @gtg http-server-query test={testID}&bla={blaID}
	// @gtg http-server-header X-Auth-Token {token}
	// @gtg http-server-json-tag pretty ThePretty
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-header X-Auth-ID {id}
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/someelsedata
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	// @gtg http-server-response-content-encoding utf-8
	// @gtg http-server-response-json-tag cool cool
	// @gtg http-server-response-json-tag nothing TheNothing
	// @gtg http-server-response-json-tag id id
	GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
}
