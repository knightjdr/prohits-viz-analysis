package dimensions

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Margins and size", func() {
	It("should add plot margins and calculate plot size", func() {
		dims := &Heatmap{
			CellSize:     10,
			ColumnNumber: 10,
			Ratio:        0.5,
			RowNumber:    20,
		}
		columns := []string{"a", "aaa", "aa"}
		rows := []string{"aa", "aaaa", "a"}

		expected := &Heatmap{
			CellSize:     10,
			ColumnNumber: 10,
			FontSize:     6,
			LeftMargin:   23,
			PlotHeight:   200,
			PlotWidth:    100,
			Ratio:        0.5,
			RowNumber:    20,
			SvgHeight:    217,
			SvgWidth:     123,
			TopMargin:    17,
		}

		calculateMarginsAndSize(dims, columns, rows)
		Expect(dims).To(Equal(expected))
	})
})
