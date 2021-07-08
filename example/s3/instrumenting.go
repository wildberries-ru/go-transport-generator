// Package s3 ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package s3

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Service
}

// CreateMultipartUpload ...
func (s *instrumentingMiddleware) CreateMultipartUpload(ctx context.Context, bucket string, key string) (data v1.CreateMultipartUploadData, errorFlag bool, errorText string, additionalErrors *v1.AdditionalErrors, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "CreateMultipartUpload",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.CreateMultipartUpload(ctx, bucket, key)
}

// UploadPartDocument ...
func (s *instrumentingMiddleware) UploadPartDocument(ctx context.Context, bucket string, key string, uploadID string, partNumber int64, document []byte) (err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "UploadPartDocument",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.UploadPartDocument(ctx, bucket, key, uploadID, partNumber, document)
}

// CompleteUpload ...
func (s *instrumentingMiddleware) CompleteUpload(ctx context.Context, bucket string, key string, uploadID string) (err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "CompleteUpload",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.CompleteUpload(ctx, bucket, key, uploadID)
}

// UploadDocument ...
func (s *instrumentingMiddleware) UploadDocument(ctx context.Context, bucket string, key string, document []byte) (err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "UploadDocument",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.UploadDocument(ctx, bucket, key, document)
}

// DownloadDocument ...
func (s *instrumentingMiddleware) DownloadDocument(ctx context.Context, bucket string, key string) (document []byte, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "DownloadDocument",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.DownloadDocument(ctx, bucket, key)
}

// GetToken ...
func (s *instrumentingMiddleware) GetToken(ctx context.Context, authToken *string, scope string, grantType string) (token string, expiresIn int, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "GetToken",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetToken(ctx, authToken, scope, grantType)
}

// GetBranches ...
func (s *instrumentingMiddleware) GetBranches(ctx context.Context, authToken *string, supplierID *string) (branches []int, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "GetBranches",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetBranches(ctx, authToken, supplierID)
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(
	metricsNamespace string,
	metricsSubsystem string,
	metricsNameCount string,
	metricsNameCountHelp string,
	metricsNameDuration string,
	metricsNameDurationHelp string,
	labels []string,
	svc Service,
) Service {
	reqCount := kitprometheus.NewCounterFrom(
		prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      metricsNameCount,
			Help:      metricsNameCountHelp,
		},
		labels,
	)
	reqDuration := kitprometheus.NewSummaryFrom(
		prometheus.SummaryOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      metricsNameDuration,
			Help:      metricsNameDurationHelp,
		},
		labels,
	)
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
