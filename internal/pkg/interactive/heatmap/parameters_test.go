package heatmap

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

var _ = Describe("Parameters", func() {
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
			"abundanceColumn":    "AvgSpec",
			"analysisType":       "dotplot",
			"clustering":         "hierarchical",
			"clusteringMethod":   "complete",
			"clusteringOptimize": true,
			"conditionColumn":    "Bait",
			"controlColumn":      "ctrl",
			"distance":           "canberra",
			"files":              []string{"file1", "file2"},
			"imageType":          "heatmap",
			"logBase":            "2",
			"normalization":      "total",
			"readoutColumn":      "Prey",
			"scoreColumn":        "bfdr",
			"scoreType":          "lte",
			"xLabel":             "Prey",
			"yLabel":             "Bait",
		}
		expectedString, _ := json.Marshal(expectedParameters)
		expected := fmt.Sprintf("\"parameters\": %s", expectedString)
		Expect(parseParameters("heatmap", settings)).To(Equal(expected))
	})
})
