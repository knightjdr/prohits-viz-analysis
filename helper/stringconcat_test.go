package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConcat(t *testing.T) {
	slice := []string{"a", "bc", "d\ne"}

	// TEST1: concatenates strings
	want := "abcd\ne"
	assert.Equal(t, want, StringConcat(slice), "Strings not concatenated")

	// TEST2: empty slice
	want = ""
	assert.Equal(t, want, StringConcat([]string{}), "Empty slice not concatenated")
}
