package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapKeys(t *testing.T) {
	hash := map[string]float64{
		"a":   1,
		"abc": 2,
		"c":   3,
	}
	expected := []string{"a", "abc", "c"}
	result := MapKeysFloat64(hash)
	assert.Equal(t, expected, result, "Map keys not returned for map of float64")
}
