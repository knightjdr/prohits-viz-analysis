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
				{"readout1", "condition1"}: {4, 0.01},
				{"readout3", "condition1"}: {7.5, 0.02},
				{"readout2", "condition1"}: {23, 0},
				{"readout3", "condition3"}: {7, 0.01},
				{"readout1", "condition3"}: {14.3, 0.08},
				{"readout2", "condition2"}: {17.8, 0.01},
				{"readout1", "condition2"}: {2, 0.01},
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

	It("should parse a table and keep replicates for readouts (i.e. rename conditions)", func() {
		settings := ConversionSettings{
			KeepReps:       true,
			ScoreType:      "lte",
			SeparateRepsBy: "column",
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
				"condition1_R1": 0,
				"condition1_R2": 1,
				"condition2_R1": 3,
				"condition3_R1": 2,
			},
			readoutCondition: map[readoutCondition]readoutData{
				{"readout1", "condition1_R1"}: {5, 0.01},
				{"readout1", "condition1_R2"}: {3, 0.01},
				{"readout3", "condition1_R1"}: {10, 0.02},
				{"readout3", "condition1_R2"}: {5, 0.02},
				{"readout2", "condition1_R1"}: {23, 0},
				{"readout3", "condition3_R1"}: {7, 0.01},
				{"readout1", "condition3_R1"}: {14.3, 0.08},
				{"readout2", "condition2_R1"}: {17.8, 0.01},
				{"readout1", "condition2_R1"}: {2, 0.01},
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

	It("should parse a table and keep replicates for conditions (i.e. rename readouts)", func() {
		settings := ConversionSettings{
			KeepReps:       true,
			ScoreType:      "lte",
			SeparateRepsBy: "row",
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
				{"readout1_R1", "condition1"}: {5, 0.01},
				{"readout1_R2", "condition1"}: {3, 0.01},
				{"readout3_R1", "condition1"}: {10, 0.02},
				{"readout3_R2", "condition1"}: {5, 0.02},
				{"readout2_R1", "condition1"}: {23, 0},
				{"readout3_R1", "condition3"}: {7, 0.01},
				{"readout1_R1", "condition3"}: {14.3, 0.08},
				{"readout2_R1", "condition2"}: {17.8, 0.01},
				{"readout1_R1", "condition2"}: {2, 0.01},
			},
			readouts: map[string]int{
				"readout1_R1": 0,
				"readout1_R2": 1,
				"readout3_R1": 2,
				"readout3_R2": 3,
				"readout2_R1": 4,
			},
			worstScore: 0.08,
		}
		Expect(parseTable(table, settings)).To(Equal(expected))
	})
})
