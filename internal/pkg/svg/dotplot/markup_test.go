package dotplot

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write markup", func() {
	It("should write all markup", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		h := &Dotplot{
			Annotations: types.Annotations{
				FontSize: 15,
				List: map[string]types.Annotation{
					"a": {
						Position: types.AnnotationPosition{X: 0.5, Y: 0.2},
						Text:     "a",
					},
				},
			},
			CellSize:   20,
			LeftMargin: 50,
			Markers: types.Markers{
				Color: "#000000",
				List: map[string]types.Marker{
					"a": {Height: 2, Width: 2, X: 0, Y: 1},
				},
			},
			PlotHeight: 200,
			PlotWidth:  100,
			SvgHeight:  250,
			SvgWidth:   150,
			TopMargin:  50,
			XLabel:     "Bait",
			YLabel:     "Prey",
		}

		expected := "\t<g transform=\"translate(50, 50)\">\n" +
			"\t\t<rect y=\"20\" x=\"0\" width=\"40\" height=\"40\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(50, 50)\">\n" +
			"\t\t<text y=\"40\" x=\"50\" font-size=\"15\" text-anchor=\"middle\">a</text>\n" +
			"\t</g>\n" +
			"\t<text y=\"10\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n" +
			"\t<text y=\"150\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 150)\">Prey</text>\n"
		writeMarkup(h, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
