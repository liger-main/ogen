// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// HealthzGetParams is parameters of GET /healthz operation.
type HealthzGetParams struct {
	Eq       string
	Plus     string
	Question string
	And      string
	Percent  string
}

func unpackHealthzGetParams(packed middleware.Parameters) (params HealthzGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "=",
			In:   "query",
		}
		params.Eq = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "+",
			In:   "query",
		}
		params.Plus = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "question?",
			In:   "query",
		}
		params.Question = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "and&",
			In:   "query",
		}
		params.And = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "percent%",
			In:   "query",
		}
		params.Percent = packed[key].(string)
	}
	return params
}

func decodeHealthzGetParams(args [0]string, r *http.Request) (params HealthzGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: =.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "=",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Eq = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "=",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: +.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "+",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Plus = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "+",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: question?.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "question?",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Question = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "question?",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: and&.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "and&",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.And = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "and&",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: percent%.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "percent%",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Percent = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "percent%",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// SameNameParams is parameters of sameName operation.
type SameNameParams struct {
	pathPath  string
	queryPath string
}

func unpackSameNameParams(packed middleware.Parameters) (params SameNameParams) {
	{
		key := middleware.ParameterKey{
			Name: "path",
			In:   "path",
		}
		params.pathPath = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "path",
			In:   "query",
		}
		params.queryPath = packed[key].(string)
	}
	return params
}

func decodeSameNameParams(args [1]string, r *http.Request) (params SameNameParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: path.
	if err := func() error {
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "path",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.pathPath = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "path",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: path.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "path",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.queryPath = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "path",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
