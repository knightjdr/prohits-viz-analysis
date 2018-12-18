package interactive

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

func rowData(imageType string, matrices *typedef.Matrices) []map[string]interface{} {
	// Create row data.
	numCols := len(matrices.Conditions)
	numRows := len(matrices.Readouts)
	data := make([]map[string]interface{}, numRows)

	if imageType == "dotplot" {
		for i, row := range matrices.Abundance {
			rowslice := make([]map[string]float64, numCols)
			for j, value := range row {
				rowslice[j] = map[string]float64{
					"ratio": helper.TruncateFloat(matrices.Ratio[i][j], 2),
					"score": helper.TruncateFloat(matrices.Score[i][j], 2),
					"value": helper.TruncateFloat(value, 2),
				}
			}
			data[i] = map[string]interface{}{
				"name": matrices.Readouts[i],
				"data": rowslice,
			}
		}
	} else {
		for i, row := range matrices.Abundance {
			rowslice := make([]map[string]float64, numCols)
			for j, value := range row {
				rowslice[j] = map[string]float64{
					"value": helper.TruncateFloat(value, 2),
				}
			}
			data[i] = map[string]interface{}{
				"name": matrices.Readouts[i],
				"data": rowslice,
			}
		}
	}

	return data
}
