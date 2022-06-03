// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeDataCreateRequest(r *http.Request, span trace.Span) (
	req OptData,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptData
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}
