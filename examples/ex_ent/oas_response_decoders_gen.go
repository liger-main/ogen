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

func decodeCreatePetResponse(resp *http.Response, span trace.Span) (res CreatePetRes, err error) {
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
			var response PetCreate
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
			var response R400
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
	case 409:
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
			var response R409
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeCreatePetCategoriesResponse(resp *http.Response, span trace.Span) (res CreatePetCategoriesRes, err error) {
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
			var response PetCategoriesCreate
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
			var response R400
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
	case 409:
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
			var response R409
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeCreatePetFriendsResponse(resp *http.Response, span trace.Span) (res CreatePetFriendsRes, err error) {
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
			var response PetFriendsCreate
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
			var response R400
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
	case 409:
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
			var response R409
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeCreatePetOwnerResponse(resp *http.Response, span trace.Span) (res CreatePetOwnerRes, err error) {
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
			var response PetOwnerCreate
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
			var response R400
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
	case 409:
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
			var response R409
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeDeletePetResponse(resp *http.Response, span trace.Span) (res DeletePetRes, err error) {
	switch resp.StatusCode {
	case 204:
		return &DeletePetNoContent{}, nil
	case 400:
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
			var response R400
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
			var response R404
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeDeletePetOwnerResponse(resp *http.Response, span trace.Span) (res DeletePetOwnerRes, err error) {
	switch resp.StatusCode {
	case 204:
		return &DeletePetOwnerNoContent{}, nil
	case 400:
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
			var response R400
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
			var response R404
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeListPetResponse(resp *http.Response, span trace.Span) (res ListPetRes, err error) {
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
			var response ListPetOKApplicationJSON
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
			var response R400
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
			var response R404
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeListPetCategoriesResponse(resp *http.Response, span trace.Span) (res ListPetCategoriesRes, err error) {
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
			var response ListPetCategoriesOKApplicationJSON
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
			var response R400
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
			var response R404
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeListPetFriendsResponse(resp *http.Response, span trace.Span) (res ListPetFriendsRes, err error) {
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
			var response ListPetFriendsOKApplicationJSON
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
			var response R400
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
			var response R404
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeReadPetResponse(resp *http.Response, span trace.Span) (res ReadPetRes, err error) {
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
			var response PetRead
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
			var response R400
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
			var response R404
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeReadPetOwnerResponse(resp *http.Response, span trace.Span) (res ReadPetOwnerRes, err error) {
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
			var response PetOwnerRead
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
			var response R400
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
			var response R404
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeUpdatePetResponse(resp *http.Response, span trace.Span) (res UpdatePetRes, err error) {
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
			var response PetUpdate
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
			var response R400
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
			var response R404
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
	case 500:
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
			var response R500
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
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
