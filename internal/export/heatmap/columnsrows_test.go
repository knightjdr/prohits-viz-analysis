package heatmap

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Column and row names", func() {
	It("should return column and row names based on order", func() {
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
		}

		expectedColumns := []string{"column2", "column3", "column1"}
		expectedRows := []string{"row1", "row2", "row4"}

		actualColumns, actualRows := getColumnsAndRows(data)
		Expect(actualColumns).To(Equal(expectedColumns), "should return column names in order")
		Expect(actualRows).To(Equal(expectedRows), "should return row names in order")
	})
})
