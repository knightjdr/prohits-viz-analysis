package dotplot

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Score converter", func() {
	It("should create a function for converting an lte scores to an index", func() {
		d := &Dotplot{
			PrimaryFilter:   0.01,
			ScoreType:       "lte",
			SecondaryFilter: 0.05,
		}

		testFunction := defineScoreConverter(d, 100)

		tests := []float64{0, 0.01, 0.04, 0.05, 0.06, 1}
		expected := []int{100, 100, 50, 50, 25, 25}

		for i, test := range tests {
			Expect(testFunction(test)).To(Equal(expected[i]))
		}
	})

	It("should create a function for converting a gte scores to an index", func() {
		d := &Dotplot{
			PrimaryFilter:   0.9,
			ScoreType:       "gte",
			SecondaryFilter: 0.8,
		}

		testFunction := defineScoreConverter(d, 100)

		tests := []float64{1, 0.91, 0.9, 0.89, 0.8, 0.79, 0}
		expected := []int{100, 100, 100, 50, 50, 25, 25}

		for i, test := range tests {
			Expect(testFunction(test)).To(Equal(expected[i]))
		}
	})
})

var _ = Describe("Max allowable index", func() {
	It("should return input number when even", func() {
		tests := []int{2, 46, 100}

		expected := []int{2, 46, 100}
		for i, test := range tests {
			Expect(defineMaxIndex(test)).To(Equal(expected[i]))
		}
	})

	It("should return input number - 1 when odd", func() {
		tests := []int{3, 47, 101}

		expected := []int{2, 46, 100}
		for i, test := range tests {
			Expect(defineMaxIndex(test)).To(Equal(expected[i]))
		}
	})
})
