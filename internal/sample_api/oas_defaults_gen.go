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
	"go.opentelemetry.io/otel/codes"
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
	_ = codes.Unset
)

// setDefaults set default value of fields.
func (s *DefaultTest) setDefaults() {
	{
		val := string("required")

		s.Required = val
	}
	{
		val := string("str")

		s.Str.SetTo(val)
	}
	{
		s.NullStr.Null = true
	}
	{
		val := DefaultTestEnum("big")

		s.Enum.SetTo(val)
	}
	{
		val, _ := json.DecodeUUID(jx.DecodeStr("\"123e4567-e89b-12d3-a456-426614174000\""))

		s.UUID.SetTo(val)
	}
	{
		val, _ := json.DecodeIP(jx.DecodeStr("\"1.1.1.1\""))

		s.IP.SetTo(val)
	}
	{
		val, _ := json.DecodeIP(jx.DecodeStr("\"1.1.1.1\""))

		s.IPV4.SetTo(val)
	}
	{
		val, _ := json.DecodeIP(jx.DecodeStr("\"2001:db8:85a3::8a2e:370:7334\""))

		s.IPV6.SetTo(val)
	}
	{
		val, _ := json.DecodeURI(jx.DecodeStr("\"s3://foo/baz\""))

		s.URI.SetTo(val)
	}
	{
		val, _ := json.DecodeDate(jx.DecodeStr("\"2011-10-10\""))

		s.Birthday.SetTo(val)
	}
	{
		val, _ := json.DecodeDuration(jx.DecodeStr("\"5s\""))

		s.Rate.SetTo(val)
	}
	{
		val := string("foo@example.com")

		s.Email.SetTo(val)
	}
	{
		val := string("example.org")

		s.Hostname.SetTo(val)
	}
	{
		val := string("1-2")

		s.Format.SetTo(val)
	}
	{
		val, _ := jx.DecodeStr("\"aGVsbG8sIHdvcmxkIQ==\"").Base64()

		s.Base64 = val
	}
}
