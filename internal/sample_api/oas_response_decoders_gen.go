// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"mime"
	"net/http"
	"path"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/validate"
)

func decodeDataGetFormatResponse(resp *http.Response, span trace.Span) (res string, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response string
			if err := func() error {
				v, err := d.Str()
				response = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeDefaultTestResponse(resp *http.Response, span trace.Span) (res int32, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response int32
			if err := func() error {
				v, err := d.Int32()
				response = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeErrorGetResponse(resp *http.Response, span trace.Span) (res ErrorStatusCode, err error) {
	switch resp.StatusCode {
	default:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
}
func decodeFoobarGetResponse(resp *http.Response, span trace.Span) (res FoobarGetRes, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		return &NotFound{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeFoobarPostResponse(resp *http.Response, span trace.Span) (res FoobarPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		return &NotFound{}, nil
	default:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
}
func decodeFoobarPutResponse(resp *http.Response, span trace.Span) (res FoobarPutDef, err error) {
	switch resp.StatusCode {
	default:
		return FoobarPutDef{
			StatusCode: resp.StatusCode,
		}, nil
	}
}
func decodeGetHeaderResponse(resp *http.Response, span trace.Span) (res Hash, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Hash
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeNoAdditionalPropertiesTestResponse(resp *http.Response, span trace.Span) (res NoAdditionalPropertiesTest, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response NoAdditionalPropertiesTest
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeNullableDefaultResponseResponse(resp *http.Response, span trace.Span) (res NilIntStatusCode, err error) {
	switch resp.StatusCode {
	default:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response NilInt
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return NilIntStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
}
func decodeOneofBugResponse(resp *http.Response, span trace.Span) (res OneofBugOK, err error) {
	switch resp.StatusCode {
	case 200:
		return OneofBugOK{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodePatternRecursiveMapGetResponse(resp *http.Response, span trace.Span) (res PatternRecursiveMap, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response PatternRecursiveMap
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodePetCreateResponse(resp *http.Response, span trace.Span) (res Pet, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodePetFriendsNamesByIDResponse(resp *http.Response, span trace.Span) (res []string, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response []string
			if err := func() error {
				response = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					response = append(response, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodePetGetResponse(resp *http.Response, span trace.Span) (res PetGetRes, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response PetGetDef
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &PetGetDefStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
}
func decodePetGetAvatarByIDResponse(resp *http.Response, span trace.Span) (res PetGetAvatarByIDRes, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/octet-stream":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			response := PetGetAvatarByIDOK{Data: bytes.NewReader(b)}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		return &NotFound{}, nil
	default:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
}
func decodePetGetAvatarByNameResponse(resp *http.Response, span trace.Span) (res PetGetAvatarByNameRes, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/octet-stream":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			response := PetGetAvatarByNameOK{Data: bytes.NewReader(b)}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		return &NotFound{}, nil
	default:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
}
func decodePetGetByNameResponse(resp *http.Response, span trace.Span) (res Pet, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodePetNameByIDResponse(resp *http.Response, span trace.Span) (res string, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response string
			if err := func() error {
				v, err := d.Str()
				response = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodePetUpdateNameAliasPostResponse(resp *http.Response, span trace.Span) (res PetUpdateNameAliasPostDef, err error) {
	switch resp.StatusCode {
	default:
		return PetUpdateNameAliasPostDef{
			StatusCode: resp.StatusCode,
		}, nil
	}
}
func decodePetUpdateNamePostResponse(resp *http.Response, span trace.Span) (res PetUpdateNamePostDef, err error) {
	switch resp.StatusCode {
	default:
		return PetUpdateNamePostDef{
			StatusCode: resp.StatusCode,
		}, nil
	}
}
func decodePetUploadAvatarByIDResponse(resp *http.Response, span trace.Span) (res PetUploadAvatarByIDRes, err error) {
	switch resp.StatusCode {
	case 200:
		return &PetUploadAvatarByIDOK{}, nil
	case 404:
		return &NotFound{}, nil
	default:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
}
func decodeRecursiveArrayGetResponse(resp *http.Response, span trace.Span) (res RecursiveArray, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response RecursiveArray
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeRecursiveMapGetResponse(resp *http.Response, span trace.Span) (res RecursiveMap, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response RecursiveMap
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeSecurityTestResponse(resp *http.Response, span trace.Span) (res string, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response string
			if err := func() error {
				v, err := d.Str()
				response = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeStringIntMapGetResponse(resp *http.Response, span trace.Span) (res StringIntMap, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response StringIntMap
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeTestContentParameterResponse(resp *http.Response, span trace.Span) (res string, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response string
			if err := func() error {
				v, err := d.Str()
				response = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeTestFloatValidationResponse(resp *http.Response, span trace.Span) (res TestFloatValidationOK, err error) {
	switch resp.StatusCode {
	case 200:
		return TestFloatValidationOK{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeTestFormURLEncodedResponse(resp *http.Response, span trace.Span) (res TestFormURLEncodedOK, err error) {
	switch resp.StatusCode {
	case 200:
		return TestFormURLEncodedOK{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeTestMultipartResponse(resp *http.Response, span trace.Span) (res TestMultipartOK, err error) {
	switch resp.StatusCode {
	case 200:
		return TestMultipartOK{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeTestMultipartUploadResponse(resp *http.Response, span trace.Span) (res TestMultipartUploadOK, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response TestMultipartUploadOK
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeTestObjectQueryParameterResponse(resp *http.Response, span trace.Span) (res TestObjectQueryParameterOK, err error) {
	switch resp.StatusCode {
	case 200:
		match := func(pattern, value string) bool {
			ok, _ := path.Match(pattern, value)
			return ok
		}
		_ = match
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.DecodeBytes(buf.Bytes())
			var response TestObjectQueryParameterOK
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
