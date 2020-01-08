package sync

import (
	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/matrix/frontend"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create matrices", func() {
	It("should create matrices for dotplot data", func() {
		data := &minimap{
			ColumnOrder: []int{1, 2, 0},
			ImageType:   "dotplot",
			RowOrder:    []int{0, 1, 3},
			RowDB: []Row{
				{
					Name: "row1",
					Data: []Cell{
						{Ratio: 0.33, Score: 0.05, Value: 1},
						{Ratio: 0.67, Score: 0.01, Value: 2},
						{Ratio: 1, Score: 0, Value: 3},
					},
				},
				{
					Name: "row2",
					Data: []Cell{
						{Ratio: 0.67, Score: 0.05, Value: 4},
						{Ratio: 0.83, Score: 0.01, Value: 5},
						{Ratio: 1, Score: 0, Value: 6},
					},
				},
				{
					Name: "row3",
					Data: []Cell{
						{Ratio: 0.78, Score: 0.05, Value: 7},
						{Ratio: 0.89, Score: 0.05, Value: 8},
						{Ratio: 1, Score: 0.01, Value: 9},
					},
				},
				{
					Name: "row4",
					Data: []Cell{
						{Ratio: 0.83, Score: 0.01, Value: 10},
						{Ratio: 0.92, Score: 0.01, Value: 11},
						{Ratio: 1, Score: 0, Value: 12},
					},
				},
			},
		}

		expected := &types.Matrices{
			Abundance: [][]float64{
				{2, 3, 1},
				{5, 6, 4},
				{11, 12, 10},
			},
			Ratio: [][]float64{
				{0.67, 1, 0.33},
				{0.83, 1, 0.67},
				{0.92, 1, 0.83},
			},
			Score: [][]float64{
				{0.01, 0, 0.05},
				{0.01, 0, 0.05},
				{0.01, 0, 0.01},
			},
		}

		Expect(createMatrices(data)).To(Equal(expected))
	})

	It("should create matrices for heatmap data", func() {
		data := &minimap{
			ColumnOrder: []int{1, 2, 0},
			ImageType:   "heatmap",
			RowOrder:    []int{0, 1, 3},
			RowDB: []Row{
				{
					Name: "row1",
					Data: []Cell{{Value: 1}, {Value: 2}, {Value: 3}},
				},
				{
					Name: "row2",
					Data: []Cell{{Value: 4}, {Value: 5}, {Value: 6}},
				},
				{
					Name: "row3",
					Data: []Cell{{Value: 7}, {Value: 8}, {Value: 9}},
				},
				{
					Name: "row4",
					Data: []Cell{{Value: 10}, {Value: 11}, {Value: 12}},
				},
			},
		}

		expected := &types.Matrices{
			Abundance: [][]float64{
				{2, 3, 1},
				{5, 6, 4},
				{11, 12, 10},
			},
		}

		Expect(createMatrices(data)).To(Equal(expected))
	})
})
