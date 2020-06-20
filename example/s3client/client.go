package s3client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Client S3 implementation
// @gtg log metrics
type Client interface {
	CreateBucketWithContext(ctx aws.Context, input *s3.CreateBucketInput, opts ...request.Option) (output *s3.CreateBucketOutput, err error)
	DeleteBucketWithContext(ctx aws.Context, input *s3.DeleteBucketInput, opts ...request.Option) (output *s3.DeleteBucketOutput, err error)
	CreateMultipartUploadWithContext(ctx aws.Context, input *s3.CreateMultipartUploadInput, opts ...request.Option) (output *s3.CreateMultipartUploadOutput, err error)
	UploadPartWithContext(ctx aws.Context, input *s3.UploadPartInput, opts ...request.Option) (output *s3.UploadPartOutput, err error)
	CompleteMultipartUploadWithContext(ctx aws.Context, input *s3.CompleteMultipartUploadInput, opts ...request.Option) (output *s3.CompleteMultipartUploadOutput, err error)
	PutObjectWithContext(ctx aws.Context, input *s3.PutObjectInput, opts ...request.Option) (output *s3.PutObjectOutput, err error)
	DeleteObjectWithContext(ctx aws.Context, input *s3.DeleteObjectInput, opts ...request.Option) (output *s3.DeleteObjectOutput, err error)
	GetObjectWithContext(ctx aws.Context, input *s3.GetObjectInput, opts ...request.Option) (output *s3.GetObjectOutput, err error)
}
