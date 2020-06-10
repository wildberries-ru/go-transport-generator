// Package s3 ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package s3

import (
	"context"
	"strconv"
	"time"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
	"github.com/go-kit/kit/metrics"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Service
}

func (s *instrumentingMiddleware) CreateMultipartUpload(ctx context.Context, bucket string, key string) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error) {
	defer s.recordMetrics("CreateMultipartUpload", time.Now(), err)
	return s.svc.CreateMultipartUpload(ctx, bucket, key)
}

func (s *instrumentingMiddleware) UploadPartDocument(ctx context.Context, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error) {
	defer s.recordMetrics("UploadPartDocument", time.Now(), err)
	return s.svc.UploadPartDocument(ctx, bucket, key, uploadID, partNumber, document)
}

func (s *instrumentingMiddleware) CompleteUpload(ctx context.Context, bucket string, key string, uploadID string) (err error) {
	defer s.recordMetrics("CompleteUpload", time.Now(), err)
	return s.svc.CompleteUpload(ctx, bucket, key, uploadID)
}

func (s *instrumentingMiddleware) UploadDocument(ctx context.Context, bucket string, key string, document []byte) (err error) {
	defer s.recordMetrics("UploadDocument", time.Now(), err)
	return s.svc.UploadDocument(ctx, bucket, key, document)
}

func (s *instrumentingMiddleware) DownloadDocument(ctx context.Context, bucket string, key string) (document []byte, err error) {
	defer s.recordMetrics("DownloadDocument", time.Now(), err)
	return s.svc.DownloadDocument(ctx, bucket, key)
}

func (s *instrumentingMiddleware) recordMetrics(method string, startTime time.Time, err error) {
	labels := []string{
		"method", method,
		"error", strconv.FormatBool(err != nil),
	}
	s.reqCount.With(labels...).Add(1)
	s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc Service) Service {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
