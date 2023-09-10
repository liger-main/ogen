// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeCreatePetsResponse(response *CreatePetsCreated, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(201)
	span.SetStatus(codes.Ok, http.StatusText(201))

	return nil
}

func encodeListPetsResponse(response *PetsHeaders, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "x-next" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "x-next",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				if val, ok := response.XNext.Get(); ok {
					return e.EncodeValue(conv.StringToString(val))
				}
				return nil
			}); err != nil {
				return errors.Wrap(err, "encode x-next header")
			}
		}
	}
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeShowPetByIdResponse(response *Pet, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeErrorResponse(response *ErrorStatusCode, w http.ResponseWriter, span trace.Span) error {
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

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil

}
