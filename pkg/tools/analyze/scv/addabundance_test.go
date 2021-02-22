package scv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Add abundances for SCV", func() {
	It("should add abundance columns to data", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "conditionA", "readout": "readoutA", "score": "0.01", "abundance": "10", "FoldChange": "5"},
				{"condition": "conditionA", "readout": "readoutB", "score": "0.02", "abundance": "15", "FoldChange": "1"},
				{"condition": "conditionA", "readout": "readoutC", "score": "0", "abundance": "20", "FoldChange": "10"},
				{"condition": "conditionB", "readout": "readoutA", "score": "0.01", "abundance": "10", "FoldChange": "7"},
				{"condition": "conditionB", "readout": "readoutB", "score": "0", "abundance": "20", "FoldChange": "10"},
				{"condition": "conditionB", "readout": "readoutC", "score": "0.01", "abundance": "4", "FoldChange": "2"},
			},
			Settings: types.Settings{
				Abundance:      "Abundance",
				MinAbundance:   5,
				OtherAbundance: []string{"FoldChange"},
				PrimaryFilter:  0.01,
				ScoreType:      "lte",
			},
		}
		data := make(map[string]map[string]map[string]float64, 0)

		expected := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutA": {
					"Abundance":  10,
					"FoldChange": 5,
				},
				"readoutC": {
					"Abundance":  20,
					"FoldChange": 10,
				},
			},
			"conditionB": {
				"readoutA": {
					"Abundance":  10,
					"FoldChange": 7,
				},
				"readoutB": {
					"Abundance":  20,
					"FoldChange": 10,
				},
			},
		}

		addAbundance(data, analysis)
		Expect(data).To(Equal(expected))
	})
})
