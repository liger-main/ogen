package json

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/go-json-experiment/json"
)

// LineColumn returns the line and column of the location.
//
// If offset is invalid, line and column are 0 and ok is false.
func LineColumn(offset int64, data []byte) (line, column int64, ok bool) {
	if offset < 0 || int64(len(data)) <= offset {
		return 0, 0, false
	}

	{
		unread := data[offset:]
		trimmed := bytes.TrimLeft(unread, "\x20\t\r\n,:")
		if len(trimmed) != len(unread) {
			// Skip leading whitespace, because decoder does not do it.
			offset += int64(len(unread) - len(trimmed))
		}
	}

	lines := data[:offset]
	// Lines count from 1.
	line = int64(bytes.Count(lines, []byte("\n"))) + 1
	// Find offset from last newline.
	lastNL := int64(bytes.LastIndexByte(lines, '\n'))
	column = offset - lastNL
	return line, column, true
}

// Location is a JSON value location.
type Location struct {
	JSONPointer  string `json:"-"`
	Offset       int64  `json:"-"`
	Line, Column int64  `json:"-"`
}

// String implements fmt.Stringer.
func (l Location) String() string {
	if l.Line == 0 {
		return strconv.Quote(l.JSONPointer)
	}
	return fmt.Sprintf("line %d:%d", l.Line, l.Column)
}

// Locatable is an interface for JSON value location store.
type Locatable interface {
	setLocation(Location)
	Location() (Location, bool)
}

// Locator stores the Location of a JSON value.
type Locator struct {
	location Location
	set      bool
}

func (l *Locator) setLocation(loc Location) {
	l.location = loc
	l.set = true
}

// Location returns the location of the value if it is set.
func (l Locator) Location() (Location, bool) {
	return l.location, l.set
}

// LocationUnmarshaler is json.Unmarshalers that sets the location.
func LocationUnmarshaler(data []byte) *json.Unmarshalers {
	return json.UnmarshalFuncV2(func(opts json.UnmarshalOptions, d *json.Decoder, l Locatable) error {
		if _, ok := l.(*Locator); ok {
			return nil
		}

		offset := d.InputOffset()
		line, column, _ := LineColumn(offset, data)
		l.setLocation(Location{
			JSONPointer: d.StackPointer(),
			Offset:      offset,
			Line:        line,
			Column:      column,
		})
		return json.SkipFunc
	})
}
