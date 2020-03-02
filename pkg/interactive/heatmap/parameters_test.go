package heatmap

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Parameters", func() {
	It("should parse settings for correlation and return string", func() {
		settings := types.Settings{
			Abundance:                "AvgSpec",
			Clustering:               "hierarchical",
			ClusteringMethod:         "complete",
			ClusteringOptimize:       true,
			Condition:                "Bait",
			ConditionAbundanceFilter: 5,
			ConditionScoreFilter:     0.1,
			Correlation:              "pearson",
			Control:                  "ctrl",
			Distance:                 "canberra",
			Files:                    []string{"file1", "file2"},
			LogBase:                  "2",
			Normalization:            "total",
			Readout:                  "Prey",
			ReadoutAbundanceFilter:   10,
			ReadoutScoreFilter:       0.01,
			Score:                    "bfdr",
			ScoreType:                "lte",
			Type:                     "dotplot",
			XLabel:                   "Prey",
			YLabel:                   "Bait",
		}

		expectedParameters := map[string]interface{}{
			"abundanceColumn":           "AvgSpec",
			"analysisType":              "dotplot",
			"clustering":                "hierarchical",
			"clusteringMethod":          "complete",
			"clusteringOptimize":        true,
			"conditionAbundanceFilter":  5,
			"conditionColumn":           "Bait",
			"conditionScoreFilter":      0.1,
			"controlColumn":             "ctrl",
			"correlation":               "pearson",
			"distance":                  "canberra",
			"files":                     []string{"file1", "file2"},
			"imageType":                 "heatmap",
			"IgnoreSourceTargetMatches": false,
			"logBase":                   "2",
			"minConditions":             0,
			"mockConditionAbundance":    false,
			"normalization":             "total",
			"parsimoniousReadouts":      false,
			"readoutAbundanceFilter":    10,
			"readoutColumn":             "Prey",
			"readoutScoreFilter":        0.01,
			"scoreColumn":               "bfdr",
			"scoreType":                 "lte",
			"useReplicates":             false,
			"xLabel":                    "Prey",
			"yLabel":                    "Bait",
		}
		expectedString, _ := json.Marshal(expectedParameters)
		expected := fmt.Sprintf("\"parameters\": %s", expectedString)
		Expect(parseParameters("correlation", settings)).To(Equal(expected))
	})

	It("should parse settings for heatmap and return string", func() {
		settings := types.Settings{
			Abundance:          "AvgSpec",
			Clustering:         "hierarchical",
			ClusteringMethod:   "complete",
			ClusteringOptimize: true,
			Condition:          "Bait",
			Control:            "ctrl",
			Distance:           "canberra",
			Files:              []string{"file1", "file2"},
			LogBase:            "2",
			Normalization:      "total",
			Readout:            "Prey",
			Score:              "bfdr",
			ScoreType:          "lte",
			Type:               "dotplot",
			XLabel:             "Prey",
			YLabel:             "Bait",
		}

		expectedParameters := map[string]interface{}{
			"abundanceColumn":        "AvgSpec",
			"analysisType":           "dotplot",
			"clustering":             "hierarchical",
			"clusteringMethod":       "complete",
			"clusteringOptimize":     true,
			"conditionColumn":        "Bait",
			"controlColumn":          "ctrl",
			"distance":               "canberra",
			"files":                  []string{"file1", "file2"},
			"imageType":              "heatmap",
			"logBase":                "2",
			"minConditions":          0,
			"mockConditionAbundance": false,
			"normalization":          "total",
			"parsimoniousReadouts":   false,
			"readoutColumn":          "Prey",
			"scoreColumn":            "bfdr",
			"scoreType":              "lte",
			"xLabel":                 "Prey",
			"yLabel":                 "Bait",
		}
		expectedString, _ := json.Marshal(expectedParameters)
		expected := fmt.Sprintf("\"parameters\": %s", expectedString)
		Expect(parseParameters("heatmap", settings)).To(Equal(expected))
	})
})
