package float_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/float"
)

var _ = Describe("Remove trailing zeros", func() {
	It("should remove trailing zeros from float", func() {
		tests := []float64{1, 1.0, 1.0000, 1.1000, 1.00001}

		expected := []string{"1", "1", "1", "1.1", "1.00001"}

		for i, test := range tests {
			Expect(RemoveTrailingZeros(test)).To(Equal(expected[i]))
		}
	})
})
