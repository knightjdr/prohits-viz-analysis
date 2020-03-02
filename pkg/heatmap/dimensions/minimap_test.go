package dimensions

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Minimap", func() {
	It("should calculate plot and svg dimensions", func() {
		dims := &Heatmap{
			CellSize:     20,
			ColumnNumber: 20,
			RowNumber:    85,
		}

		expected := &Heatmap{
			CellSize:     20,
			ColumnNumber: 20,
			RowNumber:    85,
			PlotHeight:   1700,
			PlotWidth:    400,
			SvgHeight:    1700,
			SvgWidth:     400,
		}

		setMinimapDimensions(dims)
		Expect(dims).To(Equal(expected))
	})
})
