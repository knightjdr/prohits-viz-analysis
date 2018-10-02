package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	// TEST1: filter typical data slice.
	data := []map[string]string{
		{"condition": "a", "readout": "b", "abundance": "10|7", "score": "0.5"},
		{"condition": "c", "readout": "d", "abundance": "15", "score": "0.1"},
		{"condition": "e", "readout": "f", "abundance": "2|6", "score": "0.5"},
		{"condition": "g", "readout": "d", "abundance": "5", "score": "0.8"},
	}
	want := []map[string]interface{}{
		{"condition": "a", "readout": "b", "abundance": "10|7", "score": 0.5},
		{"condition": "c", "readout": "d", "abundance": "15", "score": 0.1},
		{"condition": "g", "readout": "d", "abundance": "5", "score": 0.8},
	}
	filtered, err := Score(data, 0.5, 5, "lte")
	assert.Nil(t, err, "Valid scores should not produce an error")
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly")

	// TEST2: Numeric conversion error returns an error.
	data = []map[string]string{
		{"condition": "a", "readout": "b", "score": "x"},
		{"condition": "c", "readout": "d", "score": "y"},
	}
	filtered, err = Score(data, 0.5, 5, "lte")
	assert.NotNil(t, err, "Numeric conversion of score column should return error")
}
