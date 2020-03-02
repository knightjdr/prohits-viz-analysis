package minimap

import (
	"encoding/base64"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Draw heatmap", func() {
	It("should draw a heatmap to a file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		data := &Data{
			Filename: "test/heatmap.png",
			Matrices: &types.Matrices{
				Abundance: [][]float64{
					{0, 50},
					{50, 0},
				},
			},
			Settings: types.Settings{
				AbundanceCap: 50,
				FillColor:    "blue",
				InvertColor:  false,
				MinAbundance: 0,
			},
		}

		createHeatmap(data)

		expected := "iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAIAAAADnC86AAAAR0lEQVR4nOzWsQnAQAxDUTlk/5WdHXSQ5t7vzQNXenc3bTNT3z715WFgMBgMBoPB/9ePpiQne+2+V4PBYDAYDL4Y/gIAAP//cy8GT45ahrIAAAAASUVORK5CYII="

		pngContent, _ := afero.ReadFile(fs.Instance, "test/heatmap.png")
		actual := base64.StdEncoding.EncodeToString(pngContent)
		Expect(actual).To(Equal(expected))
	})
})
