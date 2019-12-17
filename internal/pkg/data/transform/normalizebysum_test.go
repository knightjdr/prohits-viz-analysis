package transform

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Normalize by total sum", func() {
	It("should adjust abundances based on total sum per condition", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "condition1", "readout": "readout1", "abundance": "5"},
				{"condition": "condition1", "readout": "readout2", "abundance": "5"},
				{"condition": "condition2", "readout": "readout1", "abundance": "1"},
				{"condition": "condition2", "readout": "readout2", "abundance": "1.5"},
				{"condition": "condition3", "readout": "readout1", "abundance": "1|2"},
				{"condition": "condition3", "readout": "readout2", "abundance": "3|4"},
			},
		}

		expected := []map[string]string{
			{"condition": "condition1", "readout": "readout1", "abundance": "2.5"},
			{"condition": "condition1", "readout": "readout2", "abundance": "2.5"},
			{"condition": "condition2", "readout": "readout1", "abundance": "2"},
			{"condition": "condition2", "readout": "readout2", "abundance": "3"},
			{"condition": "condition3", "readout": "readout1", "abundance": "1|2"},
			{"condition": "condition3", "readout": "readout2", "abundance": "3|4"},
		}

		normalizeByTotalSum(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})

var _ = Describe("Sum abundance by condition", func() {
	It("should return a map of conditions with summed abundances", func() {
		data := []map[string]string{
			{"condition": "condition1", "readout": "readout1", "abundance": "5"},
			{"condition": "condition1", "readout": "readout2", "abundance": "5"},
			{"condition": "condition2", "readout": "readout1", "abundance": "1"},
			{"condition": "condition2", "readout": "readout2", "abundance": "1.5"},
			{"condition": "condition3", "readout": "readout1", "abundance": "1|2"},
			{"condition": "condition3", "readout": "readout2", "abundance": "3|4"},
		}

		expected := map[string]float64{
			"condition1": 10,
			"condition2": 2.5,
			"condition3": 5,
		}
		Expect(sumAbundanceByCondition(data)).To(Equal(expected))
	})
})
