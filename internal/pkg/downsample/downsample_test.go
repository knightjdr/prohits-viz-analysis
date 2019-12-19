package downsample

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Downsample", func() {
	It("should downsample a matrix based on input scale", func() {
		matrix := [][]float64{
			{1, 1, 2, 2, 2},
			{2, 2, 2, 3, 3},
			{1, 1, 4, 1, 1},
			{4, 2, 2, 3, 1},
			{5, 3, 4, 1, 3},
		}
		scale := 2.5

		expected := [][]float64{
			{1.6, 2.24},
			{3.04, 2.08},
		}
		Expect(downsample(matrix, scale)).To(Equal(expected))
	})
})

var _ = Describe("Create down sample matrix", func() {
	It("should create subgrid", func() {
		matrix := make([][]float64, 10)
		matrix[0] = make([]float64, 6)
		scale := 2.0

		expected := [][]float64{
			make([]float64, 3),
			make([]float64, 3),
			make([]float64, 3),
			make([]float64, 3),
			make([]float64, 3),
		}
		Expect(createDownsampledMatrix(matrix, scale)).To(Equal(expected))
	})
})

var _ = Describe("Initialize subgrid", func() {
	It("should create subgrid", func() {
		scale := 4.3

		expected := gridParameters{
			multipliers: []float64{1, 1, 1, 1, 0.3},
			scale:       scale,
			startIndex:  0,
			totalCells:  5,
		}
		actual := initializeSubgrid(scale)
		for i, value := range actual.multipliers {
			Expect(value).To(BeNumerically("~", expected.multipliers[i], 0.001), "should set multipliers")
		}
		Expect(actual.scale).To(BeNumerically("~", expected.scale, 0.001), "should set scale")
		Expect(actual.startIndex).To(Equal(expected.startIndex), "should set start index")
		Expect(actual.totalCells).To(Equal(expected.totalCells), "should set total cell number")
	})
})

var _ = Describe("Update subgrid", func() {
	It("should update based on previous subgrid", func() {
		subgrid := gridParameters{
			multipliers: []float64{1, 1, 1, 1, 0.3},
			scale:       4.3,
			startIndex:  0,
			totalCells:  5,
		}

		expected := gridParameters{
			multipliers: []float64{0.7, 1, 1, 1, 0.6},
			scale:       4.3,
			startIndex:  4,
			totalCells:  5,
		}
		actual := updateSubgrid(subgrid, 1)
		for i, value := range actual.multipliers {
			Expect(value).To(BeNumerically("~", expected.multipliers[i], 0.001), "should update multipliers")
		}
		Expect(actual.scale).To(BeNumerically("~", expected.scale, 0.001), "should set scale")
		Expect(actual.startIndex).To(Equal(expected.startIndex), "should update start index")
		Expect(actual.totalCells).To(Equal(expected.totalCells), "should set total cell number")
	})
})

var _ = Describe("Add full cell multipliers", func() {
	It("should update based on previous subgrid", func() {
		expected := []float64{1, 1, 1, 1}
		Expect(addFullCellMultipliers(4)).To(Equal(expected))
	})
})

var _ = Describe("Average subgrid", func() {
	It("should average subgrid values", func() {
		matrix := [][]float64{
			{1, 1, 2, 2, 2},
			{2, 2, 2, 3, 3},
			{1, 1, 4, 1, 1},
		}
		columnSubgrid := gridParameters{
			multipliers: []float64{0.5, 1, 1},
			startIndex:  2,
			totalCells:  3,
		}
		rowSubgrid := gridParameters{
			multipliers: []float64{1, 1, 0.5},
			startIndex:  0,
			totalCells:  3,
		}

		expected := 2.24
		Expect(averageSubgridValues(matrix, rowSubgrid, columnSubgrid)).To(BeNumerically("~", expected, 0.001))
	})

	It("should average subgrid values when subgrid parameters exceed matrix bounds", func() {
		matrix := [][]float64{
			{1, 1, 2, 2, 2},
			{2, 2, 2, 3, 3},
			{1, 1, 4, 1, 1},
		}
		columnSubgrid := gridParameters{
			multipliers: []float64{0.5, 1, 1},
			startIndex:  3,
			totalCells:  3,
		}
		rowSubgrid := gridParameters{
			multipliers: []float64{0.5, 1, 1},
			startIndex:  1,
			totalCells:  3,
		}

		expected := 1.667
		Expect(averageSubgridValues(matrix, rowSubgrid, columnSubgrid)).To(BeNumerically("~", expected, 0.001))
	})
})
