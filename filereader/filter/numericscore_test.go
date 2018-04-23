package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumericScore(t *testing.T) {
	// TEST1: convert typical data slice
	data := []map[string]string{
		{"column1": "a", "column2": "b", "score": "0.5"},
		{"column1": "c", "column2": "d", "score": "0.1"},
	}
	want := []map[string]interface{}{
		{"column1": "a", "column2": "b", "score": 0.5},
		{"column1": "c", "column2": "d", "score": 0.1},
	}
	filtered, err := NumericScore(data)
	assert.Nil(t, err, "Valid scores should not produce an error")
	assert.Equal(t, want, filtered, "Data slice is not being converted correctly")

	// TEST2: non-numeric score column should return err
	data = []map[string]string{
		{"column1": "a", "column2": "b", "score": "x"},
		{"column1": "c", "column2": "d", "score": "y"},
	}
	filtered, err = NumericScore(data)
	assert.NotNil(t, err, "Non-numeric score column should return error")

	// TEST3: Invalid score columns after the first row should get converted to zeros.
	data = []map[string]string{
		{"column1": "a", "column2": "b", "score": "0.5"},
		{"column1": "c", "column2": "d", "score": "x"},
	}
	want = []map[string]interface{}{
		{"column1": "a", "column2": "b", "score": 0.5},
		{"column1": "c", "column2": "d", "score": 0},
	}
	filtered, err = NumericScore(data)
	assert.Nil(t, err, "Valid scores should not produce an error")
	assert.Equal(t, want, filtered, "Data slice is not being converted correctly")
}
