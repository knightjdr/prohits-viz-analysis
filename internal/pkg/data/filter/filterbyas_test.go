package filter

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter by abundance and score", func() {
	It("should filter analysis data by abundance and score, keeping all readouts that pass at least once", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				map[string]string{"abundance": "10", "readout": "readoutA", "score": "0.05"},
				map[string]string{"abundance": "11", "readout": "readoutB", "score": "0.04"},
				map[string]string{"abundance": "11", "readout": "readoutC", "score": "0.06"},
				map[string]string{"abundance": "9", "readout": "readoutD", "score": "0.04"},
				map[string]string{"abundance": "10", "readout": "readoutE", "score": "0.05"},
				map[string]string{"abundance": "9", "readout": "readoutE", "score": "0.06"},
			},
			Settings: types.Settings{
				MinAbundance:  10,
				PrimaryFilter: 0.05,
				ScoreType:     "lte",
			},
		}

		expected := []map[string]string{
			map[string]string{"abundance": "10", "readout": "readoutA", "score": "0.05"},
			map[string]string{"abundance": "11", "readout": "readoutB", "score": "0.04"},
			map[string]string{"abundance": "10", "readout": "readoutE", "score": "0.05"},
			map[string]string{"abundance": "9", "readout": "readoutE", "score": "0.06"},
		}

		filterByAbundanceAndScore(analysis)
		Expect(analysis.Data).To(Equal(expected))
	})
})

var _ = Describe("Parse score", func() {
	It("should parse a float", func() {
		score := "0.1"
		expected := 0.1
		Expect(parseScore(score)).To(Equal(expected))
	})

	It("should return 0 when score cannot be parsed", func() {
		score := "a"
		expected := float64(0)
		Expect(parseScore(score)).To(Equal(expected))
	})
})

var _ = Describe("Remove readouts filtered for abundance and score", func() {
	It("should should remove entries with readout not in list", func() {
		analysis := &types.Analysis{
			Data: []map[string]string{
				map[string]string{"condition": "conditionA", "readout": "readoutA"},
				map[string]string{"condition": "conditionB", "readout": "readoutB"},
				map[string]string{"condition": "conditionA", "readout": "readoutC"},
				map[string]string{"condition": "conditionB", "readout": "readoutC"},
				map[string]string{"condition": "conditionC", "readout": "readoutA"},
				map[string]string{"condition": "conditionC", "readout": "readoutB"},
			},
		}
		passingReadouts := map[string]bool{
			"readoutA": true,
			"readoutB": true,
		}

		expected := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA"},
			map[string]string{"condition": "conditionB", "readout": "readoutB"},
			map[string]string{"condition": "conditionC", "readout": "readoutA"},
			map[string]string{"condition": "conditionC", "readout": "readoutB"},
		}

		removeReadouts(analysis, passingReadouts)
		Expect(analysis.Data).To(Equal(expected))
	})
})
