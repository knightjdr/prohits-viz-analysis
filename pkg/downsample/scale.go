package downsample

type dimensions struct {
	Columns int
	Rows    int
}

func calculateScale(matrix [][]float64, maxDimension int) float64 {
	matrixDimensions := dimensions{
		Columns: len(matrix[0]),
		Rows:    len(matrix),
	}

	if matrixDimensions.Rows >= matrixDimensions.Columns {
		return float64(matrixDimensions.Rows) / float64(maxDimension)
	}
	return float64(matrixDimensions.Columns) / float64(maxDimension)
}
