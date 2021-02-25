package scv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Add specificity for SCV", func() {
	It("should add specificity to data", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "conditionA", "readout": "readoutX", "abundance": "10", "score": "0.01"},
				{"condition": "conditionA", "readout": "readoutY", "abundance": "20", "score": "0.01"},
				{"condition": "conditionB", "readout": "readoutX", "abundance": "30", "score": "0"},
				{"condition": "conditionC", "readout": "readoutY", "abundance": "15", "score": "0.02"},
				{"condition": "conditionC", "readout": "readoutZ", "abundance": "25", "score": "0.01"},
			},
			Settings: types.Settings{
				Specificity: true,
			},
		}
		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {},
				"readoutY": {},
			},
			"conditionB": {
				"readoutX": {},
			},
			"conditionC": {
				"readoutY": {},
				"readoutZ": {},
			},
		}

		expected := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {
					"Specificity": 0.67,
				},
				"readoutY": {
					"Specificity": 2.67,
				},
			},
			"conditionB": {
				"readoutX": {
					"Specificity": 6,
				},
			},
			"conditionC": {
				"readoutY": {
					"Specificity": 1.5,
				},
				"readoutZ": {
					"Specificity": 1.5,
				},
			},
		}

		addSpecificity(data, analysis)
		Expect(data).To(Equal(expected))
	})

	It("should not add specificity to data when not requested", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "conditionA", "readout": "readoutX", "abundance": "10", "score": "0.01"},
				{"condition": "conditionA", "readout": "readoutY", "abundance": "20", "score": "0.01"},
				{"condition": "conditionB", "readout": "readoutX", "abundance": "30", "score": "0"},
				{"condition": "conditionC", "readout": "readoutY", "abundance": "15", "score": "0.02"},
				{"condition": "conditionC", "readout": "readoutZ", "abundance": "25", "score": "0.01"},
			},
			Settings: types.Settings{
				Specificity: false,
			},
		}
		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {},
				"readoutY": {},
			},
			"conditionB": {
				"readoutX": {},
			},
			"conditionC": {
				"readoutY": {},
				"readoutZ": {},
			},
		}

		expected := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {},
				"readoutY": {},
			},
			"conditionB": {
				"readoutX": {},
			},
			"conditionC": {
				"readoutY": {},
				"readoutZ": {},
			},
		}

		addSpecificity(data, analysis)
		Expect(data).To(Equal(expected))
	})

	It("should not add specificity to data when there is a single condition", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "conditionA", "readout": "readoutX", "abundance": "10", "score": "0.01"},
				{"condition": "conditionA", "readout": "readoutY", "abundance": "20", "score": "0.01"},
			},
			Settings: types.Settings{
				Specificity: true,
			},
		}
		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {},
				"readoutY": {},
			},
		}

		expected := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {},
				"readoutY": {},
			},
		}

		addSpecificity(data, analysis)
		Expect(data).To(Equal(expected))
	})
})
