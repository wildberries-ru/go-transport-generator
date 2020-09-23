// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"
	"net/http/pprof"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const (
	httpMethodCreateMultipartUpload = "POST"
	uriPathCreateMultipartUpload    = "/api/v1/multipart/:bucket/:key"
	httpMethodUploadPartDocument    = "PATCH"
	uriPathUploadPartDocument       = "/api/v1/multipart/:bucket/:key"
	httpMethodCompleteUpload        = "PUT"
	uriPathCompleteUpload           = "/api/v1/multipart/:bucket/:key"
	httpMethodUploadDocument        = "POST"
	uriPathUploadDocument           = "/api/v1/doc/:bucket/:key"
	httpMethodDownloadDocument      = "GET"
	uriPathDownloadDocument         = "/api/v1/doc/:bucket/:key"
	httpMethodGetToken              = "POST"
	uriPathGetToken                 = "/token"
	httpMethodGetBranches           = "GET"
	uriPathGetBranches              = "/api/v1/branches"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type errorCreator func(err error) error

// New ...
func New(router *fasthttprouter.Router, svc service, decodeJSONErrorCreator errorCreator, encodeJSONErrorCreator errorCreator, decodeTypeIntErrorCreator errorCreator, errorProcessor errorProcessor) {

	createMultipartUploadTransport := NewCreateMultipartUploadTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodCreateMultipartUpload, uriPathCreateMultipartUpload, NewCreateMultipartUpload(createMultipartUploadTransport, svc, errorProcessor))

	uploadPartDocumentTransport := NewUploadPartDocumentTransport()
	router.Handle(httpMethodUploadPartDocument, uriPathUploadPartDocument, NewUploadPartDocument(uploadPartDocumentTransport, svc, errorProcessor))

	completeUploadTransport := NewCompleteUploadTransport()
	router.Handle(httpMethodCompleteUpload, uriPathCompleteUpload, NewCompleteUpload(completeUploadTransport, svc, errorProcessor))

	uploadDocumentTransport := NewUploadDocumentTransport()
	router.Handle(httpMethodUploadDocument, uriPathUploadDocument, NewUploadDocument(uploadDocumentTransport, svc, errorProcessor))

	downloadDocumentTransport := NewDownloadDocumentTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodDownloadDocument, uriPathDownloadDocument, NewDownloadDocument(downloadDocumentTransport, svc, errorProcessor))

	getTokenTransport := NewGetTokenTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodGetToken, uriPathGetToken, NewGetToken(getTokenTransport, svc, errorProcessor))

	getBranchesTransport := NewGetBranchesTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodGetBranches, uriPathGetBranches, NewGetBranches(getBranchesTransport, svc, errorProcessor))

	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
}
