package gen

import (
	"fmt"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/internal/bitset"
	"github.com/ogen-go/ogen/openapi"
)

func (g *Generator) generateSecurityAPIKey(s *ir.Security, spec openapi.SecurityScheme) (*ir.Security, error) {
	security := spec.Security
	if name := security.Name; name == "" {
		return nil, errors.Errorf(`invalid "apiKey" name %q`, name)
	}
	s.Format = ir.APIKeySecurityFormat
	s.ParameterName = security.Name
	s.Type.Fields = append(s.Type.Fields, &ir.Field{
		Name: "APIKey",
		Type: ir.Primitive(ir.String, nil),
	})

	switch in := security.In; in {
	case "query":
		s.Kind = ir.QuerySecurity
	case "header":
		s.Kind = ir.HeaderSecurity
		vetHeaderParameterName(g.log, s.ParameterName, spec.Security)
	case "cookie":
		s.Kind = ir.CookieSecurity
	default:
		return nil, errors.Errorf(`unknown "in" value %q`, in)
	}
	return s, nil
}

func (g *Generator) generateSecurityOauth2(s *ir.Security) *ir.Security {
	s.Format = ir.Oauth2SecurityFormat
	s.Kind = ir.HeaderSecurity

	s.Type.Fields = append(s.Type.Fields,
		&ir.Field{
			Name: "Token",
			Type: ir.Primitive(ir.String, nil),
		},
		&ir.Field{
			Name: "Scopes",
			Type: ir.Array(ir.Primitive(ir.String, nil), ir.NilOptional, nil),
		},
	)
	return s
}

func (g *Generator) generateSecurityHTTP(s *ir.Security, spec openapi.SecurityScheme) (*ir.Security, error) {
	security := spec.Security
	s.Kind = ir.HeaderSecurity
	switch scheme := strings.ToLower(security.Scheme); scheme {
	case "basic":
		s.Format = ir.BasicHTTPSecurityFormat
		s.Type.Fields = append(s.Type.Fields,
			&ir.Field{
				Name: "Username",
				Type: ir.Primitive(ir.String, nil),
			},
			&ir.Field{
				Name: "Password",
				Type: ir.Primitive(ir.String, nil),
			},
		)
	case "bearer":
		s.Format = ir.BearerSecurityFormat
		s.Type.Fields = append(s.Type.Fields,
			&ir.Field{
				Name: "Token",
				Type: ir.Primitive(ir.String, nil),
			},
		)
	default:
		return nil, errors.Wrapf(&ErrNotImplemented{Name: "http security scheme"}, "unsupported scheme %q", scheme)
	}
	return s, nil
}

func (g *Generator) generateSecurity(ctx *genctx, operationName string, spec openapi.SecurityScheme) (r *ir.Security, rErr error) {
	security := spec.Security

	typeName, err := pascalNonEmpty(spec.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "security name: %q", spec.Name)
	}

	t := &ir.Type{
		Kind: ir.KindStruct,
		Name: typeName,
	}
	s := &ir.Security{
		Type:        t,
		Description: security.Description,
	}

	// Do not create a type for custom security.
	if security.XOgenCustomSecurity {
		s.Format = ir.CustomSecurityFormat
		return s, nil
	}

	defer func() {
		_, typeAlreadyGenerated := ctx.global.types[typeName]
		if rErr == nil && !typeAlreadyGenerated {
			if err := ctx.saveType(t); err != nil {
				rErr = err
			}
		}
	}()

	switch typ := security.Type; typ {
	case "apiKey":
		return g.generateSecurityAPIKey(s, spec)
	case "http":
		return g.generateSecurityHTTP(s, spec)
	case "oauth2":
		return g.generateSecurityOauth2(s), nil
	case "openIdConnect", "mutualTLS":
		return nil, &ErrNotImplemented{Name: fmt.Sprintf("%s security", typ)}
	default:
		return nil, errors.Errorf("unknown security type %q", typ)
	}
}

func (g *Generator) generateSecurities(
	ctx *genctx,
	operationName string,
	spec openapi.SecurityRequirements,
) (r ir.SecurityRequirements, _ error) {
	indexes := map[string]int{}
	for idx, requirement := range spec {
		var set bitset.Bitset

		if err := func() error {
			for _, scheme := range requirement.Schemes {
				s, err := g.generateSecurity(ctx, operationName, scheme)
				s.Scopes = map[string][]string{
					operationName: scheme.Scopes,
				}
				key := scheme.Name
				if err != nil {
					return errors.Wrapf(err, "security scheme %q", key)
				}
				if curr, ok := g.securities[key]; ok {
					for k, v := range curr.Scopes {
						s.Scopes[k] = v
					}
				}
				g.securities[key] = s

				idx, ok := indexes[key]
				if !ok {
					r.Securities = append(r.Securities, s)
					idx = len(r.Securities) - 1
					indexes[key] = idx
				}
				set.Set(idx, true)
			}
			return nil
		}(); err != nil {
			// Skip entire requirement if at least one security is not implemented.
			err = errors.Wrapf(err, "security requirement %d", idx)
			if err := g.trySkip(err, "Skipping security", requirement); err != nil {
				return r, err
			}
			continue
		}

		r.Requirements = append(r.Requirements, set)
	}
	return r, nil
}
