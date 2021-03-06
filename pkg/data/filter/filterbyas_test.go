package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter by abundance and score", func() {
	It("should filter analysis data by abundance and score, keeping all readouts that pass at least once", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				{"abundance": "10", "readout": "readoutA", "score": "0.05"},
				{"abundance": "11", "readout": "readoutB", "score": "0.04"},
				{"abundance": "11", "readout": "readoutC", "score": "0.06"},
				{"abundance": "9", "readout": "readoutD", "score": "0.04"},
				{"abundance": "10", "readout": "readoutE", "score": "0.05"},
				{"abundance": "9", "readout": "readoutE", "score": "0.06"},
			},
			Settings: types.Settings{
				MinAbundance:  10,
				PrimaryFilter: 0.05,
				ScoreType:     "lte",
			},
		}

		expected := []map[string]string{
			{"abundance": "10", "readout": "readoutA", "score": "0.05"},
			{"abundance": "11", "readout": "readoutB", "score": "0.04"},
			{"abundance": "10", "readout": "readoutE", "score": "0.05"},
			{"abundance": "9", "readout": "readoutE", "score": "0.06"},
		}

		byAbundanceAndScore(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})

var _ = Describe("Add readout", func() {
	It("should add a readout that is not already a key", func() {
		passingReadouts := map[string]map[string]bool{}
		row := map[string]string{
			"condition": "conditionA",
			"readout":   "readoutA",
		}

		expected := map[string]map[string]bool{
			"readoutA": {
				"conditionA": true,
			},
		}
		addReadout(&passingReadouts, row)
		Expect(passingReadouts).To(Equal(expected))
	})

	It("should add a readout that is already a key", func() {
		passingReadouts := map[string]map[string]bool{
			"readoutA": {
				"conditionA": true,
			},
		}
		row := map[string]string{
			"condition": "conditionB",
			"readout":   "readoutA",
		}

		expected := map[string]map[string]bool{
			"readoutA": {
				"conditionA": true,
				"conditionB": true,
			},
		}
		addReadout(&passingReadouts, row)
		Expect(passingReadouts).To(Equal(expected))
	})
})

var _ = Describe("Filter readouts for abundance and score", func() {
	It("should should remove entries with readout not in list", func() {
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
				MinConditions:                1,
				ParsimoniousReadoutFiltering: true,
			},
		}
		passingReadouts := map[string]map[string]bool{
			"readoutA": {
				"conditionA": true,
				"conditionB": true,
			},
			"readoutB": {
				"conditionC": true,
			},
		}

		expected := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA"},
			{"condition": "conditionC", "readout": "readoutB"},
		}

		filterFailingReadouts(analysis, passingReadouts)
		Expect(analysis.Data).To(Equal(expected))
	})
})

var _ = Describe("Filter by readout", func() {
	It("should return a function when ParsimoniousReadoutFiltering is true", func() {
		passingReadouts := map[string]map[string]bool{
			"readoutA": {
				"conditionA": true,
				"conditionB": true,
			},
			"readoutB": {
				"conditionA": true,
			},
		}
		settings := types.Settings{
			MinConditions:                2,
			ParsimoniousReadoutFiltering: true,
		}

		shouldRemoveReadout := filterByReadout(passingReadouts, settings)
		Expect(shouldRemoveReadout("conditionA", "readoutA")).To(BeFalse(), "should not remove readout")
		Expect(shouldRemoveReadout("conditionC", "readoutA")).To(BeTrue(), "should remove readout not passing for specific condition")
		Expect(shouldRemoveReadout("conditionA", "readoutB")).To(BeTrue(), "should remove readout not passing min condition filter")
	})

	It("should return a function when ParsimoniousReadoutFiltering is false", func() {
		passingReadouts := map[string]map[string]bool{
			"readoutA": {
				"conditionA": true,
				"conditionB": true,
			},
			"readoutB": {
				"conditionA": true,
			},
		}
		settings := types.Settings{
			MinConditions:                2,
			ParsimoniousReadoutFiltering: false,
		}

		shouldRemoveReadout := filterByReadout(passingReadouts, settings)
		Expect(shouldRemoveReadout("conditionA", "readoutA")).To(BeFalse(), "should not remove readout passing for specific condition")
		Expect(shouldRemoveReadout("conditionC", "readoutA")).To(BeFalse(), "should not remove readout not passing for specific condition")
		Expect(shouldRemoveReadout("conditionA", "readoutB")).To(BeTrue(), "should remove readout not passing min condition filter")
	})
})
