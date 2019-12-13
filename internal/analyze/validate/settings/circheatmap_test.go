package settings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

var _ = Describe("CircHeatmap validation", func() {
	It("should validate settings and append abundance columns to column map", func() {
		settings := &types.CircHeatmap{
			File: types.File{
				Abundance:     "avgSpec",
				Condition:     "bait",
				Control:       "ctrl",
				Files:         []string{"file.txt"},
				Readout:       "prey",
				ReadoutLength: "preyLength",
				Score:         "fdr",
			},
			OtherAbundance: []string{"column1", "column2"},
		}

		expectedColumnMap := map[string]string{
			"abundance":     "avgSpec",
			"column1":       "column1",
			"column2":       "column2",
			"condition":     "bait",
			"control":       "ctrl",
			"readout":       "prey",
			"readoutLength": "preyLength",
			"score":         "fdr",
		}
		expectedSettings := settings
		actualColumnMap, acutalSettings := validateCircHeatmapSettings(settings)
		Expect(actualColumnMap).To(Equal(expectedColumnMap), "should return column map")
		Expect(acutalSettings).To(Equal(expectedSettings), "should return settings")
	})
})
