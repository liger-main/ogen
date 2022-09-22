// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeAddPetResponse(response AddPetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
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
		return errors.Errorf("/pets"+`: unexpected response type: %T`, response)
	}
}
func encodeDeletePetResponse(response DeletePetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeletePetNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
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
func encodeFindPetByIDResponse(response FindPetByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
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
func encodeFindPetsResponse(response FindPetsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *FindPetsOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
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
		return errors.Errorf("/pets"+`: unexpected response type: %T`, response)
	}
}
