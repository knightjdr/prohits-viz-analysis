package dimensions

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cellsize", func() {
	It("should calculate cellsize from ratio", func() {
		dims := &Heatmap{
			ColumnNumber: 120,
			Ratio:        0.9375,
			RowNumber:    150,
		}

		expected := 18

		calculateCellSize(dims)
		Expect(dims.CellSize).To(Equal(expected))
	})
})
