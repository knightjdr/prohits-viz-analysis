package correlation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kendall method", func() {
	It("should calculate correlation coefficient", func() {
		tests := []map[string][]float64{
			{
				"x": []float64{0.25, 5, 1, 3, 8},
				"y": []float64{0.5, 10, 2, 6, 16},
			},
			{
				"x": []float64{1, 5, 1, 10},
				"y": []float64{10, 2, 10, 1},
			},
			{
				"x": []float64{0.25, 5, 1, 3, 8},
				"y": []float64{8, 4, 9, 0.33, 7},
			},
		}

		expected := []float64{1, -1, -0.2}

		for i, test := range tests {
			Expect(Kendall(test["x"], test["y"])).To(BeNumerically("~", expected[i], 0.001))
		}
	})
})

var _ = Describe("Dense rank", func() {
	It("should return the dense rank for a slice", func() {
		tests := [][]float64{
			[]float64{1, 2, 3, 3, 4, 4, 1},
			[]float64{1, 1, 3, 3, 4, 4, 4},
		}

		expected := [][]int{
			[]int{1, 2, 3, 3, 4, 4, 5},
			[]int{1, 1, 2, 2, 3, 3, 3},
		}

		for i, test := range tests {
			Expect(getDenseRank(test)).To(Equal(expected[i]))
		}
	})
})

var _ = Describe("Count rank ties", func() {
	It("should return the dense rank for a slice", func() {
		tests := [][]int{
			[]int{1, 2, 3, 3, 4, 4, 5},
			[]int{1, 1, 2, 2, 3, 3, 3},
		}

		expected := []int{2, 5}

		for i, test := range tests {
			Expect(countRankTie(test)).To(Equal(expected[i]))
		}
	})
})

var _ = Describe("Count ties", func() {
	It("should return the dense rank for a slice", func() {
		tests := [][]int{
			[]int{1, 2, 3, 3, 4, 4, 5},
			[]int{1, 1, 2, 2, 3, 3, 3},
		}

		expected := []int{29, 11}

		for i, test := range tests {
			Expect(countTies(test)).To(Equal(expected[i]))
		}
	})
})

var _ = Describe("Get diff between non ties", func() {
	It("should return the dense rank for a slice", func() {
		tests := []map[string][]int{
			{
				"x": []int{1, 2, 3, 3, 4, 4, 5},
				"y": []int{1, 1, 2, 2, 3, 3, 3},
			},
			{
				"x": []int{1, 2, 3, 3, 4, 4, 5},
				"y": []int{1, 3, 3, 1, 2, 2, 3},
			},
		}

		expected := [][]int{
			[]int{1, 1, 2, 2, 1},
			[]int{1, 1, 1, 1, 2, 1},
		}

		for i, test := range tests {
			Expect(getDiffNonTies(test["x"], test["y"])).To(Equal(expected[i]))
		}
	})
})

var _ = Describe("Discordant pairs", func() {
	It("should return the number of discordant pairs", func() {
		tests := []map[string][]int{
			{
				"x": []int{1, 2, 3, 3, 4, 4, 5},
				"y": []int{1, 1, 2, 2, 3, 3, 3},
			},
			{
				"x": []int{1, 2, 3, 3, 4, 4, 5},
				"y": []int{1, 3, 3, 1, 2, 2, 3},
			},
		}

		expected := []int{0, 5}
		for i, test := range tests {
			Expect(discordant(test["x"], test["y"])).To(Equal(expected[i]))
		}
	})
})

var _ = Describe("Diff between non ties", func() {
	It("should return the diff between", func() {
		tests := []map[string][]int{
			{
				"x": []int{1, 2, 3, 3, 4, 4, 5},
				"y": []int{1, 1, 2, 2, 3, 3, 3},
			},
			{
				"x": []int{1, 2, 3, 3, 4, 4, 5},
				"y": []int{1, 3, 3, 1, 2, 2, 3},
			},
		}

		expected := [][]int{
			[]int{1, 1, 2, 2, 1},
			[]int{1, 1, 1, 1, 2, 1},
		}
		for i, test := range tests {
			Expect(getDiffNonTies(test["x"], test["y"])).To(Equal(expected[i]))
		}
	})
})
