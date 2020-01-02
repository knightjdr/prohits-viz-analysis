package minimap

import (
	"encoding/base64"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
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
		dims := &dimensions.Heatmap{
			CellSize:   2,
			PlotHeight: 4,
			PlotWidth:  4,
		}

		createHeatmap(data, dims)

		expected := "iVBORw0KGgoAAAANSUhEUgAAAAQAAAAECAIAAAAmkwkpAAAAHElEQVR4nGL6DwYMYMDEgAQYIRREEkUGEAAA//9Y1Aj/wbgG8AAAAABJRU5ErkJggg=="

		pngContent, _ := afero.ReadFile(fs.Instance, "test/heatmap.png")
		actual := base64.StdEncoding.EncodeToString(pngContent)
		Expect(actual).To(Equal(expected))
	})
})
