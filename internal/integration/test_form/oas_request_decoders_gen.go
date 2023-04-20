// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeOnlyFormRequest(r *http.Request) (
	req *OnlyFormReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		form, err := ht.ParseForm(r)
		if err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}

		var request OnlyFormReq
		defined := func(name string) bool {
			switch name {
			case "field":
				// Form parameter.
				return true
			default:
				return false
			}
		}

		for k := range form {
			if !defined(k) {
				return req, close, errors.Errorf("unexpected field %q", k)
			}
		}
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "field",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					request.Field = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"field\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeOnlyMultipartFileRequest(r *http.Request) (
	req *OnlyMultipartFileReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request OnlyMultipartFileReq
		defined := func(name string) bool {
			switch name {
			case "file":
				// File parameter.
				return true
			default:
				return false
			}
		}

		for k := range form {
			if !defined(k) {
				return req, close, errors.Errorf("unexpected field %q", k)
			}
		}
		for k := range r.MultipartForm.File {
			if !defined(k) {
				return req, close, errors.Errorf("unexpected field %q", k)
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["file"]
				if !ok || len(files) < 1 {
					return validate.ErrFieldRequired
				}
				fh := files[0]

				f, err := fh.Open()
				if err != nil {
					return errors.Wrap(err, "open")
				}
				closers = append(closers, f.Close)
				request.File = ht.MultipartFile{
					Name:   fh.Filename,
					File:   f,
					Header: fh.Header,
				}
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"file\"")
			}
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeOnlyMultipartFormRequest(r *http.Request) (
	req *OnlyMultipartFormReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request OnlyMultipartFormReq
		defined := func(name string) bool {
			switch name {
			case "field":
				// Form parameter.
				return true
			default:
				return false
			}
		}

		for k := range form {
			if !defined(k) {
				return req, close, errors.Errorf("unexpected field %q", k)
			}
		}
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "field",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					request.Field = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"field\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestFormURLEncodedRequest(r *http.Request) (
	req *TestForm,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		form, err := ht.ParseForm(r)
		if err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}

		var request TestForm
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIDVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ID.SetTo(requestDotIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal uuid.UUID
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "description",
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

					request.Description = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"description\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "array",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotArrayVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							requestDotArrayVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.Array = append(request.Array, requestDotArrayVal)
						return nil
					})
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"array\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "object",
				Style:   uri.QueryStyleForm,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotObjectVal TestFormObject
					if err := func() error {
						return requestDotObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					request.Object.SetTo(requestDotObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"object\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "deepObject",
				Style:   uri.QueryStyleDeepObject,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotDeepObjectVal TestFormDeepObject
					if err := func() error {
						return requestDotDeepObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					request.DeepObject.SetTo(requestDotDeepObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"deepObject\"")
				}
			}
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestMultipartRequest(r *http.Request) (
	req *TestFormMultipart,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request TestFormMultipart
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIDVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ID.SetTo(requestDotIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal uuid.UUID
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "description",
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

					request.Description = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"description\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "array",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotArrayVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							requestDotArrayVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.Array = append(request.Array, requestDotArrayVal)
						return nil
					})
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"array\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "object",
				Style:   uri.QueryStyleForm,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotObjectVal TestFormMultipartObject
					if err := func() error {
						return requestDotObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					request.Object.SetTo(requestDotObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"object\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "deepObject",
				Style:   uri.QueryStyleDeepObject,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotDeepObjectVal TestFormMultipartDeepObject
					if err := func() error {
						return requestDotDeepObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					request.DeepObject.SetTo(requestDotDeepObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"deepObject\"")
				}
			}
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestMultipartUploadRequest(r *http.Request) (
	req *TestMultipartUploadReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request TestMultipartUploadReq
		defined := func(name string) bool {
			switch name {
			case "orderId":
				// Form parameter.
				return true
			case "userId":
				// Form parameter.
				return true
			case "file":
				// File parameter.
				return true
			case "optional_file":
				// File parameter.
				return true
			case "files":
				// File parameter.
				return true
			default:
				return false
			}
		}

		for k := range form {
			if !defined(k) {
				return req, close, errors.Errorf("unexpected field %q", k)
			}
		}
		for k := range r.MultipartForm.File {
			if !defined(k) {
				return req, close, errors.Errorf("unexpected field %q", k)
			}
		}
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "orderId",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOrderIdVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotOrderIdVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.OrderId.SetTo(requestDotOrderIdVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"orderId\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "userId",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUserIdVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotUserIdVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UserId.SetTo(requestDotUserIdVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"userId\"")
				}
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["file"]
				if !ok || len(files) < 1 {
					return validate.ErrFieldRequired
				}
				fh := files[0]

				f, err := fh.Open()
				if err != nil {
					return errors.Wrap(err, "open")
				}
				closers = append(closers, f.Close)
				request.File = ht.MultipartFile{
					Name:   fh.Filename,
					File:   f,
					Header: fh.Header,
				}
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"file\"")
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["optional_file"]
				if !ok || len(files) < 1 {
					return nil
				}
				fh := files[0]

				f, err := fh.Open()
				if err != nil {
					return errors.Wrap(err, "open")
				}
				closers = append(closers, f.Close)
				request.OptionalFile.SetTo(ht.MultipartFile{
					Name:   fh.Filename,
					File:   f,
					Header: fh.Header,
				})
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"optional_file\"")
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["files"]
				_ = ok
				request.Files = make([]ht.MultipartFile, 0, len(files))
				for _, fh := range files {
					f, err := fh.Open()
					if err != nil {
						return errors.Wrap(err, "open")
					}
					closers = append(closers, f.Close)

					request.Files = append(request.Files, ht.MultipartFile{
						Name:   fh.Filename,
						File:   f,
						Header: fh.Header,
					})
				}
				if err := func() error {
					if err := (validate.Array{
						MinLength:    0,
						MinLengthSet: false,
						MaxLength:    5,
						MaxLengthSet: true,
					}).ValidateLength(len(request.Files)); err != nil {
						return errors.Wrap(err, "array")
					}
					return nil
				}(); err != nil {
					return errors.Wrap(err, "validate")
				}
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"files\"")
			}
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestReuseFormOptionalSchemaRequest(r *http.Request) (
	req OptSharedRequestMultipart,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	if _, ok := r.Header["Content-Type"]; !ok && r.ContentLength == 0 {
		return req, close, nil
	}
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, nil
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request OptSharedRequestMultipart
		{
			var optForm SharedRequestMultipart
			q := uri.NewQueryDecoder(form)
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "filename",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var optFormDotFilenameVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							optFormDotFilenameVal = c
							return nil
						}(); err != nil {
							return err
						}
						optForm.Filename.SetTo(optFormDotFilenameVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"filename\"")
					}
				}
			}
			{
				if err := func() error {
					files, ok := r.MultipartForm.File["file"]
					if !ok || len(files) < 1 {
						return nil
					}
					fh := files[0]

					f, err := fh.Open()
					if err != nil {
						return errors.Wrap(err, "open")
					}
					closers = append(closers, f.Close)
					optForm.File.SetTo(ht.MultipartFile{
						Name:   fh.Filename,
						File:   f,
						Header: fh.Header,
					})
					return nil
				}(); err != nil {
					return req, close, errors.Wrap(err, "decode \"file\"")
				}
			}
			request = OptSharedRequestMultipart{
				Value: optForm,
				Set:   true,
			}
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestReuseFormSchemaRequest(r *http.Request) (
	req *SharedRequestMultipart,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request SharedRequestMultipart
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Filename.SetTo(requestDotFilenameVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"filename\"")
				}
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["file"]
				if !ok || len(files) < 1 {
					return nil
				}
				fh := files[0]

				f, err := fh.Open()
				if err != nil {
					return errors.Wrap(err, "open")
				}
				closers = append(closers, f.Close)
				request.File.SetTo(ht.MultipartFile{
					Name:   fh.Filename,
					File:   f,
					Header: fh.Header,
				})
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"file\"")
			}
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestShareFormSchemaRequest(r *http.Request) (
	req TestShareFormSchemaReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request SharedRequest
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		// Form values may be unused.
		form := url.Values(r.MultipartForm.Value)
		_ = form

		var request SharedRequestMultipart
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Filename.SetTo(requestDotFilenameVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"filename\"")
				}
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["file"]
				if !ok || len(files) < 1 {
					return nil
				}
				fh := files[0]

				f, err := fh.Open()
				if err != nil {
					return errors.Wrap(err, "open")
				}
				closers = append(closers, f.Close)
				request.File.SetTo(ht.MultipartFile{
					Name:   fh.Filename,
					File:   f,
					Header: fh.Header,
				})
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"file\"")
			}
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}
