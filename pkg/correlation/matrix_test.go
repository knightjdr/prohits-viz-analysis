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
				{1, 5, 10},
				{2, 10, 20},
				{10, 2, 1},
			},
			Method: "kendall",
		}

		expected := [][]float64{
			{1, 1, -1},
			{1, 1, -1},
			{-1, -1, 1},
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
				{1, 5, 10},
				{2, 10, 20},
				{10, 2, 1},
			},
			Method: "pearson",
		}

		expected := [][]float64{
			{1, 1, -0.884},
			{1, 1, -0.884},
			{-0.884, -0.884, 1},
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
				{1, 5, 10},
				{2, 10, 20},
				{10, 2, 1},
			},
			Method: "spearman",
		}

		expected := [][]float64{
			{1, 1, -1},
			{1, 1, -1},
			{-1, -1, 1},
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
				{1, 5, 10},
				{2, 10, 20},
				{10, 2, 1},
			},
			Method: "pearson",
		}

		expected := [][]float64{
			{1, -0.719, -0.793},
			{-0.719, 1, 0.994},
			{-0.793, 0.994, 1},
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
				{1, 5, 10, 10, 15},
				{2, 10, 20, 12, 8},
				{10, 2, 1, 4, 8},
			},
			Method: "pearson",
			Rows:   []string{"a", "b", "c"},
		}

		expected := [][]float64{
			{1, -0.756, 0.982},
			{-0.756, 1, -0.954},
			{0.982, -0.954, 1},
		}

		actual := data.Correlate()
		for i, row := range actual {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.001))
			}
		}
	})

	It("should calculate correlation coefficient and ignore source target pairs with redunant column names", func() {
		data := Data{
			Columns:                   []string{"a", "b", "c", "d", "e", "a"},
			IgnoreSourceTargetMatches: true,
			Matrix: [][]float64{
				{1, 5, 10, 10, 8, 15},
				{2, 10, 20, 7, 4, 8},
				{10, 2, 1, 4, 9, 8},
			},
			Method: "pearson",
			Rows:   []string{"a", "b", "c"},
		}
		expected := [][]float64{
			{1, 0.645, 0.386},
			{0.645, 1, -0.679},
			{0.386, -0.679, 1},
		}

		actual := data.Correlate()
		for i, row := range actual {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.001))
			}
		}
	})

	It("should calculate correlation coefficient and ignore source target pairs with redunant row names", func() {
		data := Data{
			Columns:                   []string{"a", "b", "c"},
			Dimension:                 "column",
			IgnoreSourceTargetMatches: true,
			Matrix: [][]float64{
				{1, 2, 10},
				{5, 10, 2},
				{10, 20, 1},
				{10, 7, 4},
				{8, 4, 9},
				{15, 8, 8},
			},
			Method: "pearson",
			Rows:   []string{"a", "b", "c", "d", "e", "a"},
		}
		expected := [][]float64{
			{1, 0.645, 0.386},
			{0.645, 1, -0.679},
			{0.386, -0.679, 1},
		}

		actual := data.Correlate()
		for i, row := range actual {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.001))
			}
		}
	})
})
