package dotplot

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dots", func() {
	It("should draw dots", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		d := &Dotplot{
			AbundanceCap: 50,
			CellSize:     20,
			EdgeColor:    "blue",
			FillColor:    "blue",
			LeftMargin:   50,
			Matrices: &types.Matrices{
				Abundance: [][]float64{
					{25, 5, 50.2},
					{100, 30, 7},
					{5, 2.3, 8},
				},
				Ratio: [][]float64{
					{1, 0.5, 0.3},
					{1, 0.3, 0.1},
					{0.5, 0.25, 1},
				},
				Score: [][]float64{
					{0.01, 0, 0.02},
					{0, 0.01, 0.01},
					{0.02, 0.1, 0.01},
				},
			},
			NumColors:       101,
			PrimaryFilter:   0.01,
			Ratio:           1,
			ScoreType:       "lte",
			SecondaryFilter: 0.05,
			TopMargin:       50,
		}

		expected := "\t<g id=\"minimap\" transform=\"translate(50, 50)\">\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"10\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#ccd9ff\" cy=\"10\" cx=\"30\" r=\"4.250000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"10\" cx=\"50\" r=\"2.550000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"30\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0033cc\" cy=\"30\" cx=\"30\" r=\"2.550000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#b8c9ff\" cy=\"30\" cx=\"50\" r=\"0.850000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#ccd9ff\" cy=\"50\" cx=\"10\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#e6ecff\" cy=\"50\" cx=\"30\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#adc2ff\" cy=\"50\" cx=\"50\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t</g>\n"
		writeDots(d, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
