// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"math/big"
	"net/http"
	"regexp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

var regexMap = map[string]*regexp.Regexp{
	"^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$": regexp.MustCompile("^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$"),
	"^\\d-\\d$": regexp.MustCompile("^\\d-\\d$"),
	"foo.*":     regexp.MustCompile("foo.*"),
	"string_.*": regexp.MustCompile("string_.*"),
}
var ratMap = map[string]*big.Rat{
	"10": func() *big.Rat {
		r, ok := new(big.Rat).SetString("10")
		if !ok {
			panic(fmt.Sprintf("rat %q: can't parse", "10"))
		}
		return r
	}(),
	"5": func() *big.Rat {
		r, ok := new(big.Rat).SetString("5")
		if !ok {
			panic(fmt.Sprintf("rat %q: can't parse", "5"))
		}
		return r
	}(),
}
var (
	// Allocate option closure once.
	serverSpanKind = trace.WithSpanKind(trace.SpanKindServer)
)

// ErrorHandler is error handler.
type ErrorHandler = ogenerrors.ErrorHandler

type config struct {
	TracerProvider     trace.TracerProvider
	Tracer             trace.Tracer
	MeterProvider      metric.MeterProvider
	Meter              metric.Meter
	NotFound           http.HandlerFunc
	MethodNotAllowed   func(w http.ResponseWriter, r *http.Request, allowed string)
	ErrorHandler       ErrorHandler
	Prefix             string
	Middleware         Middleware
	MaxMultipartMemory int64
}

func newConfig(opts ...Option) config {
	cfg := config{
		TracerProvider: otel.GetTracerProvider(),
		MeterProvider:  metric.NewNoopMeterProvider(),
		NotFound:       http.NotFound,
		MethodNotAllowed: func(w http.ResponseWriter, r *http.Request, allowed string) {
			w.Header().Set("Allow", allowed)
			w.WriteHeader(http.StatusMethodNotAllowed)
		},
		ErrorHandler:       ogenerrors.DefaultErrorHandler,
		Middleware:         nil,
		MaxMultipartMemory: 32 << 20, // 32 MB
	}
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	cfg.Tracer = cfg.TracerProvider.Tracer(otelogen.Name,
		trace.WithInstrumentationVersion(otelogen.SemVersion()),
	)
	cfg.Meter = cfg.MeterProvider.Meter(otelogen.Name)
	return cfg
}

type baseServer struct {
	cfg      config
	requests syncint64.Counter
	errors   syncint64.Counter
	duration syncint64.Histogram
}

func (s baseServer) notFound(w http.ResponseWriter, r *http.Request) {
	s.cfg.NotFound(w, r)
}

func (s baseServer) notAllowed(w http.ResponseWriter, r *http.Request, allowed string) {
	s.cfg.MethodNotAllowed(w, r, allowed)
}

func (cfg config) baseServer() (s baseServer, err error) {
	s = baseServer{cfg: cfg}
	if s.requests, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerRequestCount); err != nil {
		return s, err
	}
	if s.errors, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerErrorsCount); err != nil {
		return s, err
	}
	if s.duration, err = s.cfg.Meter.SyncInt64().Histogram(otelogen.ServerDuration); err != nil {
		return s, err
	}
	return s, nil
}

// Option is config option.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
//
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.TracerProvider = provider
		}
	})
}

// WithMeterProvider specifies a meter provider to use for creating a meter.
//
// If none is specified, the metric.NewNoopMeterProvider is used.
func WithMeterProvider(provider metric.MeterProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.MeterProvider = provider
		}
	})
}

// WithNotFound specifies Not Found handler to use.
func WithNotFound(notFound http.HandlerFunc) Option {
	return optionFunc(func(cfg *config) {
		if notFound != nil {
			cfg.NotFound = notFound
		}
	})
}

// WithMethodNotAllowed specifies Method Not Allowed handler to use.
func WithMethodNotAllowed(methodNotAllowed func(w http.ResponseWriter, r *http.Request, allowed string)) Option {
	return optionFunc(func(cfg *config) {
		if methodNotAllowed != nil {
			cfg.MethodNotAllowed = methodNotAllowed
		}
	})
}

// WithErrorHandler specifies error handler to use.
func WithErrorHandler(h ErrorHandler) Option {
	return optionFunc(func(cfg *config) {
		if h != nil {
			cfg.ErrorHandler = h
		}
	})
}

// WithPathPrefix specifies server path prefix.
func WithPathPrefix(prefix string) Option {
	return optionFunc(func(cfg *config) {
		cfg.Prefix = prefix
	})
}

// WithMiddleware specifies middlewares to use.
func WithMiddleware(m ...Middleware) Option {
	return optionFunc(func(cfg *config) {
		switch len(m) {
		case 0:
			cfg.Middleware = nil
		case 1:
			cfg.Middleware = m[0]
		default:
			cfg.Middleware = middleware.ChainMiddlewares(m...)
		}
	})
}

// WithMaxMultipartMemory specifies limit of memory for storing file parts.
// File parts which can't be stored in memory will be stored on disk in temporary files.
func WithMaxMultipartMemory(max int64) Option {
	return optionFunc(func(cfg *config) {
		if max > 0 {
			cfg.MaxMultipartMemory = max
		}
	})
}
