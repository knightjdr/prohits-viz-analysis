package math_test

import (
	. "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Round", func() {
	It("should round to specified precision", func() {
		tests := [7]map[string]float64{
			{"number": 10.032453, "precision": 0.01, "expected": 10.03},
			{"number": 2.183, "precision": 0.1, "expected": 2.2},
			{"number": -156.789235, "precision": 0.0001, "expected": -156.7892},
			{"number": 10.032453, "precision": 1, "expected": 10},
			{"number": 10.032453, "precision": 0.005, "expected": 10.03},
			{"number": 10.05, "precision": 0.01, "expected": 10.05},
			{"number": 0.7142857142857143, "precision": 0.01, "expected": 0.71},
		}
		for _, test := range tests {
			Expect(Round(test["number"], test["precision"])).To(Equal(test["expected"]), "should round to specified precision")
		}
	})
})
