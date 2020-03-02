package transform

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Normalize by readout", func() {
	It("should adjust abundances based on abundance of readout", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "condition1", "readout": "readout1", "abundance": "5"},
				{"condition": "condition1", "readout": "readout2", "abundance": "5"},
				{"condition": "condition2", "readout": "readout1", "abundance": "1"},
				{"condition": "condition2", "readout": "readout2", "abundance": "1.5"},
				{"condition": "condition3", "readout": "readout1", "abundance": "1|2"},
			},
			Settings: types.Settings{
				NormalizationReadout: "readout2",
			},
		}

		expected := []map[string]string{
			{"condition": "condition1", "readout": "readout1", "abundance": "1.5"},
			{"condition": "condition1", "readout": "readout2", "abundance": "1.5"},
			{"condition": "condition2", "readout": "readout1", "abundance": "1"},
			{"condition": "condition2", "readout": "readout2", "abundance": "1.5"},
			{"condition": "condition3", "readout": "readout1", "abundance": "1|2"},
		}

		normalizeByReadout(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})

var _ = Describe("Get readout abundance by condition", func() {
	It("should return a map of conditions with readout abundances", func() {
		data := []map[string]string{
			{"condition": "condition1", "readout": "readout1", "abundance": "5"},
			{"condition": "condition1", "readout": "readout2", "abundance": "5"},
			{"condition": "condition2", "readout": "readout1", "abundance": "1"},
			{"condition": "condition2", "readout": "readout2", "abundance": "1.5"},
			{"condition": "condition3", "readout": "readout1", "abundance": "1|2"},
		}

		expected := map[string]float64{
			"condition1": 5,
			"condition2": 1.5,
			"condition3": 0,
		}
		Expect(getReadoutAbundance(data, "readout2")).To(Equal(expected))
	})
})
