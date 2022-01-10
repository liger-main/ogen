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

// Encode implements json.Marshaler.
func (s Book) Encode(e *jx.Encoder) {
	e.ObjStart()
	if s.ID.Set {
		e.FieldStart("id")
		s.ID.Encode(e)
	}
	if s.MediaID.Set {
		e.FieldStart("media_id")
		s.MediaID.Encode(e)
	}
	if s.Images.Set {
		e.FieldStart("images")
		s.Images.Encode(e)
	}
	if s.Title.Set {
		e.FieldStart("title")
		s.Title.Encode(e)
	}
	if s.Tags != nil {
		e.FieldStart("tags")
		e.ArrStart()
		for _, elem := range s.Tags {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	if s.Scanlator.Set {
		e.FieldStart("scanlator")
		s.Scanlator.Encode(e)
	}
	if s.UploadDate.Set {
		e.FieldStart("upload_date")
		s.UploadDate.Encode(e)
	}
	if s.NumPages.Set {
		e.FieldStart("num_pages")
		s.NumPages.Encode(e)
	}
	if s.NumFavorites.Set {
		e.FieldStart("num_favorites")
		s.NumFavorites.Encode(e)
	}
	e.ObjEnd()
}

// Decode decodes Book from json.
func (s *Book) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode Book to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			s.ID.Reset()
			if err := s.ID.Decode(d); err != nil {
				return err
			}
		case "media_id":
			s.MediaID.Reset()
			if err := s.MediaID.Decode(d); err != nil {
				return err
			}
		case "images":
			s.Images.Reset()
			if err := s.Images.Decode(d); err != nil {
				return err
			}
		case "title":
			s.Title.Reset()
			if err := s.Title.Decode(d); err != nil {
				return err
			}
		case "tags":
			s.Tags = nil
			if err := d.Arr(func(d *jx.Decoder) error {
				var elem Tag
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Tags = append(s.Tags, elem)
				return nil
			}); err != nil {
				return err
			}
		case "scanlator":
			s.Scanlator.Reset()
			if err := s.Scanlator.Decode(d); err != nil {
				return err
			}
		case "upload_date":
			s.UploadDate.Reset()
			if err := s.UploadDate.Decode(d); err != nil {
				return err
			}
		case "num_pages":
			s.NumPages.Reset()
			if err := s.NumPages.Decode(d); err != nil {
				return err
			}
		case "num_favorites":
			s.NumFavorites.Reset()
			if err := s.NumFavorites.Decode(d); err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s GetBookForbidden) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes GetBookForbidden from json.
func (s *GetBookForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode GetBookForbidden to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s GetPageCoverImageForbidden) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes GetPageCoverImageForbidden from json.
func (s *GetPageCoverImageForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode GetPageCoverImageForbidden to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s GetPageImageForbidden) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes GetPageImageForbidden from json.
func (s *GetPageImageForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode GetPageImageForbidden to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s GetPageThumbnailImageForbidden) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes GetPageThumbnailImageForbidden from json.
func (s *GetPageThumbnailImageForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode GetPageThumbnailImageForbidden to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s Image) Encode(e *jx.Encoder) {
	e.ObjStart()
	if s.T.Set {
		e.FieldStart("t")
		s.T.Encode(e)
	}
	if s.W.Set {
		e.FieldStart("w")
		s.W.Encode(e)
	}
	if s.H.Set {
		e.FieldStart("h")
		s.H.Encode(e)
	}
	e.ObjEnd()
}

// Decode decodes Image from json.
func (s *Image) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode Image to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "t":
			s.T.Reset()
			if err := s.T.Decode(d); err != nil {
				return err
			}
		case "w":
			s.W.Reset()
			if err := s.W.Decode(d); err != nil {
				return err
			}
		case "h":
			s.H.Reset()
			if err := s.H.Decode(d); err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s Images) Encode(e *jx.Encoder) {
	e.ObjStart()
	if s.Pages != nil {
		e.FieldStart("pages")
		e.ArrStart()
		for _, elem := range s.Pages {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	if s.Cover.Set {
		e.FieldStart("cover")
		s.Cover.Encode(e)
	}
	if s.Thumbnail.Set {
		e.FieldStart("thumbnail")
		s.Thumbnail.Encode(e)
	}
	e.ObjEnd()
}

// Decode decodes Images from json.
func (s *Images) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode Images to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "pages":
			s.Pages = nil
			if err := d.Arr(func(d *jx.Decoder) error {
				var elem Image
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Pages = append(s.Pages, elem)
				return nil
			}); err != nil {
				return err
			}
		case "cover":
			s.Cover.Reset()
			if err := s.Cover.Decode(d); err != nil {
				return err
			}
		case "thumbnail":
			s.Thumbnail.Reset()
			if err := s.Thumbnail.Decode(d); err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes Image as json.
func (o OptImage) Encode(e *jx.Encoder) {
	o.Value.Encode(e)
}

// Decode decodes Image from json.
func (o *OptImage) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptImage to nil`)
	}
	switch d.Next() {
	case jx.Object:
		o.Set = true
		if err := o.Value.Decode(d); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptImage`, d.Next())
	}
}

// Encode encodes Images as json.
func (o OptImages) Encode(e *jx.Encoder) {
	o.Value.Encode(e)
}

// Decode decodes Images from json.
func (o *OptImages) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptImages to nil`)
	}
	switch d.Next() {
	case jx.Object:
		o.Set = true
		if err := o.Value.Decode(d); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptImages`, d.Next())
	}
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptInt to nil`)
	}
	switch d.Next() {
	case jx.Number:
		o.Set = true
		v, err := d.Int()
		if err != nil {
			return err
		}
		o.Value = int(v)
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptInt`, d.Next())
	}
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptString to nil`)
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		v, err := d.Str()
		if err != nil {
			return err
		}
		o.Value = string(v)
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptString`, d.Next())
	}
}

