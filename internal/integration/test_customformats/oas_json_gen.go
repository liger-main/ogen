// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	custom0 "github.com/ogen-go/ogen/internal/integration/customformats/phonetype"
	custom1 "github.com/ogen-go/ogen/internal/integration/customformats/rgbatype"
	"github.com/ogen-go/ogen/validate"
)

// Encode encodes custom0.Phone as json.
func (o OptPhone) Encode(e *jx.Encoder, format func(*jx.Encoder, custom0.Phone)) {
	if !o.Set {
		return
	}
	formatPhone().EncodeJSON(e, o.Value)
}

// Decode decodes custom0.Phone from json.
func (o *OptPhone) Decode(d *jx.Decoder, format func(*jx.Decoder) (custom0.Phone, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptPhone to nil")
	}
	o.Set = true
	v, err := formatPhone().DecodeJSON(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptPhone) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, formatPhone().EncodeJSON)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptPhone) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, formatPhone().DecodeJSON)
}

// Encode encodes custom1.RGBA as json.
func (o OptRgba) Encode(e *jx.Encoder, format func(*jx.Encoder, custom1.RGBA)) {
	if !o.Set {
		return
	}
	formatRgba().EncodeJSON(e, o.Value)
}

// Decode decodes custom1.RGBA from json.
func (o *OptRgba) Decode(d *jx.Decoder, format func(*jx.Decoder) (custom1.RGBA, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRgba to nil")
	}
	o.Set = true
	v, err := formatRgba().DecodeJSON(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRgba) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, formatRgba().EncodeJSON)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRgba) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, formatRgba().DecodeJSON)
}

// Encode implements json.Marshaler.
func (s *User) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *User) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("id")
		e.Int64(s.ID)
	}
	{

		e.FieldStart("phone")
		formatPhone().EncodeJSON(e, s.Phone)
	}
	{
		if s.HomePhone.Set {
			e.FieldStart("home_phone")
			s.HomePhone.Encode(e, formatPhone().EncodeJSON)
		}
	}
	{

		e.FieldStart("profile_color")
		formatRgba().EncodeJSON(e, s.ProfileColor)
	}
	{
		if s.BackgroundColor.Set {
			e.FieldStart("background_color")
			s.BackgroundColor.Encode(e, formatRgba().EncodeJSON)
		}
	}
}

var jsonFieldsNameOfUser = [5]string{
	0: "id",
	1: "phone",
	2: "home_phone",
	3: "profile_color",
	4: "background_color",
}

// Decode decodes User from json.
func (s *User) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode User to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "phone":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := formatPhone().DecodeJSON(d)
				s.Phone = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"phone\"")
			}
		case "home_phone":
			if err := func() error {
				s.HomePhone.Reset()
				if err := s.HomePhone.Decode(d, formatPhone().DecodeJSON); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"home_phone\"")
			}
		case "profile_color":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := formatRgba().DecodeJSON(d)
				s.ProfileColor = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_color\"")
			}
		case "background_color":
			if err := func() error {
				s.BackgroundColor.Reset()
				if err := s.BackgroundColor.Decode(d, formatRgba().DecodeJSON); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"background_color\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode User")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUser) {
					name = jsonFieldsNameOfUser[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *User) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *User) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
