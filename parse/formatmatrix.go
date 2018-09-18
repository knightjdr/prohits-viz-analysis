package parse

// FormatMatrix converts json data slice to matrices for svg creation.
func FormatMatrix(data *Data) ([][]float64, [][]float64, [][]float64) {
	// Define matrix dimensions.
	colNum := len(data.Rows[0].Data)
	rowNum := len(data.Rows)

	// Init matrices.
	abundance := make([][]float64, rowNum)
	ratios := make([][]float64, rowNum)
	score := make([][]float64, rowNum)

	// Define matrices.
	for i, row := range data.Rows {
		abundance[i] = make([]float64, colNum)
		ratios[i] = make([]float64, colNum)
		score[i] = make([]float64, colNum)
		for j, cell := range row.Data {
			abundance[i][j] = cell.Value
			ratios[i][j] = cell.Ratio
			score[i][j] = cell.Score
		}
	}
	return abundance, ratios, score
}
