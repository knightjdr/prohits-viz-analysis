package transform

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

var _ = Describe("Normalization", func() {
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
				Normalization:        "readout",
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

		normalize(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})

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
			Settings: types.Settings{
				Normalization: "total",
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

		normalize(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})

	It("should not adjust abundances when not required", func() {
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
			{"condition": "condition1", "readout": "readout1", "abundance": "5"},
			{"condition": "condition1", "readout": "readout2", "abundance": "5"},
			{"condition": "condition2", "readout": "readout1", "abundance": "1"},
			{"condition": "condition2", "readout": "readout2", "abundance": "1.5"},
			{"condition": "condition3", "readout": "readout1", "abundance": "1|2"},
			{"condition": "condition3", "readout": "readout2", "abundance": "3|4"},
		}

		normalize(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})
