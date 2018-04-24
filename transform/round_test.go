package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	// TEST1: rounding.
	tests := [7]map[string]float64{
		{"number": 10.032453, "precision": 0.01, "want": 10.03},
		{"number": 2.183, "precision": 0.1, "want": 2.2},
		{"number": -156.789235, "precision": 0.0001, "want": -156.7892},
		{"number": 10.032453, "precision": 1, "want": 10},
		{"number": 10.032453, "precision": 0.005, "want": 10.03},
		{"number": 10.05, "precision": 0.01, "want": 10.05},
		{"number": 0.7142857142857143, "precision": 0.01, "want": 0.71},
	}
	for _, test := range tests {
		assert.Equal(t, test["want"], Round(test["number"], test["precision"]), "Rounding error")
	}
}