// Encode encodes TagType as json.
func (o OptTagType) Encode(e *jx.Encoder) {
	e.Str(string(o.Value))
}

// Decode decodes TagType from json.
func (o *OptTagType) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptTagType to nil`)
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		if err := o.Value.Decode(d); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptTagType`, d.Next())
	}
}

// Encode encodes Title as json.
func (o OptTitle) Encode(e *jx.Encoder) {
	o.Value.Encode(e)
}

// Decode decodes Title from json.
func (o *OptTitle) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptTitle to nil`)
	}
	switch d.Next() {
	case jx.Object:
		o.Set = true
		if err := o.Value.Decode(d); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptTitle`, d.Next())
	}
}

// Encode implements json.Marshaler.
func (s SearchByTagIDForbidden) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes SearchByTagIDForbidden from json.
func (s *SearchByTagIDForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode SearchByTagIDForbidden to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes SearchByTagIDOKApplicationJSON as json.
func (s SearchByTagIDOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []SearchResponse(s)
	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes SearchByTagIDOKApplicationJSON from json.
func (s *SearchByTagIDOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode SearchByTagIDOKApplicationJSON to nil`)
	}
	var unwrapped []SearchResponse
	if err := func() error {
		unwrapped = nil
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem SearchResponse
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = SearchByTagIDOKApplicationJSON(unwrapped)
	return nil
}

// Encode implements json.Marshaler.
func (s SearchForbidden) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes SearchForbidden from json.
func (s *SearchForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode SearchForbidden to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes SearchOKApplicationJSON as json.
func (s SearchOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []SearchResponse(s)
	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes SearchOKApplicationJSON from json.
func (s *SearchOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode SearchOKApplicationJSON to nil`)
	}
	var unwrapped []SearchResponse
	if err := func() error {
		unwrapped = nil
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem SearchResponse
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = SearchOKApplicationJSON(unwrapped)
	return nil
}

// Encode implements json.Marshaler.
func (s SearchResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	if s.Result != nil {
		e.FieldStart("result")
		e.ArrStart()
		for _, elem := range s.Result {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	if s.NumPages.Set {
		e.FieldStart("num_pages")
		s.NumPages.Encode(e)
	}
	if s.PerPage.Set {
		e.FieldStart("per_page")
		s.PerPage.Encode(e)
	}
	e.ObjEnd()
}

// Decode decodes SearchResponse from json.
func (s *SearchResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode SearchResponse to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			s.Result = nil
			if err := d.Arr(func(d *jx.Decoder) error {
				var elem Book
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Result = append(s.Result, elem)
				return nil
			}); err != nil {
				return err
			}
		case "num_pages":
			s.NumPages.Reset()
			if err := s.NumPages.Decode(d); err != nil {
				return err
			}
		case "per_page":
			s.PerPage.Reset()
			if err := s.PerPage.Decode(d); err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s Tag) Encode(e *jx.Encoder) {
	e.ObjStart()
	if s.ID.Set {
		e.FieldStart("id")
		s.ID.Encode(e)
	}
	if s.Type.Set {
		e.FieldStart("type")
		s.Type.Encode(e)
	}
	if s.Name.Set {
		e.FieldStart("name")
		s.Name.Encode(e)
	}
	if s.URL.Set {
		e.FieldStart("url")
		s.URL.Encode(e)
	}
	if s.Count.Set {
		e.FieldStart("count")
		s.Count.Encode(e)
	}
	e.ObjEnd()
}

// Decode decodes Tag from json.
func (s *Tag) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode Tag to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			s.ID.Reset()
			if err := s.ID.Decode(d); err != nil {
				return err
			}
		case "type":
			s.Type.Reset()
			if err := s.Type.Decode(d); err != nil {
				return err
			}
		case "name":
			s.Name.Reset()
			if err := s.Name.Decode(d); err != nil {
				return err
			}
		case "url":
			s.URL.Reset()
			if err := s.URL.Decode(d); err != nil {
				return err
			}
		case "count":
			s.Count.Reset()
			if err := s.Count.Decode(d); err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes TagType as json.
func (s TagType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes TagType from json.
func (s *TagType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode TagType to nil`)
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch TagType(v) {
	case TagTypeParody:
		*s = TagTypeParody
	case TagTypeCharacter:
		*s = TagTypeCharacter
	case TagTypeTag:
		*s = TagTypeTag
	case TagTypeArtist:
		*s = TagTypeArtist
	case TagTypeGroup:
		*s = TagTypeGroup
	case TagTypeCategory:
		*s = TagTypeCategory
	case TagTypeLanguage:
		*s = TagTypeLanguage
	default:
		*s = TagType(v)
	}

	return nil
}

// Encode implements json.Marshaler.
func (s Title) Encode(e *jx.Encoder) {
	e.ObjStart()
	if s.English.Set {
		e.FieldStart("english")
		s.English.Encode(e)
	}
	if s.Japanese.Set {
		e.FieldStart("japanese")
		s.Japanese.Encode(e)
	}
	if s.Pretty.Set {
		e.FieldStart("pretty")
		s.Pretty.Encode(e)
	}
	e.ObjEnd()
}

// Decode decodes Title from json.
func (s *Title) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode Title to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "english":
			s.English.Reset()
			if err := s.English.Decode(d); err != nil {
				return err
			}
		case "japanese":
			s.Japanese.Reset()
			if err := s.Japanese.Decode(d); err != nil {
				return err
			}
		case "pretty":
			s.Pretty.Reset()
			if err := s.Pretty.Decode(d); err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}
