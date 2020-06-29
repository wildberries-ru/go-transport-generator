// Package service ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger log.Logger
	svc    SomeService
}

// UploadDocument ...
func (s *loggingMiddleware) UploadDocument(ctx context.Context, token *string, name string, extension string, categoryID string, supplierID *int64, contractID *int64, data multipart.File) (err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "UploadDocument",
			"token", token,

			"extension", extension,
			"categoryID", categoryID,
			"supplierID", supplierID,
			"contractID", contractID,
			"data", data,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.UploadDocument(ctx, token, name, extension, categoryID, supplierID, contractID, data)
}

// GetWarehouses ...
func (s *loggingMiddleware) GetWarehouses(ctx context.Context) (pets map[string]v1.Detail, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetWarehouses",

			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetWarehouses(ctx)
}

// GetDetails ...
func (s *loggingMiddleware) GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetDetails",

			"fileID", fileID,
			"someID", someID,
			"token", token,
			"det", det,
			"ns", ns,
			"id", id,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetDetails(ctx, namespace, detail, fileID, someID, token)
}

// GetDetailsEmbedStruct ...
func (s *loggingMiddleware) GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetDetailsEmbedStruct",
			"namespace", namespace,
			"detail", detail,
			"response", response,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetDetailsEmbedStruct(ctx, namespace, detail)
}

// GetDetailsListEmbedStruct ...
func (s *loggingMiddleware) GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetDetailsListEmbedStruct",
			"namespace", namespace,
			"detail", detail,
			"details", details,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetDetailsListEmbedStruct(ctx, namespace, detail)
}

// PutDetails ...
func (s *loggingMiddleware) PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "PutDetails",
			"namespace", namespace,
			"detail", detail,
			"testID", testID,
			"blaID", blaID,
			"token", token,
			"pretty", pretty,
			"yang", yang,
			"cool", cool,
			"nothing", nothing,
			"id", id,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.PutDetails(ctx, namespace, detail, testID, blaID, token, pretty, yang)
}

// GetSomeElseDataUtf8 ...
func (s *loggingMiddleware) GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetSomeElseDataUtf8",
			"cool", cool,
			"nothing", nothing,
			"id", id,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetSomeElseDataUtf8(ctx)
}

// GetFile ...
func (s *loggingMiddleware) GetFile(ctx context.Context) (data []byte, fileName string, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetFile",
			"data", data,
			"fileName", fileName,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetFile(ctx)
}

func (s *loggingMiddleware) wrap(err error) log.Logger {
	lvl := level.Debug
	if err != nil {
		lvl = level.Error
	}
	return lvl(s.logger)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger, svc SomeService) SomeService {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
