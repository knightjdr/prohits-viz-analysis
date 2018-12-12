package interactive

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
)

func rowData(
	imageType string,
	abundance, ratios, scores [][]float64,
	columns, rows []string,
) []map[string]interface{} {
	// Create row data.
	numCols := len(columns)
	numRows := len(rows)
	data := make([]map[string]interface{}, numRows)

	if imageType == "dotplot" {
		for i, row := range abundance {
			rowslice := make([]map[string]float64, numCols)
			for j, value := range row {
				rowslice[j] = map[string]float64{
					"ratio": helper.TruncateFloat(ratios[i][j], 2),
					"score": helper.TruncateFloat(scores[i][j], 2),
					"value": helper.TruncateFloat(value, 2),
				}
			}
			data[i] = map[string]interface{}{
				"name": rows[i],
				"data": rowslice,
			}
		}
	} else {
		for i, row := range abundance {
			rowslice := make([]map[string]float64, numCols)
			for j, value := range row {
				rowslice[j] = map[string]float64{
					"value": helper.TruncateFloat(value, 2),
				}
			}
			data[i] = map[string]interface{}{
				"name": rows[i],
				"data": rowslice,
			}
		}
	}

	return data
}
