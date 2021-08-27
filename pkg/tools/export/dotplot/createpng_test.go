package dotplot

import (
	"encoding/base64"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/export/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create PNG", func() {
	It("should create png via svg", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		oldConvert := convertSVG
		convertSVG = func(string, string, string) {}
		defer func() { convertSVG = oldConvert }()

		data := &heatmap.Heatmap{
			Annotations: types.Annotations{
				FontSize: 16,
				List: map[string]types.Annotation{
					"a": {Text: "a", Position: types.AnnotationPosition{X: 0.5, Y: 0.25}},
				},
			},
			Markers: types.Markers{
				Color: "#000000",
				List: map[string]types.Marker{
					"a": {Height: 1, Width: 2, X: 0.25, Y: 0.5},
				},
			},
			Settings: types.Settings{
				EdgeColor:       "blue",
				FillColor:       "blue",
				FillMax:         4,
				FillMin:         0,
				PrimaryFilter:   0.01,
				ScoreType:       "lte",
				SecondaryFilter: 0.05,
				XLabel:          "Bait",
				YLabel:          "Prey",
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
			Ratio: [][]float64{
				{0.5, 0.5, 1, 1},
				{0.67, 0.67, 0.67, 1},
				{0.25, 0.25, 1, 0.25},
				{1, 0.5, 0.5, 0.75},
			},
			Readouts: []string{"prey1", "prey2", "prey3", "prey4"},
			Score: [][]float64{
				{0.05, 0.05, 0.01, 0.01},
				{0.01, 0.01, 0.01, 0},
				{0.1, 0.1, 0, 0.1},
				{0, 0.05, 0.05, 0.01},
			},
		}
		settings := Settings{
			DownsampleThreshold: 5,
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
			"\t\t<circle fill=\"#809fff\" cy=\"10\" cx=\"10\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"10\" cx=\"30\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"10\" cx=\"50\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"10\" cx=\"70\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"30\" cx=\"10\" r=\"5.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"30\" cx=\"30\" r=\"5.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"30\" cx=\"50\" r=\"5.700000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#002080\" cy=\"30\" cx=\"70\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"50\" cx=\"10\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"50\" cx=\"30\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"50\" cx=\"50\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#809fff\" cy=\"50\" cx=\"70\" r=\"2.130000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"70\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"70\" cx=\"30\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#0040ff\" cy=\"70\" cx=\"50\" r=\"4.250000\" stroke=\"#0040ff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#002080\" cy=\"70\" cx=\"70\" r=\"6.380000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t</g>\n" +
			"\t<rect fill=\"none\" y=\"57\" x=\"57\" width=\"80\" height=\"80\" stroke=\"#000000\" stroke-width=\"0.5\" />\n" +
			"\t<g transform=\"translate(57, 57)\">\n" +
			"\t\t<rect y=\"40\" x=\"20\" width=\"40\" height=\"20\" stroke=\"#000000\" stroke-width=\"1\" fill=\"none\"/>\n" +
			"\t</g>\n" +
			"\t<g transform=\"translate(57, 57)\">\n" +
			"\t\t<text y=\"20\" x=\"40\" font-size=\"16\" text-anchor=\"middle\">a</text>\n" +
			"\t</g>\n" +
			"\t<text y=\"10\" x=\"97\" font-size=\"12\" text-anchor=\"middle\">Bait</text>\n" +
			"\t<text y=\"97\" x=\"10\" font-size=\"12\" text-anchor=\"middle\" transform=\"rotate(-90, 10, 97)\">Prey</text>\n" +
			"</svg>\n"

		createPNG(data, matrices, settings)
		svgContent, _ := afero.ReadFile(fs.Instance, "svg/dotplot.svg")
		Expect(string(svgContent)).To(Equal(expected))
	})

	It("should create downsampled png", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := &heatmap.Heatmap{
			Settings: types.Settings{
				EdgeColor:       "blue",
				FillColor:       "blue",
				FillMax:         4,
				FillMin:         0,
				PrimaryFilter:   0.01,
				ScoreType:       "lte",
				SecondaryFilter: 0.05,
				XLabel:          "Bait",
				YLabel:          "Prey",
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
			Ratio: [][]float64{
				{0.5, 0.5, 1, 1},
				{0.67, 0.67, 0.67, 1},
				{0.25, 0.25, 1, 0.25},
				{1, 0.5, 0.5, 0.75},
			},
			Readouts: []string{"prey1", "prey2", "prey3", "prey4"},
			Score: [][]float64{
				{0.05, 0.05, 0.01, 0.01},
				{0.01, 0.01, 0.01, 0},
				{0.1, 0.1, 0, 0.1},
				{0, 0.05, 0.05, 0.01},
			},
		}
		settings := Settings{
			DownsampleThreshold: 2,
		}

		expected := "iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAIAAAADnC86AAAAU0lEQVR4nOzWsQmAYAyE0Siu" +
			"JTiRI9rYuYUTOILucMLf5F3/8SBVlnV/K9153XE7x+XPgcFgMBgMBo/fVFv+c9VzxGm/U4PBYDAYDG4MfwEA" +
			"AP//OmgHVhr0NIsAAAAASUVORK5CYII="

		createPNG(data, matrices, settings)
		pngContent, _ := afero.ReadFile(fs.Instance, "png/dotplot.png")
		Expect(base64.StdEncoding.EncodeToString(pngContent)).To(Equal(expected))
	})
})
