package scatter

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write axes", func() {
	It("should write axes", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		axisLength := 100.00
		s := &Scatter{
			Axes: Axes{
				X: Line{X1: 0, X2: 100, Y1: 100, Y2: 100},
				Y: Line{X1: 0, X2: 0, Y1: 100, Y2: 0},
			},
			PlotSize: 250,
			Ticks: Ticks{
				X:      []float64{0, 50, 100},
				XLabel: []string{"0", "10", "20"},
				Y:      []float64{0, 33.33, 66.67, 100},
				YLabel: []string{"0", "10", "20", "30"},
			},
			XLabel: "x-axis",
			YLabel: "y-axis",
		}

		expected := "\t\t<g>\n" +
			"\t\t\t<g transform=\"translate(100 0)\">\n" +
			"\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"100\" y1=\"100\" y2=\"100\"/>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"50\" x2=\"50\" y1=\"100\" y2=\"110\" />\n" +
			"\t\t\t\t\t<text font-size=\"12\" text-anchor=\"middle\" x=\"50\" y=\"128\">10</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"100\" x2=\"100\" y1=\"100\" y2=\"110\" />\n" +
			"\t\t\t\t\t<text font-size=\"12\" text-anchor=\"middle\" x=\"100\" y=\"128\">20</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t\t<g transform=\"translate(0 100)\">\n" +
			"\t\t\t\t\t<text text-anchor=\"middle\" x=\"50\" y=\"70\">x-axis</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t\t<g transform=\"translate(100 0)\">\n" +
			"\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"0\" y1=\"100\" y2=\"0\"/>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"-10\" y1=\"66.67\" y2=\"66.67\" />\n" +
			"\t\t\t\t\t<text dy=\"4\" font-size=\"12\" text-anchor=\"end\" x=\"-13\" y=\"66.67\">10</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"-10\" y1=\"33.33\" y2=\"33.33\" />\n" +
			"\t\t\t\t\t<text dy=\"4\" font-size=\"12\" text-anchor=\"end\" x=\"-13\" y=\"33.33\">20</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t\t<g>\n" +
			"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"-10\" y1=\"0\" y2=\"0\" />\n" +
			"\t\t\t\t\t<text dy=\"4\" font-size=\"12\" text-anchor=\"end\" x=\"-13\" y=\"0\">30</text>\n" +
			"\t\t\t\t</g>\n" +
			"\t\t\t</g>\n" +
			"\t\t\t<text dy=\"15\" text-anchor=\"middle\" transform=\"rotate(-90, 0, 50)\" x=\"0\" y=\"50\">y-axis</text>\n" +
			"\t\t</g>\n"
		writeAxes(s, axisLength, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
