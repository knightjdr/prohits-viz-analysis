package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetRange(t *testing.T) {
	rangeFunc := SetRange(-1, 1, 0, 100)
	tests := [6]map[string]float64{
		{"number": 0, "want": 50},
		{"number": 2.183, "want": 100},
		{"number": -1, "want": 0},
		{"number": -10.5, "want": 0},
		{"number": -0.5, "want": 25},
		{"number": 0.5, "want": 75},
	}
	for _, test := range tests {
		assert.InDelta(t, test["want"], rangeFunc(test["number"]), 0.001, "Input number not mapped to correct output range")
	}
}
