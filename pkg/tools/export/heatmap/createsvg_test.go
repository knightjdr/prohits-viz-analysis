package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create SVG", func() {
	It("should create svg", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := &Heatmap{
			Annotations: types.Annotations{
				FontSize: 16,
				List: map[string]types.Annotation{
					"a": {Text: "a", Position: types.AnnotationPosition{X: 0.5, Y: 0.25}},
				},
			},
			Markers: types.Markers{
				Color: "#000000",
				List: map[string]types.Marker{
					"a": types.Marker{Height: 1, Width: 2, X: 0.25, Y: 0.5},
				},
			},
			Settings: types.Settings{
				AbundanceCap: 4,
				FillColor:    "blue",
				MinAbundance: 0,
				XLabel:       "Bait",
				YLabel:       "Prey",
			},
		}
		matrices := &types.Matrices{
			Abundance: [][]float64{
				{1, 1, 2, 2},
				{2, 2, 2, 3},
				{1, 1, 4, 1},
				{4, 2, 2, 3},
			},
			Conditions: []string{"bait1", "bait2", "bait3", "bait4"},
			Readouts:   []string{"prey1", "prey2", "prey3", "prey4"},
		}

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"137\" height=\"137\" viewBox=\"0 0 137 137\">\n" +
			"\t<g transform=\"translate(57)\">\n" +
			"\t\t<text y=\"55\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 55)\">bait1</text>\n" +
			"\t\t<text y=\"55\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 55)\">bait2</text>\n" +
			"\t\t<text y=\"55\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 55)\">bait3</text>\n" +
			"\t\t<text y=\"55\" x=\"66\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 66, 55)\">bait4</text>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(0, 57)\">\n" +
			"\t\t<text y=\"15\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
			"\t\t<text y=\"35\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
			"\t\t<text y=\"55\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
			"\t\t<text y=\"75\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey4</text>\n" +
			"\t</g>\n" +
			"\t<g id=\"minimap\" transform=\"translate(57, 57)\">\n" +
			"\t\t<rect fill=\"#809fff\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#809fff\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"60\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#002080\" y=\"20\" x=\"60\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#809fff\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#809fff\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#809fff\" y=\"40\" x=\"60\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#000000\" y=\"60\" x=\"0\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"60\" x=\"20\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#0040ff\" y=\"60\" x=\"40\" width=\"20\" height=\"20\" />\n" +
			"\t\t<rect fill=\"#002080\" y=\"60\" x=\"60\" width=\"20\" height=\"20\" />\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(57, 57)\">\n" +
			"\t\t<rect y=\"40\" x=\"20\" width=\"40\" height=\"20\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(57, 57)\">\n" +
			"\t\t<text y=\"20\" x=\"40\" font-size=\"16\" text-anchor=\"middle\">a</text>\n" +
			"\t</g>\n" +
			"\t<text y=\"10\" x=\"97\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n" +
			"\t<text y=\"97\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 97)\">Prey</text>\n" +
			"</svg>\n"

		createSVG(data, matrices)
		svgContent, _ := afero.ReadFile(fs.Instance, "svg/heatmap.svg")
		Expect(string(svgContent)).To(Equal(expected))
	})
})
