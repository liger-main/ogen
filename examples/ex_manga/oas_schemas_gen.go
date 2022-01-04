// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
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

// Ref: #/components/schemas/Book
type Book struct {
	ID           OptInt    `json:"id"`
	MediaID      OptInt    `json:"media_id"`
	Images       OptImages `json:"images"`
	Title        OptTitle  `json:"title"`
	Tags         []Tag     `json:"tags"`
	Scanlator    OptString `json:"scanlator"`
	UploadDate   OptInt    `json:"upload_date"`
	NumPages     OptInt    `json:"num_pages"`
	NumFavorites OptInt    `json:"num_favorites"`
}

func (*Book) getBookRes() {}

// GetBookForbidden is response for GetBook operation.
type GetBookForbidden struct{}

func (*GetBookForbidden) getBookRes() {}

// GetPageCoverImageForbidden is response for GetPageCoverImage operation.
type GetPageCoverImageForbidden struct{}

func (*GetPageCoverImageForbidden) getPageCoverImageRes() {}

type GetPageCoverImageOKImage struct {
	Data io.Reader
}

func (s GetPageCoverImageOKImage) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

func (*GetPageCoverImageOKImage) getPageCoverImageRes() {}

// GetPageImageForbidden is response for GetPageImage operation.
type GetPageImageForbidden struct{}

func (*GetPageImageForbidden) getPageImageRes() {}

type GetPageImageOKImage struct {
	Data io.Reader
}

func (s GetPageImageOKImage) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

func (*GetPageImageOKImage) getPageImageRes() {}

// GetPageThumbnailImageForbidden is response for GetPageThumbnailImage operation.
type GetPageThumbnailImageForbidden struct{}

func (*GetPageThumbnailImageForbidden) getPageThumbnailImageRes() {}

type GetPageThumbnailImageOKImage struct {
	Data io.Reader
}

func (s GetPageThumbnailImageOKImage) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

func (*GetPageThumbnailImageOKImage) getPageThumbnailImageRes() {}

// Ref: #/components/schemas/Image
type Image struct {
	T OptString `json:"t"`
	W OptInt    `json:"w"`
	H OptInt    `json:"h"`
}

// Ref: #/components/schemas/Images
type Images struct {
	Pages     []Image  `json:"pages"`
	Cover     OptImage `json:"cover"`
	Thumbnail OptImage `json:"thumbnail"`
}

// NewOptImage returns new OptImage with value set to v.
func NewOptImage(v Image) OptImage {
	return OptImage{
		Value: v,
		Set:   true,
	}
}

// OptImage is optional Image.
type OptImage struct {
	Value Image
	Set   bool
}

// IsSet returns true if OptImage was set.
func (o OptImage) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptImage) Reset() {
	var v Image
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptImage) SetTo(v Image) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptImage) Get() (v Image, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptImage) Or(d Image) Image {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptImages returns new OptImages with value set to v.
func NewOptImages(v Images) OptImages {
	return OptImages{
		Value: v,
		Set:   true,
	}
}

// OptImages is optional Images.
type OptImages struct {
	Value Images
	Set   bool
}

// IsSet returns true if OptImages was set.
func (o OptImages) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptImages) Reset() {
	var v Images
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptImages) SetTo(v Images) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptImages) Get() (v Images, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptImages) Or(d Images) Images {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTagType returns new OptTagType with value set to v.
func NewOptTagType(v TagType) OptTagType {
	return OptTagType{
		Value: v,
		Set:   true,
	}
}

// OptTagType is optional TagType.
type OptTagType struct {
	Value TagType
	Set   bool
}

// IsSet returns true if OptTagType was set.
func (o OptTagType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTagType) Reset() {
	var v TagType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTagType) SetTo(v TagType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTagType) Get() (v TagType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTagType) Or(d TagType) TagType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTitle returns new OptTitle with value set to v.
func NewOptTitle(v Title) OptTitle {
	return OptTitle{
		Value: v,
		Set:   true,
	}
}

// OptTitle is optional Title.
type OptTitle struct {
	Value Title
	Set   bool
}

// IsSet returns true if OptTitle was set.
func (o OptTitle) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTitle) Reset() {
	var v Title
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTitle) SetTo(v Title) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTitle) Get() (v Title, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTitle) Or(d Title) Title {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// SearchByTagIDForbidden is response for SearchByTagID operation.
type SearchByTagIDForbidden struct{}

func (*SearchByTagIDForbidden) searchByTagIDRes() {}

type SearchByTagIDOKApplicationJSON []SearchResponse

func (SearchByTagIDOKApplicationJSON) searchByTagIDRes() {}

// SearchForbidden is response for Search operation.
type SearchForbidden struct{}

func (*SearchForbidden) searchRes() {}

type SearchOKApplicationJSON []SearchResponse

func (SearchOKApplicationJSON) searchRes() {}

// Ref: #/components/schemas/SearchResponse
type SearchResponse struct {
	Result   []Book `json:"result"`
	NumPages OptInt `json:"num_pages"`
	PerPage  OptInt `json:"per_page"`
}

// Ref: #/components/schemas/Tag
type Tag struct {
	ID    OptInt     `json:"id"`
	Type  OptTagType `json:"type"`
	Name  OptString  `json:"name"`
	URL   OptString  `json:"url"`
	Count OptInt     `json:"count"`
}

type TagType string

const (
	TagTypeParody    TagType = "parody"
	TagTypeCharacter TagType = "character"
	TagTypeTag       TagType = "tag"
	TagTypeArtist    TagType = "artist"
	TagTypeGroup     TagType = "group"
	TagTypeCategory  TagType = "category"
	TagTypeLanguage  TagType = "language"
)

// Ref: #/components/schemas/Title
type Title struct {
	English  OptString `json:"english"`
	Japanese OptString `json:"japanese"`
	Pretty   OptString `json:"pretty"`
}
