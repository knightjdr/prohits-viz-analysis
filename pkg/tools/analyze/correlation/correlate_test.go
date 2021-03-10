package correlation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Correlate conditions", func() {
	var testCorrelationData = []map[string]string{
		{"readout": "a", "condition": "a", "abundance": "4|5", "score": "0.01"},
		{"readout": "b", "condition": "a", "abundance": "7", "score": "0.01"},
		{"readout": "c", "condition": "a", "abundance": "15", "score": "0.01"},
		{"readout": "d", "condition": "a", "abundance": "8", "score": "0.01"},
		{"readout": "e", "condition": "a", "abundance": "1", "score": "0.01"},
		{"readout": "a", "condition": "b", "abundance": "10", "score": "0.02"},
		{"readout": "b", "condition": "b", "abundance": "10", "score": "0"},
		{"readout": "c", "condition": "b", "abundance": "17", "score": "0"},
		{"readout": "d", "condition": "b", "abundance": "19", "score": "0"},
		{"readout": "e", "condition": "b", "abundance": "4", "score": "0"},
		{"readout": "a", "condition": "c", "abundance": "23|12", "score": "0"},
		{"readout": "b", "condition": "c", "abundance": "3|6", "score": "0"},
		{"readout": "c", "condition": "c", "abundance": "20|25", "score": "0"},
		{"readout": "d", "condition": "c", "abundance": "13|14", "score": "0"},
		{"readout": "e", "condition": "c", "abundance": "15|15", "score": "0"},
	}

	Describe("averaging replicates", func() {
		It("should return correlation", func() {
			analysis := &types.Analysis{
				Data: testCorrelationData,
				Settings: types.Settings{
					Correlation:               "pearson",
					IgnoreSourceTargetMatches: false,
					ReadoutAbundanceFilter:    5,
					ReadoutScoreFilter:        0.01,
					ScoreType:                 "lte",
					UseReplicates:             false,
				},
			}

			expected := correlationData{
				labels: []string{
					"a",
					"b",
					"c",
				},
				matrix: [][]float64{
					{1, 0.799, 0.384},
					{0.799, 1, 0.270},
					{0.384, 0.270, 1},
				},
			}

			actual := correlateConditions(analysis)

			Expect(actual.labels).To(Equal(expected.labels), "should return correct labels")
			for i, row := range actual.matrix {
				for j, value := range row {
					Expect(value).To(BeNumerically("~", expected.matrix[i][j], 0.001))
				}
			}
		})

		It("should return correlation ignoring source target matches", func() {
			analysis := &types.Analysis{
				Data: testCorrelationData,
				Settings: types.Settings{
					Correlation:               "pearson",
					IgnoreSourceTargetMatches: true,
					ReadoutAbundanceFilter:    5,
					ReadoutScoreFilter:        0.01,
					ScoreType:                 "lte",
					UseReplicates:             false,
				},
			}

			expected := correlationData{
				labels: []string{
					"a",
					"b",
					"c",
				},
				matrix: [][]float64{
					{1, 0.798, -0.500},
					{0.798, 1, -0.475},
					{-0.500, -0.475, 1},
				},
			}

			actual := correlateConditions(analysis)

			Expect(actual.labels).To(Equal(expected.labels), "should return correct labels")
			for i, row := range actual.matrix {
				for j, value := range row {
					Expect(value).To(BeNumerically("~", expected.matrix[i][j], 0.001))
				}
			}
		})
	})

	Describe("keeping replicates", func() {
		It("should return correlation", func() {
			analysis := &types.Analysis{
				Data: testCorrelationData,
				Settings: types.Settings{
					Correlation:               "pearson",
					IgnoreSourceTargetMatches: false,
					ReadoutAbundanceFilter:    5,
					ReadoutScoreFilter:        0.01,
					ScoreType:                 "lte",
					UseReplicates:             true,
				},
			}

			expected := correlationData{
				labels: []string{
					"a",
					"b",
					"c",
				},
				matrix: [][]float64{
					{1, 0.842, 0.026},
					{0.842, 1, 0.067},
					{0.026, 0.067, 1},
				},
			}

			actual := correlateConditions(analysis)

			Expect(actual.labels).To(Equal(expected.labels), "should return correct labels")
			for i, row := range actual.matrix {
				for j, value := range row {
					Expect(value).To(BeNumerically("~", expected.matrix[i][j], 0.001))
				}
			}
		})

		It("should return correlation ignoring source target pairs", func() {
			analysis := &types.Analysis{
				Data: testCorrelationData,
				Settings: types.Settings{
					Correlation:               "pearson",
					IgnoreSourceTargetMatches: true,
					ReadoutAbundanceFilter:    5,
					ReadoutScoreFilter:        0.01,
					ScoreType:                 "lte",
					UseReplicates:             true,
				},
			}

			expected := correlationData{
				labels: []string{
					"a",
					"b",
					"c",
				},
				matrix: [][]float64{
					{1, 0.902, -0.368},
					{0.902, 1, 0.205},
					{-0.368, 0.205, 1},
				},
			}

			actual := correlateConditions(analysis)

			Expect(actual.labels).To(Equal(expected.labels), "should return correct labels")
			for i, row := range actual.matrix {
				for j, value := range row {
					Expect(value).To(BeNumerically("~", expected.matrix[i][j], 0.001))
				}
			}
		})
	})
})

