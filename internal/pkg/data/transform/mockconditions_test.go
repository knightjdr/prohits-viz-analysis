package transform

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mock condition abundances", func() {
	It("should mock abundances when missing", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "12", "condition": "conditionA", "readout": "readoutA", "score": "0.05"},
				{"abundance": "1", "condition": "conditionA", "readout": "readoutB", "score": "0.01"},
				{"abundance": "10", "condition": "conditionA", "readout": "readoutC", "score": "0.01"},
				{"abundance": "10|5", "condition": "conditionB", "readout": "readoutA", "score": "0.01"},
				{"abundance": "10|5|2.5", "condition": "conditionB", "readout": "readoutB", "score": "0.01"},
				{"abundance": "10|8", "condition": "conditionB", "readout": "readoutC", "score": "0.01"},
				{"abundance": "10", "condition": "conditionC", "readout": "readoutA", "score": "0.01"},
				{"abundance": "1", "condition": "conditionC", "readout": "conditionC", "score": "0.01"},
				{"abundance": "12", "condition": "conditionC", "readout": "readoutC", "score": "0.01"},
			},
			Settings: types.Settings{
				MockConditionAbundance: true,
				PrimaryFilter:          0.01,
				ScoreType:              "lte",
			},
		}

		expected := []map[string]string{
			{"abundance": "12", "condition": "conditionA", "readout": "readoutA", "score": "0.05"},
			{"abundance": "1", "condition": "conditionA", "readout": "readoutB", "score": "0.01"},
			{"abundance": "10", "condition": "conditionA", "readout": "readoutC", "score": "0.01"},
			{"abundance": "10|5", "condition": "conditionB", "readout": "readoutA", "score": "0.01"},
			{"abundance": "10|5|2.5", "condition": "conditionB", "readout": "readoutB", "score": "0.01"},
			{"abundance": "10|8", "condition": "conditionB", "readout": "readoutC", "score": "0.01"},
			{"abundance": "10", "condition": "conditionC", "readout": "readoutA", "score": "0.01"},
			{"abundance": "1", "condition": "conditionC", "readout": "conditionC", "score": "0.01"},
			{"abundance": "12", "condition": "conditionC", "readout": "readoutC", "score": "0.01"},
			{"abundance": "10", "condition": "conditionA", "readout": "conditionA", "score": "0.01"},
			{"abundance": "9", "condition": "conditionB", "readout": "conditionB", "score": "0.01"},
		}
		mockConditionAbundance(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})

	It("should not mock abundances", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "12", "condition": "conditionA", "readout": "readoutA", "score": "0.05"},
				{"abundance": "1", "condition": "conditionA", "readout": "readoutB", "score": "0.01"},
				{"abundance": "10", "condition": "conditionA", "readout": "readoutC", "score": "0.01"},
				{"abundance": "10|5", "condition": "conditionB", "readout": "readoutA", "score": "0.01"},
				{"abundance": "10|5|2.5", "condition": "conditionB", "readout": "readoutB", "score": "0.01"},
				{"abundance": "10|8", "condition": "conditionB", "readout": "readoutC", "score": "0.01"},
				{"abundance": "10", "condition": "conditionC", "readout": "readoutA", "score": "0.01"},
				{"abundance": "1", "condition": "conditionC", "readout": "conditionC", "score": "0.01"},
				{"abundance": "12", "condition": "conditionC", "readout": "readoutC", "score": "0.01"},
			},
			Settings: types.Settings{
				MockConditionAbundance: false,
				PrimaryFilter:          0.01,
				ScoreType:              "lte",
			},
		}

		expected := []map[string]string{
			{"abundance": "12", "condition": "conditionA", "readout": "readoutA", "score": "0.05"},
			{"abundance": "1", "condition": "conditionA", "readout": "readoutB", "score": "0.01"},
			{"abundance": "10", "condition": "conditionA", "readout": "readoutC", "score": "0.01"},
			{"abundance": "10|5", "condition": "conditionB", "readout": "readoutA", "score": "0.01"},
			{"abundance": "10|5|2.5", "condition": "conditionB", "readout": "readoutB", "score": "0.01"},
			{"abundance": "10|8", "condition": "conditionB", "readout": "readoutC", "score": "0.01"},
			{"abundance": "10", "condition": "conditionC", "readout": "readoutA", "score": "0.01"},
			{"abundance": "1", "condition": "conditionC", "readout": "conditionC", "score": "0.01"},
			{"abundance": "12", "condition": "conditionC", "readout": "readoutC", "score": "0.01"},
		}
		mockConditionAbundance(analysis)

		Expect(analysis.Data).To(Equal(expected))
	})
})
