// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

// WriteJSON implements json.Marshaler.
func (s Error) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("code")
	j.WriteInt64(s.Code)
	field.Write("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes Error json value to io.Writer.
func (s Error) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Error json value from io.Reader.
func (s *Error) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Error from json stream.
func (s *Error) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "code":
			s.Code = i.ReadInt64()
			return i.Error == nil
		case "message":
			s.Message = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s ErrorStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes ErrorStatusCode json value to io.Writer.
func (s ErrorStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads ErrorStatusCode json value from io.Reader.
func (s *ErrorStatusCode) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads ErrorStatusCode from json stream.
func (s *ErrorStatusCode) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPutDefault) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes FoobarPutDefault json value to io.Writer.
func (s FoobarPutDefault) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads FoobarPutDefault json value from io.Reader.
func (s *FoobarPutDefault) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads FoobarPutDefault from json stream.
func (s *FoobarPutDefault) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON writes json value of bool to json stream.
func (o NilBool) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteBool(bool(o.Value))
}

// ReadJSON reads json value of bool from json iterator.
func (o *NilBool) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: bool, JSONType: BoolValue, format "", read "ReadBool".
	switch i.WhatIsNext() {
	case json.BoolValue:
		o.Null = false
		o.Value = bool(i.ReadBool())
		return i.Error
	case json.NilValue:
		var v bool
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilBool", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of time.Duration to json stream.
func (o NilDuration) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	json.WriteDuration(j, o.Value)
}

// ReadJSON reads json value of time.Duration from json iterator.
func (o *NilDuration) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: time.Duration, JSONType: StringValue, format "Duration", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Null = false
		v, err := json.ReadDuration(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v time.Duration
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilDuration", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of float32 to json stream.
func (o NilFloat32) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteFloat32(float32(o.Value))
}

// ReadJSON reads json value of float32 from json iterator.
func (o *NilFloat32) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: float32, JSONType: NumberValue, format "", read "ReadFloat32".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Null = false
		o.Value = float32(i.ReadFloat32())
		return i.Error
	case json.NilValue:
		var v float32
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilFloat32", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of float64 to json stream.
func (o NilFloat64) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteFloat64(float64(o.Value))
}

// ReadJSON reads json value of float64 from json iterator.
func (o *NilFloat64) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: float64, JSONType: NumberValue, format "", read "ReadFloat64".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Null = false
		o.Value = float64(i.ReadFloat64())
		return i.Error
	case json.NilValue:
		var v float64
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilFloat64", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of net.IP to json stream.
func (o NilIP) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	json.WriteIP(j, o.Value)
}

// ReadJSON reads json value of net.IP from json iterator.
func (o *NilIP) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: net.IP, JSONType: StringValue, format "IP", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Null = false
		v, err := json.ReadIP(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v net.IP
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilIP", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int to json stream.
func (o NilInt) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteInt(int(o.Value))
}

// ReadJSON reads json value of int from json iterator.
func (o *NilInt) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: int, JSONType: NumberValue, format "", read "ReadInt".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Null = false
		o.Value = int(i.ReadInt())
		return i.Error
	case json.NilValue:
		var v int
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilInt", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int32 to json stream.
func (o NilInt32) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteInt32(int32(o.Value))
}

// ReadJSON reads json value of int32 from json iterator.
func (o *NilInt32) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: int32, JSONType: NumberValue, format "", read "ReadInt32".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Null = false
		o.Value = int32(i.ReadInt32())
		return i.Error
	case json.NilValue:
		var v int32
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilInt32", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int64 to json stream.
func (o NilInt64) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteInt64(int64(o.Value))
}

// ReadJSON reads json value of int64 from json iterator.
func (o *NilInt64) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: int64, JSONType: NumberValue, format "", read "ReadInt64".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Null = false
		o.Value = int64(i.ReadInt64())
		return i.Error
	case json.NilValue:
		var v int64
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilInt64", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of string to json stream.
func (o NilString) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *NilString) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: string, JSONType: StringValue, format "", read "ReadString".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Null = false
		o.Value = string(i.ReadString())
		return i.Error
	case json.NilValue:
		var v string
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilString", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of time.Time to json stream.
func (o NilTime) WriteJSON(j *json.Stream, format func(*json.Stream, time.Time)) {
	if o.Null {
		j.WriteNil()
		return
	}
	format(j, o.Value)
}

