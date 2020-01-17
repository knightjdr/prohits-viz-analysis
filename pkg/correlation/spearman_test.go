package correlation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spearman method", func() {
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

		expected := []float64{1, -1, -0.5}

		for i, test := range tests {
			Expect(spearman(test["x"], test["y"])).To(BeNumerically("~", expected[i], 0.001))
		}
	})
})

var _ = Describe("Rank array for spearman", func() {
	It("should calculate correlation coefficient", func() {
		tests := [][]float64{
			[]float64{2, 4.5, 6, 4.5, 3},
			[]float64{5, 5, 5, 3, 8, 3.5},
		}

		expected := [][]float64{
			[]float64{1, 3.5, 5, 3.5, 2},
			[]float64{4, 4, 4, 1, 6, 2},
		}

		for i, test := range tests {
			Expect(rankArray(test)).To(Equal(expected[i]))
		}
	})
})
