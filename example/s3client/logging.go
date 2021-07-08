// Package s3client ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package s3client

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/wildberries-ru/go-transport-generator/log/logger"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger logger.Logger
	svc    Client
}

// CreateBucketWithContext ...
func (s *loggingMiddleware) CreateBucketWithContext(ctx aws.Context, input *s3.CreateBucketInput, opts ...request.Option) (output *s3.CreateBucketOutput, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"input":   input,
				"opts":    opts,
				"output":  output,
				"elapsed": time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("CreateBucketWithContext")
		} else {
			lg.Error("CreateBucketWithContext")
		}
	}(time.Now())
	return s.svc.CreateBucketWithContext(ctx, input, opts...)
}

// DeleteBucketWithContext ...
func (s *loggingMiddleware) DeleteBucketWithContext(ctx aws.Context, input *s3.DeleteBucketInput, opts ...request.Option) (output *s3.DeleteBucketOutput, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"input":   input,
				"opts":    opts,
				"output":  output,
				"elapsed": time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("DeleteBucketWithContext")
		} else {
			lg.Error("DeleteBucketWithContext")
		}
	}(time.Now())
	return s.svc.DeleteBucketWithContext(ctx, input, opts...)
}

// CreateMultipartUploadWithContext ...
func (s *loggingMiddleware) CreateMultipartUploadWithContext(ctx aws.Context, input *s3.CreateMultipartUploadInput, opts ...request.Option) (output *s3.CreateMultipartUploadOutput, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"input":   input,
				"opts":    opts,
				"output":  output,
				"elapsed": time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("CreateMultipartUploadWithContext")
		} else {
			lg.Error("CreateMultipartUploadWithContext")
		}
	}(time.Now())
	return s.svc.CreateMultipartUploadWithContext(ctx, input, opts...)
}

// UploadPartWithContext ...
func (s *loggingMiddleware) UploadPartWithContext(ctx aws.Context, input *s3.UploadPartInput, opts ...request.Option) (output *s3.UploadPartOutput, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"input":   input,
				"opts":    opts,
				"output":  output,
				"elapsed": time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("UploadPartWithContext")
		} else {
			lg.Error("UploadPartWithContext")
		}
	}(time.Now())
	return s.svc.UploadPartWithContext(ctx, input, opts...)
}

// CompleteMultipartUploadWithContext ...
func (s *loggingMiddleware) CompleteMultipartUploadWithContext(ctx aws.Context, input *s3.CompleteMultipartUploadInput, opts ...request.Option) (output *s3.CompleteMultipartUploadOutput, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"input":   input,
				"opts":    opts,
				"output":  output,
				"elapsed": time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("CompleteMultipartUploadWithContext")
		} else {
			lg.Error("CompleteMultipartUploadWithContext")
		}
	}(time.Now())
	return s.svc.CompleteMultipartUploadWithContext(ctx, input, opts...)
}

// PutObjectWithContext ...
func (s *loggingMiddleware) PutObjectWithContext(ctx aws.Context, input *s3.PutObjectInput, opts ...request.Option) (output *s3.PutObjectOutput, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"input":   input,
				"opts":    opts,
				"output":  output,
				"elapsed": time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("PutObjectWithContext")
		} else {
			lg.Error("PutObjectWithContext")
		}
	}(time.Now())
	return s.svc.PutObjectWithContext(ctx, input, opts...)
}

// DeleteObjectWithContext ...
func (s *loggingMiddleware) DeleteObjectWithContext(ctx aws.Context, input *s3.DeleteObjectInput, opts ...request.Option) (output *s3.DeleteObjectOutput, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"input":   input,
				"opts":    opts,
				"output":  output,
				"elapsed": time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("DeleteObjectWithContext")
		} else {
			lg.Error("DeleteObjectWithContext")
		}
	}(time.Now())
	return s.svc.DeleteObjectWithContext(ctx, input, opts...)
}

// GetObjectWithContext ...
func (s *loggingMiddleware) GetObjectWithContext(ctx aws.Context, input *s3.GetObjectInput, opts ...request.Option) (output *s3.GetObjectOutput, err error) {
	defer func(begin time.Time) {
		lg := s.logger.WithError(err).WithFields(
			map[string]interface{}{
				"input":   input,
				"opts":    opts,
				"output":  output,
				"elapsed": time.Since(begin),
			},
		)
		if err == nil {
			lg.Debug("GetObjectWithContext")
		} else {
			lg.Error("GetObjectWithContext")
		}
	}(time.Now())
	return s.svc.GetObjectWithContext(ctx, input, opts...)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger logger.Logger, svc Client) Client {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
