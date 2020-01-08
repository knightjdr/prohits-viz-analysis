package frontend_test

import (
	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/matrix/frontend"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Column names", func() {
	It("should return column and row names based on order", func() {
		columnDB := []string{"column1", "column2", "column3"}
		columnOrder := []int{1, 2, 0}

		expected := []string{"column2", "column3", "column1"}

		Expect(GetColumnNames(columnDB, columnOrder)).To(Equal(expected))
	})
})

var _ = Describe("Row names", func() {
	It("should return column and row names based on order", func() {
		rowDB := []Row{
			{Name: "row1", Data: []Cell{{Value: 1}, {Value: 2}, {Value: 3}}},
			{Name: "row2", Data: []Cell{{Value: 4}, {Value: 5}, {Value: 6}}},
			{Name: "row3", Data: []Cell{{Value: 7}, {Value: 8}, {Value: 9}}},
			{Name: "row4", Data: []Cell{{Value: 10}, {Value: 11}, {Value: 12}}},
		}
		rowOrder := []int{0, 1, 3}

		expected := []string{"row1", "row2", "row4"}

		Expect(GetRowNames(rowDB, rowOrder)).To(Equal(expected))
	})
})
