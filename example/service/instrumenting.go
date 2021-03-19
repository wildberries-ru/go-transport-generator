// Package service ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package service

import (
	"context"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         SomeService
}

// UploadDocument ...
func (s *instrumentingMiddleware) UploadDocument(ctx context.Context, token *string, name string, extension string, categoryID string, supplierID *int64, contractID *int64, data *multipart.FileHeader) (err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "UploadDocument",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.UploadDocument(ctx, token, name, extension, categoryID, supplierID, contractID, data)
}

// GetWarehouses ...
func (s *instrumentingMiddleware) GetWarehouses(ctx context.Context) (pets map[string]v1.Detail, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "GetWarehouses",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetWarehouses(ctx)
}

// GetDetails ...
func (s *instrumentingMiddleware) GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error) {
	defer func(startTime time.Time) {

		var _someID string
		if someID != nil {

			_someID = strconv.Itoa(int(*someID))

		} else {
			_someID = "empty"
		}

		var _token string
		if token != nil {

			_token = *token

		} else {
			_token = "empty"
		}

		labels := []string{
			"method", "GetDetails",
			"error", strconv.FormatBool(err != nil),

			"detail", detail,

			"fileID", strconv.Itoa(int(fileID)),

			"namespace", namespace,

			"someID", _someID,

			"token", _token,
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetDetails(ctx, namespace, detail, fileID, someID, token)
}

// GetDetailsEmbedStruct ...
func (s *instrumentingMiddleware) GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "GetDetailsEmbedStruct",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetDetailsEmbedStruct(ctx, namespace, detail)
}

// GetDetailsListEmbedStruct ...
func (s *instrumentingMiddleware) GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "GetDetailsListEmbedStruct",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetDetailsListEmbedStruct(ctx, namespace, detail)
}

// PutDetails ...
func (s *instrumentingMiddleware) PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "PutDetails",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.PutDetails(ctx, namespace, detail, testID, blaID, token, pretty, yang)
}

// GetSomeElseDataUtf8 ...
func (s *instrumentingMiddleware) GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "GetSomeElseDataUtf8",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetSomeElseDataUtf8(ctx)
}

// GetFile ...
func (s *instrumentingMiddleware) GetFile(ctx context.Context) (data []byte, fileName string, err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "GetFile",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetFile(ctx)
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc SomeService) SomeService {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
