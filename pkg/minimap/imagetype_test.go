package minimap

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Image type", func() {
	It("should return 'dotplot' when the image type is requested and the matrices are less than 500x500", func() {
		abundance := make([][]float64, 500)
		abundance[0] = make([]float64, 500)
		data := &Data{
			ImageType: "dotplot",
			Matrices: &types.Matrices{
				Abundance: abundance,
			},
		}

		expected := "dotplot"
		Expect(defineImageType(data)).To(Equal(expected))
	})

	It("should return 'heatmap' when the matrices have > 500 rows", func() {
		abundance := make([][]float64, 501)
		abundance[0] = make([]float64, 500)
		data := &Data{
			ImageType: "dotplot",
			Matrices: &types.Matrices{
				Abundance: abundance,
			},
		}

		expected := "heatmap"
		Expect(defineImageType(data)).To(Equal(expected))
	})

	It("should return 'heatmap' when the matrices have > 500 columns", func() {
		abundance := make([][]float64, 500)
		abundance[0] = make([]float64, 501)
		data := &Data{
			ImageType: "dotplot",
			Matrices: &types.Matrices{
				Abundance: abundance,
			},
		}

		expected := "heatmap"
		Expect(defineImageType(data)).To(Equal(expected))
	})
})
