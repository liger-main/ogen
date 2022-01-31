// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
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
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func encodeGetBookResponse(response GetBookRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Book:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *GetBookForbidden:
		w.WriteHeader(403)
		return nil
	default:
		return errors.Errorf(`/api/gallery/{book_id}: unexpected response type: %T`, response)
	}
}

func encodeSearchResponse(response SearchRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SearchOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *SearchForbidden:
		w.WriteHeader(403)
		return nil
	default:
		return errors.Errorf(`/api/galleries/search: unexpected response type: %T`, response)
	}
}

func encodeSearchByTagIDResponse(response SearchByTagIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SearchByTagIDOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *SearchByTagIDForbidden:
		w.WriteHeader(403)
		return nil
	default:
		return errors.Errorf(`/api/galleries/tagged: unexpected response type: %T`, response)
	}
}
