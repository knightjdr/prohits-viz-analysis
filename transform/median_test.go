package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedian(t *testing.T) {
	// TEST1: calculate various medians for MedianInt.
	type testInt struct {
		slice []int
		want  float64
	}
	intTests := [4]testInt{
		testInt{[]int{3, 6, 10}, float64(6)},
		testInt{[]int{10, 3, 6}, float64(6)},
		testInt{[]int{3, 6, 7, 10}, 6.5},
		testInt{[]int{10, 3, 7, 6}, 6.5},
	}
	for _, test := range intTests {
		assert.Equal(
			t,
			test.want,
			MedianInt(test.slice),
			"MedianInt is not calculating correct value",
		)
	}
	// TEST2: calculate various medians for MedianFloat.
	type testFloat struct {
		slice []float64
		want  float64
	}
	floatTests := [4]testFloat{
		testFloat{[]float64{2.5, 5, 10.5}, float64(5)},
		testFloat{[]float64{10.5, 2.5, 5}, float64(5)},
		testFloat{[]float64{3.5, 6, 7, 10}, 6.5},
		testFloat{[]float64{10, 7, 3.5, 6}, 6.5},
	}
	for _, test := range floatTests {
		assert.Equal(
			t,
			test.want,
			MedianFloat(test.slice),
			"MedianFloat is not calculating correct value",
		)
	}
}
