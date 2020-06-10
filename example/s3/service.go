package s3

import (
	"context"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// Service ...
// @gtg http-server log metrics http-client mock http-errors swagger
type Service interface {
	// @gtg http-server-method POST
	// @gtg http-server-uri-path /api/v1/multipart/{bucket}/{key}
	// @gtg http-server-response-status 201
	// @gtg http-server-response-content-type application/json
	// @gtg http-server-response-json-tag data data
	// @gtg http-server-response-json-tag errorFlag error
	// @gtg http-server-response-json-tag additionalErrors additionalErrors
	// @gtg http-server-response-json-tag errorText errorText
	CreateMultipartUpload(ctx context.Context, bucket, key string) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error)
	// @gtg http-server-method PATCH
	// @gtg http-server-uri-path /api/v1/multipart/{bucket}/{key}
	// @gtg http-server-response-status 200
	// @gtg http-server-response-content-type application/json
	UploadPartDocument(ctx context.Context, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error)
	// @gtg http-server-method PUT
	// @gtg http-server-uri-path /api/v1/multipart/{bucket}/{key}
	// @gtg http-server-response-status 200
	// @gtg http-server-response-content-type application/json
	CompleteUpload(ctx context.Context, bucket string, key string, uploadID string) (err error)
	// @gtg http-server-method POST
	// @gtg http-server-uri-path /api/v1/doc/{bucket}/{key}
	// @gtg http-server-response-status 201
	// @gtg http-server-response-content-type application/json
	UploadDocument(ctx context.Context, bucket string, key string, document []byte) (err error)
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/doc/{bucket}/{key}
	// @gtg http-server-response-status 200
	// @gtg http-server-response-content-type application/json
	DownloadDocument(ctx context.Context, bucket string, key string) (document []byte, err error)
}
