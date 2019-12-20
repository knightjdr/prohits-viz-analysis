package heatmap

import (
	"bytes"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
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
		h.AbundanceCap = 50
		h.CellSize = 2
		h.Height = 4
		h.MinAbundance = 0
		h.NumColors = 11
		h.Width = 4

		matrix := [][]float64{
			{0, 50},
			{50, 0},
		}
		h.Draw(matrix, "heatmap.png")

		expected := []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 4, 0, 0, 0, 4, 8, 2, 0, 0, 0, 38, 147, 9, 41, 0, 0, 0, 28, 73, 68, 65, 84, 120, 156, 98, 250, 15, 6, 12, 96, 192, 196, 128, 4, 24, 33, 20, 68, 18, 69, 6, 16, 0, 0, 255, 255, 88, 212, 8, 255, 193, 184, 6, 240, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}

		pngContent, _ := afero.ReadFile(fs.Instance, "heatmap.png")
		Expect(bytes.Compare(expected, pngContent)).To(Equal(0))
	})
})

var _ = Describe("Get gradient index", func() {
	It("should return a function for computing the slice index", func() {
		h := &Heatmap{
			AbundanceCap: 50,
			MinAbundance: 0,
			NumColors:    11,
		}

		testFunc := getGradientIndex(h)
		tests := []float64{-1, 60, 25, 10}

		expected := []int{0, 10, 5, 2}
		for i, test := range tests {
			Expect(testFunc(test)).To(Equal(expected[i]))
		}
	})
})
