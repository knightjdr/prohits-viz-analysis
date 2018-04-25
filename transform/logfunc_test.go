package transform

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogFunc(t *testing.T) {
	type testValues struct {
		input float64
		want  float64
	}

	// TEST1: log base 2.
	filterFunc := LogFunc("2")
	tests := [4]testValues{
		testValues{0, 0},
		testValues{1, 0},
		testValues{2, 1},
		testValues{4, 2},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, filterFunc(test.input), "Log base 2 is not calculating correctly")
	}

	// TEST2: log base 10.
	filterFunc = LogFunc("10")
	tests = [4]testValues{
		testValues{0, 0},
		testValues{1, 0},
		testValues{10, 1},
		testValues{100, 2},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, filterFunc(test.input), "Log base 10 is not calculating correctly")
	}

	// TEST3: log base e.
	filterFunc = LogFunc("e")
	tests = [4]testValues{
		testValues{0, 0},
		testValues{1, 0},
		testValues{math.E, 1},
		testValues{math.SqrtE, 0.5},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, filterFunc(test.input), "Natural log is not calculating correctly")
	}
}
