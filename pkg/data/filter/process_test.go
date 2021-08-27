package filter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Process", func() {
	It("should filter data", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "10", "condition": "conditionA", "readout": "readoutA", "score": "0.05"},
				{"abundance": "9", "condition": "conditionA", "readout": "readoutA", "score": "0.05"},
				{"abundance": "10", "condition": "conditionC", "readout": "readoutA", "score": "0.05"},
				{"abundance": "10", "condition": "conditionA", "readout": "readoutC", "score": "0.05"},
				{"abundance": "10", "condition": "conditionA", "readout": "readoutA", "score": "0.06"},
				{"abundance": "11", "condition": "conditionB", "readout": "readoutB", "score": "0.04"},
			},
			Settings: types.Settings{
				ConditionClustering: "none",
				ConditionList:       []string{"conditionA", "conditionB"},
				Clustering:          "none",
				MinAbundance:        10,
				PrimaryFilter:       0.05,
				ReadoutClustering:   "none",
				ReadoutList:         []string{"readoutA", "readoutB"},
				ScoreType:           "lte",
				Type:                "dotplot",
			},
		}

		expected := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "10", "condition": "conditionA", "readout": "readoutA", "score": "0.05"},
				{"abundance": "9", "condition": "conditionA", "readout": "readoutA", "score": "0.05"},
				{"abundance": "10", "condition": "conditionA", "readout": "readoutA", "score": "0.06"},
				{"abundance": "11", "condition": "conditionB", "readout": "readoutB", "score": "0.04"},
			},
			Settings: types.Settings{
				ConditionClustering: "none",
				ConditionList:       []string{"conditionA", "conditionB"},
				Clustering:          "none",
				MinAbundance:        10,
				PrimaryFilter:       0.05,
				ReadoutClustering:   "none",
				ReadoutList:         []string{"readoutA", "readoutB"},
				ScoreType:           "lte",
				Type:                "dotplot",
			},
		}

		filter.Process(analysis)
		Expect(analysis).To(Equal(expected))
	})
})
