package service

import (
	"context"
	"mime/multipart"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// SomeService ...
// @gtg http-server log metrics mock http-errors
type SomeService interface {
	// @gtg http-server-method POST
	// @gtg http-server-uri-path /api/v1/document
	// @gtg http-server-content-type multipart/form-data
	// @gtg http-server-header Authorization {token}
	// @gtg http-server-value-tag name name
	// @gtg http-server-value-tag extension extension
	// @gtg http-server-value-tag categoryID categoryID
	// @gtg http-server-value-tag contractID contractID
	// @gtg http-server-value-tag supplierID supplierID
	// @gtg log-ignore name
	// @gtg http-server-file-tag data document
	// @gtg http-server-response-content-type application/json
	// @gtg http-server-response-status 201
	UploadDocument(ctx context.Context, token *string, name string, extension string, categoryID string, supplierID *int64, contractID *int64, data multipart.File) (err error)
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/getWarehouses
	// @gtg http-server-content-type application/json
	// @gtg log-ignore pets
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	// @gtg http-server-response-body pets
	GetWarehouses(ctx context.Context) (pets map[string]v1.Detail, err error)
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/namespaces/{namespace}/details/{detail}
	// @gtg http-server-query file={fileID}&some={someID}
	// @gtg http-server-header X-Auth-Token {token}
	// @gtg http-server-content-type application/json
	// @gtg log-ignore namespace, detail
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
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/file
	// @gtg http-server-response-header Content-Type application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-file data fileName
	GetFile(ctx context.Context) (data []byte, fileName string, err error)
}
