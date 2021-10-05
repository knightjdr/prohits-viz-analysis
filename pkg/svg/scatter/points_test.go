package scatter

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write points", func() {
	It("should write points", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		axisLength := 100.00
		s := &Scatter{
			Plot: []types.ScatterPoint{
				{Color: "#ff0000", Label: "point1", X: 10, Y: 75},
				{Color: "#00ff00", Label: "point2", X: 63, Y: 12},
			},
		}

		expected := "\t\t<g transform=\"translate(100 0)\">\n" +
			"\t\t\t<defs>\n" +
			"\t\t\t\t<clipPath id=\"plot_points_clip\">\n" +
			"\t\t\t\t\t<rect height=\"100\" width=\"100\" x=\"0\" y=\"0\" />\n" +
			"\t\t\t\t</clipPath>\n" +
			"\t\t\t</defs>\n" +
			"\t\t\t<g clip-path=\"url(#plot_points_clip)\">\n" +
			"\t\t\t\t<circle cx=\"10\" cy=\"25\" fill=\"#ff0000\" r=\"4\"><title>point1</title></circle>\n" +
			"\t\t\t\t<circle cx=\"63\" cy=\"88\" fill=\"#00ff00\" r=\"4\"><title>point2</title></circle>\n" +
			"\t\t\t</g>\n" +
			"\t\t</g>\n"
		writePoints(s, axisLength, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
