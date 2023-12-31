// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// PublishEvent invokes publishEvent operation.
	//
	// POST /event
	PublishEvent(ctx context.Context, request OptEvent) (*Event, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Operations = &Client{}

type errorHandler interface {
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

var _ Handler = struct {
	errorHandler
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// PublishEvent invokes publishEvent operation.
//
// POST /event
func (c *Client) PublishEvent(ctx context.Context, request OptEvent) (*Event, error) {
	res, err := c.sendPublishEvent(ctx, request)
	return res, err
}

func (c *Client) sendPublishEvent(ctx context.Context, request OptEvent) (res *Event, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("publishEvent"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/event"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "PublishEvent",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	paramsByName := map[string]interface{}{}
	_ = paramsByName

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/event"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodePublishEventRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer func() {
		if err != nil {
			resp.Body.Close()
		}
	}()

	stage = "DecodeResponse"
	result, err := decodePublishEventResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// WebhookClient implements webhook client.
type WebhookClient struct {
	baseClient
}

// NewWebhookClient initializes new WebhookClient.
func NewWebhookClient(opts ...ClientOption) (*WebhookClient, error) {
	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &WebhookClient{
		baseClient: c,
	}, nil
}

// StatusWebhook invokes statusWebhook operation.
func (c *WebhookClient) StatusWebhook(ctx context.Context, targetURL string) (*StatusWebhookOK, error) {
	res, err := c.sendStatusWebhook(ctx, targetURL)
	return res, err
}

func (c *WebhookClient) sendStatusWebhook(ctx context.Context, targetURL string) (res *StatusWebhookOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("statusWebhook"),
		otelogen.WebhookName("status"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "StatusWebhook",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	paramsByName := map[string]interface{}{}
	_ = paramsByName

	stage = "BuildURL"
	u, err := url.Parse(targetURL)
	if err != nil {
		return res, errors.Wrap(err, "parse target URL")
	}
	trimTrailingSlashes(u)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer func() {
		if err != nil {
			resp.Body.Close()
		}
	}()

	stage = "DecodeResponse"
	result, err := decodeStatusWebhookResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateDelete invokes DELETE update operation.
func (c *WebhookClient) UpdateDelete(ctx context.Context, targetURL string) (UpdateDeleteRes, error) {
	res, err := c.sendUpdateDelete(ctx, targetURL)
	return res, err
}

func (c *WebhookClient) sendUpdateDelete(ctx context.Context, targetURL string) (res UpdateDeleteRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.WebhookName("update"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "UpdateDelete",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	paramsByName := map[string]interface{}{}
	_ = paramsByName

	stage = "BuildURL"
	u, err := url.Parse(targetURL)
	if err != nil {
		return res, errors.Wrap(err, "parse target URL")
	}
	trimTrailingSlashes(u)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer func() {
		if err != nil {
			resp.Body.Close()
		}
	}()

	stage = "DecodeResponse"
	result, err := decodeUpdateDeleteResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateWebhook invokes updateWebhook operation.
func (c *WebhookClient) UpdateWebhook(ctx context.Context, targetURL string, request OptEvent, params UpdateWebhookParams) (UpdateWebhookRes, error) {
	res, err := c.sendUpdateWebhook(ctx, targetURL, request, params)
	return res, err
}

func (c *WebhookClient) sendUpdateWebhook(ctx context.Context, targetURL string, request OptEvent, params UpdateWebhookParams) (res UpdateWebhookRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateWebhook"),
		otelogen.WebhookName("update"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "UpdateWebhook",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	paramsByName := map[string]interface{}{}
	_ = paramsByName

	stage = "BuildURL"
	u, err := url.Parse(targetURL)
	if err != nil {
		return res, errors.Wrap(err, "parse target URL")
	}
	trimTrailingSlashes(u)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "event_type" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "event_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			paramsByName["EventType"] = params.EventType
			return e.EncodeValue(conv.StringToString(params.EventType))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateWebhookRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Webhook-Token",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			paramsByName["XWebhookToken"] = params.XWebhookToken
			if val, ok := params.XWebhookToken.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer func() {
		if err != nil {
			resp.Body.Close()
		}
	}()

	stage = "DecodeResponse"
	result, err := decodeUpdateWebhookResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
