package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFts(t *testing.T) {
	slice := []float64{0.62, 0.21, 4.4554}

	// TEST1: no precision
	want := []string{"0.62", "0.21", "4.4554"}
	assert.Equal(t, want, ConvertFts(slice, -1), "Slice not converted to strings")

	// TEST2: with precision
	want = []string{"0.62", "0.21", "4.46"}
	assert.Equal(t, want, ConvertFts(slice, 2), "Slice not converted to strings with precision")
}
