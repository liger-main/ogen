// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DataGetParams is parameters of dataGet operation.
type DataGetParams struct {
	ID  int
	Key string
}

func unpackDataGetParams(packed middleware.Parameters) (params DataGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "key",
			In:   "path",
		}
		params.Key = packed[key].(string)
	}
	return params
}

func decodeDataGetParams(args [2]string, r *http.Request) (params DataGetParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.ID)); err != nil {
					return errors.Wrap(err, "int")
				}
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
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: key.
	if err := func() error {
		param := args[1]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "key",
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

				params.Key = c
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
			Name: "key",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DataGetIDParams is parameters of dataGetID operation.
type DataGetIDParams struct {
	ID int
}

func unpackDataGetIDParams(packed middleware.Parameters) (params DataGetIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeDataGetIDParams(args [1]string, r *http.Request) (params DataGetIDParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.ID)); err != nil {
					return errors.Wrap(err, "int")
				}
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
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
