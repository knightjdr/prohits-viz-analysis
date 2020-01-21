package convert

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse table", func() {
	It("should parse a table and average replicates", func() {
		settings := ConversionSettings{
			ScoreType: "lte",
		}
		table := &[]map[string]string{
			{"condition": "condition1", "readout": "readout1", "abundance": "5|3", "score": "0.01"},
			{"condition": "condition1", "readout": "readout3", "abundance": "10|5", "score": "0.02"},
			{"condition": "condition1", "readout": "readout2", "abundance": "23", "score": "0"},
			{"condition": "condition3", "readout": "readout3", "abundance": "7", "score": "0.01"},
			{"condition": "condition3", "readout": "readout1", "abundance": "14.3", "score": "0.08"},
			{"condition": "condition2", "readout": "readout2", "abundance": "17.8", "score": "0.01"},
			{"condition": "condition2", "readout": "readout1", "abundance": "2", "score": "0.01"},
		}

		expected := &tableData{
			conditions: map[string]int{
				"condition1": 0,
				"condition2": 2,
				"condition3": 1,
			},
			readoutCondition: map[readoutCondition]readoutData{
				readoutCondition{"readout1", "condition1"}: readoutData{4, 0.01},
				readoutCondition{"readout3", "condition1"}: readoutData{7.5, 0.02},
				readoutCondition{"readout2", "condition1"}: readoutData{23, 0},
				readoutCondition{"readout3", "condition3"}: readoutData{7, 0.01},
				readoutCondition{"readout1", "condition3"}: readoutData{14.3, 0.08},
				readoutCondition{"readout2", "condition2"}: readoutData{17.8, 0.01},
				readoutCondition{"readout1", "condition2"}: readoutData{2, 0.01},
			},
			readouts: map[string]int{
				"readout1": 0,
				"readout2": 2,
				"readout3": 1,
			},
			worstScore: 0.08,
		}
		Expect(parseTable(table, settings)).To(Equal(expected))
	})

	It("should parse a table and keep replicates", func() {
		settings := ConversionSettings{
			KeepReps:  true,
			ScoreType: "lte",
		}
		table := &[]map[string]string{
			{"condition": "condition1", "readout": "readout1", "abundance": "5|3", "score": "0.01"},
			{"condition": "condition1", "readout": "readout3", "abundance": "10|5", "score": "0.02"},
			{"condition": "condition1", "readout": "readout2", "abundance": "23", "score": "0"},
			{"condition": "condition3", "readout": "readout3", "abundance": "7", "score": "0.01"},
			{"condition": "condition3", "readout": "readout1", "abundance": "14.3", "score": "0.08"},
			{"condition": "condition2", "readout": "readout2", "abundance": "17.8", "score": "0.01"},
			{"condition": "condition2", "readout": "readout1", "abundance": "2", "score": "0.01"},
		}

		expected := &tableData{
			conditions: map[string]int{
				"condition1R1": 0,
				"condition1R2": 1,
				"condition2R1": 3,
				"condition3R1": 2,
			},
			readoutCondition: map[readoutCondition]readoutData{
				readoutCondition{"readout1", "condition1R1"}: readoutData{5, 0.01},
				readoutCondition{"readout1", "condition1R2"}: readoutData{3, 0.01},
				readoutCondition{"readout3", "condition1R1"}: readoutData{10, 0.02},
				readoutCondition{"readout3", "condition1R2"}: readoutData{5, 0.02},
				readoutCondition{"readout2", "condition1R1"}: readoutData{23, 0},
				readoutCondition{"readout3", "condition3R1"}: readoutData{7, 0.01},
				readoutCondition{"readout1", "condition3R1"}: readoutData{14.3, 0.08},
				readoutCondition{"readout2", "condition2R1"}: readoutData{17.8, 0.01},
				readoutCondition{"readout1", "condition2R1"}: readoutData{2, 0.01},
			},
			readouts: map[string]int{
				"readout1": 0,
				"readout2": 2,
				"readout3": 1,
			},
			worstScore: 0.08,
		}
		Expect(parseTable(table, settings)).To(Equal(expected))
	})
})
