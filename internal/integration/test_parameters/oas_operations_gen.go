// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Operations defines a contract for the operations described by OpenAPI v3 specification.
type Operations interface {
	// ComplicatedParameterNameGet implements GET /complicatedParameterName operation.
	//
	// GET /complicatedParameterName
	ComplicatedParameterNameGet(ctx context.Context, params ComplicatedParameterNameGetParams) error
	// ContentParameters implements contentParameters operation.
	//
	// GET /contentParameters/{path}
	ContentParameters(ctx context.Context, params ContentParametersParams) (*ContentParameters, error)
	// CookieParameter implements cookieParameter operation.
	//
	// Test for cookie param.
	//
	// GET /cookieParameter
	CookieParameter(ctx context.Context, params CookieParameterParams) (*Value, error)
	// HeaderParameter implements headerParameter operation.
	//
	// Test for header param.
	//
	// GET /headerParameter
	HeaderParameter(ctx context.Context, params HeaderParameterParams) (*Value, error)
	// ObjectCookieParameter implements objectCookieParameter operation.
	//
	// GET /objectCookieParameter
	ObjectCookieParameter(ctx context.Context, params ObjectCookieParameterParams) (*OneLevelObject, error)
	// ObjectQueryParameter implements objectQueryParameter operation.
	//
	// GET /objectQueryParameter
	ObjectQueryParameter(ctx context.Context, params ObjectQueryParameterParams) (*ObjectQueryParameterOK, error)
	// PathParameter implements pathParameter operation.
	//
	// Test for path param.
	//
	// GET /pathParameter/{value}
	PathParameter(ctx context.Context, params PathParameterParams) (*Value, error)
	// SameName implements sameName operation.
	//
	// Parameters with different location, but with the same name.
	//
	// GET /same_name/{param}
	SameName(ctx context.Context, params SameNameParams) error
	// SimilarNames implements similarNames operation.
	//
	// Parameters with different location, but with similar names.
	//
	// GET /similarNames
	SimilarNames(ctx context.Context, params SimilarNamesParams) error
}
