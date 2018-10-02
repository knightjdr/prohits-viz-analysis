package columnparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MappedTerms struct {
	key1 string
	key2 string
}

func TestMapLines(t *testing.T) {
	// values are column indexes to map to the key names
	headerMap := map[string]int{
		"key1": 0,
		"key2": 2,
	}
	// array of lines from a csv file
	lines := [][]string{
		{"condition1", "x", "readout1"},
		{"condition2", "x", "readout2"},
	}
	want := []map[string]string{
		{"key1": "condition1", "key2": "readout1"},
		{"key1": "condition2", "key2": "readout2"},
	}
	data := MapLines(lines, headerMap)
	assert.Equal(t, want, data)
}
