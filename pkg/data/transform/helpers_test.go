package transform

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Adjust abundance by multiplier", func() {
	It("should adjust abundances by multipliers ", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"readout": "readout1", "abundance": "10", "readoutLength": "2"},
				{"readout": "readout2", "abundance": "1", "readoutLength": "5"},
				{"readout": "readout1", "abundance": "10|5", "readoutLength": "2"},
				{"readout": "readout3", "abundance": "10|5|2.5", "readoutLength": "10"},
			},
		}
		multiplier := map[string]float64{
			"readout1": 2.5,
			"readout2": 1,
			"readout3": 0.5,
		}

		expected := []map[string]string{
			{"readout": "readout1", "abundance": "25", "readoutLength": "2"},
			{"readout": "readout2", "abundance": "1", "readoutLength": "5"},
			{"readout": "readout1", "abundance": "25|12.5", "readoutLength": "2"},
			{"readout": "readout3", "abundance": "5|2.5|1.25", "readoutLength": "10"},
		}

		adjustAbundanceByMultiplier(analysis, "readout", multiplier)
		Expect(analysis.Data).To(Equal(expected))
	})
})

var _ = Describe("Median from map", func() {
	It("should return median length", func() {
		dict := map[string]float64{
			"readout1": 2,
			"readout2": 5,
			"readout3": 10,
		}

		expected := float64(5)
		Expect(calculateMedian(dict)).To(Equal(expected))
	})
})

var _ = Describe("Calculate multiplers", func() {
	It("should return a map with multipliers", func() {
		dict := map[string]float64{
			"readout1": 2,
			"readout2": 5,
			"readout3": 10,
			"readout4": 0,
		}
		median := float64(5)

		expected := map[string]float64{
			"readout1": 2.5,
			"readout2": 1,
			"readout3": 0.5,
			"readout4": 1,
		}
		Expect(calculateMultipliers(dict, median)).To(Equal(expected))
	})
})
