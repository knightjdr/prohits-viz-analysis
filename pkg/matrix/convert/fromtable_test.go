package convert_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/matrix/convert"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Convert from table", func() {
	It("should generate matrices from a table", func() {
		conversionSettings := ConversionSettings{
			CalculateRatios: true,
			Resort:          true,
			ScoreType:       "lte",
		}
		table := &[]map[string]string{
			{"condition": "condition1", "readout": "readout1", "abundance": "5", "score": "0.01"},
			{"condition": "condition1", "readout": "readout3", "abundance": "10", "score": "0.02"},
			{"condition": "condition1", "readout": "readout2", "abundance": "23", "score": "0"},
			{"condition": "condition3", "readout": "readout3", "abundance": "7", "score": "0.01"},
			{"condition": "condition3", "readout": "readout1", "abundance": "14.3", "score": "0.08"},
			{"condition": "condition2", "readout": "readout2", "abundance": "17.8", "score": "0.01"},
			{"condition": "condition2", "readout": "readout1", "abundance": "2", "score": "0.01"},
		}

		expected := &types.Matrices{
			Abundance: [][]float64{
				{5, 2, 14.3},
				{23, 17.8, 0},
				{10, 0, 7},
			},
			Conditions: []string{"condition1", "condition2", "condition3"},
			Ratio: [][]float64{
				{0.35, 0.14, 1},
				{1, 0.77, 0},
				{1, 0, 0.7},
			},
			Readouts: []string{"readout1", "readout2", "readout3"},
			Score: [][]float64{
				{0.01, 0.01, 0.08},
				{0, 0.01, 0.08},
				{0.02, 0.08, 0.01},
			},
		}

		actual := FromTable(table, conversionSettings)
		Expect(actual.Conditions).To(Equal(expected.Conditions), "should return conditions")
		Expect(actual.Readouts).To(Equal(expected.Readouts), "should return readouts")
		for i, row := range actual.Abundance {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected.Abundance[i][j], 0.01), "should return abundance matrix")
			}
		}
		for i, row := range actual.Ratio {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected.Ratio[i][j], 0.01), "should return ratio matrix")
			}
		}
		for i, row := range actual.Score {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected.Score[i][j], 0.01), "should return score matrix")
			}
		}
	})
})
