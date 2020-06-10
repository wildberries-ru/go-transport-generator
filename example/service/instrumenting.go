// Package service ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package service

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
	svc         SomeService
}

func (s *instrumentingMiddleware) GetWarehouses(ctx context.Context) (pets []v1.Detail, err error) {
	defer s.recordMetrics("GetWarehouses", time.Now(), err)
	return s.svc.GetWarehouses(ctx)
}

func (s *instrumentingMiddleware) GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error) {
	defer s.recordMetrics("GetDetails", time.Now(), err)
	return s.svc.GetDetails(ctx, namespace, detail, fileID, someID, token)
}

func (s *instrumentingMiddleware) GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error) {
	defer s.recordMetrics("GetDetailsEmbedStruct", time.Now(), err)
	return s.svc.GetDetailsEmbedStruct(ctx, namespace, detail)
}

func (s *instrumentingMiddleware) GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error) {
	defer s.recordMetrics("GetDetailsListEmbedStruct", time.Now(), err)
	return s.svc.GetDetailsListEmbedStruct(ctx, namespace, detail)
}

func (s *instrumentingMiddleware) PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer s.recordMetrics("PutDetails", time.Now(), err)
	return s.svc.PutDetails(ctx, namespace, detail, testID, blaID, token, pretty, yang)
}

func (s *instrumentingMiddleware) GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	defer s.recordMetrics("GetSomeElseDataUtf8", time.Now(), err)
	return s.svc.GetSomeElseDataUtf8(ctx)
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
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc SomeService) SomeService {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
