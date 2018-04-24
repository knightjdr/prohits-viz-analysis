package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumericScore(t *testing.T) {
	// TEST1: convert typical data slice.
	data := []map[string]string{
		{"score": "0.5"},
		{"score": "0.1"},
	}
	want := []map[string]interface{}{
		{"score": 0.5},
		{"score": 0.1},
	}
	filtered, err := NumericScore(data)
	assert.Nil(t, err, "Valid scores should not produce an error")
	assert.Equal(t, want, filtered, "Data slice is not being converted correctly")

	// TEST2: non-numeric score column should return err.
	data = []map[string]string{
		{"score": "x"},
		{"score": "y"},
	}
	filtered, err = NumericScore(data)
	assert.NotNil(t, err, "Non-numeric score column should return error")

	// TEST3: Invalid score columns after the first row should get converted to zeros.
	data = []map[string]string{
		{"score": "0.5"},
		{"score": "x"},
	}
	want = []map[string]interface{}{
		{"score": 0.5},
		{"score": 0},
	}
	filtered, err = NumericScore(data)
	assert.Nil(t, err, "Valid scores should not produce an error")
	assert.Equal(t, want, filtered, "Data slice is not being converted correctly")
}
