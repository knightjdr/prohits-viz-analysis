package settings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("CircHeatmap validation", func() {
	It("should validate settings and append abundance columns to column map", func() {
		analysis := &types.Analysis{
			Settings: types.Settings{
				Abundance:      "avgSpec",
				Condition:      "bait",
				Control:        "ctrl",
				Files:          []string{"file.txt"},
				OtherAbundance: []string{"column1", "column2"},
				Readout:        "prey",
				ReadoutLength:  "preyLength",
				Score:          "fdr",
			},
		}

		expected := &types.Analysis{
			Columns: map[string]string{
				"abundance":     "avgSpec",
				"column1":       "column1",
				"column2":       "column2",
				"condition":     "bait",
				"control":       "ctrl",
				"readout":       "prey",
				"readoutLength": "preyLength",
				"score":         "fdr",
			},
			Settings: types.Settings{
				Abundance:      "avgSpec",
				Condition:      "bait",
				Control:        "ctrl",
				Files:          []string{"file.txt"},
				OtherAbundance: []string{"column1", "column2"},
				Readout:        "prey",
				ReadoutLength:  "preyLength",
				Score:          "fdr",
			},
		}
		validateCircHeatmapSettings(analysis)
		Expect(analysis).To(Equal(expected))
	})
})
