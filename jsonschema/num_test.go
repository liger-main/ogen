package jsonschema

import (
	"fmt"
	"testing"

	"github.com/go-json-experiment/json"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestNum_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		data    string
		value   Num
		wantErr bool
	}{
		{`0`, Num(`0`), false},
		{`1e1`, Num(`10`), false},
		{`0x0a`, Num(`10`), false},
		// Invalid YAML.
		{`"`, nil, true},
		{`0ee1`, nil, true},
		// Invalid type.
		{`{}`, nil, true},
		{`"100"`, nil, true},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			a := require.New(t)

			var val Num
			err := yaml.Unmarshal([]byte(tt.data), &val)
			if tt.wantErr {
				a.Error(err)
				t.Log("Input:", tt.data)
				t.Log("Error:", err)
				return
			}
			a.NoError(err)
			a.Equal(tt.value, val)
		})
	}
}

func TestNum_UnmarshalNextJSON(t *testing.T) {
	tests := []struct {
		data    string
		value   Num
		wantErr bool
	}{
		{`0`, Num(`0`), false},
		{`1e1`, Num(`1e1`), false},
		// Invalid JSON.
		{``, nil, true},
		{`"`, nil, true},
		{`0ee1`, nil, true},
		// Invalid type.
		{`{}`, nil, true},
		{`"100"`, nil, true},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			a := require.New(t)

			var val Num
			err := json.Unmarshal([]byte(tt.data), &val)
			if tt.wantErr {
				a.Error(err)
				t.Log("Input:", tt.data)
				t.Log("Error:", err)
				return
			}
			a.NoError(err)
			a.Equal(tt.value, val)

			data, err := json.Marshal(val)
			a.NoError(err)
			a.JSONEq(tt.data, string(data))
		})
	}
}