package heatmap_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

var _ = Describe("Draw", func() {
	It("should draw a heatmap to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		heatmap := Initialize()
		heatmap.AbundanceCap = 50
		heatmap.Annotations = types.Annotations{
			FontSize: 15,
			List: map[string]types.Annotation{
				"a": {
					Position: types.AnnotationPosition{X: 0.5, Y: 0.2},
					Text:     "a",
				},
			},
		}
		heatmap.CellSize = 20
		heatmap.Columns = []string{"bait1", "bait2", "bait3"}
		heatmap.LeftMargin = 50
		heatmap.Markers = types.Markers{
			Color: "#000000",
			List: map[string]types.Marker{
				"a": {Height: 2, Width: 2, X: 0, Y: 1},
			},
		}
		heatmap.Matrix = [][]float64{
			{25, 5, 50.2},
			{100, 30, 7},
			{5, 2.3, 8},
		}
		heatmap.PlotHeight = 200
		heatmap.PlotWidth = 100
		heatmap.Rows = []string{"prey1", "prey2", "prey3"}
		heatmap.SvgHeight = 250
		heatmap.SvgWidth = 150
		heatmap.TopMargin = 50
		heatmap.XLabel = "Bait"
		heatmap.YLabel = "Prey"

		heatmap.Draw("test/heatmap.svg")

		actual, _ := afero.ReadFile(fs.Instance, "test/heatmap.svg")
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
			"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#ccd9ff\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0033cc\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#b8c9ff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#ccd9ff\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#e6ecff\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#adc2ff\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(50, 50)\">\n" +
			"\t\t<rect y=\"20\" x=\"0\" width=\"40\" height=\"40\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
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
