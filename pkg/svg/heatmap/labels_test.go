package heatmap_test

import (
	"strings"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/svg/heatmap"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write column and row labels", func() {
	It("should write column and row names", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			CellSize:   20,
			Columns:    []string{"bait1", "bait2", "bait3"},
			FontSize:   12,
			LeftMargin: 50,
			Rows:       []string{"prey1", "prey2", "prey3"},
			TopMargin:  50,
		}

		expected := "\t<g transform=\"translate(50)\">\n" +
			"\t\t<text y=\"48\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 48)\">bait1</text>\n" +
			"\t\t<text y=\"48\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 48)\">bait2</text>\n" +
			"\t\t<text y=\"48\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 48)\">bait3</text>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(0, 50)\">\n" +
			"\t\t<text y=\"15\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
			"\t\t<text y=\"35\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
			"\t\t<text y=\"55\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
			"\t</g>\n"
		WriteLabels(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
