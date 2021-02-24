package circheatmap

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Parameters scv", func() {
	It("should parse settings", func() {
		settings := types.Settings{
			Abundance:     "AvgSpec",
			Condition:     "Bait",
			Control:       "ctrl",
			Files:         []string{"file1", "file2"},
			Normalization: "total",
			Readout:       "Prey",
			Score:         "bfdr",
			ScoreType:     "lte",
			Type:          "scv",
		}

		expectedParameters := map[string]interface{}{
			"abundanceColumn": "AvgSpec",
			"analysisType":    "scv",
			"conditionColumn": "Bait",
			"controlColumn":   "ctrl",
			"files":           []string{"file1", "file2"},
			"imageType":       "circheatmap",
			"normalization":   "total",
			"readoutColumn":   "Prey",
			"scoreColumn":     "bfdr",
			"scoreType":       "lte",
		}
		expectedString, _ := json.Marshal(expectedParameters)
		expected := fmt.Sprintf("\"parameters\": %s", expectedString)
		Expect(parseParameters(settings)).To(Equal(expected))
	})
})
