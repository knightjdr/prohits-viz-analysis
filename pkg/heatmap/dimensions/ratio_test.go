package dimensions

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ratio", func() {
	It("should calculate ratio when number of rows and columns is <= max", func() {
		dims := &Heatmap{}
		matrix := make([][]float64, 150)
		matrix[0] = make([]float64, 112)

		expected := &Heatmap{
			ColumnNumber: 112,
			Ratio:        1,
			RowNumber:    150,
		}

		calculateRatio(dims, matrix)
		Expect(dims).To(Equal(expected))
	})

	It("should calculate ratio when number of rows is > max", func() {
		dims := &Heatmap{}
		matrix := make([][]float64, 200)
		matrix[0] = make([]float64, 112)

		expected := &Heatmap{
			ColumnNumber: 112,
			Ratio:        0.75,
			RowNumber:    200,
		}

		calculateRatio(dims, matrix)
		Expect(dims).To(Equal(expected))
	})

	It("should calculate ratio when number of columns is > max", func() {
		dims := &Heatmap{}
		matrix := make([][]float64, 150)
		matrix[0] = make([]float64, 120)

		expected := &Heatmap{
			ColumnNumber: 120,
			Ratio:        0.9375,
			RowNumber:    150,
		}

		calculateRatio(dims, matrix)
		Expect(dims).To(Equal(expected))
	})

	It("should set ratio to the min when the calculated ratio would be too small", func() {
		dims := &Heatmap{}
		matrix := make([][]float64, 3001)
		matrix[0] = make([]float64, 112)

		expected := &Heatmap{
			ColumnNumber: 112,
			Ratio:        0.05,
			RowNumber:    3001,
		}

		calculateRatio(dims, matrix)
		Expect(dims).To(Equal(expected))
	})
})
