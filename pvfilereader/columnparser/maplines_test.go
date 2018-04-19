package columnparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MappedTerms struct {
	key1 string
	key2 string
}

func TestMaplines(t *testing.T) {
	// values are column indexes to map to the key names
	headerMap := map[string]int{
		"key1": 0,
		"key2": 2,
	}
	// array of lines from a csv file
	lines := [][]string{
		{"bait1", "x", "prey1"},
		{"bait2", "x", "prey2"},
	}
	want := []map[string]string{
		{"key1": "bait1", "key2": "prey1"},
		{"key1": "bait2", "key2": "prey2"},
	}
	data := Maplines(lines, headerMap)
	assert.Equal(t, want, data)
}
