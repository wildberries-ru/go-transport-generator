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
	reqCountUploadDocument    metrics.Counter
	reqDurationUploadDocument metrics.Histogram

	reqCountGetWarehouses    metrics.Counter
	reqDurationGetWarehouses metrics.Histogram

	reqCountGetDetails    metrics.Counter
	reqDurationGetDetails metrics.Histogram

	reqCountGetDetailsEmbedStruct    metrics.Counter
	reqDurationGetDetailsEmbedStruct metrics.Histogram

	reqCountGetDetailsListEmbedStruct    metrics.Counter
	reqDurationGetDetailsListEmbedStruct metrics.Histogram

	reqCountPutDetails    metrics.Counter
	reqDurationPutDetails metrics.Histogram

	reqCountGetSomeElseDataUtf8    metrics.Counter
	reqDurationGetSomeElseDataUtf8 metrics.Histogram

	reqCountGetFile    metrics.Counter
	reqDurationGetFile metrics.Histogram

	svc SomeService
}

// UploadDocument ...
func (s *instrumentingMiddleware) UploadDocument(ctx context.Context, token *string, name string, extension string, categoryID string, supplierID *int64, contractID *int64, data *multipart.FileHeader) (err error) {
	defer func(startTime time.Time) {

		labels := []string{
			"method", "UploadDocument",
			"error", strconv.FormatBool(err != nil),
		}
		s.reqCountUploadDocument.With(labels...).Add(1)
		s.reqDurationUploadDocument.With(labels...).Observe(time.Since(startTime).Seconds())
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
		s.reqCountGetWarehouses.With(labels...).Add(1)
		s.reqDurationGetWarehouses.With(labels...).Observe(time.Since(startTime).Seconds())
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
		s.reqCountGetDetails.With(labels...).Add(1)
		s.reqDurationGetDetails.With(labels...).Observe(time.Since(startTime).Seconds())
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
		s.reqCountGetDetailsEmbedStruct.With(labels...).Add(1)
		s.reqDurationGetDetailsEmbedStruct.With(labels...).Observe(time.Since(startTime).Seconds())
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
		s.reqCountGetDetailsListEmbedStruct.With(labels...).Add(1)
		s.reqDurationGetDetailsListEmbedStruct.With(labels...).Observe(time.Since(startTime).Seconds())
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
		s.reqCountPutDetails.With(labels...).Add(1)
		s.reqDurationPutDetails.With(labels...).Observe(time.Since(startTime).Seconds())
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
		s.reqCountGetSomeElseDataUtf8.With(labels...).Add(1)
		s.reqDurationGetSomeElseDataUtf8.With(labels...).Observe(time.Since(startTime).Seconds())
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
		s.reqCountGetFile.With(labels...).Add(1)
		s.reqDurationGetFile.With(labels...).Observe(time.Since(startTime).Seconds())
	}(time.Now())
	return s.svc.GetFile(ctx)
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(
	metricsNamespace string,
	metricsSubsystem string,
	metricsNameCount string,
	metricsNameCountHelp string,
	metricsNameDuration string,
	metricsNameDurationHelp string,
	svc SomeService,
) SomeService {

	reqCountUploadDocument := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameCount,
		Help:      metricsNameCountHelp,
	}, []string{
		"method",
		"error",
	})
	reqDurationUploadDocument := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameDuration,
		Help:      metricsNameDurationHelp,
	}, []string{
		"method",
		"error",
	})

	reqCountGetWarehouses := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameCount,
		Help:      metricsNameCountHelp,
	}, []string{
		"method",
		"error",
	})
	reqDurationGetWarehouses := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameDuration,
		Help:      metricsNameDurationHelp,
	}, []string{
		"method",
		"error",
	})

	reqCountGetDetails := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameCount,
		Help:      metricsNameCountHelp,
	}, []string{
		"method",
		"error",

		"detail",
		"fileID",
		"namespace",
		"someID",
		"token",
	})
	reqDurationGetDetails := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameDuration,
		Help:      metricsNameDurationHelp,
	}, []string{
		"method",
		"error",

		"detail",
		"fileID",
		"namespace",
		"someID",
		"token",
	})

	reqCountGetDetailsEmbedStruct := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameCount,
		Help:      metricsNameCountHelp,
	}, []string{
		"method",
		"error",
	})
	reqDurationGetDetailsEmbedStruct := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameDuration,
		Help:      metricsNameDurationHelp,
	}, []string{
		"method",
		"error",
	})

	reqCountGetDetailsListEmbedStruct := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameCount,
		Help:      metricsNameCountHelp,
	}, []string{
		"method",
		"error",
	})
	reqDurationGetDetailsListEmbedStruct := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameDuration,
		Help:      metricsNameDurationHelp,
	}, []string{
		"method",
		"error",
	})

	reqCountPutDetails := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameCount,
		Help:      metricsNameCountHelp,
	}, []string{
		"method",
		"error",
	})
	reqDurationPutDetails := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameDuration,
		Help:      metricsNameDurationHelp,
	}, []string{
		"method",
		"error",
	})

	reqCountGetSomeElseDataUtf8 := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameCount,
		Help:      metricsNameCountHelp,
	}, []string{
		"method",
		"error",
	})
	reqDurationGetSomeElseDataUtf8 := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameDuration,
		Help:      metricsNameDurationHelp,
	}, []string{
		"method",
		"error",
	})

	reqCountGetFile := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameCount,
		Help:      metricsNameCountHelp,
	}, []string{
		"method",
		"error",
	})
	reqDurationGetFile := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      metricsNameDuration,
		Help:      metricsNameDurationHelp,
	}, []string{
		"method",
		"error",
	})

	return &instrumentingMiddleware{

		reqCountUploadDocument:    reqCountUploadDocument,
		reqDurationUploadDocument: reqDurationUploadDocument,

		reqCountGetWarehouses:    reqCountGetWarehouses,
		reqDurationGetWarehouses: reqDurationGetWarehouses,

		reqCountGetDetails:    reqCountGetDetails,
		reqDurationGetDetails: reqDurationGetDetails,

		reqCountGetDetailsEmbedStruct:    reqCountGetDetailsEmbedStruct,
		reqDurationGetDetailsEmbedStruct: reqDurationGetDetailsEmbedStruct,

		reqCountGetDetailsListEmbedStruct:    reqCountGetDetailsListEmbedStruct,
		reqDurationGetDetailsListEmbedStruct: reqDurationGetDetailsListEmbedStruct,

		reqCountPutDetails:    reqCountPutDetails,
		reqDurationPutDetails: reqDurationPutDetails,

		reqCountGetSomeElseDataUtf8:    reqCountGetSomeElseDataUtf8,
		reqDurationGetSomeElseDataUtf8: reqDurationGetSomeElseDataUtf8,

		reqCountGetFile:    reqCountGetFile,
		reqDurationGetFile: reqDurationGetFile,

		svc: svc,
	}
}
