// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	"go.opentelemetry.io/otel/metric/instrument/syncint64"

	"github.com/ogen-go/ogen/otelogen"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// NullableStrings implements nullableStrings operation.
	//
	// Nullable strings.
	//
	// POST /nullableStrings
	NullableStrings(ctx context.Context, req string) (NullableStringsOK, error)
	// ObjectsWithConflictingArrayProperty implements objectsWithConflictingArrayProperty operation.
	//
	// Objects with conflicting array property.
	//
	// POST /objectsWithConflictingArrayProperty
	ObjectsWithConflictingArrayProperty(ctx context.Context, req ObjectsWithConflictingArrayPropertyReq) (ObjectsWithConflictingArrayPropertyOK, error)
	// ObjectsWithConflictingProperties implements objectsWithConflictingProperties operation.
	//
	// Objects with conflicting properties.
	//
	// POST /objectsWithConflictingProperties
	ObjectsWithConflictingProperties(ctx context.Context, req ObjectsWithConflictingPropertiesReq) (ObjectsWithConflictingPropertiesOK, error)
	// ReferencedAllof implements referencedAllof operation.
	//
	// Referenced allOf.
	//
	// POST /referencedAllof
	ReferencedAllof(ctx context.Context, req ReferencedAllofReq) (ReferencedAllofOK, error)
	// ReferencedAllofOptional implements referencedAllofOptional operation.
	//
	// Referenced allOf, but requestBody is not required.
	//
	// POST /referencedAllofOptional
	ReferencedAllofOptional(ctx context.Context, req ReferencedAllofOptionalReq) (ReferencedAllofOptionalOK, error)
	// SimpleInteger implements simpleInteger operation.
	//
	// Simple integers with validation.
	//
	// POST /simpleInteger
	SimpleInteger(ctx context.Context, req int) (SimpleIntegerOK, error)
	// SimpleObjects implements simpleObjects operation.
	//
	// Simple objects.
	//
	// POST /simpleObjects
	SimpleObjects(ctx context.Context, req SimpleObjectsReq) (SimpleObjectsOK, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config

	requests syncint64.Counter
	errors   syncint64.Counter
	duration syncint64.Histogram
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...Option) (*Server, error) {
	s := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	var err error
	if s.requests, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerRequestCount); err != nil {
		return nil, err
	}
	if s.errors, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerErrorsCount); err != nil {
		return nil, err
	}
	if s.duration, err = s.cfg.Meter.SyncInt64().Histogram(otelogen.ServerDuration); err != nil {
		return nil, err
	}
	return s, nil
}
