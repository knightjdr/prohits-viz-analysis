package circheatmap

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write header", func() {
	It("should write svg tag and dimensions", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		c := &CircHeatmap{
			Dimensions: CircHeatmapDimensions{
				Center:   500,
				PlotSize: 1000,
			},
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
			" xml:space=\"preserve\" width=\"1000\" height=\"1000\" viewBox=\"-500 -500 1000 1000\">\n"
		writeHeader(c, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
