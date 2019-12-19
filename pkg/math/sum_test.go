package math_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

var _ = Describe("Sum float", func() {
	It("should sum values", func() {
		values := []float64{1.3, 4.6, 7, 3.13}

		expected := 16.03
		Expect(SumFloat(values)).To(BeNumerically("~", expected, 0.001))
	})
})
