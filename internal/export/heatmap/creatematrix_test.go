package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create matrix", func() {
	It("should read json", func() {
		data := &heatmap{
			ColumnDB:    []string{"column1", "column2", "column3"},
			ColumnOrder: []int{1, 2, 0},
			RowDB: []rows{
				{Name: "row1", Data: []cell{{Value: 1}, {Value: 2}, {Value: 3}}},
				{Name: "row2", Data: []cell{{Value: 4}, {Value: 5}, {Value: 6}}},
				{Name: "row3", Data: []cell{{Value: 7}, {Value: 8}, {Value: 9}}},
				{Name: "row4", Data: []cell{{Value: 10}, {Value: 11}, {Value: 12}}},
			},
			RowOrder: []int{0, 1, 3},
			Settings: types.Settings{
				AbundanceCap: 50,
				ScoreType:    "lte",
			},
		}

		expected := [][]float64{
			{2, 3, 1},
			{5, 6, 4},
			{11, 12, 10},
		}
		Expect(createMatrix(data)).To(Equal(expected))
	})
})
