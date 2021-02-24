package settings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("SCV validation", func() {
	It("should validate settings and append additional abundance columns to column map", func() {
		analysis := &types.Analysis{
			Settings: types.Settings{
				Abundance:      "avgSpec",
				AbundanceCap:   50,
				Condition:      "bait",
				Control:        "ctrl",
				Files:          []string{"file.txt"},
				MinAbundance:   10,
				OtherAbundance: []string{"column1", "column2"},
				Readout:        "prey",
				ReadoutLength:  "preyLength",
				Score:          "fdr",
				ScoreType:      "lte",
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
				AbundanceCap:   50,
				Condition:      "bait",
				Control:        "ctrl",
				Files:          []string{"file.txt"},
				MinAbundance:   10,
				OtherAbundance: []string{"column1", "column2"},
				Readout:        "prey",
				ReadoutLength:  "preyLength",
				Score:          "fdr",
				ScoreType:      "lte",
			},
		}
		validateSCVSettings(analysis)
		Expect(analysis).To(Equal(expected))
	})

	It("should validate settings and append abundance and map columns", func() {
		analysis := &types.Analysis{
			Settings: types.Settings{
				Abundance:          "avgSpec",
				Condition:          "bait",
				ConditionMapColumn: "baitid",
				Control:            "ctrl",
				Files:              []string{"file.txt"},
				OtherAbundance:     []string{"column1", "column2"},
				Readout:            "prey",
				ReadoutLength:      "preyLength",
				ReadoutMapColumn:   "preyid",
				Score:              "fdr",
			},
		}

		expected := &types.Analysis{
			Columns: map[string]string{
				"abundance":     "avgSpec",
				"baitid":        "baitid",
				"column1":       "column1",
				"column2":       "column2",
				"condition":     "bait",
				"control":       "ctrl",
				"preyid":        "preyid",
				"readout":       "prey",
				"readoutLength": "preyLength",
				"score":         "fdr",
			},
			Settings: types.Settings{
				Abundance:          "avgSpec",
				Condition:          "bait",
				ConditionMapColumn: "baitid",
				Control:            "ctrl",
				Files:              []string{"file.txt"},
				OtherAbundance:     []string{"column1", "column2"},
				Readout:            "prey",
				ReadoutLength:      "preyLength",
				ReadoutMapColumn:   "preyid",
				Score:              "fdr",
			},
		}
		validateSCVSettings(analysis)
		Expect(analysis).To(Equal(expected))
	})
})
