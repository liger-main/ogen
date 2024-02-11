// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Operations defines a contract for the operations described by OpenAPI v3 specification.
type defaultOperations interface {
	// CustomSecurity implements customSecurity operation.
	//
	// GET /customSecurity
	CustomSecurity(ctx context.Context) error
	// DisjointSecurity implements disjointSecurity operation.
	//
	// GET /disjointSecurity
	DisjointSecurity(ctx context.Context) error
	// IntersectSecurity implements intersectSecurity operation.
	//
	// GET /intersectSecurity
	IntersectSecurity(ctx context.Context) error
	// OptionalSecurity implements optionalSecurity operation.
	//
	// GET /optionalSecurity
	OptionalSecurity(ctx context.Context) error
}

type Operations interface {
	defaultOperations
}