// ReadJSON reads json value of time.Time from json iterator.
func (o *NilTime) ReadJSON(i *json.Iterator, format func(*json.Iterator) (time.Time, error)) error {
	// FormatCustom: true, Type: time.Time, JSONType: StringValue, format "", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Null = false
		v, err := format(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v time.Time
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilTime", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of url.URL to json stream.
func (o NilURL) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	json.WriteURI(j, o.Value)
}

// ReadJSON reads json value of url.URL from json iterator.
func (o *NilURL) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: url.URL, JSONType: StringValue, format "URI", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Null = false
		v, err := json.ReadURI(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v url.URL
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilURL", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of uuid.UUID to json stream.
func (o NilUUID) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	json.WriteUUID(j, o.Value)
}

// ReadJSON reads json value of uuid.UUID from json iterator.
func (o *NilUUID) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: uuid.UUID, JSONType: StringValue, format "UUID", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Null = false
		v, err := json.ReadUUID(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v uuid.UUID
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilUUID", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of bool to json stream.
func (o OptionalBool) WriteJSON(j *json.Stream) {
	j.WriteBool(bool(o.Value))
}

// ReadJSON reads json value of bool from json iterator.
func (o *OptionalBool) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: bool, JSONType: BoolValue, format "", read "ReadBool".
	switch i.WhatIsNext() {
	case json.BoolValue:
		o.Set = true
		o.Value = bool(i.ReadBool())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalBool", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of time.Duration to json stream.
func (o OptionalDuration) WriteJSON(j *json.Stream) {
	json.WriteDuration(j, o.Value)
}

// ReadJSON reads json value of time.Duration from json iterator.
func (o *OptionalDuration) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: time.Duration, JSONType: StringValue, format "Duration", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		v, err := json.ReadDuration(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalDuration", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of float32 to json stream.
func (o OptionalFloat32) WriteJSON(j *json.Stream) {
	j.WriteFloat32(float32(o.Value))
}

// ReadJSON reads json value of float32 from json iterator.
func (o *OptionalFloat32) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: float32, JSONType: NumberValue, format "", read "ReadFloat32".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Value = float32(i.ReadFloat32())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalFloat32", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of float64 to json stream.
func (o OptionalFloat64) WriteJSON(j *json.Stream) {
	j.WriteFloat64(float64(o.Value))
}

// ReadJSON reads json value of float64 from json iterator.
func (o *OptionalFloat64) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: float64, JSONType: NumberValue, format "", read "ReadFloat64".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Value = float64(i.ReadFloat64())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalFloat64", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of net.IP to json stream.
func (o OptionalIP) WriteJSON(j *json.Stream) {
	json.WriteIP(j, o.Value)
}

// ReadJSON reads json value of net.IP from json iterator.
func (o *OptionalIP) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: net.IP, JSONType: StringValue, format "IP", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		v, err := json.ReadIP(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalIP", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int to json stream.
func (o OptionalInt) WriteJSON(j *json.Stream) {
	j.WriteInt(int(o.Value))
}

// ReadJSON reads json value of int from json iterator.
func (o *OptionalInt) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: int, JSONType: NumberValue, format "", read "ReadInt".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Value = int(i.ReadInt())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalInt", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int32 to json stream.
func (o OptionalInt32) WriteJSON(j *json.Stream) {
	j.WriteInt32(int32(o.Value))
}

// ReadJSON reads json value of int32 from json iterator.
func (o *OptionalInt32) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: int32, JSONType: NumberValue, format "", read "ReadInt32".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Value = int32(i.ReadInt32())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalInt32", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int64 to json stream.
func (o OptionalInt64) WriteJSON(j *json.Stream) {
	j.WriteInt64(int64(o.Value))
}

// ReadJSON reads json value of int64 from json iterator.
func (o *OptionalInt64) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: int64, JSONType: NumberValue, format "", read "ReadInt64".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Value = int64(i.ReadInt64())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalInt64", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of bool to json stream.
func (o OptionalNilBool) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteBool(bool(o.Value))
}

// ReadJSON reads json value of bool from json iterator.
func (o *OptionalNilBool) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: bool, JSONType: BoolValue, format "", read "ReadBool".
	switch i.WhatIsNext() {
	case json.BoolValue:
		o.Set = true
		o.Null = false
		o.Value = bool(i.ReadBool())
		return i.Error
	case json.NilValue:
		var v bool
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilBool", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of time.Duration to json stream.
func (o OptionalNilDuration) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	json.WriteDuration(j, o.Value)
}

// ReadJSON reads json value of time.Duration from json iterator.
func (o *OptionalNilDuration) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: time.Duration, JSONType: StringValue, format "Duration", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Null = false
		v, err := json.ReadDuration(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v time.Duration
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilDuration", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of float32 to json stream.
func (o OptionalNilFloat32) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteFloat32(float32(o.Value))
}

// ReadJSON reads json value of float32 from json iterator.
func (o *OptionalNilFloat32) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: float32, JSONType: NumberValue, format "", read "ReadFloat32".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Null = false
		o.Value = float32(i.ReadFloat32())
		return i.Error
	case json.NilValue:
		var v float32
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilFloat32", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of float64 to json stream.
func (o OptionalNilFloat64) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteFloat64(float64(o.Value))
}

// ReadJSON reads json value of float64 from json iterator.
func (o *OptionalNilFloat64) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: float64, JSONType: NumberValue, format "", read "ReadFloat64".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Null = false
		o.Value = float64(i.ReadFloat64())
		return i.Error
	case json.NilValue:
		var v float64
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilFloat64", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of net.IP to json stream.
func (o OptionalNilIP) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	json.WriteIP(j, o.Value)
}

// ReadJSON reads json value of net.IP from json iterator.
func (o *OptionalNilIP) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: net.IP, JSONType: StringValue, format "IP", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Null = false
		v, err := json.ReadIP(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v net.IP
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilIP", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int to json stream.
func (o OptionalNilInt) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteInt(int(o.Value))
}

// ReadJSON reads json value of int from json iterator.
func (o *OptionalNilInt) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: int, JSONType: NumberValue, format "", read "ReadInt".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Null = false
		o.Value = int(i.ReadInt())
		return i.Error
	case json.NilValue:
		var v int
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilInt", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int32 to json stream.
func (o OptionalNilInt32) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteInt32(int32(o.Value))
}

// ReadJSON reads json value of int32 from json iterator.
func (o *OptionalNilInt32) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: int32, JSONType: NumberValue, format "", read "ReadInt32".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Null = false
		o.Value = int32(i.ReadInt32())
		return i.Error
	case json.NilValue:
		var v int32
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilInt32", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int64 to json stream.
func (o OptionalNilInt64) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteInt64(int64(o.Value))
}

