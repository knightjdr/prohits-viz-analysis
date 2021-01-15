package transform_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/data/transform"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Transform abundance", func() {
	It("should transform abundances", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "75", "condition": "condition1", "control": "1|1", "readout": "readout1", "readoutLength": "100"},
				{"abundance": "125", "condition": "condition1", "control": "2|2", "readout": "readout2", "readoutLength": "300"},
				{"abundance": "180", "condition": "condition2", "control": "2|2", "readout": "readout1", "readoutLength": "100"},
				{"abundance": "17", "condition": "condition2", "control": "1|1", "readout": "readout2", "readoutLength": "300"},
			},
			Settings: types.Settings{
				Control:       "ctrl",
				LogBase:       "2",
				Normalization: "total",
				ReadoutLength: "PreyLength",
				Type:          "dotplot",
			},
		}
		expected := []map[string]string{
			{"abundance": "7.58", "condition": "condition1", "control": "1|1", "readout": "readout1", "readoutLength": "100"},
			{"abundance": "6.73", "condition": "condition1", "control": "2|2", "readout": "readout2", "readoutLength": "300"},
			{"abundance": "8.18", "condition": "condition2", "control": "2|2", "readout": "readout1", "readoutLength": "100"},
			{"abundance": "3.12", "condition": "condition2", "control": "1|1", "readout": "readout2", "readoutLength": "300"},
		}

		Abundance(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})
