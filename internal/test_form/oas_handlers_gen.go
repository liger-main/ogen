// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// HandleTestFormURLEncodedRequest handles testFormURLEncoded operation.
//
// POST /testFormURLEncoded
func (s *Server) handleTestFormURLEncodedRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testFormURLEncoded"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TestFormURLEncoded",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "TestFormURLEncoded",
			ID:   "testFormURLEncoded",
		}
	)
	request, close, err := s.decodeTestFormURLEncodedRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response TestFormURLEncodedOK
	if m := s.cfg.Middleware; m != nil {
		mreq := MiddlewareRequest{
			Context:       ctx,
			OperationName: "TestFormURLEncoded",
			OperationID:   "testFormURLEncoded",
			Body:          request,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = TestForm
			Params   = struct{}
			Response = TestFormURLEncodedOK
		)
		response, err = hookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			struct{}{},
			mreq,
			func(ctx context.Context, params Params, request Request) (Response, error) {
				return s.h.TestFormURLEncoded(ctx, request)
			},
		)
	} else {
		response, err = s.h.TestFormURLEncoded(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestFormURLEncodedResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleTestMultipartRequest handles testMultipart operation.
//
// POST /testMultipart
func (s *Server) handleTestMultipartRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testMultipart"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TestMultipart",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "TestMultipart",
			ID:   "testMultipart",
		}
	)
	request, close, err := s.decodeTestMultipartRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response TestMultipartOK
	if m := s.cfg.Middleware; m != nil {
		mreq := MiddlewareRequest{
			Context:       ctx,
			OperationName: "TestMultipart",
			OperationID:   "testMultipart",
			Body:          request,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = TestForm
			Params   = struct{}
			Response = TestMultipartOK
		)
		response, err = hookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			struct{}{},
			mreq,
			func(ctx context.Context, params Params, request Request) (Response, error) {
				return s.h.TestMultipart(ctx, request)
			},
		)
	} else {
		response, err = s.h.TestMultipart(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestMultipartResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleTestMultipartUploadRequest handles testMultipartUpload operation.
//
// POST /testMultipartUpload
func (s *Server) handleTestMultipartUploadRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testMultipartUpload"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TestMultipartUpload",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "TestMultipartUpload",
			ID:   "testMultipartUpload",
		}
	)
	request, close, err := s.decodeTestMultipartUploadRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response TestMultipartUploadOK
	if m := s.cfg.Middleware; m != nil {
		mreq := MiddlewareRequest{
			Context:       ctx,
			OperationName: "TestMultipartUpload",
			OperationID:   "testMultipartUpload",
			Body:          request,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = TestMultipartUploadReqForm
			Params   = struct{}
			Response = TestMultipartUploadOK
		)
		response, err = hookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			struct{}{},
			mreq,
			func(ctx context.Context, params Params, request Request) (Response, error) {
				return s.h.TestMultipartUpload(ctx, request)
			},
		)
	} else {
		response, err = s.h.TestMultipartUpload(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestMultipartUploadResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleTestShareFormSchemaRequest handles testShareFormSchema operation.
//
// POST /testShareFormSchema
func (s *Server) handleTestShareFormSchemaRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testShareFormSchema"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TestShareFormSchema",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "TestShareFormSchema",
			ID:   "testShareFormSchema",
		}
	)
	request, close, err := s.decodeTestShareFormSchemaRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response TestShareFormSchemaOK
	if m := s.cfg.Middleware; m != nil {
		mreq := MiddlewareRequest{
			Context:       ctx,
			OperationName: "TestShareFormSchema",
			OperationID:   "testShareFormSchema",
			Body:          request,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = TestShareFormSchemaReq
			Params   = struct{}
			Response = TestShareFormSchemaOK
		)
		response, err = hookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			struct{}{},
			mreq,
			func(ctx context.Context, params Params, request Request) (Response, error) {
				return s.h.TestShareFormSchema(ctx, request)
			},
		)
	} else {
		response, err = s.h.TestShareFormSchema(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestShareFormSchemaResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