// ReadJSON reads json value of int64 from json iterator.
func (o *OptionalNilInt64) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: int64, JSONType: NumberValue, format "", read "ReadInt64".
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Null = false
		o.Value = int64(i.ReadInt64())
		return i.Error
	case json.NilValue:
		var v int64
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilInt64", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of string to json stream.
func (o OptionalNilString) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *OptionalNilString) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: string, JSONType: StringValue, format "", read "ReadString".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Null = false
		o.Value = string(i.ReadString())
		return i.Error
	case json.NilValue:
		var v string
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilString", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of time.Time to json stream.
func (o OptionalNilTime) WriteJSON(j *json.Stream, format func(*json.Stream, time.Time)) {
	if o.Null {
		j.WriteNil()
		return
	}
	format(j, o.Value)
}

// ReadJSON reads json value of time.Time from json iterator.
func (o *OptionalNilTime) ReadJSON(i *json.Iterator, format func(*json.Iterator) (time.Time, error)) error {
	// FormatCustom: true, Type: time.Time, JSONType: StringValue, format "", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Null = false
		v, err := format(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v time.Time
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilTime", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of url.URL to json stream.
func (o OptionalNilURL) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	json.WriteURI(j, o.Value)
}

// ReadJSON reads json value of url.URL from json iterator.
func (o *OptionalNilURL) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: url.URL, JSONType: StringValue, format "URI", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Null = false
		v, err := json.ReadURI(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v url.URL
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilURL", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of uuid.UUID to json stream.
func (o OptionalNilUUID) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	json.WriteUUID(j, o.Value)
}

// ReadJSON reads json value of uuid.UUID from json iterator.
func (o *OptionalNilUUID) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: uuid.UUID, JSONType: StringValue, format "UUID", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Null = false
		v, err := json.ReadUUID(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	case json.NilValue:
		var v uuid.UUID
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilUUID", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of PetType to json stream.
func (o OptionalPetType) WriteJSON(j *json.Stream) {
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of PetType from json iterator.
func (o *OptionalPetType) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: PetType, JSONType: StringValue, format "", read "ReadString".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Value = PetType(i.ReadString())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalPetType", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of string to json stream.
func (o OptionalString) WriteJSON(j *json.Stream) {
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *OptionalString) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: string, JSONType: StringValue, format "", read "ReadString".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Value = string(i.ReadString())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalString", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of time.Time to json stream.
func (o OptionalTime) WriteJSON(j *json.Stream, format func(*json.Stream, time.Time)) {
	format(j, o.Value)
}

// ReadJSON reads json value of time.Time from json iterator.
func (o *OptionalTime) ReadJSON(i *json.Iterator, format func(*json.Iterator) (time.Time, error)) error {
	// FormatCustom: true, Type: time.Time, JSONType: StringValue, format "", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		v, err := format(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalTime", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of url.URL to json stream.
func (o OptionalURL) WriteJSON(j *json.Stream) {
	json.WriteURI(j, o.Value)
}

// ReadJSON reads json value of url.URL from json iterator.
func (o *OptionalURL) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: url.URL, JSONType: StringValue, format "URI", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		v, err := json.ReadURI(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalURL", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of uuid.UUID to json stream.
func (o OptionalUUID) WriteJSON(j *json.Stream) {
	json.WriteUUID(j, o.Value)
}

// ReadJSON reads json value of uuid.UUID from json iterator.
func (o *OptionalUUID) ReadJSON(i *json.Iterator) error {
	// FormatCustom: false, Type: uuid.UUID, JSONType: StringValue, format "UUID", read "".
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		v, err := json.ReadUUID(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalUUID", i.WhatIsNext())
	}
	return nil
}

// WriteJSON implements json.Marshaler.
func (s Pet) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("birthday")
	json.WriteDate(j, s.Birthday)
	// Unsupported kind "pointer" for field "Friends".
	field.Write("id")
	j.WriteInt64(s.ID)
	field.Write("ip")
	json.WriteIP(j, s.IP)
	field.Write("ip_v4")
	json.WriteIP(j, s.IPV4)
	field.Write("ip_v6")
	json.WriteIP(j, s.IPV6)
	field.Write("name")
	j.WriteString(s.Name)
	field.Write("nickname")
	s.Nickname.WriteJSON(j)
	if s.NullStr.Set {
		field.Write("nullStr")
		s.NullStr.WriteJSON(j)
	}
	field.Write("rate")
	json.WriteDuration(j, s.Rate)
	if s.Tag.Set {
		field.Write("tag")
		s.Tag.WriteJSON(j)
	}
	// Unsupported kind "pointer" for field "TestArray1".
	if s.TestDate.Set {
		field.Write("testDate")
		s.TestDate.WriteJSON(j, json.WriteDate)
	}
	if s.TestDateTime.Set {
		field.Write("testDateTime")
		s.TestDateTime.WriteJSON(j, json.WriteDateTime)
	}
	if s.TestDuration.Set {
		field.Write("testDuration")
		s.TestDuration.WriteJSON(j)
	}
	if s.TestFloat1.Set {
		field.Write("testFloat1")
		s.TestFloat1.WriteJSON(j)
	}
	if s.TestInteger1.Set {
		field.Write("testInteger1")
		s.TestInteger1.WriteJSON(j)
	}
	if s.TestTime.Set {
		field.Write("testTime")
		s.TestTime.WriteJSON(j, json.WriteTime)
	}
	if s.Type.Set {
		field.Write("type")
		s.Type.WriteJSON(j)
	}
	field.Write("uri")
	json.WriteURI(j, s.URI)
	field.Write("unique_id")
	json.WriteUUID(j, s.UniqueID)
	j.WriteObjectEnd()
}

// WriteJSONTo writes Pet json value to io.Writer.
func (s Pet) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Pet json value from io.Reader.
func (s *Pet) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Pet from json stream.
func (s *Pet) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "birthday":
			v, err := json.ReadDate(i)
			if err != nil {
				i.ReportError("Field Birthday", err.Error())
				return false
			}
			s.Birthday = v
			return true
		case "friends":
			// Unsupported kind "pointer" for field "Friends".
			i.Skip()
			return true
		case "id":
			s.ID = i.ReadInt64()
			return i.Error == nil
		case "ip":
			v, err := json.ReadIP(i)
			if err != nil {
				i.ReportError("Field IP", err.Error())
				return false
			}
			s.IP = v
			return true
		case "ip_v4":
			v, err := json.ReadIP(i)
			if err != nil {
				i.ReportError("Field IPV4", err.Error())
				return false
			}
			s.IPV4 = v
			return true
		case "ip_v6":
			v, err := json.ReadIP(i)
			if err != nil {
				i.ReportError("Field IPV6", err.Error())
				return false
			}
			s.IPV6 = v
			return true
		case "name":
			s.Name = i.ReadString()
			return i.Error == nil
		case "nickname":
			if err := s.Nickname.ReadJSON(i); err != nil {
				i.ReportError("Field Nickname", err.Error())
				return false
			}
			return true
		case "nullStr":
			s.NullStr.Reset()
			if err := s.NullStr.ReadJSON(i); err != nil {
				i.ReportError("Field NullStr", err.Error())
				return false
			}
			return true
		case "rate":
			v, err := json.ReadDuration(i)
			if err != nil {
				i.ReportError("Field Rate", err.Error())
				return false
			}
			s.Rate = v
			return true
		case "tag":
			s.Tag.Reset()
			if err := s.Tag.ReadJSON(i); err != nil {
				i.ReportError("Field Tag", err.Error())
				return false
			}
			return true
		case "testArray1":
			// Unsupported kind "pointer" for field "TestArray1".
			i.Skip()
			return true
		case "testDate":
			s.TestDate.Reset()
			if err := s.TestDate.ReadJSON(i, json.ReadDate); err != nil {
				i.ReportError("Field TestDate", err.Error())
				return false
			}
			return true
		case "testDateTime":
			s.TestDateTime.Reset()
			if err := s.TestDateTime.ReadJSON(i, json.ReadDateTime); err != nil {
				i.ReportError("Field TestDateTime", err.Error())
				return false
			}
			return true
		case "testDuration":
			s.TestDuration.Reset()
			if err := s.TestDuration.ReadJSON(i); err != nil {
				i.ReportError("Field TestDuration", err.Error())
				return false
			}
			return true
		case "testFloat1":
			s.TestFloat1.Reset()
			if err := s.TestFloat1.ReadJSON(i); err != nil {
				i.ReportError("Field TestFloat1", err.Error())
				return false
			}
			return true
		case "testInteger1":
			s.TestInteger1.Reset()
			if err := s.TestInteger1.ReadJSON(i); err != nil {
				i.ReportError("Field TestInteger1", err.Error())
				return false
			}
			return true
		case "testTime":
			s.TestTime.Reset()
			if err := s.TestTime.ReadJSON(i, json.ReadTime); err != nil {
				i.ReportError("Field TestTime", err.Error())
				return false
			}
			return true
		case "type":
			s.Type.Reset()
			if err := s.Type.ReadJSON(i); err != nil {
				i.ReportError("Field Type", err.Error())
				return false
			}
			return true
		case "uri":
			v, err := json.ReadURI(i)
			if err != nil {
				i.ReportError("Field URI", err.Error())
				return false
			}
			s.URI = v
			return true
		case "unique_id":
			v, err := json.ReadUUID(i)
			if err != nil {
				i.ReportError("Field UniqueID", err.Error())
				return false
			}
			s.UniqueID = v
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetDefault) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetDefault json value to io.Writer.
func (s PetGetDefault) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PetGetDefault json value from io.Reader.
func (s *PetGetDefault) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PetGetDefault from json stream.
func (s *PetGetDefault) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "message":
			s.Message = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetDefaultStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetDefaultStatusCode json value to io.Writer.
func (s PetGetDefaultStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PetGetDefaultStatusCode json value from io.Reader.
func (s *PetGetDefaultStatusCode) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PetGetDefaultStatusCode from json stream.
func (s *PetGetDefaultStatusCode) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}
