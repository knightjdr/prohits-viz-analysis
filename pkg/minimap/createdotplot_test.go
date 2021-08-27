package minimap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
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
		convertSVG = func(string, string, string) {}
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
				EdgeColor:       "blue",
				FillColor:       "blue",
				FillMax:         50,
				InvertColor:     false,
				MinAbundance:    0,
				PrimaryFilter:   0.01,
				ScoreType:       "lte",
				SecondaryFilter: 0.05,
			},
		}

		createDotplot(data)

		expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
			" xml:space=\"preserve\" width=\"40\" height=\"40\" viewBox=\"0 0 40 40\">\n" +
			"\t<g id=\"minimap\" transform=\"translate(0, 0)\">\n" +
			"\t\t<circle fill=\"#ffffff\" cy=\"10\" cx=\"10\" r=\"0.000000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"10\" cx=\"30\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#000000\" cy=\"30\" cx=\"10\" r=\"8.500000\" stroke=\"#000000\" stroke-width=\"2.000000\"/>\n" +
			"\t\t<circle fill=\"#ffffff\" cy=\"30\" cx=\"30\" r=\"0.000000\" stroke=\"#809fff\" stroke-width=\"2.000000\"/>\n" +
			"\t</g>\n" +
			"</svg>\n"

		actual, _ := afero.ReadFile(fs.Instance, "test/dotplot.svg")
		Expect(string(actual)).To(Equal(expected))
	})
})
