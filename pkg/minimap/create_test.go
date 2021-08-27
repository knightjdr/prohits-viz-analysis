package minimap

import (
	"encoding/base64"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create minimap", func() {
	It("should draw a dotplot to a file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		oldConvert := convertSVG
		convertSVG = func(string, string, string) {}
		defer func() { convertSVG = oldConvert }()

		data := &Data{
			Filename:  "test/dotplot.png",
			ImageType: "dotplot",
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

		Create(data)

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

	It("should draw a heatmap to a file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		data := &Data{
			Filename:  "test/heatmap.png",
			ImageType: "heatmap",
			Matrices: &types.Matrices{
				Abundance: [][]float64{
					{0, 50},
					{50, 0},
				},
			},
			Settings: types.Settings{
				FillColor:    "blue",
				FillMax:      50,
				InvertColor:  false,
				MinAbundance: 0,
			},
		}

		Create(data)

		expected := "iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAIAAAADnC86AAAAR0lEQVR4nOzWsQnAQAxDUTlk/5WdHXSQ5t7vzQNXenc3bTNT3z715WFgMBgMBoPB/9ePpiQne+2+V4PBYDAYDL4Y/gIAAP//cy8GT45ahrIAAAAASUVORK5CYII="

		pngContent, _ := afero.ReadFile(fs.Instance, "test/heatmap.png")
		actual := base64.StdEncoding.EncodeToString(pngContent)
		Expect(actual).To(Equal(expected))
	})
})
