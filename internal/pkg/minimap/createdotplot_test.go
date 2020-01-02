package minimap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Draw dotplot", func() {
	It("should draw a dotplot to a file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		oldConvert := convertSVG
		convertSVG = func(string, string) {}
		defer func() { convertSVG = oldConvert }()

		data := &Data{
			Filename: "test/dotplot.png",
			Matrices: &types.Matrices{
				Abundance: [][]float64{
					{0, 50},
					{50, 0},
				},
				Ratio: [][]float64{
					{0, 1},
					{1, 0},
				},
				Score: [][]float64{
					{1, 0},
					{0, 1},
				},
			},
			Settings: types.Settings{
				AbundanceCap:    50,
				EdgeColor:       "blue",
				FillColor:       "blue",
				InvertColor:     false,
				MinAbundance:    0,
				PrimaryFilter:   0.01,
				ScoreType:       "lte",
				SecondaryFilter: 0.05,
			},
		}
		dims := &dimensions.Heatmap{
			CellSize:   2,
			PlotHeight: 4,
			PlotWidth:  4,
			Ratio:      0.05,
			SvgHeight:  4,
			SvgWidth:   4,
		}

		createDotplot(data, dims)

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
			" xml:space=\"preserve\" width=\"4\" height=\"4\" viewBox=\"0 0 4 4\">\n" +
			"\t<g id=\"minimap\" transform=\"translate(0, 0)\">\n" +
			"\t\t<circle fill=\"#ffffff\" cy=\"1\" cx=\"1\" r=\"0.000000\" stroke=\"#809fff\" stroke-width=\"0.100000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"1\" cx=\"3\" r=\"0.930000\" stroke=\"#000000\" stroke-width=\"0.100000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"3\" cx=\"1\" r=\"0.930000\" stroke=\"#000000\" stroke-width=\"0.100000\"/>\n" +
			"\t\t<circle fill=\"#ffffff\" cy=\"3\" cx=\"3\" r=\"0.000000\" stroke=\"#809fff\" stroke-width=\"0.100000\"/>\n" +
			"\t</g>\n" +
			"</svg>\n"

		actual, _ := afero.ReadFile(fs.Instance, "test/dotplot.svg")
		Expect(string(actual)).To(Equal(expected))
	})
})
