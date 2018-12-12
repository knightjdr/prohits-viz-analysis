package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContains(t *testing.T) {
	testSlice := []string{"a", "c", "d"}

	// TEST: slice contains tested values
	shouldContain := []string{"a", "c", "d"}
	for _, value := range shouldContain {
		assert.True(t, SliceContains(value, testSlice), "Slice should contain value but does not")
	}

	// TEST: slice contains tested values
	shouldNotContain := []string{"aa", "b", "something"}
	for _, value := range shouldNotContain {
		assert.False(t, SliceContains(value, testSlice), "Slice should not contain value but does")
	}
}
