package scatter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Infer settings", func() {
	It("should infer settings for Binary plot", func() {
		plotSettings := &jsonSettings{
			Plot:  "Binary",
			Score: "FDR",
		}

		expected := types.Settings{
			Score: "FDR",
			Type:  "condition-condition",
		}
		Expect(inferSettings(plotSettings)).To(Equal(expected))
	})

	It("should infer settings for Binary-biDist plot", func() {
		plotSettings := &jsonSettings{
			Plot:      "Binary-biDist",
			Score:     "FDR",
			ScoreType: 1,
		}

		expected := types.Settings{
			Score:     "FDR",
			ScoreType: "gte",
			Type:      "condition-condition",
		}
		Expect(inferSettings(plotSettings)).To(Equal(expected))
	})

	It("should infer settings for Specificity plot", func() {
		plotSettings := &jsonSettings{
			Plot:   "Specificity",
			Score:  "FDR",
			XLabel: "AvgSpec",
		}

		expected := types.Settings{
			Abundance: "AvgSpec",
			Score:     "FDR",
			Type:      "specificity",
		}
		Expect(inferSettings(plotSettings)).To(Equal(expected))
	})
})

var _ = Describe("Determine score type", func() {
	It("should identify current score type from previous nomenclature", func() {
		tests := []int{1, 0}

		expected := []string{"gte", "lte"}
		for i, test := range tests {
			Expect(convertScoreTypeFromInt(test)).To(Equal(expected[i]))
		}
	})
})

var _ = Describe("Determine analysis type", func() {
	It("should identify current analysis name from previous nomenclature", func() {
		tests := []string{"Binary", "Binary-biDist", "Specificity", "Transition"}

		expected := []string{"condition-condition", "condition-condition", "specificity", ""}
		for i, test := range tests {
			Expect(determineAnalysisType(test)).To(Equal(expected[i]))
		}
	})
})
