// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
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
)

func NewFoobarGetHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := DecodeFoobarGetParams(r)
		if err != nil {
			respondError(w, err)
			return
		}

		response, err := s.FoobarGet(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := EncodeFoobarGetResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewFoobarPutHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := s.FoobarPut(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := EncodeFoobarPutResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewFoobarPostHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := DecodeFoobarPostRequest(r)
		if err != nil {
			respondError(w, err)
			return
		}

		response, err := s.FoobarPost(r.Context(), request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := EncodeFoobarPostResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewPetGetHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := DecodePetGetParams(r)
		if err != nil {
			respondError(w, err)
			return
		}

		response, err := s.PetGet(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := EncodePetGetResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewPetCreateHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := DecodePetCreateRequest(r)
		if err != nil {
			respondError(w, err)
			return
		}

		response, err := s.PetCreate(r.Context(), request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := EncodePetCreateResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewPetGetByNameHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := DecodePetGetByNameParams(r)
		if err != nil {
			respondError(w, err)
			return
		}

		response, err := s.PetGetByName(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := EncodePetGetByNameResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func respondError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
}
