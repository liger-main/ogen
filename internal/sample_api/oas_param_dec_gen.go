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

func decodeDataGetFormatParams(args [5]string, r *http.Request) (DataGetFormatParams, error) {
	var (
		params DataGetFormatParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	// Decode path: foo.
	{
		param := args[1]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "foo",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Foo = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: foo: not specified")
		}
	}
	// Decode path: bar.
	{
		param := args[2]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "bar",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Bar = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: bar: not specified")
		}
	}
	// Decode path: baz.
	{
		param := args[3]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "baz",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Baz = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: baz: not specified")
		}
	}
	// Decode path: kek.
	{
		param := args[4]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "kek",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Kek = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: kek: not specified")
		}
	}
	return params, nil
}

func decodeDefaultTestParams(args [0]string, r *http.Request) (DefaultTestParams, error) {
	var (
		params    DefaultTestParams
		queryArgs = r.URL.Query()
	)
	// Set default value for query: default.
	{
		val := int32(10)

		params.Default.SetTo(val)
	}
	// Decode query: default.
	{
		values, ok := queryArgs["default"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDefaultVal int32
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(s)
					if err != nil {
						return err
					}

					paramsDefaultVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Default.SetTo(paramsDefaultVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: default: parse")
			}
		}
	}
	return params, nil
}

func decodeFoobarGetParams(args [0]string, r *http.Request) (FoobarGetParams, error) {
	var (
		params    FoobarGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: inlinedParam.
	{
		values, ok := queryArgs["inlinedParam"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(s)
				if err != nil {
					return err
				}

				params.InlinedParam = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: inlinedParam: parse")
			}
		} else {
			return params, errors.New("query: inlinedParam: not specified")
		}
	}
	// Decode query: skip.
	{
		values, ok := queryArgs["skip"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt32(s)
				if err != nil {
					return err
				}

				params.Skip = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: skip: parse")
			}
		} else {
			return params, errors.New("query: skip: not specified")
		}
	}
	return params, nil
}

func decodeGetHeaderParams(args [0]string, r *http.Request) (GetHeaderParams, error) {
	var (
		params GetHeaderParams
	)
	// Decode header: x-auth-token.
	{
		param := r.Header.Get("x-auth-token")
		if len(param) > 0 {
			d := uri.NewHeaderDecoder(uri.HeaderDecoderConfig{
				Value:   param,
				Explode: false,
			})
			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.XAuthToken = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "header: x-auth-token: parse")
			}
		} else {
			return params, errors.New("header: x-auth-token: not specified")
		}
	}
	return params, nil
}

func decodePetFriendsNamesByIDParams(args [1]string, r *http.Request) (PetFriendsNamesByIDParams, error) {
	var (
		params PetFriendsNamesByIDParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodePetGetParams(args [0]string, r *http.Request) (PetGetParams, error) {
	var (
		params    PetGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: petID.
	{
		values, ok := queryArgs["petID"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(s)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: petID: parse")
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:       true,
					Min:          1337,
					MaxSet:       false,
					Max:          0,
					MinExclusive: false,
					MaxExclusive: false,
				}).Validate(int64(params.PetID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: petID: invalid")
			}
		} else {
			return params, errors.New("query: petID: not specified")
		}
	}
	// Decode header: x-tags.
	{
		param := r.Header.Get("x-tags")
		if len(param) > 0 {
			d := uri.NewHeaderDecoder(uri.HeaderDecoderConfig{
				Value:   param,
				Explode: false,
			})
			if err := func() error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsXTagsVal uuid.UUID
					if err := func() error {
						s, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(s)
						if err != nil {
							return err
						}

						paramsXTagsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.XTags = append(params.XTags, paramsXTagsVal)
					return nil
				})
			}(); err != nil {
				return params, errors.Wrap(err, "header: x-tags: parse")
			}
		} else {
			return params, errors.New("header: x-tags: not specified")
		}
	}
	// Decode header: x-scope.
	{
		param := r.Header.Get("x-scope")
		if len(param) > 0 {
			d := uri.NewHeaderDecoder(uri.HeaderDecoderConfig{
				Value:   param,
				Explode: false,
			})
			if err := func() error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsXScopeVal string
					if err := func() error {
						s, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(s)
						if err != nil {
							return err
						}

						paramsXScopeVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.XScope = append(params.XScope, paramsXScopeVal)
					return nil
				})
			}(); err != nil {
				return params, errors.Wrap(err, "header: x-scope: parse")
			}
		} else {
			return params, errors.New("header: x-scope: not specified")
		}
	}
	// Decode query: token.
	{
		values, ok := queryArgs["token"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Token = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: token: parse")
			}
		} else {
			return params, errors.New("query: token: not specified")
		}
	}
	return params, nil
}

func decodePetGetAvatarByIDParams(args [0]string, r *http.Request) (PetGetAvatarByIDParams, error) {
	var (
		params    PetGetAvatarByIDParams
		queryArgs = r.URL.Query()
	)
	// Decode query: petID.
	{
		values, ok := queryArgs["petID"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(s)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: petID: parse")
			}
		} else {
			return params, errors.New("query: petID: not specified")
		}
	}
	return params, nil
}

func decodePetGetAvatarByNameParams(args [1]string, r *http.Request) (PetGetAvatarByNameParams, error) {
	var (
		params PetGetAvatarByNameParams
	)
	// Decode path: name.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "name",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Name = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: name: not specified")
		}
	}
	return params, nil
}

func decodePetGetByNameParams(args [1]string, r *http.Request) (PetGetByNameParams, error) {
	var (
		params PetGetByNameParams
	)
	// Decode path: name.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "name",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Name = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: name: not specified")
		}
	}
	return params, nil
}

func decodePetNameByIDParams(args [1]string, r *http.Request) (PetNameByIDParams, error) {
	var (
		params PetNameByIDParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodePetUploadAvatarByIDParams(args [0]string, r *http.Request) (PetUploadAvatarByIDParams, error) {
	var (
		params    PetUploadAvatarByIDParams
		queryArgs = r.URL.Query()
	)
	// Decode query: petID.
	{
		values, ok := queryArgs["petID"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(s)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: petID: parse")
			}
		} else {
			return params, errors.New("query: petID: not specified")
		}
	}
	return params, nil
}

func decodeTestObjectQueryParameterParams(args [0]string, r *http.Request) (TestObjectQueryParameterParams, error) {
	var (
		params    TestObjectQueryParameterParams
		queryArgs = r.URL.Query()
	)
	// Decode query: formObject.
	{
		values, ok := queryArgs["formObject"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsFormObjectVal TestObjectQueryParameterFormObject
				if err := func() error {
					return paramsFormObjectVal.decodeURI(d)
				}(); err != nil {
					return err
				}
				params.FormObject.SetTo(paramsFormObjectVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: formObject: parse")
			}
		}
	}
	// Decode query: deepObject.
	{
		values, ok := queryArgs["deepObject"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleDeepObject,
				Explode: true,
			})

			if err := func() error {
				var paramsDeepObjectVal TestObjectQueryParameterDeepObject
				if err := func() error {
					return paramsDeepObjectVal.decodeURI(d)
				}(); err != nil {
					return err
				}
				params.DeepObject.SetTo(paramsDeepObjectVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: deepObject: parse")
			}
		}
	}
	return params, nil
}
