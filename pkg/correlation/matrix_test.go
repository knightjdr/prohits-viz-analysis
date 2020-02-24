package correlation_test

import (
	. "github.com/knightjdr/prohits-viz-analysis/pkg/correlation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Correlation", func() {
	It("should calculate pearson kendall coefficient", func() {
		data := Data{
			Matrix: [][]float64{
				[]float64{1, 5, 10},
				[]float64{2, 10, 20},
				[]float64{10, 2, 1},
			},
			Method: "kendall",
		}

		expected := [][]float64{
			[]float64{1, 1, -1},
			[]float64{1, 1, -1},
			[]float64{-1, -1, 1},
		}

		actual := data.Correlate()
		for i, row := range actual {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.001))
			}
		}
	})

	It("should calculate pearson correlation coefficient", func() {
		data := Data{
			Matrix: [][]float64{
				[]float64{1, 5, 10},
				[]float64{2, 10, 20},
				[]float64{10, 2, 1},
			},
			Method: "pearson",
		}

		expected := [][]float64{
			[]float64{1, 1, -0.884},
			[]float64{1, 1, -0.884},
			[]float64{-0.884, -0.884, 1},
		}

		actual := data.Correlate()
		for i, row := range actual {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.001))
			}
		}
	})

	It("should calculate spearman correlation coefficient", func() {
		data := Data{
			Matrix: [][]float64{
				[]float64{1, 5, 10},
				[]float64{2, 10, 20},
				[]float64{10, 2, 1},
			},
			Method: "spearman",
		}

		expected := [][]float64{
			[]float64{1, 1, -1},
			[]float64{1, 1, -1},
			[]float64{-1, -1, 1},
		}

		actual := data.Correlate()
		for i, row := range actual {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.001))
			}
		}
	})

	It("should calculate correlation coefficient between columns", func() {
		data := Data{
			Dimension: "column",
			Matrix: [][]float64{
				[]float64{1, 5, 10},
				[]float64{2, 10, 20},
				[]float64{10, 2, 1},
			},
			Method: "pearson",
		}

		expected := [][]float64{
			[]float64{1, -0.719, -0.793},
			[]float64{-0.719, 1, 0.994},
			[]float64{-0.793, 0.994, 1},
		}

		actual := data.Correlate()
		for i, row := range actual {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.001))
			}
		}
	})

	It("should calculate correlation coefficient and ignore source target pairs", func() {
		data := Data{
			Columns:                   []string{"a", "b", "c", "d", "e"},
			IgnoreSourceTargetMatches: true,
			Matrix: [][]float64{
				[]float64{1, 5, 10, 10, 15},
				[]float64{2, 10, 20, 12, 8},
				[]float64{10, 2, 1, 4, 8},
			},
			Method: "pearson",
			Rows:   []string{"a", "b", "c"},
		}

		expected := [][]float64{
			[]float64{1, -0.756, 0.982},
			[]float64{-0.756, 1, -0.954},
			[]float64{0.982, -0.954, 1},
		}

		actual := data.Correlate()
		for i, row := range actual {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.001))
			}
		}
	})
})
