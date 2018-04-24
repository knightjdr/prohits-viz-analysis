package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceToMap(t *testing.T) {
	data := []string{"a", "b", "c"}

	// TEST1: convert slice to map.
	want := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}
	assert.Equal(
		t,
		want,
		SliceToMap(data),
		"Slice is not being mapped correctly",
	)
}
