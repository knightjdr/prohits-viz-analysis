package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicetomap(t *testing.T) {
	data := []string{"a", "b", "c"}

	// TEST1: convert slice to map
	want := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}
	assert.Equal(
		t,
		want,
		Slicetomap(data),
		"Slice is not being mapped correctly",
	)
}
