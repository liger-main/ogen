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

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DataGetFormat implements dataGetFormat operation.
	//
	// Retrieve data.
	//
	// GET /name/{id}/{foo}1234{bar}-{baz}!{kek}
	DataGetFormat(ctx context.Context, params DataGetFormatParams) (string, error)
	// DefaultTest implements defaultTest operation.
	//
	// POST /defaultTest
	DefaultTest(ctx context.Context, req DefaultTest, params DefaultTestParams) (int32, error)
	// ErrorGet implements errorGet operation.
	//
	// Returns error.
	//
	// GET /error
	ErrorGet(ctx context.Context) (ErrorStatusCode, error)
	// FoobarGet implements foobarGet operation.
	//
	// Dumb endpoint for testing things.
	//
	// GET /foobar
	FoobarGet(ctx context.Context, params FoobarGetParams) (FoobarGetRes, error)
	// FoobarPost implements foobarPost operation.
	//
	// Dumb endpoint for testing things.
	//
	// POST /foobar
	FoobarPost(ctx context.Context, req OptPet) (FoobarPostRes, error)
	// FoobarPut implements  operation.
	//
	// PUT /foobar
	FoobarPut(ctx context.Context) (FoobarPutDefStatusCode, error)
	// GetHeader implements getHeader operation.
	//
	// Test for header param.
	//
	// GET /test/header
	GetHeader(ctx context.Context, params GetHeaderParams) (Hash, error)
	// MultipleGenericResponses implements multipleGenericResponses operation.
	//
	// GET /multipleGenericResponses
	MultipleGenericResponses(ctx context.Context) (MultipleGenericResponsesRes, error)
	// NullableDefaultResponse implements nullableDefaultResponse operation.
	//
	// GET /nullableDefaultResponse
	NullableDefaultResponse(ctx context.Context) (NilIntStatusCode, error)
	// OctetStreamBinaryStringSchema implements octetStreamBinaryStringSchema operation.
	//
	// GET /octetStreamBinaryStringSchema
	OctetStreamBinaryStringSchema(ctx context.Context) (OctetStreamBinaryStringSchemaOK, error)
	// OctetStreamEmptySchema implements octetStreamEmptySchema operation.
	//
	// GET /octetStreamEmptySchema
	OctetStreamEmptySchema(ctx context.Context) (OctetStreamEmptySchemaOK, error)
	// OneofBug implements oneofBug operation.
	//
	// POST /oneofBug
	OneofBug(ctx context.Context, req OneOfBugs) (OneofBugOK, error)
	// PetCreate implements petCreate operation.
	//
	// Creates pet.
	//
	// POST /pet
	PetCreate(ctx context.Context, req OptPet) (Pet, error)
	// PetFriendsNamesByID implements petFriendsNamesByID operation.
	//
	// Returns names of all friends of pet.
	//
	// GET /pet/friendNames/{id}
	PetFriendsNamesByID(ctx context.Context, params PetFriendsNamesByIDParams) ([]string, error)
	// PetGet implements petGet operation.
	//
	// Returns pet from the system that the user has access to.
	//
	// GET /pet
	PetGet(ctx context.Context, params PetGetParams) (PetGetRes, error)
	// PetGetAvatarByID implements petGetAvatarByID operation.
	//
	// Returns pet avatar by id.
	//
	// GET /pet/avatar
	PetGetAvatarByID(ctx context.Context, params PetGetAvatarByIDParams) (PetGetAvatarByIDRes, error)
	// PetGetAvatarByName implements petGetAvatarByName operation.
	//
	// Returns pet's avatar by name.
	//
	// GET /pet/{name}/avatar
	PetGetAvatarByName(ctx context.Context, params PetGetAvatarByNameParams) (PetGetAvatarByNameRes, error)
	// PetGetByName implements petGetByName operation.
	//
	// Returns pet by name from the system that the user has access to.
	//
	// GET /pet/{name}
	PetGetByName(ctx context.Context, params PetGetByNameParams) (Pet, error)
	// PetNameByID implements petNameByID operation.
	//
	// Returns pet name by pet id.
	//
	// GET /pet/name/{id}
	PetNameByID(ctx context.Context, params PetNameByIDParams) (string, error)
	// PetUpdateNameAliasPost implements  operation.
	//
	// POST /pet/updateNameAlias
	PetUpdateNameAliasPost(ctx context.Context, req OptPetName) (PetUpdateNameAliasPostDefStatusCode, error)
	// PetUpdateNamePost implements  operation.
	//
	// POST /pet/updateName
	PetUpdateNamePost(ctx context.Context, req OptString) (PetUpdateNamePostDefStatusCode, error)
	// PetUploadAvatarByID implements petUploadAvatarByID operation.
	//
	// Uploads pet avatar by id.
	//
	// POST /pet/avatar
	PetUploadAvatarByID(ctx context.Context, req PetUploadAvatarByIDReq, params PetUploadAvatarByIDParams) (PetUploadAvatarByIDRes, error)
	// RecursiveMapGet implements  operation.
	//
	// GET /recursiveMap
	RecursiveMapGet(ctx context.Context) (RecursiveMap, error)
	// TestObjectQueryParameter implements testObjectQueryParameter operation.
	//
	// GET /testObjectQueryParameter
	TestObjectQueryParameter(ctx context.Context, params TestObjectQueryParameterParams) (TestObjectQueryParameterOK, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config
}

func NewServer(h Handler, opts ...Option) *Server {
	srv := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	return srv
}
