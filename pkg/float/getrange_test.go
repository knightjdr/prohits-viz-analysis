package float_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/float"
)

var _ = Describe("Join", func() {
	It("should join a slice of float64", func() {
		mapToOutputRange := GetRange(-1, 1, 0, 100)
		tests := []map[string]float64{
			{"value": 0, "expected": 50},
			{"value": 2.183, "expected": 100},
			{"value": -1, "expected": 0},
			{"value": -10.5, "expected": 0},
			{"value": -0.5, "expected": 25},
			{"value": 0.5, "expected": 75},
		}

		for _, test := range tests {
			Expect(mapToOutputRange(test["value"])).To(Equal(test["expected"]))
		}
	})
})
