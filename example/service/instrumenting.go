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
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         SomeService
}

// UploadDocument ...
func (s *instrumentingMiddleware) UploadDocument(ctx context.Context, token *string, name []string, extension string, categoryID string, supplierID []int64, contractID *bool, data *multipart.FileHeader) (err error) {
	defer func(startTime time.Time) {

		var _token string
		if token != nil {

			_token = *token

		} else {
			_token = "empty"
		}

		labels := []string{
			"method", "UploadDocument",
			"error", strconv.FormatBool(err != nil),

			"token", _token,
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.UploadDocument(ctx, token, name, extension, categoryID, supplierID, contractID, data)
}

// GetWarehouses ...
func (s *instrumentingMiddleware) GetWarehouses(ctx context.Context, token *string) (pets map[string]v1.Detail, someCookie *string, err error) {
	defer func(startTime time.Time) {

		var _token string
		if token != nil {

			_token = *token

		} else {
			_token = "empty"
		}

		labels := []string{
			"method", "GetWarehouses",
			"error", strconv.FormatBool(err != nil),

			"token", _token,
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetWarehouses(ctx, token)
}

// GetDetails ...
func (s *instrumentingMiddleware) GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error) {
	defer func(startTime time.Time) {

		var _token string
		if token != nil {

			_token = *token

		} else {
			_token = "empty"
		}

		labels := []string{
			"method", "GetDetails",
			"error", strconv.FormatBool(err != nil),

			"token", _token,
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetDetails(ctx, namespace, detail, fileID, someID, token)
}

// GetDetailsEmbedStruct ...
func (s *instrumentingMiddleware) GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string, token *string) (response v1.GetDetailsEmbedStructResponse, err error) {
	defer func(startTime time.Time) {

		var _token string
		if token != nil {

			_token = *token

		} else {
			_token = "empty"
		}

		labels := []string{
			"method", "GetDetailsEmbedStruct",
			"error", strconv.FormatBool(err != nil),

			"token", _token,
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetDetailsEmbedStruct(ctx, namespace, detail, token)
}

// GetDetailsListEmbedStruct ...
func (s *instrumentingMiddleware) GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string, token *string) (details []v1.Detail, err error) {
	defer func(startTime time.Time) {

		var _token string
		if token != nil {

			_token = *token

		} else {
			_token = "empty"
		}

		labels := []string{
			"method", "GetDetailsListEmbedStruct",
			"error", strconv.FormatBool(err != nil),

			"token", _token,
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetDetailsListEmbedStruct(ctx, namespace, detail, token)
}

// PutDetails ...
func (s *instrumentingMiddleware) PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer func(startTime time.Time) {

		var _token string
		if token != nil {

			_token = *token

		} else {
			_token = "empty"
		}

		labels := []string{
			"method", "PutDetails",
			"error", strconv.FormatBool(err != nil),

			"token", _token,
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.PutDetails(ctx, namespace, detail, testID, blaID, token, pretty, yang)
}

// GetSomeElseDataUtf8 ...
func (s *instrumentingMiddleware) GetSomeElseDataUtf8(ctx context.Context, token *string) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer func(startTime time.Time) {

		var _token string
		if token != nil {

			_token = *token

		} else {
			_token = "empty"
		}

		labels := []string{
			"method", "GetSomeElseDataUtf8",
			"error", strconv.FormatBool(err != nil),

			"token", _token,
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetSomeElseDataUtf8(ctx, token)
}

// GetFile ...
func (s *instrumentingMiddleware) GetFile(ctx context.Context, token *string) (data []byte, fileName string, err error) {
	defer func(startTime time.Time) {

		var _token string
		if token != nil {

			_token = *token

		} else {
			_token = "empty"
		}

		labels := []string{
			"method", "GetFile",
			"error", strconv.FormatBool(err != nil),

			"token", _token,
		}
		s.reqCount.With(labels...).Add(1)
		s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetFile(ctx, token)
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
	svc SomeService,
) SomeService {
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
