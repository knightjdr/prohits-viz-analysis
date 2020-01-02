package heatmap

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write row names", func() {
	It("should write row names", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Heatmap{
			CellSize:   20,
			FontSize:   12,
			LeftMargin: 50,
			Rows:       []string{"prey1", "prey2", "prey3"},
			TopMargin:  50,
		}

		expected := "\t<g transform=\"translate(0, 50)\">\n" +
			"\t\t<text y=\"15\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
			"\t\t<text y=\"35\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
			"\t\t<text y=\"55\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
			"\t</g>\n"
		writeRowNames(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
