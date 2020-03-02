package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter by conditions and readouts", func() {
	It("should filter out entries not satisfying both requested conditions and readouts", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				map[string]string{"condition": "conditionA", "readout": "readoutA"},
				map[string]string{"condition": "conditionB", "readout": "readoutB"},
				map[string]string{"condition": "conditionA", "readout": "readoutC"},
				map[string]string{"condition": "conditionB", "readout": "readoutC"},
				map[string]string{"condition": "conditionC", "readout": "readoutA"},
				map[string]string{"condition": "conditionC", "readout": "readoutB"},
			},
			Settings: types.Settings{
				ConditionClustering: "none",
				ConditionList:       []string{"conditionA", "conditionB"},
				ReadoutClustering:   "none",
				ReadoutList:         []string{"readoutA", "readoutB"},
			},
		}

		expected := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA"},
			map[string]string{"condition": "conditionB", "readout": "readoutB"},
		}

		filterByConditionsAndReadouts(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})

	It("should filter by conditions only", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				map[string]string{"condition": "conditionA", "readout": "readoutA"},
				map[string]string{"condition": "conditionB", "readout": "readoutB"},
				map[string]string{"condition": "conditionA", "readout": "readoutC"},
				map[string]string{"condition": "conditionB", "readout": "readoutC"},
				map[string]string{"condition": "conditionC", "readout": "readoutA"},
				map[string]string{"condition": "conditionC", "readout": "readoutB"},
			},
			Settings: types.Settings{
				ConditionClustering: "none",
				ConditionList:       []string{"conditionA", "conditionB"},
				ReadoutClustering:   "hierarchical",
				ReadoutList:         []string{"readoutA", "readoutB"},
			},
		}

		expected := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA"},
			map[string]string{"condition": "conditionB", "readout": "readoutB"},
			map[string]string{"condition": "conditionA", "readout": "readoutC"},
			map[string]string{"condition": "conditionB", "readout": "readoutC"},
		}

		filterByConditionsAndReadouts(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})

	It("should filter by readouts only", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				map[string]string{"condition": "conditionA", "readout": "readoutA"},
				map[string]string{"condition": "conditionB", "readout": "readoutB"},
				map[string]string{"condition": "conditionA", "readout": "readoutC"},
				map[string]string{"condition": "conditionB", "readout": "readoutC"},
				map[string]string{"condition": "conditionC", "readout": "readoutA"},
				map[string]string{"condition": "conditionC", "readout": "readoutB"},
			},
			Settings: types.Settings{
				ConditionClustering: "hierarchical",
				ConditionList:       []string{"conditionA", "conditionB"},
				ReadoutClustering:   "none",
				ReadoutList:         []string{"readoutA", "readoutB"},
			},
		}

		expected := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA"},
			map[string]string{"condition": "conditionB", "readout": "readoutB"},
			map[string]string{"condition": "conditionC", "readout": "readoutA"},
			map[string]string{"condition": "conditionC", "readout": "readoutB"},
		}

		filterByConditionsAndReadouts(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})
