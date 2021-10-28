// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
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
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
)

func decodeGetBookParams(r *http.Request) (GetBookParams, error) {
	var params GetBookParams
	// Decode param "book_id" located in "Path".
	if err := func() error {
		param := chi.URLParam(r, "book_id")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'book_id' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "book_id",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeInt()
		if err != nil {
			return err
		}

		params.BookID = int(v)
		return nil
	}(); err != nil {
		return params, err
	}
	return params, nil
}

func decodeGetPageCoverImageParams(r *http.Request) (GetPageCoverImageParams, error) {
	var params GetPageCoverImageParams
	// Decode param "media_id" located in "Path".
	if err := func() error {
		param := chi.URLParam(r, "media_id")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'media_id' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "media_id",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeInt()
		if err != nil {
			return err
		}

		params.MediaID = int(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param "format" located in "Path".
	if err := func() error {
		param := chi.URLParam(r, "format")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'format' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "format",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeString()
		if err != nil {
			return err
		}

		params.Format = string(v)
		return nil
	}(); err != nil {
		return params, err
	}
	return params, nil
}

func decodeGetPageImageParams(r *http.Request) (GetPageImageParams, error) {
	var params GetPageImageParams
	// Decode param "media_id" located in "Path".
	if err := func() error {
		param := chi.URLParam(r, "media_id")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'media_id' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "media_id",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeInt()
		if err != nil {
			return err
		}

		params.MediaID = int(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param "page" located in "Path".
	if err := func() error {
		param := chi.URLParam(r, "page")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'page' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "page",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeInt()
		if err != nil {
			return err
		}

		params.Page = int(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param "format" located in "Path".
	if err := func() error {
		param := chi.URLParam(r, "format")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'format' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "format",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeString()
		if err != nil {
			return err
		}

		params.Format = string(v)
		return nil
	}(); err != nil {
		return params, err
	}
	return params, nil
}

func decodeGetPageThumbnailImageParams(r *http.Request) (GetPageThumbnailImageParams, error) {
	var params GetPageThumbnailImageParams
	// Decode param "media_id" located in "Path".
	if err := func() error {
		param := chi.URLParam(r, "media_id")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'media_id' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "media_id",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeInt()
		if err != nil {
			return err
		}

		params.MediaID = int(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param "page" located in "Path".
	if err := func() error {
		param := chi.URLParam(r, "page")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'page' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "page",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeInt()
		if err != nil {
			return err
		}

		params.Page = int(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param "format" located in "Path".
	if err := func() error {
		param := chi.URLParam(r, "format")
		if len(param) == 0 {
			return fmt.Errorf("path parameter 'format' not specified")
		}

		d := uri.NewPathDecoder(uri.PathDecoderConfig{
			Param:   "format",
			Value:   param,
			Style:   uri.PathStyleSimple,
			Explode: false,
		})

		v, err := d.DecodeString()
		if err != nil {
			return err
		}

		params.Format = string(v)
		return nil
	}(); err != nil {
		return params, err
	}
	return params, nil
}

func decodeSearchParams(r *http.Request) (SearchParams, error) {
	var params SearchParams
	// Decode param "query" located in "Query".
	if err := func() error {
		values, ok := r.URL.Query()["query"]
		if !ok {
			return fmt.Errorf("query parameter 'query' not specified")
		}

		d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
			Values:  values,
			Style:   uri.QueryStyleForm,
			Explode: true,
		})

		v, err := d.DecodeString()
		if err != nil {
			return err
		}

		params.Query = string(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param "page" located in "Query".
	if err := func() error {
		values, ok := r.URL.Query()["page"]
		if !ok {
			return nil
		}

		d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
			Values:  values,
			Style:   uri.QueryStyleForm,
			Explode: true,
		})

		v, err := d.DecodeInt()
		if err != nil {
			return err
		}

		params.Page = int(v)
		return nil
	}(); err != nil {
		return params, err
	}
	return params, nil
}

func decodeSearchByTagIDParams(r *http.Request) (SearchByTagIDParams, error) {
	var params SearchByTagIDParams
	// Decode param "tag_id" located in "Query".
	if err := func() error {
		values, ok := r.URL.Query()["tag_id"]
		if !ok {
			return fmt.Errorf("query parameter 'tag_id' not specified")
		}

		d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
			Values:  values,
			Style:   uri.QueryStyleForm,
			Explode: true,
		})

		v, err := d.DecodeInt()
		if err != nil {
			return err
		}

		params.TagID = int(v)
		return nil
	}(); err != nil {
		return params, err
	}
	// Decode param "page" located in "Query".
	if err := func() error {
		values, ok := r.URL.Query()["page"]
		if !ok {
			return nil
		}

		d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
			Values:  values,
			Style:   uri.QueryStyleForm,
			Explode: true,
		})

		v, err := d.DecodeInt()
		if err != nil {
			return err
		}

		params.Page = int(v)
		return nil
	}(); err != nil {
		return params, err
	}
	return params, nil
}
