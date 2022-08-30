// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
)

type AllRequestBodiesApplicationJSON SimpleObject

func (*AllRequestBodiesApplicationJSON) allRequestBodiesReq() {}

type AllRequestBodiesApplicationXWwwFormUrlencoded SimpleObject

func (*AllRequestBodiesApplicationXWwwFormUrlencoded) allRequestBodiesReq() {}

type AllRequestBodiesMultipartFormData SimpleObject

func (*AllRequestBodiesMultipartFormData) allRequestBodiesReq() {}

type AllRequestBodiesOK struct {
	Data io.Reader
}

func (s AllRequestBodiesOK) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

type AllRequestBodiesOptionalApplicationJSON OptSimpleObject

func (*AllRequestBodiesOptionalApplicationJSON) allRequestBodiesOptionalReq() {}

type AllRequestBodiesOptionalApplicationXWwwFormUrlencoded OptSimpleObject

func (*AllRequestBodiesOptionalApplicationXWwwFormUrlencoded) allRequestBodiesOptionalReq() {}

type AllRequestBodiesOptionalMultipartFormData OptSimpleObject

func (*AllRequestBodiesOptionalMultipartFormData) allRequestBodiesOptionalReq() {}

type AllRequestBodiesOptionalOK struct {
	Data io.Reader
}

func (s AllRequestBodiesOptionalOK) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

type AllRequestBodiesOptionalReqApplicationOctetStream struct {
	Data io.Reader
}

func (s AllRequestBodiesOptionalReqApplicationOctetStream) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

func (*AllRequestBodiesOptionalReqApplicationOctetStream) allRequestBodiesOptionalReq() {}

type AllRequestBodiesOptionalReqTextPlain struct {
	Data io.Reader
}

func (s AllRequestBodiesOptionalReqTextPlain) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

func (*AllRequestBodiesOptionalReqTextPlain) allRequestBodiesOptionalReq() {}

type AllRequestBodiesReqApplicationOctetStream struct {
	Data io.Reader
}

func (s AllRequestBodiesReqApplicationOctetStream) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

func (*AllRequestBodiesReqApplicationOctetStream) allRequestBodiesReq() {}

type AllRequestBodiesReqTextPlain struct {
	Data io.Reader
}

func (s AllRequestBodiesReqTextPlain) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

func (*AllRequestBodiesReqTextPlain) allRequestBodiesReq() {}

type MaskContentTypeOptionalReq struct {
	Data io.Reader
}

func (s MaskContentTypeOptionalReq) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

// MaskContentTypeOptionalReqWithContentType wraps MaskContentTypeOptionalReq with Content-Type.
type MaskContentTypeOptionalReqWithContentType struct {
	ContentType string
	Content     MaskContentTypeOptionalReq
}

type MaskContentTypeReq struct {
	Data io.Reader
}

func (s MaskContentTypeReq) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

// MaskContentTypeReqWithContentType wraps MaskContentTypeReq with Content-Type.
type MaskContentTypeReqWithContentType struct {
	ContentType string
	Content     MaskContentTypeReq
}

// Ref: #/components/schemas/MaskResponse
type MaskResponse struct {
	ContentType string "json:\"contentType\""
	Content     string "json:\"content\""
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSimpleObject returns new OptSimpleObject with value set to v.
func NewOptSimpleObject(v SimpleObject) OptSimpleObject {
	return OptSimpleObject{
		Value: v,
		Set:   true,
	}
}

// OptSimpleObject is optional SimpleObject.
type OptSimpleObject struct {
	Value SimpleObject
	Set   bool
}

// IsSet returns true if OptSimpleObject was set.
func (o OptSimpleObject) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSimpleObject) Reset() {
	var v SimpleObject
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSimpleObject) SetTo(v SimpleObject) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSimpleObject) Get() (v SimpleObject, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSimpleObject) Or(d SimpleObject) SimpleObject {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/SimpleObject
type SimpleObject struct {
	Name string "json:\"name\""
	Age  OptInt "json:\"age\""
}
