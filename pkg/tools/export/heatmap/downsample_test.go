package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Downsample", func() {
	It("should not downsample", func() {
		data := &Heatmap{
			Annotations: types.Annotations{
				FontSize: 16,
				List: map[string]types.Annotation{
					"a": {Text: "a", Position: types.AnnotationPosition{X: 0.5, Y: 0.25}},
					"b": {Text: "b", Position: types.AnnotationPosition{X: 0.5, Y: 0.75}},
				},
			},
			Markers: types.Markers{
				Color: "#000000",
				List: map[string]types.Marker{
					"a": types.Marker{Height: 3, Width: 4, X: 0.25, Y: 0.5},
				},
			},
		}
		downsampleTreshold := 4
		matrices := &types.Matrices{
			Abundance: [][]float64{
				{1, 1, 2, 2},
				{2, 2, 2, 3},
				{1, 1, 4, 1},
				{4, 2, 2, 3},
			},
		}

		expectedData := &Heatmap{
			Annotations: types.Annotations{
				FontSize: 16,
				List: map[string]types.Annotation{
					"a": {Text: "a", Position: types.AnnotationPosition{X: 0.5, Y: 0.25}},
					"b": {Text: "b", Position: types.AnnotationPosition{X: 0.5, Y: 0.75}},
				},
			},
			Markers: types.Markers{
				Color: "#000000",
				List: map[string]types.Marker{
					"a": types.Marker{Height: 3, Width: 4, X: 0.25, Y: 0.5},
				},
			},
		}
		expectedMatrices := &types.Matrices{
			Abundance: [][]float64{
				{1, 1, 2, 2},
				{2, 2, 2, 3},
				{1, 1, 4, 1},
				{4, 2, 2, 3},
			},
		}

		DownsampleData(data, matrices, downsampleTreshold)
		Expect(data).To(Equal(expectedData), "should not alter data")
		Expect(matrices).To(Equal(expectedMatrices), "should not alter matrices")
	})

	It("should downsample", func() {
		data := &Heatmap{
			Annotations: types.Annotations{
				FontSize: 16,
				List: map[string]types.Annotation{
					"a": {Text: "a", Position: types.AnnotationPosition{X: 0.5, Y: 0.25}},
					"b": {Text: "b", Position: types.AnnotationPosition{X: 0.5, Y: 0.75}},
				},
			},
			Markers: types.Markers{
				Color: "#000000",
				List: map[string]types.Marker{
					"a": types.Marker{Height: 3, Width: 4, X: 0.25, Y: 0.5},
				},
			},
		}
		downsampleTreshold := 2
		matrices := &types.Matrices{
			Abundance: [][]float64{
				{1, 1, 2, 2},
				{2, 2, 2, 3},
				{1, 1, 4, 1},
				{4, 2, 2, 3},
			},
		}

		expectedData := &Heatmap{
			Annotations: types.Annotations{
				FontSize: 16,
				List: map[string]types.Annotation{
					"a": {Text: "a", Position: types.AnnotationPosition{X: 0.25, Y: 0.125}},
					"b": {Text: "b", Position: types.AnnotationPosition{X: 0.25, Y: 0.375}},
				},
			},
			Markers: types.Markers{
				Color: "#000000",
				List: map[string]types.Marker{
					"a": types.Marker{Height: 2, Width: 2, X: 0.125, Y: 0.25},
				},
			},
		}
		expectedMatrices := &types.Matrices{
			Abundance: [][]float64{
				{1.5, 2.25},
				{2, 2.5},
			},
		}

		DownsampleData(data, matrices, downsampleTreshold)
		Expect(data).To(Equal(expectedData), "should adjust data annotations and markers")
		Expect(matrices).To(Equal(expectedMatrices), "should downsample matrices")
	})
})
