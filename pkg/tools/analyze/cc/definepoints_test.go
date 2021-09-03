package cc

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Define points", func() {
	It("should return an array of points for scatter plot with gte score type", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "conditionA", "readout": "readoutA", "abundance": "1", "score": "0.05"},
				{"condition": "conditionA", "readout": "readoutC", "abundance": "2", "score": "0.03"},
				{"condition": "conditionC", "readout": "readoutA", "abundance": "3", "score": "0.01"},
				{"condition": "conditionC", "readout": "readoutB", "abundance": "4", "score": "0.02"},
			},
			Settings: types.Settings{
				ConditionX:      "conditionA",
				ConditionY:      "conditionC",
				PrimaryFilter:   0.03,
				ScoreType:       "gte",
				SecondaryFilter: 0.01,
				Type:            "condition-condition",
			},
		}

		expected := []types.ScatterPoint{
			{Label: "readoutA", X: 1, Y: 3, Color: "#0066cc"},
			{Label: "readoutB", X: 0, Y: 4, Color: "#99ccff"},
			{Label: "readoutC", X: 2, Y: 0, Color: "#0066cc"},
		}
		Expect(definePoints(analysis)).To(Equal(expected))
	})

	It("should return an array of points for scatter plot lte score type", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "conditionA", "readout": "readoutA", "abundance": "1", "score": "0.05"},
				{"condition": "conditionA", "readout": "readoutC", "abundance": "2", "score": "0.03"},
				{"condition": "conditionC", "readout": "readoutA", "abundance": "3", "score": "0.01"},
				{"condition": "conditionC", "readout": "readoutB", "abundance": "4", "score": "0.02"},
			},
			Settings: types.Settings{
				ConditionX:      "conditionA",
				ConditionY:      "conditionC",
				PrimaryFilter:   0.01,
				ScoreType:       "lte",
				SecondaryFilter: 0.05,
				Type:            "condition-condition",
			},
		}

		expected := []types.ScatterPoint{
			{Label: "readoutA", X: 1, Y: 3, Color: "#0066cc"},
			{Label: "readoutB", X: 0, Y: 4, Color: "#99ccff"},
			{Label: "readoutC", X: 2, Y: 0, Color: "#99ccff"},
		}
		Expect(definePoints(analysis)).To(Equal(expected))
	})

	It("should return an array of points for scatter plot including positive and negative abundances", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "conditionA", "readout": "readoutA", "abundance": "1", "score": "0.05"},
				{"condition": "conditionA", "readout": "readoutC", "abundance": "-2", "score": "0.03"},
				{"condition": "conditionC", "readout": "readoutA", "abundance": "-3", "score": "0.01"},
				{"condition": "conditionC", "readout": "readoutB", "abundance": "4", "score": "0.02"},
			},
			Settings: types.Settings{
				ConditionX:      "conditionA",
				ConditionY:      "conditionC",
				PrimaryFilter:   0.01,
				ScoreType:       "lte",
				SecondaryFilter: 0.05,
				Type:            "condition-condition",
			},
		}

		expected := []types.ScatterPoint{
			{Label: "readoutA", X: 1, Y: -3, Color: "#0066cc"},
			{Label: "readoutB", X: 0, Y: 4, Color: "#99ccff"},
			{Label: "readoutC", X: -2, Y: 0, Color: "#99ccff"},
		}
		Expect(definePoints(analysis)).To(Equal(expected))
	})
})