var _ = Describe("Correlate readouts", func() {
	var testCorrelationData = []map[string]string{
		{"condition": "a", "readout": "a", "abundance": "4|5", "score": "0.01"},
		{"condition": "a", "readout": "b", "abundance": "10", "score": "0.02"},
		{"condition": "a", "readout": "c", "abundance": "23|12", "score": "0"},
		{"condition": "a", "readout": "d", "abundance": "12", "score": "0.02"},
		{"condition": "a", "readout": "e", "abundance": "4", "score": "0"},
		{"condition": "b", "readout": "a", "abundance": "7", "score": "0.01"},
		{"condition": "b", "readout": "b", "abundance": "10", "score": "0"},
		{"condition": "b", "readout": "c", "abundance": "3|6", "score": "0"},
		{"condition": "b", "readout": "d", "abundance": "12", "score": "0.02"},
		{"condition": "b", "readout": "e", "abundance": "4", "score": "0"},
		{"condition": "c", "readout": "b", "abundance": "17", "score": "0"},
		{"condition": "c", "readout": "a", "abundance": "15", "score": "0.01"},
		{"condition": "c", "readout": "c", "abundance": "20|25", "score": "0"},
		{"condition": "d", "readout": "a", "abundance": "8", "score": "0.01"},
		{"condition": "d", "readout": "b", "abundance": "19", "score": "0"},
		{"condition": "d", "readout": "c", "abundance": "13|14", "score": "0"},
		{"condition": "e", "readout": "a", "abundance": "1", "score": "0.01"},
		{"condition": "e", "readout": "b", "abundance": "4", "score": "0"},
		{"condition": "e", "readout": "c", "abundance": "15|15", "score": "0"},
	}

	Describe("averaging replicates", func() {
		It("should return correlation", func() {
			analysis := &types.Analysis{
				Data: testCorrelationData,
				Settings: types.Settings{
					Correlation:               "pearson",
					IgnoreSourceTargetMatches: false,
					ReadoutAbundanceFilter:    5,
					ReadoutScoreFilter:        0.01,
					ScoreType:                 "lte",
					UseReplicates:             false,
				},
			}

			expected := correlationData{
				labels: []string{
					"a",
					"b",
					"c",
				},
				matrix: [][]float64{
					{1, 0.799, 0.384},
					{0.799, 1, 0.270},
					{0.384, 0.270, 1},
				},
			}

			actual := correlateReadouts(analysis)

			Expect(actual.labels).To(Equal(expected.labels), "should return correct labels")
			for i, row := range actual.matrix {
				for j, value := range row {
					Expect(value).To(BeNumerically("~", expected.matrix[i][j], 0.001))
				}
			}
		})

		It("should return correlation ignoring source target matches", func() {
			analysis := &types.Analysis{
				Data: testCorrelationData,
				Settings: types.Settings{
					Correlation:               "pearson",
					IgnoreSourceTargetMatches: true,
					ReadoutAbundanceFilter:    5,
					ReadoutScoreFilter:        0.01,
					ScoreType:                 "lte",
					UseReplicates:             false,
				},
			}

			expected := correlationData{
				labels: []string{
					"a",
					"b",
					"c",
				},
				matrix: [][]float64{
					{1, 0.798, -0.500},
					{0.798, 1, -0.475},
					{-0.500, -0.475, 1},
				},
			}

			actual := correlateReadouts(analysis)

			Expect(actual.labels).To(Equal(expected.labels), "should return correct labels")
			for i, row := range actual.matrix {
				for j, value := range row {
					Expect(value).To(BeNumerically("~", expected.matrix[i][j], 0.001))
				}
			}
		})
	})

	Describe("keeping replicates", func() {
		It("should return correlation", func() {
			analysis := &types.Analysis{
				Data: testCorrelationData,
				Settings: types.Settings{
					Correlation:               "pearson",
					IgnoreSourceTargetMatches: false,
					ReadoutAbundanceFilter:    5,
					ReadoutScoreFilter:        0.01,
					ScoreType:                 "lte",
					UseReplicates:             true,
				},
			}

			expected := correlationData{
				labels: []string{
					"a",
					"b",
					"c",
				},
				matrix: [][]float64{
					{1, 0.842, 0.026},
					{0.842, 1, 0.067},
					{0.026, 0.067, 1},
				},
			}

			actual := correlateReadouts(analysis)

			Expect(actual.labels).To(Equal(expected.labels), "should return correct labels")
			for i, row := range actual.matrix {
				for j, value := range row {
					Expect(value).To(BeNumerically("~", expected.matrix[i][j], 0.001))
				}
			}
		})

		It("should return correlation ignoring source target pairs", func() {
			analysis := &types.Analysis{
				Data: testCorrelationData,
				Settings: types.Settings{
					Correlation:               "pearson",
					IgnoreSourceTargetMatches: true,
					ReadoutAbundanceFilter:    5,
					ReadoutScoreFilter:        0.01,
					ScoreType:                 "lte",
					UseReplicates:             true,
				},
			}

			expected := correlationData{
				labels: []string{
					"a",
					"b",
					"c",
				},
				matrix: [][]float64{
					{1, 0.902, -0.368},
					{0.902, 1, 0.205},
					{-0.368, 0.205, 1},
				},
			}

			actual := correlateReadouts(analysis)

			Expect(actual.labels).To(Equal(expected.labels), "should return correct labels")
			for i, row := range actual.matrix {
				for j, value := range row {
					Expect(value).To(BeNumerically("~", expected.matrix[i][j], 0.001))
				}
			}
		})
	})
})

var _ = Describe("Strip replicate information", func() {
	It("should not strip replicate information (only from end)", func() {
		names := []string{"a", "b", "c_R1", "c_R2", "c_R1x", "d_R1", "d_R2"}
		useReplicates := false

		expected := names
		Expect(stripReplicates(names, useReplicates)).To(Equal(expected))
	})

	It("should strip replicate information (only from end)", func() {
		names := []string{"a", "b", "c_R1", "c_R2", "c_R1x", "d_R1", "d_R2"}
		useReplicates := true

		expected := []string{"a", "b", "c", "c", "c_R1x", "d", "d"}
		Expect(stripReplicates(names, useReplicates)).To(Equal(expected))
	})
})
