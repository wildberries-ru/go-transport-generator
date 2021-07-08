// Package s3client ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package s3client

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Client
}

// CreateBucketWithContext ...
func (s *instrumentingMiddleware) CreateBucketWithContext(ctx aws.Context, input *s3.CreateBucketInput, opts ...request.Option) (output *s3.CreateBucketOutput, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "CreateBucketWithContext",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.CreateBucketWithContext(ctx, input, opts...)
}

// DeleteBucketWithContext ...
func (s *instrumentingMiddleware) DeleteBucketWithContext(ctx aws.Context, input *s3.DeleteBucketInput, opts ...request.Option) (output *s3.DeleteBucketOutput, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "DeleteBucketWithContext",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.DeleteBucketWithContext(ctx, input, opts...)
}

// CreateMultipartUploadWithContext ...
func (s *instrumentingMiddleware) CreateMultipartUploadWithContext(ctx aws.Context, input *s3.CreateMultipartUploadInput, opts ...request.Option) (output *s3.CreateMultipartUploadOutput, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "CreateMultipartUploadWithContext",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.CreateMultipartUploadWithContext(ctx, input, opts...)
}

// UploadPartWithContext ...
func (s *instrumentingMiddleware) UploadPartWithContext(ctx aws.Context, input *s3.UploadPartInput, opts ...request.Option) (output *s3.UploadPartOutput, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "UploadPartWithContext",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.UploadPartWithContext(ctx, input, opts...)
}

// CompleteMultipartUploadWithContext ...
func (s *instrumentingMiddleware) CompleteMultipartUploadWithContext(ctx aws.Context, input *s3.CompleteMultipartUploadInput, opts ...request.Option) (output *s3.CompleteMultipartUploadOutput, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "CompleteMultipartUploadWithContext",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.CompleteMultipartUploadWithContext(ctx, input, opts...)
}

// PutObjectWithContext ...
func (s *instrumentingMiddleware) PutObjectWithContext(ctx aws.Context, input *s3.PutObjectInput, opts ...request.Option) (output *s3.PutObjectOutput, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "PutObjectWithContext",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.PutObjectWithContext(ctx, input, opts...)
}

// DeleteObjectWithContext ...
func (s *instrumentingMiddleware) DeleteObjectWithContext(ctx aws.Context, input *s3.DeleteObjectInput, opts ...request.Option) (output *s3.DeleteObjectOutput, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "DeleteObjectWithContext",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.DeleteObjectWithContext(ctx, input, opts...)
}

// GetObjectWithContext ...
func (s *instrumentingMiddleware) GetObjectWithContext(ctx aws.Context, input *s3.GetObjectInput, opts ...request.Option) (output *s3.GetObjectOutput, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "GetObjectWithContext",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetObjectWithContext(ctx, input, opts...)
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
	svc Client,
) Client {
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
