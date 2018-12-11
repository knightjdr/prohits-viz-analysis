package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestControlSubtraction(t *testing.T) {
	data := []map[string]string{
		{"abundance": "10", "control": "2|3|10"},
		{"abundance": "1", "control": "2|3|10"},
		{"abundance": "10|5", "control": "2|3|10"},
		{"abundance": "10|5|2.5", "control": "2"},
	}
	want := []map[string]string{
		{"abundance": "5", "control": "2|3|10"},
		{"abundance": "0", "control": "2|3|10"},
		{"abundance": "5|0", "control": "2|3|10"},
		{"abundance": "8|3|0.5", "control": "2"},
	}

	// TEST1: different abundance and control column formats.
	transformed := ControlSubtraction(data, "controlColumn")
	assert.Equal(t, want, transformed, "Control subtraction is not correct")

	// TEST2: when control subtraction is not requested, return original data.
	transformed = ControlSubtraction(data, "")
	assert.Equal(t, data, transformed, "No control subtraction should return input data")
}
