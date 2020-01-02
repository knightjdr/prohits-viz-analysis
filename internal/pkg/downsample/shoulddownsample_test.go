package downsample_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
)

var _ = Describe("Should downsample", func() {
	It("should not downsample a matrix with dimensions <= 1000", func() {
		matrix := make([][]float64, 1000)
		matrix[0] = make([]float64, 1000)
		Expect(Should(matrix, 1000)).To(BeFalse())
	})

	It("should downsample a matrix with rows > 1000", func() {
		matrix := make([][]float64, 1001)
		matrix[0] = make([]float64, 1000)
		Expect(Should(matrix, 1000)).To(BeTrue())
	})

	It("should downsample a matrix with columns > 1000", func() {
		matrix := make([][]float64, 1000)
		matrix[0] = make([]float64, 1001)
		Expect(Should(matrix, 1000)).To(BeTrue())
	})

	It("should downsample a matrix when supplied threshold is zero/not supplied", func() {
		matrix := make([][]float64, 1000)
		matrix[0] = make([]float64, 1001)
		Expect(Should(matrix, 0)).To(BeTrue())
	})
})
