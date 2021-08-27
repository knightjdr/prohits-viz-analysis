// Package correlation calculates the correlation between matrices or vectors.
package correlation

import (
	"sort"

	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix"
	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
)

// Data correlation settings.
type Data struct {
	Columns                   []string
	Dimension                 string // Either "column" or "row" (default).
	IgnoreSourceTargetMatches bool
	Matrix                    [][]float64
	Method                    string
	Rows                      []string
}

// Correlate calculates the correlation between the columns or rows.
func (data *Data) Correlate() [][]float64 {
	matrix, columns, rows := transposeMatrix(data)

	n := len(matrix)
	correlation := make([][]float64, n)

	for i := 0; i < n; i++ {
		correlation[i] = make([]float64, n)
	}

	calculateStatistic := getStatistic(data.Method)
	filterData := getDataFilter(matrix, data.IgnoreSourceTargetMatches, columns, rows)

	for i := 0; i < n; i++ {
		correlation[i][i] = 1
		for j := i + 1; j < n; j++ {
			x, y := filterData(i, j)
			value := calculateStatistic(x, y)
			correlation[i][j] = value
			correlation[j][i] = value
		}
	}

	return correlation
}

func transposeMatrix(data *Data) ([][]float64, []string, []string) {
	if data.Dimension == "column" {
		return matrix.Transpose(data.Matrix), data.Rows, data.Columns
	}
	return data.Matrix, data.Columns, data.Rows
}

func getStatistic(method string) func([]float64, []float64) float64 {
	switch method {
	case "kendall":
		return Kendall
	case "spearman":
		return Spearman
	default:
		return Pearson
	}
}

func getDataFilter(matrix [][]float64, IgnoreSourceTargetMatches bool, columns, rows []string) func(int, int) ([]float64, []float64) {
	if IgnoreSourceTargetMatches {
		columnNameToIndicies := make(map[string][]int)
		for index, value := range columns {
			if _, ok := columnNameToIndicies[value]; !ok {
				columnNameToIndicies[value] = make([]int, 0)
			}
			columnNameToIndicies[value] = append(columnNameToIndicies[value], index)
		}

		addIndex := func(indices *[]int, name string) {
			if _, ok := columnNameToIndicies[name]; ok {
				(*indices) = append((*indices), columnNameToIndicies[name]...)
			}
		}

		return func(i, j int) ([]float64, []float64) {
			ignoreIndices := make([]int, 0)
			addIndex(&ignoreIndices, rows[i])
			addIndex(&ignoreIndices, rows[j])
			sort.Ints(ignoreIndices)
			ignoreIndices = slice.ReverseInt(ignoreIndices)

			x := make([]float64, len(matrix[i]))
			copy(x, matrix[i])
			y := make([]float64, len(matrix[j]))
			copy(y, matrix[j])

			for _, indexToIgnore := range ignoreIndices {
				x = append(x[:indexToIgnore], x[indexToIgnore+1:]...)
				y = append(y[:indexToIgnore], y[indexToIgnore+1:]...)
			}

			return x, y
		}
	}

	return func(i, j int) ([]float64, []float64) {
		return matrix[i], matrix[j]
	}
}
