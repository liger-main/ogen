package parser

import (
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/internal/oas"
	"github.com/ogen-go/ogen/jsonschema"
)

type resolveCtx map[string]struct{}

type componentsResolver map[string]*ogen.Schema

func (c componentsResolver) ResolveReference(ref string) (*jsonschema.RawSchema, error) {
	const prefix = "#/components/schemas/"
	if !strings.HasPrefix(ref, prefix) {
		return nil, errors.Errorf("invalid schema reference %q", ref)
	}

	name := strings.TrimPrefix(ref, prefix)
	s, ok := c[name]
	if !ok {
		return nil, errors.New("schema not found")
	}

	return s.ToJSONSchema(), nil
}

func (p *parser) resolveRequestBody(ref string, ctx resolveCtx) (*oas.RequestBody, error) {
	const prefix = "#/components/requestBodies/"
	if !strings.HasPrefix(ref, prefix) {
		return nil, errors.Errorf("invalid requestBody reference: %q", ref)
	}

	if r, ok := p.refs.requestBodies[ref]; ok {
		return r, nil
	}

	if _, ok := ctx[ref]; ok {
		return nil, errors.Errorf("infinite recursion: %q", ref)
	}
	ctx[ref] = struct{}{}

	name := strings.TrimPrefix(ref, prefix)
	component, found := p.spec.Components.RequestBodies[name]
	if !found {
		return nil, errors.Errorf("component by reference %q not found", ref)
	}

	r, err := p.parseRequestBody(component, ctx)
	if err != nil {
		return nil, err
	}

	p.refs.requestBodies[ref] = r
	return r, nil
}

func (p *parser) resolveResponse(ref string, ctx resolveCtx) (*oas.Response, error) {
	const prefix = "#/components/responses/"
	if !strings.HasPrefix(ref, prefix) {
		return nil, errors.Errorf("invalid response reference: %q", ref)
	}

	if r, ok := p.refs.responses[ref]; ok {
		return r, nil
	}

	if _, ok := ctx[ref]; ok {
		return nil, errors.Errorf("infinite recursion: %q", ref)
	}
	ctx[ref] = struct{}{}

	name := strings.TrimPrefix(ref, prefix)
	component, found := p.spec.Components.Responses[name]
	if !found {
		return nil, errors.Errorf("component by reference %q not found", ref)
	}

	r, err := p.parseResponse(component, ctx)
	if err != nil {
		return nil, err
	}

	r.Ref = ref
	p.refs.responses[ref] = r
	return r, nil
}

func (p *parser) resolveParameter(ref string, ctx resolveCtx) (*oas.Parameter, error) {
	const prefix = "#/components/parameters/"
	if !strings.HasPrefix(ref, prefix) {
		return nil, errors.Errorf("invalid parameter reference: %q", ref)
	}

	if param, ok := p.refs.parameters[ref]; ok {
		return param, nil
	}

	if _, ok := ctx[ref]; ok {
		return nil, errors.Errorf("infinite recursion: %q", ref)
	}
	ctx[ref] = struct{}{}

	name := strings.TrimPrefix(ref, prefix)
	component, found := p.spec.Components.Parameters[name]
	if !found {
		return nil, errors.Errorf("component by reference %q not found", ref)
	}

	param, err := p.parseParameter(component, ctx)
	if err != nil {
		return nil, err
	}

	p.refs.parameters[ref] = param
	return param, nil
}

func (p *parser) resolveExample(ref string) (*oas.Example, error) {
	const prefix = "#/components/examples/"
	if !strings.HasPrefix(ref, prefix) {
		return nil, errors.Errorf("invalid parameter reference: %q", ref)
	}

	if param, ok := p.refs.examples[ref]; ok {
		return param, nil
	}

	name := strings.TrimPrefix(ref, prefix)
	component, found := p.spec.Components.Examples[name]
	if !found {
		return nil, errors.Errorf("component by reference %q not found", ref)
	}
	example := &oas.Example{
		Ref:           component.Ref,
		Summary:       component.Summary,
		Description:   component.Description,
		Value:         component.Value,
		ExternalValue: component.ExternalValue,
	}

	p.refs.examples[ref] = example
	return example, nil
}
