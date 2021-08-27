package heatmap

import (
	"encoding/base64"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
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

		h := Initialize()
		h.CellSize = 2
		h.FillMax = 50
		h.FillMin = 0
		h.Height = 4
		h.NumColors = 11
		h.Width = 4

		matrix := [][]float64{
			{0, 50},
			{50, 0},
		}
		h.Draw(matrix, "heatmap.png")

		expected := "iVBORw0KGgoAAAANSUhEUgAAAAQAAAAECAIAAAAmkwkpAAAAHElEQVR4nGL6DwYMYMDEgAQYIRREEkUGEAAA//9Y1Aj/wbgG8AAAAABJRU5ErkJggg=="

		pngContent, _ := afero.ReadFile(fs.Instance, "heatmap.png")
		Expect(base64.StdEncoding.EncodeToString((pngContent))).To(Equal(expected))
	})
})
