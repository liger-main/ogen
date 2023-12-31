// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
)

var securityScopes = map[string][]string{
	"SecurityTest": []string{
		"claim",
	},
	"TestNullableOneofs": []string{
		"claim",
	},
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// APIKey provides api_key security value.
	APIKey(ctx context.Context, operationName string, scopes []string, paramsByName map[string]interface{}) (APIKey, error)
}

func (s *Client) securityAPIKey(ctx context.Context, operationName string, paramsByName map[string]interface{}, req *http.Request) error {
	t, err := s.sec.APIKey(ctx, operationName, securityScopes[operationName], paramsByName)
	if err != nil {
		return errors.Wrap(err, "security source \"APIKey\"")
	}
	req.Header.Set("Api_key", t.APIKey)
	return nil
}
