// Package service ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package service

import (
	"context"
	"mime/multipart"
	"time"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
	"github.com/wildberries-ru/go-transport-generator/log/logger"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger logger.Logger
	svc    SomeService
}

// UploadDocument ...
func (s *loggingMiddleware) UploadDocument(ctx context.Context, token *string, name string, extension string, categoryID string, supplierID *int64, contractID *int64, data *multipart.FileHeader) (err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"token": token,

				"extension":  extension,
				"categoryID": categoryID,
				"supplierID": supplierID,
				"contractID": contractID,
				"data":       data,
				"elapsed":    time.Since(begin),
			},
		)
		if err != nil {
			lg.Debug("UploadDocument")
		} else {
			lg.Error("UploadDocument")
		}
	}(time.Now())
	return s.svc.UploadDocument(ctx, token, name, extension, categoryID, supplierID, contractID, data)
}

// GetWarehouses ...
func (s *loggingMiddleware) GetWarehouses(ctx context.Context) (pets map[string]v1.Detail, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{

				"elapsed": time.Since(begin),
			},
		)
		if err != nil {
			lg.Debug("GetWarehouses")
		} else {
			lg.Error("GetWarehouses")
		}
	}(time.Now())
	return s.svc.GetWarehouses(ctx)
}

// GetDetails ...
func (s *loggingMiddleware) GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{

				"fileID":  fileID,
				"someID":  someID,
				"token":   token,
				"det":     det,
				"ns":      ns,
				"id":      id,
				"elapsed": time.Since(begin),
			},
		)
		if err != nil {
			lg.Debug("GetDetails")
		} else {
			lg.Error("GetDetails")
		}
	}(time.Now())
	return s.svc.GetDetails(ctx, namespace, detail, fileID, someID, token)
}

// GetDetailsEmbedStruct ...
func (s *loggingMiddleware) GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"namespace": namespace,
				"detail":    detail,
				"response":  response,
				"elapsed":   time.Since(begin),
			},
		)
		if err != nil {
			lg.Debug("GetDetailsEmbedStruct")
		} else {
			lg.Error("GetDetailsEmbedStruct")
		}
	}(time.Now())
	return s.svc.GetDetailsEmbedStruct(ctx, namespace, detail)
}

// GetDetailsListEmbedStruct ...
func (s *loggingMiddleware) GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"namespace": namespace,
				"detail":    detail,
				"details":   details,
				"elapsed":   time.Since(begin),
			},
		)
		if err != nil {
			lg.Debug("GetDetailsListEmbedStruct")
		} else {
			lg.Error("GetDetailsListEmbedStruct")
		}
	}(time.Now())
	return s.svc.GetDetailsListEmbedStruct(ctx, namespace, detail)
}

// PutDetails ...
func (s *loggingMiddleware) PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"namespace": namespace,
				"detail":    detail,
				"testID":    testID,
				"blaID":     blaID,
				"token":     token,
				"pretty":    pretty,
				"yang":      yang,
				"cool":      cool,
				"nothing":   nothing,
				"id":        id,
				"elapsed":   time.Since(begin),
			},
		)
		if err != nil {
			lg.Debug("PutDetails")
		} else {
			lg.Error("PutDetails")
		}
	}(time.Now())
	return s.svc.PutDetails(ctx, namespace, detail, testID, blaID, token, pretty, yang)
}

// GetSomeElseDataUtf8 ...
func (s *loggingMiddleware) GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"cool":    cool,
				"nothing": nothing,
				"id":      id,
				"elapsed": time.Since(begin),
			},
		)
		if err != nil {
			lg.Debug("GetSomeElseDataUtf8")
		} else {
			lg.Error("GetSomeElseDataUtf8")
		}
	}(time.Now())
	return s.svc.GetSomeElseDataUtf8(ctx)
}

// GetFile ...
func (s *loggingMiddleware) GetFile(ctx context.Context) (data []byte, fileName string, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"data":     data,
				"fileName": fileName,
				"elapsed":  time.Since(begin),
			},
		)
		if err != nil {
			lg.Debug("GetFile")
		} else {
			lg.Error("GetFile")
		}
	}(time.Now())
	return s.svc.GetFile(ctx)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger logger.Logger, svc SomeService) SomeService {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
