// Package s3 ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package s3

import (
	"context"
	"time"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
	"github.com/wildberries-ru/go-transport-generator/log/logger"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger logger.Logger
	svc    Service
}

// CreateMultipartUpload ...
func (s *loggingMiddleware) CreateMultipartUpload(ctx context.Context, bucket string, key string) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"bucket":           bucket,
				"key":              key,
				"data":             data,
				"errorFlag":        errorFlag,
				"errorText":        errorText,
				"additionalErrors": additionalErrors,
				"elapsed":          time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("CreateMultipartUpload")
		} else {
			lg.Error("CreateMultipartUpload")
		}
	}(time.Now())
	return s.svc.CreateMultipartUpload(ctx, bucket, key)
}

// UploadPartDocument ...
func (s *loggingMiddleware) UploadPartDocument(ctx context.Context, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"bucket":     bucket,
				"key":        key,
				"uploadID":   uploadID,
				"partNumber": partNumber,
				"document":   document,
				"elapsed":    time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("UploadPartDocument")
		} else {
			lg.Error("UploadPartDocument")
		}
	}(time.Now())
	return s.svc.UploadPartDocument(ctx, bucket, key, uploadID, partNumber, document)
}

// CompleteUpload ...
func (s *loggingMiddleware) CompleteUpload(ctx context.Context, bucket string, key string, uploadID string) (err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"bucket":   bucket,
				"key":      key,
				"uploadID": uploadID,
				"elapsed":  time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("CompleteUpload")
		} else {
			lg.Error("CompleteUpload")
		}
	}(time.Now())
	return s.svc.CompleteUpload(ctx, bucket, key, uploadID)
}

// UploadDocument ...
func (s *loggingMiddleware) UploadDocument(ctx context.Context, bucket string, key string, document []byte) (err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"bucket":   bucket,
				"key":      key,
				"document": document,
				"elapsed":  time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("UploadDocument")
		} else {
			lg.Error("UploadDocument")
		}
	}(time.Now())
	return s.svc.UploadDocument(ctx, bucket, key, document)
}

// DownloadDocument ...
func (s *loggingMiddleware) DownloadDocument(ctx context.Context, bucket string, key string) (document []byte, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"bucket":   bucket,
				"key":      key,
				"document": document,
				"elapsed":  time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("DownloadDocument")
		} else {
			lg.Error("DownloadDocument")
		}
	}(time.Now())
	return s.svc.DownloadDocument(ctx, bucket, key)
}

// GetToken ...
func (s *loggingMiddleware) GetToken(ctx context.Context, authToken *string, scope string, grantType string) (token string, expiresIn int, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"authToken": authToken,
				"scope":     scope,
				"grantType": grantType,
				"token":     token,
				"expiresIn": expiresIn,
				"elapsed":   time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("GetToken")
		} else {
			lg.Error("GetToken")
		}
	}(time.Now())
	return s.svc.GetToken(ctx, authToken, scope, grantType)
}

// GetBranches ...
func (s *loggingMiddleware) GetBranches(ctx context.Context, authToken *string, supplierID *string) (branches []int, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"authToken":  authToken,
				"supplierID": supplierID,
				"branches":   branches,
				"elapsed":    time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("GetBranches")
		} else {
			lg.Error("GetBranches")
		}
	}(time.Now())
	return s.svc.GetBranches(ctx, authToken, supplierID)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger logger.Logger, svc Service) Service {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
