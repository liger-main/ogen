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

// PatchGuestDriveByIDParams is parameters of patchGuestDriveByID operation.
type PatchGuestDriveByIDParams struct {
	// The id of the guest drive.
	DriveID string
}

func unpackPatchGuestDriveByIDParams(packed middleware.Parameters) (params PatchGuestDriveByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "drive_id",
			In:   "path",
		}
		params.DriveID = packed[key].(string)
	}
	return params
}

func decodePatchGuestDriveByIDParams(args [1]string, r *http.Request) (params PatchGuestDriveByIDParams, _ error) {
	// Decode path: drive_id.
	if err := func() error {
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "drive_id",
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

				params.DriveID = c
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
			Name: "drive_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PatchGuestNetworkInterfaceByIDParams is parameters of patchGuestNetworkInterfaceByID operation.
type PatchGuestNetworkInterfaceByIDParams struct {
	// The id of the guest network interface.
	IfaceID string
}

func unpackPatchGuestNetworkInterfaceByIDParams(packed middleware.Parameters) (params PatchGuestNetworkInterfaceByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "iface_id",
			In:   "path",
		}
		params.IfaceID = packed[key].(string)
	}
	return params
}

func decodePatchGuestNetworkInterfaceByIDParams(args [1]string, r *http.Request) (params PatchGuestNetworkInterfaceByIDParams, _ error) {
	// Decode path: iface_id.
	if err := func() error {
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "iface_id",
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

				params.IfaceID = c
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
			Name: "iface_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PutGuestDriveByIDParams is parameters of putGuestDriveByID operation.
type PutGuestDriveByIDParams struct {
	// The id of the guest drive.
	DriveID string
}

func unpackPutGuestDriveByIDParams(packed middleware.Parameters) (params PutGuestDriveByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "drive_id",
			In:   "path",
		}
		params.DriveID = packed[key].(string)
	}
	return params
}

func decodePutGuestDriveByIDParams(args [1]string, r *http.Request) (params PutGuestDriveByIDParams, _ error) {
	// Decode path: drive_id.
	if err := func() error {
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "drive_id",
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

				params.DriveID = c
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
			Name: "drive_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PutGuestNetworkInterfaceByIDParams is parameters of putGuestNetworkInterfaceByID operation.
type PutGuestNetworkInterfaceByIDParams struct {
	// The id of the guest network interface.
	IfaceID string
}

func unpackPutGuestNetworkInterfaceByIDParams(packed middleware.Parameters) (params PutGuestNetworkInterfaceByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "iface_id",
			In:   "path",
		}
		params.IfaceID = packed[key].(string)
	}
	return params
}

func decodePutGuestNetworkInterfaceByIDParams(args [1]string, r *http.Request) (params PutGuestNetworkInterfaceByIDParams, _ error) {
	// Decode path: iface_id.
	if err := func() error {
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "iface_id",
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

				params.IfaceID = c
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
			Name: "iface_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
