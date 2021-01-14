package scatter

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Parameters", func() {
	It("should parse settings for condition-condition and return string", func() {
		settings := types.Settings{
			Abundance:     "AvgSpec",
			Condition:     "Bait",
			Control:       "ctrl",
			Files:         []string{"file1", "file2"},
			Normalization: "total",
			Readout:       "Prey",
			Score:         "bfdr",
			ScoreType:     "lte",
			Type:          "condition-condition",
		}

		expectedParameters := map[string]interface{}{
			"abundanceColumn":        "AvgSpec",
			"analysisType":           "condition-condition",
			"conditionColumn":        "Bait",
			"controlColumn":          "ctrl",
			"files":                  []string{"file1", "file2"},
			"imageType":              "scatter",
			"mockConditionAbundance": false,
			"normalization":          "total",
			"readoutColumn":          "Prey",
			"scoreColumn":            "bfdr",
			"scoreType":              "lte",
		}
		expectedString, _ := json.Marshal(expectedParameters)
		expected := fmt.Sprintf("\"parameters\": %s", expectedString)
		Expect(parseParameters("condition-condition", settings)).To(Equal(expected))
	})

	It("should parse settings for specificity and return string", func() {
		settings := types.Settings{
			Abundance:         "AvgSpec",
			Condition:         "Bait",
			Control:           "ctrl",
			Files:             []string{"file1", "file2"},
			Normalization:     "total",
			PrimaryFilter:     0.01,
			Readout:           "Prey",
			Score:             "bfdr",
			ScoreType:         "lte",
			SpecificityMetric: "fe",
			Type:              "specificity",
		}

		expectedParameters := map[string]interface{}{
			"abundanceColumn":        "AvgSpec",
			"analysisType":           "specificity",
			"conditionColumn":        "Bait",
			"controlColumn":          "ctrl",
			"files":                  []string{"file1", "file2"},
			"imageType":              "scatter",
			"mockConditionAbundance": false,
			"normalization":          "total",
			"primaryFilter":          0.01,
			"readoutColumn":          "Prey",
			"scoreColumn":            "bfdr",
			"scoreType":              "lte",
			"specificityMetric":      "fe",
		}
		expectedString, _ := json.Marshal(expectedParameters)
		expected := fmt.Sprintf("\"parameters\": %s", expectedString)
		Expect(parseParameters("specificity", settings)).To(Equal(expected))
	})
})
