// Package s3client ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package s3client

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger log.Logger
	svc    Client
}

// CreateBucketWithContext ...
func (s *loggingMiddleware) CreateBucketWithContext(ctx aws.Context, input *s3.CreateBucketInput, opts ...request.Option) (output *s3.CreateBucketOutput, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "CreateBucketWithContext",
			"input", input,
			"opts", opts,
			"output", output,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.CreateBucketWithContext(ctx, input, opts...)
}

// DeleteBucketWithContext ...
func (s *loggingMiddleware) DeleteBucketWithContext(ctx aws.Context, input *s3.DeleteBucketInput, opts ...request.Option) (output *s3.DeleteBucketOutput, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "DeleteBucketWithContext",
			"input", input,
			"opts", opts,
			"output", output,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.DeleteBucketWithContext(ctx, input, opts...)
}

// CreateMultipartUploadWithContext ...
func (s *loggingMiddleware) CreateMultipartUploadWithContext(ctx aws.Context, input *s3.CreateMultipartUploadInput, opts ...request.Option) (output *s3.CreateMultipartUploadOutput, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "CreateMultipartUploadWithContext",
			"input", input,
			"opts", opts,
			"output", output,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.CreateMultipartUploadWithContext(ctx, input, opts...)
}

// UploadPartWithContext ...
func (s *loggingMiddleware) UploadPartWithContext(ctx aws.Context, input *s3.UploadPartInput, opts ...request.Option) (output *s3.UploadPartOutput, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "UploadPartWithContext",
			"input", input,
			"opts", opts,
			"output", output,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.UploadPartWithContext(ctx, input, opts...)
}

// CompleteMultipartUploadWithContext ...
func (s *loggingMiddleware) CompleteMultipartUploadWithContext(ctx aws.Context, input *s3.CompleteMultipartUploadInput, opts ...request.Option) (output *s3.CompleteMultipartUploadOutput, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "CompleteMultipartUploadWithContext",
			"input", input,
			"opts", opts,
			"output", output,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.CompleteMultipartUploadWithContext(ctx, input, opts...)
}

// PutObjectWithContext ...
func (s *loggingMiddleware) PutObjectWithContext(ctx aws.Context, input *s3.PutObjectInput, opts ...request.Option) (output *s3.PutObjectOutput, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "PutObjectWithContext",
			"input", input,
			"opts", opts,
			"output", output,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.PutObjectWithContext(ctx, input, opts...)
}

// DeleteObjectWithContext ...
func (s *loggingMiddleware) DeleteObjectWithContext(ctx aws.Context, input *s3.DeleteObjectInput, opts ...request.Option) (output *s3.DeleteObjectOutput, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "DeleteObjectWithContext",
			"input", input,
			"opts", opts,
			"output", output,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.DeleteObjectWithContext(ctx, input, opts...)
}

// GetObjectWithContext ...
func (s *loggingMiddleware) GetObjectWithContext(ctx aws.Context, input *s3.GetObjectInput, opts ...request.Option) (output *s3.GetObjectOutput, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetObjectWithContext",
			"input", input,
			"opts", opts,
			"output", output,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetObjectWithContext(ctx, input, opts...)
}

func (s *loggingMiddleware) wrap(err error) log.Logger {
	lvl := level.Debug
	if err != nil {
		lvl = level.Error
	}
	return lvl(s.logger)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger, svc Client) Client {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
