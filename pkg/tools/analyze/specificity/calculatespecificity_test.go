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
	It("should return specificity indexed by readout and condition", func() {
		analysis := types.Analysis{
			Data: []map[string]string{
				{"condition": "a", "readout": "x", "abundance": "10"},
				{"condition": "a", "readout": "y", "abundance": "20"},
				{"condition": "b", "readout": "x", "abundance": "30"},
				{"condition": "c", "readout": "y", "abundance": "15"},
				{"condition": "c", "readout": "z", "abundance": "25"},
			},
			Settings: types.Settings{
				SpecificityMetric: "fe",
			},
		}

		expected := map[string]map[string]float64{
			"x": {
				"a": 0.67,
				"b": 6,
			},
			"y": {
				"a": 2.67,
				"c": 1.5,
			},
			"z": {
				"c": math.Inf(1),
			},
		}
		Expect(calculateSpecificity(&analysis)).To(Equal(expected))

	})

	Describe("get abundance by readout", func() {
		It("should get abundance and reproducibilty by readout and condition", func() {
			data := []map[string]string{
				{"condition": "a", "readout": "x", "abundance": "10"},
				{"condition": "a", "readout": "y", "abundance": "20"},
				{"condition": "b", "readout": "x", "abundance": "30"},
				{"condition": "c", "readout": "y", "abundance": "15"},
				{"condition": "c", "readout": "z", "abundance": "25"},
			}

			expectedAbundanceByReadout := map[string]map[string]map[string]float64{
				"x": {
					"a": {"abundance": 10, "reproducibility": 1},
					"b": {"abundance": 30, "reproducibility": 1},
				},
				"y": {
					"a": {"abundance": 20, "reproducibility": 1},
					"c": {"abundance": 15, "reproducibility": 1},
				},
				"z": {
					"c": {"abundance": 25, "reproducibility": 1},
				},
			}
			expectedNoCondition := 3

			actualAbundanceByReadout, actualNoCondition := getAbundanceByReadout(data)
			Expect(actualAbundanceByReadout).To(Equal(expectedAbundanceByReadout), "should return abundance by readout and condition")
			Expect(actualNoCondition).To(Equal(expectedNoCondition), "should return the number of conditions")

		})

		It("should get abundance and reproducibilty by readout and condition with a pipe-separated list", func() {
			data := []map[string]string{
				{"condition": "a", "readout": "x", "abundance": "10|0"},
				{"condition": "a", "readout": "y", "abundance": "20|2"},
				{"condition": "b", "readout": "x", "abundance": "0|30"},
				{"condition": "c", "readout": "y", "abundance": "15|15"},
				{"condition": "c", "readout": "z", "abundance": "25|20"},
			}

			expectedAbundanceByReadout := map[string]map[string]map[string]float64{
				"x": {
					"a": {"abundance": 5, "reproducibility": 1},
					"b": {"abundance": 15, "reproducibility": 1},
				},
				"y": {
					"a": {"abundance": 11, "reproducibility": 2},
					"c": {"abundance": 15, "reproducibility": 2},
				},
				"z": {
					"c": {"abundance": 22.5, "reproducibility": 2},
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
					},
				},
			},
			{
				condition: "a",
				values: map[string]map[string]float64{
					"a": {
						"abundance":       10,
						"reproducibility": 2,
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
					},
					"b": {"abundance": 20},
					"d": {"abundance": 50},
				},
			},
		}

		It("should calculate specificity for fe (fold enrichment)", func() {
			metric := "fe"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []float64{math.Inf(1), 0.3, 4.29}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})

		It("should calculate specificity for zscore", func() {
			metric := "zscore"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []float64{1.5, -1.02, 1.32}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})

		It("should calculate specificity for sscore", func() {
			metric := "sscore"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []float64{6.32, 3.16, 11.55}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})

		It("should calculate specificity for dscore", func() {
			metric := "dscore"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []float64{12.65, 3.16, 13.33}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})

		It("should calculate specificity for wdscore", func() {
			metric := "wdscore"
			definespecificity := getSpecificityMetric(metric, noConditions)

			expected := []float64{25.30, 3.16, 13.64}
			for i, test := range tests {
				Expect(definespecificity(test.condition, test.values)).To(Equal(expected[i]))
			}
		})
	})
})
