package correlation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pearson method", func() {
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

		expected := []float64{1, -0.915, -0.220}

		for i, test := range tests {
			Expect(pearson(test["x"], test["y"])).To(BeNumerically("~", expected[i], 0.001))
		}
	})
})
