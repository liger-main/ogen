// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
	"github.com/ogen-go/ogen/types"
	"github.com/ogen-go/ogen/uri"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = math.Mod
	_ = types.Date{}
)

func NewCachingHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodeCachingParams(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.Caching(r.Context(), params)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeCachingResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewDBHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := s.DB(r.Context())
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeDBResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewJSONHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := s.JSON(r.Context())
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeJSONResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewQueriesHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodeQueriesParams(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.Queries(r.Context(), params)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeQueriesResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewUpdatesHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodeUpdatesParams(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.Updates(r.Context(), params)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeUpdatesResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}
