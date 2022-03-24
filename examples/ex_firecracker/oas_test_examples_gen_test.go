// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/nonrecording"
	"go.opentelemetry.io/otel/trace"
)

import (
	std "encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
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
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.MeterConfig{}
	_ = syncint64.Counter(nil)
	_ = nonrecording.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

var (
	_ = std.Marshal
	_ = testing.TB(nil)
	_ = require.NoError
)

func TestBalloon_EncodeDecode(t *testing.T) {
	var typ Balloon
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Balloon
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBalloonStats_EncodeDecode(t *testing.T) {
	var typ BalloonStats
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BalloonStats
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBalloonStatsUpdate_EncodeDecode(t *testing.T) {
	var typ BalloonStatsUpdate
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BalloonStatsUpdate
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBalloonUpdate_EncodeDecode(t *testing.T) {
	var typ BalloonUpdate
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BalloonUpdate
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBootSource_EncodeDecode(t *testing.T) {
	var typ BootSource
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BootSource
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCpuTemplate_EncodeDecode(t *testing.T) {
	var typ CpuTemplate
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CpuTemplate
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestDrive_EncodeDecode(t *testing.T) {
	var typ Drive
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Drive
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestError_EncodeDecode(t *testing.T) {
	var typ Error
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Error
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestFullVmConfiguration_EncodeDecode(t *testing.T) {
	var typ FullVmConfiguration
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 FullVmConfiguration
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestInstanceActionInfo_EncodeDecode(t *testing.T) {
	var typ InstanceActionInfo
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 InstanceActionInfo
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestInstanceActionInfoActionType_EncodeDecode(t *testing.T) {
	var typ InstanceActionInfoActionType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 InstanceActionInfoActionType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestInstanceInfo_EncodeDecode(t *testing.T) {
	var typ InstanceInfo
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 InstanceInfo
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestInstanceInfoState_EncodeDecode(t *testing.T) {
	var typ InstanceInfoState
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 InstanceInfoState
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLogger_EncodeDecode(t *testing.T) {
	var typ Logger
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Logger
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLoggerLevel_EncodeDecode(t *testing.T) {
	var typ LoggerLevel
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 LoggerLevel
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMachineConfiguration_EncodeDecode(t *testing.T) {
	var typ MachineConfiguration
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MachineConfiguration
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMetrics_EncodeDecode(t *testing.T) {
	var typ Metrics
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Metrics
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMmdsConfig_EncodeDecode(t *testing.T) {
	var typ MmdsConfig
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MmdsConfig
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMmdsGetOK_EncodeDecode(t *testing.T) {
	var typ MmdsGetOK
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MmdsGetOK
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMmdsPatchReq_EncodeDecode(t *testing.T) {
	var typ MmdsPatchReq
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MmdsPatchReq
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMmdsPutReq_EncodeDecode(t *testing.T) {
	var typ MmdsPutReq
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MmdsPutReq
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestNetworkInterface_EncodeDecode(t *testing.T) {
	var typ NetworkInterface
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 NetworkInterface
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPartialDrive_EncodeDecode(t *testing.T) {
	var typ PartialDrive
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PartialDrive
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPartialNetworkInterface_EncodeDecode(t *testing.T) {
	var typ PartialNetworkInterface
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PartialNetworkInterface
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestRateLimiter_EncodeDecode(t *testing.T) {
	var typ RateLimiter
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 RateLimiter
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSnapshotCreateParams_EncodeDecode(t *testing.T) {
	var typ SnapshotCreateParams
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SnapshotCreateParams
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSnapshotCreateParamsSnapshotType_EncodeDecode(t *testing.T) {
	var typ SnapshotCreateParamsSnapshotType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SnapshotCreateParamsSnapshotType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSnapshotLoadParams_EncodeDecode(t *testing.T) {
	var typ SnapshotLoadParams
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SnapshotLoadParams
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTokenBucket_EncodeDecode(t *testing.T) {
	var typ TokenBucket
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 TokenBucket
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestVM_EncodeDecode(t *testing.T) {
	var typ VM
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 VM
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestVMState_EncodeDecode(t *testing.T) {
	var typ VMState
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 VMState
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestVsock_EncodeDecode(t *testing.T) {
	var typ Vsock
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Vsock
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
