package downsample

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scale", func() {
	It("should scale dimensions based on column length", func() {
		matrix := make([][]float64, 600)
		matrix[0] = make([]float64, 3000)

		Expect(calculateScale(matrix, 1000)).To(Equal(float64(3)))
	})

	It("should scale dimensions based on row length", func() {
		matrix := make([][]float64, 2000)
		matrix[0] = make([]float64, 500)

		Expect(calculateScale(matrix, 1000)).To(Equal(float64(2)))
	})
})
