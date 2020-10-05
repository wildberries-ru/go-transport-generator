// Package s3 ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package s3

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

// CreateMultipartUpload ...
func (s *loggingMiddleware) CreateMultipartUpload(ctx context.Context, bucket string, key string) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "CreateMultipartUpload",
			"bucket", bucket,
			"key", key,
			"data", data,
			"errorFlag", errorFlag,
			"errorText", errorText,
			"additionalErrors", additionalErrors,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.CreateMultipartUpload(ctx, bucket, key)
}

// UploadPartDocument ...
func (s *loggingMiddleware) UploadPartDocument(ctx context.Context, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "UploadPartDocument",
			"bucket", bucket,
			"key", key,
			"uploadID", uploadID,
			"partNumber", partNumber,
			"document", document,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.UploadPartDocument(ctx, bucket, key, uploadID, partNumber, document)
}

// CompleteUpload ...
func (s *loggingMiddleware) CompleteUpload(ctx context.Context, bucket string, key string, uploadID string) (err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "CompleteUpload",
			"bucket", bucket,
			"key", key,
			"uploadID", uploadID,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.CompleteUpload(ctx, bucket, key, uploadID)
}

// UploadDocument ...
func (s *loggingMiddleware) UploadDocument(ctx context.Context, bucket string, key string, document []byte) (err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "UploadDocument",
			"bucket", bucket,
			"key", key,
			"document", document,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.UploadDocument(ctx, bucket, key, document)
}

// DownloadDocument ...
func (s *loggingMiddleware) DownloadDocument(ctx context.Context, bucket string, key string) (document []byte, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "DownloadDocument",
			"bucket", bucket,
			"key", key,
			"document", document,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.DownloadDocument(ctx, bucket, key)
}

// GetToken ...
func (s *loggingMiddleware) GetToken(ctx context.Context, authToken *string, scope string, grantType string) (token string, expiresIn int, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetToken",
			"authToken", authToken,
			"scope", scope,
			"grantType", grantType,
			"token", token,
			"expiresIn", expiresIn,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetToken(ctx, authToken, scope, grantType)
}

// GetBranches ...
func (s *loggingMiddleware) GetBranches(ctx context.Context, authToken *string, supplierID *string) (branches []int, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetBranches",
			"authToken", authToken,
			"supplierID", supplierID,
			"branches", branches,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetBranches(ctx, authToken, supplierID)
}

func (s *loggingMiddleware) wrap(err error) log.Logger {
	lvl := level.Debug
	if err != nil {
		lvl = level.Error
	}
	return lvl(s.logger)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger, svc Service) Service {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
