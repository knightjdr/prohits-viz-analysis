package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncateFloat(t *testing.T) {
	// TEST1: rounding.
	tests := [7]map[string]float64{
		{"number": 10.032453, "precision": 2, "want": 10.03},
		{"number": 2.183, "precision": 1, "want": 2.2},
		{"number": -156.789235, "precision": 4, "want": -156.7892},
		{"number": 10.032453, "precision": 2, "want": 10.03},
		{"number": 10.05, "precision": 2, "want": 10.05},
		{"number": 0.7142857142857143, "precision": 2, "want": 0.71},
	}
	for _, test := range tests {
		assert.Equal(t, test["want"], TruncateFloat(test["number"], int(test["precision"])), "Rounding error")
	}
}
