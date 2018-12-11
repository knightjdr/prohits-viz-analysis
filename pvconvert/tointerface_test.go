package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInterface(t *testing.T) {
	// TEST: with score key
	csv := []map[string]string{
		{"abundance": "23", "condition": "a", "readout": "x", "score": "0.01"},
		{"abundance": "5", "condition": "b", "readout": "y", "score": "0.56"},
		{"abundance": "13.4", "condition": "c", "readout": "z", "score": "0"},
	}
	expected := []map[string]interface{}{
		{"abundance": float64(23), "condition": "a", "readout": "x", "score": float64(0.01)},
		{"abundance": float64(5), "condition": "b", "readout": "y", "score": float64(0.56)},
		{"abundance": float64(13.4), "condition": "c", "readout": "z", "score": float64(0)},
	}
	assert.Equal(
		t,
		expected,
		toInterface(csv),
		"CSV with score key not converted to interface",
	)

	// TEST: without score key
	csv = []map[string]string{
		{"abundance": "23", "condition": "a", "readout": "x"},
		{"abundance": "5", "condition": "b", "readout": "y"},
		{"abundance": "13.4", "condition": "c", "readout": "z"},
	}
	expected = []map[string]interface{}{
		{"abundance": float64(23), "condition": "a", "readout": "x", "score": float64(0)},
		{"abundance": float64(5), "condition": "b", "readout": "y", "score": float64(0)},
		{"abundance": float64(13.4), "condition": "c", "readout": "z", "score": float64(0)},
	}
	assert.Equal(
		t,
		expected,
		toInterface(csv),
		"CSV without score key not converted to interface",
	)
}
