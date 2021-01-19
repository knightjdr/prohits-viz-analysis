package specificity

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type specificityTest struct {
	condition string
	values    map[string]map[string]float64
}

var _ = Describe("Calculate specificity", func() {
	It("should return specificity indexed by condition and readout", func() {
		analysis := types.Analysis{
			Data: []map[string]string{
				{"condition": "a", "readout": "x", "abundance": "10", "score": "0.01"},
				{"condition": "a", "readout": "y", "abundance": "20", "score": "0.01"},
				{"condition": "b", "readout": "x", "abundance": "30", "score": "0"},
				{"condition": "c", "readout": "y", "abundance": "15", "score": "0.02"},
				{"condition": "c", "readout": "z", "abundance": "25", "score": "0.01"},
			},
			Settings: types.Settings{
				SpecificityMetric: "fe",
			},
		}

		expected := map[string]map[string]map[string]float64{
			"a": {
				"x": {"abundance": 10, "score": 0.01, "specificity": 0.67},
				"y": {"abundance": 20, "score": 0.01, "specificity": 2.67},
			},
			"b": {
				"x": {"abundance": 30, "score": 0, "specificity": 6},
			},
			"c": {
				"y": {"abundance": 15, "score": 0.02, "specificity": 1.5},
				"z": {"abundance": 25, "score": 0.01, "specificity": math.Inf(1)},
			},
		}

		Expect(calculateSpecificity(&analysis)).To(Equal(expected))
	})

	Describe("get abundance by readout", func() {
		It("should get abundance and reproducibilty by readout and condition", func() {
			data := []map[string]string{
				{"condition": "a", "readout": "x", "abundance": "10", "score": "0.01"},
				{"condition": "a", "readout": "y", "abundance": "20", "score": "0.01"},
				{"condition": "b", "readout": "x", "abundance": "30", "score": "0"},
				{"condition": "c", "readout": "y", "abundance": "15", "score": "0.02"},
				{"condition": "c", "readout": "z", "abundance": "25", "score": "0.01"},
			}

			expectedAbundanceByReadout := map[string]map[string]map[string]float64{
				"x": {
					"a": {"abundance": 10, "reproducibility": 1, "score": 0.01},
					"b": {"abundance": 30, "reproducibility": 1, "score": 0},
				},
				"y": {
					"a": {"abundance": 20, "reproducibility": 1, "score": 0.01},
					"c": {"abundance": 15, "reproducibility": 1, "score": 0.02},
				},
				"z": {
					"c": {"abundance": 25, "reproducibility": 1, "score": 0.01},
				},
			}
			expectedNoCondition := 3

			actualAbundanceByReadout, actualNoCondition := getAbundanceByReadout(data)
			Expect(actualAbundanceByReadout).To(Equal(expectedAbundanceByReadout), "should return abundance by readout and condition")
			Expect(actualNoCondition).To(Equal(expectedNoCondition), "should return the number of conditions")

		})

		It("should get abundance and reproducibilty by readout and condition with a pipe-separated list", func() {
			data := []map[string]string{
				{"condition": "a", "readout": "x", "abundance": "10|0", "score": "0.01"},
				{"condition": "a", "readout": "y", "abundance": "20|2", "score": "0.01"},
				{"condition": "b", "readout": "x", "abundance": "0|30", "score": "0"},
				{"condition": "c", "readout": "y", "abundance": "15|15", "score": "0.02"},
				{"condition": "c", "readout": "z", "abundance": "25|20", "score": "0.01"},
			}

			expectedAbundanceByReadout := map[string]map[string]map[string]float64{
				"x": {
					"a": {"abundance": 5, "reproducibility": 1, "score": 0.01},
					"b": {"abundance": 15, "reproducibility": 1, "score": 0},
				},
				"y": {
					"a": {"abundance": 11, "reproducibility": 2, "score": 0.01},
					"c": {"abundance": 15, "reproducibility": 2, "score": 0.02},
				},
				"z": {
					"c": {"abundance": 22.5, "reproducibility": 2, "score": 0.01},
				},
			}
			expectedNoCondition := 3

			actualAbundanceByReadout, actualNoCondition := getAbundanceByReadout(data)
			Expect(actualAbundanceByReadout).To(Equal(expectedAbundanceByReadout), "should return abundance by readout and condition")
			Expect(actualNoCondition).To(Equal(expectedNoCondition), "should return the number of conditions")
		})
	})

	Describe("using metric", func() {
		noConditions := 4
		tests := []specificityTest{
			{
				condition: "a",
				values: map[string]map[string]float64{
					"a": {
						"abundance":       10,
						"reproducibility": 2,
						"score":           0.01,
					},
				},
			},
			{
				condition: "a",
				values: map[string]map[string]float64{
					"a": {
						"abundance":       10,
						"reproducibility": 2,
						"score":           0.01,
					},
					"b": {"abundance": 20},
					"c": {"abundance": 30},
					"d": {"abundance": 50},
				},
			},
			{
				condition: "a",
				values: map[string]map[string]float64{
					"a": {
						"abundance":       100,
						"reproducibility": 2,
						"score":           0.01,
					},
					"b": {"abundance": 20},
					"d": {"abundance": 50},
				},
			},
		}

		It("should calculate specificity for fe (fold enrichment)", func() {
			metric := "fe"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []map[string]float64{
				{"abundance": 10, "score": 0.01, "specificity": math.Inf(1)},
				{"abundance": 10, "score": 0.01, "specificity": 0.3},
				{"abundance": 100, "score": 0.01, "specificity": 4.29},
			}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})

		It("should calculate specificity for zscore", func() {
			metric := "zscore"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []map[string]float64{
				{"abundance": 10, "score": 0.01, "specificity": 1.5},
				{"abundance": 10, "score": 0.01, "specificity": -1.02},
				{"abundance": 100, "score": 0.01, "specificity": 1.32},
			}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})

		It("should calculate specificity for sscore", func() {
			metric := "sscore"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []map[string]float64{
				{"abundance": 10, "score": 0.01, "specificity": 6.32},
				{"abundance": 10, "score": 0.01, "specificity": 3.16},
				{"abundance": 100, "score": 0.01, "specificity": 11.55},
			}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})

		It("should calculate specificity for dscore", func() {
			metric := "dscore"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []map[string]float64{
				{"abundance": 10, "score": 0.01, "specificity": 12.65},
				{"abundance": 10, "score": 0.01, "specificity": 3.16},
				{"abundance": 100, "score": 0.01, "specificity": 13.33},
			}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})

		It("should calculate specificity for wdscore", func() {
			metric := "wdscore"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []map[string]float64{
				{"abundance": 10, "score": 0.01, "specificity": 25.30},
				{"abundance": 10, "score": 0.01, "specificity": 3.16},
				{"abundance": 100, "score": 0.01, "specificity": 13.64},
			}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})
	})
})
