// Package frontend creates matrices from frontend format.
package frontend

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"

// Row contains heatmap row information.
type Row struct {
	Data []Cell
	Name string
}

// Cell contains individual heatmap cell information.
type Cell struct {
	Ratio float64
	Score float64
	Value float64
}

// CreateHeatmapMatrix from row database.
func CreateHeatmapMatrix(db []Row, order map[string][]int) [][]float64 {
	matrix := make([][]float64, len(order["rows"]))

	for i, rowIndex := range order["rows"] {
		matrix[i] = make([]float64, len(order["columns"]))
		for j, columnIndex := range order["columns"] {
			matrix[i][j] = db[rowIndex].Data[columnIndex].Value
		}
	}

	return matrix
}

// CreateDotplotMatrices from row database.
func CreateDotplotMatrices(db []Row, order map[string][]int) *types.Matrices {
	matrices := &types.Matrices{
		Abundance: make([][]float64, len(order["rows"])),
		Ratio:     make([][]float64, len(order["rows"])),
		Score:     make([][]float64, len(order["rows"])),
	}

	for i, rowIndex := range order["rows"] {
		matrices.Abundance[i] = make([]float64, len(order["columns"]))
		matrices.Ratio[i] = make([]float64, len(order["columns"]))
		matrices.Score[i] = make([]float64, len(order["columns"]))
		for j, columnIndex := range order["columns"] {
			matrices.Abundance[i][j] = db[rowIndex].Data[columnIndex].Value
			matrices.Ratio[i][j] = db[rowIndex].Data[columnIndex].Ratio
			matrices.Score[i][j] = db[rowIndex].Data[columnIndex].Score
		}
	}

	return matrices
}
