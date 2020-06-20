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
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Client
}

// CreateBucketWithContext ...
func (s *instrumentingMiddleware) CreateBucketWithContext(ctx aws.Context, input *s3.CreateBucketInput, opts ...request.Option) (output *s3.CreateBucketOutput, err error) {
	defer s.recordMetrics("CreateBucketWithContext", time.Now(), err)
	return s.svc.CreateBucketWithContext(ctx, input, opts...)
}

// DeleteBucketWithContext ...
func (s *instrumentingMiddleware) DeleteBucketWithContext(ctx aws.Context, input *s3.DeleteBucketInput, opts ...request.Option) (output *s3.DeleteBucketOutput, err error) {
	defer s.recordMetrics("DeleteBucketWithContext", time.Now(), err)
	return s.svc.DeleteBucketWithContext(ctx, input, opts...)
}

// CreateMultipartUploadWithContext ...
func (s *instrumentingMiddleware) CreateMultipartUploadWithContext(ctx aws.Context, input *s3.CreateMultipartUploadInput, opts ...request.Option) (output *s3.CreateMultipartUploadOutput, err error) {
	defer s.recordMetrics("CreateMultipartUploadWithContext", time.Now(), err)
	return s.svc.CreateMultipartUploadWithContext(ctx, input, opts...)
}

// UploadPartWithContext ...
func (s *instrumentingMiddleware) UploadPartWithContext(ctx aws.Context, input *s3.UploadPartInput, opts ...request.Option) (output *s3.UploadPartOutput, err error) {
	defer s.recordMetrics("UploadPartWithContext", time.Now(), err)
	return s.svc.UploadPartWithContext(ctx, input, opts...)
}

// CompleteMultipartUploadWithContext ...
func (s *instrumentingMiddleware) CompleteMultipartUploadWithContext(ctx aws.Context, input *s3.CompleteMultipartUploadInput, opts ...request.Option) (output *s3.CompleteMultipartUploadOutput, err error) {
	defer s.recordMetrics("CompleteMultipartUploadWithContext", time.Now(), err)
	return s.svc.CompleteMultipartUploadWithContext(ctx, input, opts...)
}

// PutObjectWithContext ...
func (s *instrumentingMiddleware) PutObjectWithContext(ctx aws.Context, input *s3.PutObjectInput, opts ...request.Option) (output *s3.PutObjectOutput, err error) {
	defer s.recordMetrics("PutObjectWithContext", time.Now(), err)
	return s.svc.PutObjectWithContext(ctx, input, opts...)
}

// DeleteObjectWithContext ...
func (s *instrumentingMiddleware) DeleteObjectWithContext(ctx aws.Context, input *s3.DeleteObjectInput, opts ...request.Option) (output *s3.DeleteObjectOutput, err error) {
	defer s.recordMetrics("DeleteObjectWithContext", time.Now(), err)
	return s.svc.DeleteObjectWithContext(ctx, input, opts...)
}

// GetObjectWithContext ...
func (s *instrumentingMiddleware) GetObjectWithContext(ctx aws.Context, input *s3.GetObjectInput, opts ...request.Option) (output *s3.GetObjectOutput, err error) {
	defer s.recordMetrics("GetObjectWithContext", time.Now(), err)
	return s.svc.GetObjectWithContext(ctx, input, opts...)
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
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc Client) Client {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
