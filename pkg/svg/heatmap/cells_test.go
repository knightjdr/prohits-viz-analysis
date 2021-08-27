package heatmap

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write cells", func() {
	It("should write heatmap cells", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		matrix := [][]float64{
			{25, 5, 50.2},
			{100, 30, 7},
			{5, 2.3, 8},
		}
		h := &Heatmap{
			CellSize:   20,
			FillColor:  "blueBlack",
			FillMax:    50,
			FillMin:    0,
			LeftMargin: 50,
			Matrix:     matrix,
			NumColors:  101,
			TopMargin:  50,
		}

		expected := "\t<g id=\"minimap\" transform=\"translate(50, 50)\">\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#ccd9ff\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0033cc\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#b8c9ff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#ccd9ff\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#e6ecff\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#adc2ff\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t</g>\n"
		writeCells(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
