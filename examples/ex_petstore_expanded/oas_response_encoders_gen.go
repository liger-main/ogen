// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeDeletePetResponse(response DeletePetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeletePetNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/pets/{id}"+`: unexpected response type: %T`, response)
	}
}
