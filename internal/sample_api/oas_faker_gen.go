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

// SetFake set fake values.
func (s *AnyOfTest) SetFake() {
	{

		{
			s.Medium = "string"
		}
	}
	{

		{
			s.SizeLimit.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *AnyOfTestSizeLimit) SetFake() {
	var elem int

	{
		elem = int(0)
	}
	s.SetInt(elem)
}

// SetFake set fake values.
func (s *AnyTest) SetFake() {
	{

		{
			s.Empty = []byte("null")
		}
	}
	{

		{
			s.AnyMap.SetFake()
		}
	}
	{

		{
			s.AnyArray = nil
			for i := 0; i < 0; i++ {
				var elem jx.Raw

				{
					elem = []byte("null")
				}
				s.AnyArray = append(s.AnyArray, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *AnyTestAnyMap) SetFake() {
	var (
		elem jx.Raw
		m    map[string]jx.Raw = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *ArrayTest) SetFake() {
	{

		{
			s.Required = nil
			for i := 0; i < 0; i++ {
				var elem string

				{
					elem = "string"
				}
				s.Required = append(s.Required, elem)
			}
		}
	}
	{

		{
			s.Optional = nil
			for i := 0; i < 0; i++ {
				var elem string

				{
					elem = "string"
				}
				s.Optional = append(s.Optional, elem)
			}
		}
	}
	{

		{
			s.NullableRequired = nil
			for i := 0; i < 0; i++ {
				var elem string

				{
					elem = "string"
				}
				s.NullableRequired = append(s.NullableRequired, elem)
			}
		}
	}
	{

		{
			s.NullableOptional.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Data) SetFake() {
	{

		{
			s.ID.SetFake()
		}
	}
	{

		{
			s.Description.SetFake()
		}
	}
	{

		{
			s.Email = "string"
		}
	}
	{

		{
			s.Hostname = "string"
		}
	}
	{

		{
			s.Format = "string"
		}
	}
	{

		{
			s.Base64 = []byte("[]byte")
		}
	}
	{

		{
			s.NullableEnum.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *DataDescription) SetFake() {
	var elem DescriptionDetailed

	{
		elem.SetFake()
	}
	s.SetDescriptionDetailed(elem)
}

// SetFake set fake values.
func (s *DefaultTest) SetFake() {
	{

		{
			s.Required = "string"
		}
	}
	{

		{
			s.Str.SetFake()
		}
	}
	{

		{
			s.NullStr.SetFake()
		}
	}
	{

		{
			s.Enum.SetFake()
		}
	}
	{

		{
			s.UUID.SetFake()
		}
	}
	{

		{
			s.IP.SetFake()
		}
	}
	{

		{
			s.IPV4.SetFake()
		}
	}
	{

		{
			s.IPV6.SetFake()
		}
	}
	{

		{
			s.URI.SetFake()
		}
	}
	{

		{
			s.Birthday.SetFake()
		}
	}
	{

		{
			s.Rate.SetFake()
		}
	}
	{

		{
			s.Email.SetFake()
		}
	}
	{

		{
			s.Hostname.SetFake()
		}
	}
	{

		{
			s.Format.SetFake()
		}
	}
	{

		{
			s.Base64 = []byte("[]byte")
		}
	}
}

// SetFake set fake values.
func (s *DefaultTestEnum) SetFake() {
	*s = DefaultTestEnumBig
}

// SetFake set fake values.
func (s *DescriptionDetailed) SetFake() {
	{

		{
			s.Name = "string"
		}
	}
	{

		{
			s.Count = int(0)
		}
	}
	{

		{
			s.ID.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *DescriptionSimple) SetFake() {
	{

		{
			s.Description = "string"
		}
	}
}

// SetFake set fake values.
func (s *Error) SetFake() {
	{

		{
			s.Code = int64(0)
		}
	}
	{

		{
			s.Message = "string"
		}
	}
}

// SetFake set fake values.
func (s *Hash) SetFake() {
	{

		{
			s.Raw = []byte("[]byte")
		}
	}
	{

		{
			s.Hex = "string"
		}
	}
}

// SetFake set fake values.
func (s *ID) SetFake() {
	var elem string

	{
		elem = "string"
	}
	s.SetString(elem)
}

// SetFake set fake values.
func (s *Issue143) SetFake() {
	var elem Issue1430

	{
		elem.SetFake()
	}
	s.SetIssue1430(elem)
}

// SetFake set fake values.
func (s *Issue1430) SetFake() {
	{

		{
			s.CommonMinus1 = "string"
		}
	}
	{

		{
			s.CommonMinus2 = int(0)
		}
	}
	{

		{
			s.UniqueMinus1 = "string"
		}
	}
}

// SetFake set fake values.
func (s *Issue1431) SetFake() {
	{

		{
			s.CommonMinus1 = "string"
		}
	}
	{

		{
			s.CommonMinus2 = int(0)
		}
	}
	{

		{
			s.UniqueMinus2 = "string"
		}
	}
}

// SetFake set fake values.
func (s *Issue1432) SetFake() {
	{

		{
			s.CommonMinus1 = "string"
		}
	}
	{

		{
			s.CommonMinus2 = int(0)
		}
	}
	{

		{
			s.CommonMinus3.SetFake()
		}
	}
	{

		{
			s.UniqueMinus3 = "string"
		}
	}
}

// SetFake set fake values.
func (s *Issue1433) SetFake() {
	{

		{
			s.CommonMinus3.SetFake()
		}
	}
	{

		{
			s.UniqueMinus4 = "string"
		}
	}
}

// SetFake set fake values.
func (s *MapWithProperties) SetFake() {
	{

		{
			s.Required = int(0)
		}
	}
	{

		{
			s.AdditionalProps.SetFake()
		}
	}
	{

		{
			s.Optional.SetFake()
		}
	}
	{

		{
			s.SubMap.SetFake()
		}
	}
	{

		{
			s.InlinedSubMap.SetFake()
		}
	}
	{

		{
			s.MapValidation.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *MapWithPropertiesAdditional) SetFake() {
	var (
		elem string
		m    map[string]string = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *MapWithPropertiesInlinedSubMap) SetFake() {
	var (
		elem string
		m    map[string]string = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *MaxPropertiesTest) SetFake() {
	{

		{
			s.Required = int(0)
		}
	}
	{

		{
			s.OptionalA.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *NilInt) SetFake() {
	s.Null = true
}

// SetFake set fake values.
func (s *NilNullableEnumsBoth) SetFake() {
	s.Null = true
}

// SetFake set fake values.
func (s *NilNullableEnumsOnlyNullValue) SetFake() {
	s.Null = true
}

// SetFake set fake values.
func (s *NilNullableEnumsOnlyNullable) SetFake() {
	s.Null = true
}

// SetFake set fake values.
func (s *NilString) SetFake() {
	s.Null = true
}

// SetFake set fake values.
func (s *NullableEnums) SetFake() {
	{

		{
			s.OnlyNullable.SetFake()
		}
	}
	{

		{
			s.OnlyNullValue.SetFake()
		}
	}
	{

		{
			s.Both.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *NullableEnumsBoth) SetFake() {
	*s = NullableEnumsBothAsc
}

// SetFake set fake values.
func (s *NullableEnumsOnlyNullValue) SetFake() {
	*s = NullableEnumsOnlyNullValueAsc
}

// SetFake set fake values.
func (s *NullableEnumsOnlyNullable) SetFake() {
	*s = NullableEnumsOnlyNullableAsc
}

// SetFake set fake values.
func (s *OneOfBugs) SetFake() {
	{

		{
			s.Issue143.SetFake()
		}
	}
	{

		{
			s.AdditionalMinusFields.SetFake()
		}
	}
	{

		{
			s.OneOfMinusUUIDMinusIntMinusEnum.SetFake()
		}
	}
	{

		{
			s.OneOfMinusMappingMinusReference.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OneOfMappingReference) SetFake() {
	var elem OneOfMappingReferenceB

	{
		elem.SetFake()
	}
	s.SetOneOfMappingReferenceB(elem)
}

// SetFake set fake values.
func (s *OneOfMappingReferenceA) SetFake() {
	{

		{
			s.Description.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OneOfMappingReferenceB) SetFake() {
	{

		{
			s.Code.SetFake()
		}
	}
	{

		{ // Keep pointer nil to prevent infinite recursion.
			s.Data = nil
		}
	}
	{

		{
			s.Info.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OneOfMappingReferenceBData) SetFake() {
}

// SetFake set fake values.
func (s *OneOfUUIDAndIntEnum) SetFake() {
	var elem uuid.UUID

	{
		elem = uuid.New()
	}
	s.SetUUID(elem)
}

// SetFake set fake values.
func (s *OneOfUUIDAndIntEnum1) SetFake() {
	*s = OneOfUUIDAndIntEnum10
}

// SetFake set fake values.
func (s *OneVariantHasNoUniqueFields) SetFake() {
	var elem OneVariantHasNoUniqueFields0

	{
		elem.SetFake()
	}
	s.SetOneVariantHasNoUniqueFields0(elem)
}

// SetFake set fake values.
func (s *OneVariantHasNoUniqueFields0) SetFake() {
	{

		{
			s.A = "string"
		}
	}
	{

		{
			s.B.SetFake()
		}
	}
	{

		{
			s.C = "string"
		}
	}
}

// SetFake set fake values.
func (s *OneVariantHasNoUniqueFields1) SetFake() {
	{

		{
			s.A = "string"
		}
	}
	{

		{
			s.B.SetFake()
		}
	}
	{

		{
			s.C = "string"
		}
	}
	{

		{
			s.D.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OptAnyOfTest) SetFake() {
	var elem AnyOfTest

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptAnyTest) SetFake() {
	var elem AnyTest

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptAnyTestAnyMap) SetFake() {
	var elem AnyTestAnyMap

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptArrayTest) SetFake() {
	var elem ArrayTest

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptData) SetFake() {
	var elem Data

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptDate) SetFake() {
	var elem time.Time

	{
		elem = time.Now()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptDateTime) SetFake() {
	var elem time.Time

	{
		elem = time.Now()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptDefaultTestEnum) SetFake() {
	var elem DefaultTestEnum

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptDuration) SetFake() {
	var elem time.Duration

	{
		elem = time.Duration(5 * time.Second)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptFloat64) SetFake() {
	var elem float64

	{
		elem = float64(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptID) SetFake() {
	var elem ID

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptIP) SetFake() {
	var elem net.IP

	{
		elem = net.ParseIP("127.0.0.1")
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptIPv4) SetFake() {
	var elem net.IP

	{
		elem = net.ParseIP("127.0.0.1")
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptIPv6) SetFake() {
	var elem net.IP

	{
		elem = net.ParseIP("127.0.0.1")
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptInt) SetFake() {
	var elem int

	{
		elem = int(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptInt32) SetFake() {
	var elem int32

	{
		elem = int32(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptMapWithProperties) SetFake() {
	var elem MapWithProperties

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptMapWithPropertiesInlinedSubMap) SetFake() {
	var elem MapWithPropertiesInlinedSubMap

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptMaxPropertiesTest) SetFake() {
	var elem MaxPropertiesTest

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptNilString) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilStringArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNullableEnums) SetFake() {
	var elem NullableEnums

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptOneOfMappingReference) SetFake() {
	var elem OneOfMappingReference

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptOneOfUUIDAndIntEnum) SetFake() {
	var elem OneOfUUIDAndIntEnum

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptPet) SetFake() {
	var elem Pet

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptPetName) SetFake() {
	var elem PetName

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptPetType) SetFake() {
	var elem PetType

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptString) SetFake() {
	var elem string

	{
		elem = "string"
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptStringMap) SetFake() {
	var elem StringMap

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptStringStringMap) SetFake() {
	var elem StringStringMap

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptTime) SetFake() {
	var elem time.Time

	{
		elem = time.Now()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptURI) SetFake() {
	var elem url.URL

	{
		elem = url.URL{Scheme: "https", Host: "github.com", Path: "/ogen-go/ogen"}
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptUUID) SetFake() {
	var elem uuid.UUID

	{
		elem = uuid.New()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptValidationStringMap) SetFake() {
	var elem ValidationStringMap

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *Pet) SetFake() {
	{

		{ // Keep pointer nil to prevent infinite recursion.
			s.Primary = nil
		}
	}
	{

		{
			s.ID = int64(0)
		}
	}
	{

		{
			s.UniqueID = uuid.New()
		}
	}
	{

		{
			s.Name = "string"
		}
	}
	{

		{
			s.Type.SetFake()
		}
	}
	{

		{
			s.Kind.SetFake()
		}
	}
	{

		{
			s.Tag.SetFake()
		}
	}
	{

		{
			s.IP = net.ParseIP("127.0.0.1")
		}
	}
	{

		{
			s.IPV4 = net.ParseIP("127.0.0.1")
		}
	}
	{

		{
			s.IPV6 = net.ParseIP("127.0.0.1")
		}
	}
	{

		{
			s.URI = url.URL{Scheme: "https", Host: "github.com", Path: "/ogen-go/ogen"}
		}
	}
	{

		{
			s.Birthday = time.Now()
		}
	}
	{

		{
			s.Rate = time.Duration(5 * time.Second)
		}
	}
	{

		{
			s.Nickname.SetFake()
		}
	}
	{

		{
			s.NullStr.SetFake()
		}
	}
	{

		{
			s.Friends = nil
			for i := 0; i < 0; i++ {
				var elem Pet

				{
					elem.SetFake()
				}
				s.Friends = append(s.Friends, elem)
			}
		}
	}
	{

		{
			s.Next.SetFake()
		}
	}
	{

		{
			s.TestInteger1.SetFake()
		}
	}
	{

		{
			s.TestFloat1.SetFake()
		}
	}
	{

		{
			s.TestArray1 = nil
			for i := 0; i < 0; i++ {
				var elem []string

				{
					elem = nil
					for i := 0; i < 0; i++ {
						var elemElem string

						{
							elemElem = "string"
						}
						elem = append(elem, elemElem)
					}
				}
				s.TestArray1 = append(s.TestArray1, elem)
			}
		}
	}
	{

		{
			s.TestArray2.SetFake()
		}
	}
	{

		{
			s.TestMap.SetFake()
		}
	}
	{

		{
			s.TestMapWithProps.SetFake()
		}
	}
	{

		{
			s.TestAny.SetFake()
		}
	}
	{

		{
			s.TestAnyOf.SetFake()
		}
	}
	{

		{
			s.TestMaxProperties.SetFake()
		}
	}
	{

		{
			s.TestDate.SetFake()
		}
	}
	{

		{
			s.TestDuration.SetFake()
		}
	}
	{

		{
			s.TestTime.SetFake()
		}
	}
	{

		{
			s.TestDateTime.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *PetGetDef) SetFake() {
	{

		{
			s.Message = "string"
		}
	}
}

// SetFake set fake values.
func (s *PetKind) SetFake() {
	*s = PetKindBig
}

// SetFake set fake values.
func (s *PetName) SetFake() {
	var unwrapped string

	{
		unwrapped = "string"
	}
	*s = PetName(unwrapped)
}

// SetFake set fake values.
func (s *PetType) SetFake() {
	*s = PetTypeFifa
}

// SetFake set fake values.
func (s *RecursiveArray) SetFake() {
	var unwrapped []RecursiveArray

	{
		unwrapped = nil
		for i := 0; i < 0; i++ {
			var elem RecursiveArray

			{
				elem.SetFake()
			}
			unwrapped = append(unwrapped, elem)
		}
	}
	*s = RecursiveArray(unwrapped)
}

// SetFake set fake values.
func (s *RecursiveMap) SetFake() {
	{

		{ // Keep pointer nil to prevent infinite recursion.
			s.OptionalRecursiveField = nil
		}
	}
	{

		{
			s.AdditionalProps.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *RecursiveMapAdditional) SetFake() {
	var (
		elem RecursiveMap
		m    map[string]RecursiveMap = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *StringMap) SetFake() {
	var (
		elem string
		m    map[string]string = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *StringStringMap) SetFake() {
	var (
		elem StringMap
		m    map[string]StringMap = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *TestFloatValidation) SetFake() {
	{

		{
			s.Minmax = float64(0)
		}
	}
	{

		{
			s.MultipleOf = float64(0)
		}
	}
}

// SetFake set fake values.
func (s *TestObjectQueryParameterOK) SetFake() {
	{

		{
			s.Style = "string"
		}
	}
	{

		{
			s.Min = int(0)
		}
	}
	{

		{
			s.Max = int(0)
		}
	}
	{

		{
			s.Filter = "string"
		}
	}
}

// SetFake set fake values.
func (s *ValidationStringMap) SetFake() {
	var (
		elem string
		m    map[string]string = s.init()
	)
	for i := 0; i < 1; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}
