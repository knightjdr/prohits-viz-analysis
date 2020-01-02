package minimap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Downsample", func() {
	It("should downsample matrix", func() {
		matrix := [][]float64{
			{1, 1, 2, 2, 2},
			{2, 2, 2, 3, 3},
			{1, 1, 4, 1, 1},
			{4, 2, 2, 3, 1},
			{5, 3, 4, 1, 3},
		}
		data := &Data{
			DownsampleThreshold: 2,
			Matrices: &types.Matrices{
				Abundance: matrix,
			},
		}

		expected := [][]float64{
			{1.6, 2.24},
			{3.04, 2.08},
		}
		downsampleIfNeeded(data)
		Expect(data.Matrices.Abundance).To(Equal(expected))
	})

	It("should not downsample matrix", func() {
		matrix := [][]float64{
			{1, 1, 2, 2, 2},
			{2, 2, 2, 3, 3},
			{1, 1, 4, 1, 1},
			{4, 2, 2, 3, 1},
			{5, 3, 4, 1, 3},
		}
		data := &Data{
			DownsampleThreshold: 5,
			Matrices: &types.Matrices{
				Abundance: matrix,
			},
		}

		expected := [][]float64{
			{1, 1, 2, 2, 2},
			{2, 2, 2, 3, 3},
			{1, 1, 4, 1, 1},
			{4, 2, 2, 3, 1},
			{5, 3, 4, 1, 3},
		}
		downsampleIfNeeded(data)
		Expect(data.Matrices.Abundance).To(Equal(expected))
	})
})
