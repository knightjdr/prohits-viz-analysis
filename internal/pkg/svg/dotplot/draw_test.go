package dotplot_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

var _ = Describe("Draw a dotplot", func() {
	It("should draw a dotplot to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		dotplot := Initialize()
		dotplot.AbundanceCap = 50
		dotplot.Annotations = types.Annotations{
			FontSize: 15,
			List: map[string]types.Annotation{
				"a": {
					Position: types.AnnotationPosition{X: 0.5, Y: 0.2},
					Text:     "a",
				},
			},
		}
		dotplot.BoundingBox = true
		dotplot.CellSize = 20
		dotplot.Columns = []string{"bait1", "bait2", "bait3"}
		dotplot.LeftMargin = 50
		dotplot.Markers = types.Markers{
			Color: "#000000",
			List: map[string]types.Marker{
				"a": {Height: 2, Width: 2, X: 0, Y: 0.5},
			},
		}
		dotplot.Matrices = &types.Matrices{
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
		}
		dotplot.PlotHeight = 200
		dotplot.PlotWidth = 100
		dotplot.PrimaryFilter = 0.01
		dotplot.Ratio = 1
		dotplot.Rows = []string{"prey1", "prey2", "prey3"}
		dotplot.SecondaryFilter = 0.05
		dotplot.SvgHeight = 250
		dotplot.SvgWidth = 150
		dotplot.TopMargin = 50
		dotplot.XLabel = "Bait"
		dotplot.YLabel = "Prey"

		dotplot.Draw("test/dotplot.svg")

		actual, _ := afero.ReadFile(fs.Instance, "test/dotplot.svg")
		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
			" xml:space=\"preserve\" width=\"150\" height=\"250\" viewBox=\"0 0 150 250\">\n" +
			"\t<g transform=\"translate(50)\">\n" +
			"\t\t<text y=\"48\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 48)\">bait1</text>\n" +
			"\t\t<text y=\"48\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 48)\">bait2</text>\n" +
			"\t\t<text y=\"48\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 48)\">bait3</text>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(0, 50)\">\n" +
			"\t\t<text y=\"15\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
			"\t\t<text y=\"35\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
			"\t\t<text y=\"55\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
			"\t</g>\n" +
			"\t<g id=\"minimap\" transform=\"translate(50, 50)\">\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"10\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#ccd9ff\" cy=\"10\" cx=\"30\" r=\"4.250000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"10\" cx=\"50\" r=\"2.550000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"30\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0033cc\" cy=\"30\" cx=\"30\" r=\"2.550000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#b8c9ff\" cy=\"30\" cx=\"50\" r=\"0.850000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#ccd9ff\" cy=\"50\" cx=\"10\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#e6ecff\" cy=\"50\" cx=\"30\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#adc2ff\" cy=\"50\" cx=\"50\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t</g>\n" +
			"\t<rect fill=\"none\" y=\"50\" x=\"50\" width=\"100\" height=\"200\"" +
			" stroke=\"#000000\" stroke-width=\"0.5\" />\n" +
			"\t<g transform=\"translate(50, 50)\">\n" +
			"\t\t<rect y=\"100\" x=\"0\" width=\"40\" height=\"40\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(50, 50)\">\n" +
			"\t\t<text y=\"40\" x=\"50\" font-size=\"15\" text-anchor=\"middle\">a</text>\n" +
			"\t</g>\n" +
			"\t<text y=\"10\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n" +
			"\t<text y=\"150\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 150)\">Prey</text>\n" +
			"</svg>\n"
		Expect(string(actual)).To(Equal(expected))
	})
})
