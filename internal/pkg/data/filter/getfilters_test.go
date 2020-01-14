package filter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

var _ = Describe("Define abundance and score filter", func() {
	It("should return a function for filtering by abundance and score", func() {
		settings := types.Settings{
			MinAbundance:  10,
			PrimaryFilter: 0.05,
			ScoreType:     "lte",
		}

		filter := getAbundanceAndScoreFilter(settings)
		Expect(filter(11, 0.04)).To(BeTrue(), "should return true when both abundance and score pass filters")
		Expect(filter(11, 0.06)).To(BeFalse(), "should return false when only abundance passes filter")
		Expect(filter(9, 0.04)).To(BeFalse(), "should return false when only score passes filter")
	})
})

var _ = Describe("Define condition and readout filter", func() {
	It("should return a function for filtering by condition and readout", func() {
		settings := types.Settings{
			ConditionClustering: "none",
			ConditionList:       []string{"conditionA", "conditionB"},
			ReadoutClustering:   "none",
			ReadoutList:         []string{"readoutA", "readoutB"},
		}

		filter := getConditionAndReadoutFilter(settings)
		Expect(filter("conditionA", "readoutA")).To(BeTrue(), "should return true when both condition and readout pass filters")
		Expect(filter("conditionA", "readoutC")).To(BeFalse(), "should return false when only condition passes filter")
		Expect(filter("conditionC", "readoutA")).To(BeFalse(), "should return false when only readout passes filter")
	})
})

var _ = Describe("Define abundance filter", func() {
	It("should return a function for filtering by abundance", func() {
		minAbundance := float64(10)

		filterByAbundance := defineAbundanceFilter(minAbundance)
		Expect(filterByAbundance(9)).To(BeFalse(), "should return false when value is less than filter")
		Expect(filterByAbundance(10)).To(BeTrue(), "should return true when value is equal to  filter")
		Expect(filterByAbundance(11)).To(BeTrue(), "should return true when value is greater than filter")
	})
})

var _ = Describe("Define score filter", func() {
	It("should return a function for filtering scores when smaller values are better", func() {
		filter := 0.05
		scoreType := "lte"

		filterByScore := DefineScoreFilter(scoreType, filter)
		Expect(filterByScore(0.04)).To(BeTrue(), "should return true when value is less than filter")
		Expect(filterByScore(0.05)).To(BeTrue(), "should return true when value is equal to  filter")
		Expect(filterByScore(0.06)).To(BeFalse(), "should return false when value is greater than filter")
	})

	It("should return a function for filtering scores when larger values are better", func() {
		filter := 0.05
		scoreType := "gte"

		filterByScore := DefineScoreFilter(scoreType, filter)
		Expect(filterByScore(0.04)).To(BeFalse(), "should return false when value is less than filter")
		Expect(filterByScore(0.05)).To(BeTrue(), "should return true when value is equal to  filter")
		Expect(filterByScore(0.06)).To(BeTrue(), "should return true when value is greater than filter")
	})
})

var _ = Describe("Define name filter", func() {
	It("should return a function for filtering by name in a list", func() {
		clusteringType := "none"
		names := []string{"conditionA", "conditionB"}

		filterByName := defineNameFilter(clusteringType, names)
		Expect(filterByName("conditionA")).To(BeTrue(), "should return true when value is in list")
		Expect(filterByName("conditionC")).To(BeFalse(), "should return false when value is not in list")
	})

	It("should return a function that always returns true when clustering type is not 'none'", func() {
		clusteringType := "hierarchical"
		names := []string{"conditionA", "conditionB"}

		filterByName := defineNameFilter(clusteringType, names)
		Expect(filterByName("conditionA")).To(BeTrue(), "should return true when value is in list")
		Expect(filterByName("conditionC")).To(BeTrue(), "should return true when value is not in list")
	})
})
