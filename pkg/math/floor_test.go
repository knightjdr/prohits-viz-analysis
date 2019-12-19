package math_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

var _ = Describe("Floor", func() {
	It("should floor float values", func() {
		tests := []float64{1.3, 4.6, 7, 3.13}

		expected := []int{1, 4, 7, 3}
		for i, test := range tests {
			Expect(Floor(test)).To(Equal(expected[i]))
		}
	})
})
