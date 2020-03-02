package transform

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Adjust by readout length", func() {
	It("should adjust abundances based on readout length", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"readout": "readout1", "abundance": "10", "readoutLength": "2"},
				{"readout": "readout2", "abundance": "1", "readoutLength": "5"},
				{"readout": "readout1", "abundance": "10|5", "readoutLength": "2"},
				{"readout": "readout3", "abundance": "10|5|2.5", "readoutLength": "10"},
			},
			Settings: types.Settings{
				ReadoutLength: "PreyLength",
			},
		}

		expected := []map[string]string{
			{"readout": "readout1", "abundance": "25", "readoutLength": "2"},
			{"readout": "readout2", "abundance": "1", "readoutLength": "5"},
			{"readout": "readout1", "abundance": "25|12.5", "readoutLength": "2"},
			{"readout": "readout3", "abundance": "5|2.5|1.25", "readoutLength": "10"},
		}

		adjustByReadoutLength(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})

	It("should not adjust abundances based when not requested", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"readout": "readout1", "abundance": "10", "readoutLength": "2"},
				{"readout": "readout2", "abundance": "1", "readoutLength": "5"},
				{"readout": "readout1", "abundance": "10|5", "readoutLength": "2"},
				{"readout": "readout3", "abundance": "10|5|2.5", "readoutLength": "10"},
			},
		}

		expected := []map[string]string{
			{"readout": "readout1", "abundance": "10", "readoutLength": "2"},
			{"readout": "readout2", "abundance": "1", "readoutLength": "5"},
			{"readout": "readout1", "abundance": "10|5", "readoutLength": "2"},
			{"readout": "readout3", "abundance": "10|5|2.5", "readoutLength": "10"},
		}

		adjustByReadoutLength(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})

var _ = Describe("Get length of unique readouts", func() {
	It("should return a map of lengths, ignoring uparseable values", func() {
		data := []map[string]string{
			{"readout": "readout1", "readoutLength": "a"},
			{"readout": "readout2", "readoutLength": "5"},
			{"readout": "readout1", "readoutLength": "2"},
			{"readout": "readout3", "readoutLength": "10"},
		}

		expected := map[string]float64{
			"readout1": 2,
			"readout2": 5,
			"readout3": 10,
		}
		Expect(getLengthOfUniqueReadouts(data)).To(Equal(expected))
	})
})
