package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	// TEST1: filter typical data slice
	data := []map[string]string{
		{"bait": "a", "prey": "b", "score": "0.5"},
		{"bait": "c", "prey": "d", "score": "0.1"},
		{"bait": "e", "prey": "f", "score": "0.8"},
		{"bait": "g", "prey": "d", "score": "0.8"},
	}
	want := []map[string]interface{}{
		{"bait": "a", "prey": "b", "score": 0.5},
		{"bait": "c", "prey": "d", "score": 0.1},
		{"bait": "g", "prey": "d", "score": 0.8},
	}
	filtered, err := Score(data, 0.5, "lte")
	assert.Nil(t, err, "Valid scores should not produce an error")
	assert.Equal(t, want, filtered, "Data slice is not being filtered correctly")

	// TEST2: Numeric conversion error returns an error
	data = []map[string]string{
		{"bait": "a", "prey": "b", "score": "x"},
		{"bait": "c", "prey": "d", "score": "y"},
	}
	filtered, err = Score(data, 0.5, "lte")
	assert.NotNil(t, err, "Numeric conversion of score column should return error")
}
