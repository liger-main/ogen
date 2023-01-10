// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleBasicAuth handles basicAuth security.
	HandleBasicAuth(ctx context.Context, operationName string, t BasicAuth) (context.Context, error)
	// HandleBearerToken handles bearerToken security.
	HandleBearerToken(ctx context.Context, operationName string, t BearerToken) (context.Context, error)
	// HandleHeaderKey handles headerKey security.
	HandleHeaderKey(ctx context.Context, operationName string, t HeaderKey) (context.Context, error)
	// HandleQueryKey handles queryKey security.
	HandleQueryKey(ctx context.Context, operationName string, t QueryKey) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityBasicAuth(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t BasicAuth
	if _, ok := findAuthorization(req.Header, "Basic"); !ok {
		return ctx, false, nil
	}
	username, password, ok := req.BasicAuth()
	if !ok {
		return nil, false, errors.New("invalid basic auth")
	}
	t.Username = username
	t.Password = password
	rctx, err := s.sec.HandleBasicAuth(ctx, operationName, t)
	if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}
func (s *Server) securityBearerToken(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t BearerToken
	token, ok := findAuthorization(req.Header, "Bearer")
	if !ok {
		return ctx, false, nil
	}
	t.Token = token
	rctx, err := s.sec.HandleBearerToken(ctx, operationName, t)
	if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}
func (s *Server) securityHeaderKey(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t HeaderKey
	const parameterName = "X-API-Key"
	value := req.Header.Get(parameterName)
	if value == "" {
		return ctx, false, nil
	}
	t.APIKey = value
	rctx, err := s.sec.HandleHeaderKey(ctx, operationName, t)
	if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}
func (s *Server) securityQueryKey(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t QueryKey
	const parameterName = "api_key"
	q := req.URL.Query()
	if !q.Has(parameterName) {
		return ctx, false, nil
	}
	value := q.Get(parameterName)
	t.APIKey = value
	rctx, err := s.sec.HandleQueryKey(ctx, operationName, t)
	if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// BasicAuth provides basicAuth security value.
	BasicAuth(ctx context.Context, operationName string) (BasicAuth, error)
	// BearerToken provides bearerToken security value.
	BearerToken(ctx context.Context, operationName string) (BearerToken, error)
	// HeaderKey provides headerKey security value.
	HeaderKey(ctx context.Context, operationName string) (HeaderKey, error)
	// QueryKey provides queryKey security value.
	QueryKey(ctx context.Context, operationName string) (QueryKey, error)
}

func (s *Client) securityBasicAuth(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.BasicAuth(ctx, operationName)
	if err != nil {
		if errors.Is(err, ogenerrors.ErrSkipClientSecurity) {
			return ogenerrors.ErrSkipClientSecurity
		}
		return errors.Wrap(err, "security source \"BasicAuth\"")
	}
	req.SetBasicAuth(t.Username, t.Password)
	return nil
}
func (s *Client) securityBearerToken(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.BearerToken(ctx, operationName)
	if err != nil {
		if errors.Is(err, ogenerrors.ErrSkipClientSecurity) {
			return ogenerrors.ErrSkipClientSecurity
		}
		return errors.Wrap(err, "security source \"BearerToken\"")
	}
	req.Header.Set("Authorization", "Bearer "+t.Token)
	return nil
}
func (s *Client) securityHeaderKey(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.HeaderKey(ctx, operationName)
	if err != nil {
		if errors.Is(err, ogenerrors.ErrSkipClientSecurity) {
			return ogenerrors.ErrSkipClientSecurity
		}
		return errors.Wrap(err, "security source \"HeaderKey\"")
	}
	req.Header.Set("X-API-Key", t.APIKey)
	return nil
}
func (s *Client) securityQueryKey(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.QueryKey(ctx, operationName)
	if err != nil {
		if errors.Is(err, ogenerrors.ErrSkipClientSecurity) {
			return ogenerrors.ErrSkipClientSecurity
		}
		return errors.Wrap(err, "security source \"QueryKey\"")
	}
	q := req.URL.Query()
	q.Set("api_key", t.APIKey)
	req.URL.RawQuery = q.Encode()
	return nil
}
