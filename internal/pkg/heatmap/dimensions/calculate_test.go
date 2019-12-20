package dimensions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
)

var _ = Describe("Calculate heatmap dimensions", func() {
	It("should calculate dimensions for minimap", func() {
		matrix := [][]float64{
			{25, 5, 50.2},
			{100, 30, 7},
			{5, 2.3, 8},
		}
		columns := []string{"bait1", "bait2", "bait3"}
		rows := []string{"prey1", "prey2", "prey3"}

		expected := &Heatmap{
			CellSize:     20,
			ColumnNumber: 3,
			FontSize:     0,
			LeftMargin:   0,
			PlotHeight:   60,
			PlotWidth:    60,
			Ratio:        1,
			RowNumber:    3,
			SvgHeight:    60,
			SvgWidth:     60,
			TopMargin:    0,
		}
		Expect(Calculate(matrix, columns, rows, true)).To(Equal(expected))
	})

	It("should calculate dimensions for image with column and row names", func() {
		matrix := [][]float64{
			{25, 5, 50.2},
			{100, 30, 7},
			{5, 2.3, 8},
		}
		columns := []string{"bait1", "bait2", "bait3"}
		rows := []string{"prey1", "prey2", "prey3"}

		expected := &Heatmap{
			CellSize:     20,
			ColumnNumber: 3,
			FontSize:     12,
			LeftMargin:   57,
			PlotHeight:   60,
			PlotWidth:    60,
			Ratio:        1,
			RowNumber:    3,
			SvgHeight:    117,
			SvgWidth:     117,
			TopMargin:    57,
		}
		Expect(Calculate(matrix, columns, rows, false)).To(Equal(expected))
	})
})
