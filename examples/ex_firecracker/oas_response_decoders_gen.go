// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

func decodeCreateSnapshotResponse(resp *http.Response) (res CreateSnapshotRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &CreateSnapshotNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeCreateSyncActionResponse(resp *http.Response) (res CreateSyncActionRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &CreateSyncActionNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeDescribeBalloonConfigResponse(resp *http.Response) (res DescribeBalloonConfigRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Balloon
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
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeDescribeBalloonStatsResponse(resp *http.Response) (res DescribeBalloonStatsRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response BalloonStats
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
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeDescribeInstanceResponse(resp *http.Response) (res DescribeInstanceRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response InstanceInfo
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeGetExportVmConfigResponse(resp *http.Response) (res GetExportVmConfigRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response FullVmConfiguration
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeGetMachineConfigurationResponse(resp *http.Response) (res GetMachineConfigurationRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response MachineConfiguration
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeLoadSnapshotResponse(resp *http.Response) (res LoadSnapshotRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &LoadSnapshotNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeMmdsConfigPutResponse(resp *http.Response) (res MmdsConfigPutRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &MmdsConfigPutNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeMmdsGetResponse(resp *http.Response) (res MmdsGetRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response MmdsGetOK
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
		// Code 404.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeMmdsPatchResponse(resp *http.Response) (res MmdsPatchRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &MmdsPatchNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodeMmdsPutResponse(resp *http.Response) (res MmdsPutRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &MmdsPutNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePatchBalloonResponse(resp *http.Response) (res PatchBalloonRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PatchBalloonNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePatchBalloonStatsIntervalResponse(resp *http.Response) (res PatchBalloonStatsIntervalRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PatchBalloonStatsIntervalNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePatchGuestDriveByIDResponse(resp *http.Response) (res PatchGuestDriveByIDRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PatchGuestDriveByIDNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePatchGuestNetworkInterfaceByIDResponse(resp *http.Response) (res PatchGuestNetworkInterfaceByIDRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PatchGuestNetworkInterfaceByIDNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePatchMachineConfigurationResponse(resp *http.Response) (res PatchMachineConfigurationRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PatchMachineConfigurationNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePatchVmResponse(resp *http.Response) (res PatchVmRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PatchVmNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePutBalloonResponse(resp *http.Response) (res PutBalloonRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PutBalloonNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePutGuestBootSourceResponse(resp *http.Response) (res PutGuestBootSourceRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PutGuestBootSourceNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePutGuestDriveByIDResponse(resp *http.Response) (res PutGuestDriveByIDRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PutGuestDriveByIDNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePutGuestNetworkInterfaceByIDResponse(resp *http.Response) (res PutGuestNetworkInterfaceByIDRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PutGuestNetworkInterfaceByIDNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePutGuestVsockResponse(resp *http.Response) (res PutGuestVsockRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PutGuestVsockNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePutLoggerResponse(resp *http.Response) (res PutLoggerRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PutLoggerNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePutMachineConfigurationResponse(resp *http.Response) (res PutMachineConfigurationRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PutMachineConfigurationNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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

func decodePutMetricsResponse(resp *http.Response) (res PutMetricsRes, err error) {
	switch resp.StatusCode {
	case 204:
		// Code 204.
		return &PutMetricsNoContent{}, nil
	case 400:
		// Code 400.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
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
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
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
