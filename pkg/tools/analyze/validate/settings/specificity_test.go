package settings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Specificity validation", func() {
	It("should validate settings", func() {
		analysis := &types.Analysis{
			Settings: types.Settings{
				Abundance:     "avgSpec",
				Condition:     "bait",
				Control:       "ctrl",
				Files:         []string{"file.txt"},
				Readout:       "prey",
				ReadoutLength: "preyLength",
				Score:         "fdr",
			},
		}

		expected := &types.Analysis{
			Columns: map[string]string{
				"abundance":     "avgSpec",
				"condition":     "bait",
				"control":       "ctrl",
				"readout":       "prey",
				"readoutLength": "preyLength",
				"score":         "fdr",
			},
			Settings: types.Settings{
				Abundance:     "avgSpec",
				Condition:     "bait",
				Control:       "ctrl",
				Files:         []string{"file.txt"},
				Readout:       "prey",
				ReadoutLength: "preyLength",
				Score:         "fdr",
			},
		}
		validateSpecificitySettings(analysis)
		Expect(analysis).To(Equal(expected))
	})
})
