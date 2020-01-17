package math_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

var _ = Describe("Sign", func() {
	It("should return the sign of values", func() {
		tests := []float64{1.3, -4.6, 7, 3.13, 0}

		expected := []float64{1, -1, 1, 1, 0}
		for i, test := range tests {
			Expect(Sign(test)).To(Equal(expected[i]))
		}
	})
})
