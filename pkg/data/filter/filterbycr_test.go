package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter by conditions and/or readouts", func() {
	It("should filter out entries not satisfying both requested conditions and readouts", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "conditionA", "readout": "readoutA"},
				{"condition": "conditionB", "readout": "readoutB"},
				{"condition": "conditionA", "readout": "readoutC"},
				{"condition": "conditionB", "readout": "readoutC"},
				{"condition": "conditionC", "readout": "readoutA"},
				{"condition": "conditionC", "readout": "readoutB"},
			},
			Settings: types.Settings{
				ConditionX: "conditionA",
				ConditionY: "conditionC",
				Type:       "condition-condition",
			},
		}

		expected := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA"},
			{"condition": "conditionA", "readout": "readoutC"},
			{"condition": "conditionC", "readout": "readoutA"},
			{"condition": "conditionC", "readout": "readoutB"},
		}

		byConditionsAndReadouts(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})

	Describe("Dotplot", func() {
		It("should filter out entries not satisfying both requested conditions and readouts", func() {
			analysis := &types.Analysis{
				Data: []map[string]string{
					{"condition": "conditionA", "readout": "readoutA"},
					{"condition": "conditionB", "readout": "readoutB"},
					{"condition": "conditionA", "readout": "readoutC"},
					{"condition": "conditionB", "readout": "readoutC"},
					{"condition": "conditionC", "readout": "readoutA"},
					{"condition": "conditionC", "readout": "readoutB"},
				},
				Settings: types.Settings{
					Clustering:          "none",
					ConditionClustering: "none",
					ConditionList:       []string{"conditionA", "conditionB"},
					ReadoutClustering:   "none",
					ReadoutList:         []string{"readoutA", "readoutB"},
					Type:                "dotplot",
				},
			}

			expected := []map[string]string{
				{"condition": "conditionA", "readout": "readoutA"},
				{"condition": "conditionB", "readout": "readoutB"},
			}

			byConditionsAndReadouts(analysis)
			Expect(analysis.Data).To(Equal(expected))
		})

		It("should filter by conditions only", func() {
			analysis := &types.Analysis{
				Data: []map[string]string{
					{"condition": "conditionA", "readout": "readoutA"},
					{"condition": "conditionB", "readout": "readoutB"},
					{"condition": "conditionA", "readout": "readoutC"},
					{"condition": "conditionB", "readout": "readoutC"},
					{"condition": "conditionC", "readout": "readoutA"},
					{"condition": "conditionC", "readout": "readoutB"},
				},
				Settings: types.Settings{
					Clustering:          "none",
					ConditionClustering: "none",
					ConditionList:       []string{"conditionA", "conditionB"},
					ReadoutClustering:   "hierarchical",
					ReadoutList:         []string{"readoutA", "readoutB"},
					Type:                "dotplot",
				},
			}

			expected := []map[string]string{
				{"condition": "conditionA", "readout": "readoutA"},
				{"condition": "conditionB", "readout": "readoutB"},
				{"condition": "conditionA", "readout": "readoutC"},
				{"condition": "conditionB", "readout": "readoutC"},
			}

			byConditionsAndReadouts(analysis)
			Expect(analysis.Data).To(Equal(expected))
		})

		It("should filter by readouts only", func() {
			analysis := &types.Analysis{
				Data: []map[string]string{
					{"condition": "conditionA", "readout": "readoutA"},
					{"condition": "conditionB", "readout": "readoutB"},
					{"condition": "conditionA", "readout": "readoutC"},
					{"condition": "conditionB", "readout": "readoutC"},
					{"condition": "conditionC", "readout": "readoutA"},
					{"condition": "conditionC", "readout": "readoutB"},
				},
				Settings: types.Settings{
					Clustering:          "none",
					ConditionClustering: "hierarchical",
					ConditionList:       []string{"conditionA", "conditionB"},
					ReadoutClustering:   "none",
					ReadoutList:         []string{"readoutA", "readoutB"},
					Type:                "dotplot",
				},
			}

			expected := []map[string]string{
				{"condition": "conditionA", "readout": "readoutA"},
				{"condition": "conditionB", "readout": "readoutB"},
				{"condition": "conditionC", "readout": "readoutA"},
				{"condition": "conditionC", "readout": "readoutB"},
			}

			byConditionsAndReadouts(analysis)
			Expect(analysis.Data).To(Equal(expected))
		})
	})
})
