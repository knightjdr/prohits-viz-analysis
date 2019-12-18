package convert

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Score test", func() {
	It("should return smaller value when a larger score is better", func() {
		tests := []map[string]float64{
			{"expected": 1.4, "value": 1.4},
			{"expected": 5, "value": 5},
			{"expected": 8, "value": 8},
			{"expected": 15, "value": 16},
			{"expected": 15, "value": 27},
		}
		threshold := float64(15)
		findWorseScore := getScoreTest("gte")

		for _, test := range tests {
			Expect(findWorseScore(test["value"], threshold)).To(Equal(test["expected"]))
		}
	})

	It("should return larger value when a smaller score is better", func() {
		tests := []map[string]float64{
			{"expected": 15, "value": 1.4},
			{"expected": 15, "value": 5},
			{"expected": 15, "value": 8},
			{"expected": 16, "value": 16},
			{"expected": 27, "value": 27},
		}
		threshold := float64(15)
		findWorseScore := getScoreTest("lte")

		for _, test := range tests {
			Expect(findWorseScore(test["value"], threshold)).To(Equal(test["expected"]))
		}
	})
})
