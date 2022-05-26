// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/validate"
)

func decodeAPICaptcha2chcaptchaIDGetResponse(resp *http.Response, span trace.Span) (res Captcha, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Captcha
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
func decodeAPICaptcha2chcaptchaShowGetResponse(resp *http.Response, span trace.Span) (res APICaptcha2chcaptchaShowGetRes, err error) {
	switch resp.StatusCode {
	case 200:
		return &APICaptcha2chcaptchaShowGetOK{}, nil
	case 404:
		return &APICaptcha2chcaptchaShowGetNotFound{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeAPICaptchaAppIDPublicKeyGetResponse(resp *http.Response, span trace.Span) (res Captcha, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Captcha
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
func decodeAPICaptchaInvisibleRecaptchaIDGetResponse(resp *http.Response, span trace.Span) (res Captcha, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Captcha
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
func decodeAPICaptchaInvisibleRecaptchaMobileGetResponse(resp *http.Response, span trace.Span) (res APICaptchaInvisibleRecaptchaMobileGetOK, err error) {
	switch resp.StatusCode {
	case 200:
		return APICaptchaInvisibleRecaptchaMobileGetOK{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeAPICaptchaRecaptchaIDGetResponse(resp *http.Response, span trace.Span) (res Captcha, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Captcha
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
func decodeAPICaptchaRecaptchaMobileGetResponse(resp *http.Response, span trace.Span) (res APICaptchaRecaptchaMobileGetOK, err error) {
	switch resp.StatusCode {
	case 200:
		return APICaptchaRecaptchaMobileGetOK{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeAPIDislikeGetResponse(resp *http.Response, span trace.Span) (res Like, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Like
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
func decodeAPILikeGetResponse(resp *http.Response, span trace.Span) (res Like, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Like
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
func decodeAPIMobileV2AfterBoardThreadNumGetResponse(resp *http.Response, span trace.Span) (res MobileThreadPostsAfter, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response MobileThreadPostsAfter
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
func decodeAPIMobileV2BoardsGetResponse(resp *http.Response, span trace.Span) (res Boards, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Boards
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
func decodeAPIMobileV2InfoBoardThreadGetResponse(resp *http.Response, span trace.Span) (res MobileThreadLastInfo, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response MobileThreadLastInfo
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
func decodeAPIMobileV2PostBoardNumGetResponse(resp *http.Response, span trace.Span) (res MobilePost, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response MobilePost
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
func decodeUserPassloginPostResponse(resp *http.Response, span trace.Span) (res Passcode, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Passcode
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
func decodeUserPostingPostResponse(resp *http.Response, span trace.Span) (res UserPostingPostOK, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response UserPostingPostOK
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
func decodeUserReportPostResponse(resp *http.Response, span trace.Span) (res Report, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Report
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
