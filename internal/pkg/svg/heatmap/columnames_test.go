package heatmap

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write column names", func() {
	It("should write column names", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			CellSize:   20,
			Columns:    []string{"bait1", "bait2", "bait3"},
			FontSize:   12,
			LeftMargin: 50,
			TopMargin:  50,
		}

		expected := "\t<g transform=\"translate(50)\">\n" +
			"\t\t<text y=\"48\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 48)\">bait1</text>\n" +
			"\t\t<text y=\"48\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 48)\">bait2</text>\n" +
			"\t\t<text y=\"48\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 48)\">bait3</text>\n" +
			"\t</g>\n"
		writeColumnNames(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
