package heatmap

func createMatrix(data *heatmap) [][]float64 {
	matrix := make([][]float64, len(data.RowOrder))

	for i, rowIndex := range data.RowOrder {
		matrix[i] = make([]float64, len(data.ColumnOrder))
		for j, columnIndex := range data.ColumnOrder {
			matrix[i][j] = data.RowDB[rowIndex].Data[columnIndex].Value
		}
	}

	return matrix
}
