package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/interactive"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// InteractiveDotplot creates vars for passing to the interactive file maker
// add passes those vars.
func InteractiveDotplot(
	abundance, ratios, scores [][]float64,
	columns, rows []string,
	userParams typedef.Parameters,
	pngfilename string,
) (json string) {
	// Needed parameters.
	params := map[string]interface{}{
		"abundanceColumn":  userParams.Abundance,
		"analysisType":     userParams.AnalysisType,
		"baitColumn":       userParams.Bait,
		"clustering":       userParams.Clustering,
		"clusteringMethod": userParams.ClusteringMethod,
		"colorSpace":       userParams.ColorSpace,
		"controlColumn":    userParams.Control,
		"distance":         userParams.Distance,
		"files":            userParams.Files,
		"imageType":        "dotplot",
		"logBase":          userParams.LogBase,
		"maximumAbundance": userParams.MaximumAbundance,
		"minimumAbundance": userParams.MinimumAbundance,
		"normalization":    userParams.Normalization,
		"preyColumn":       userParams.Prey,
		"primaryFilter":    userParams.PrimaryFilter,
		"scoreColumn":      userParams.Score,
		"scoreType":        userParams.ScoreType,
		"secondaryFilter":  userParams.SecondaryFilter,
	}

	// Convert png to url.
	url := interactive.Pngurl(pngfilename)

	// Create row data.
	numCols := len(columns)
	numRows := len(rows)
	data := make([]map[string]interface{}, numRows)
	for i, row := range abundance {
		rowslice := make([]map[string]float64, numCols)
		for j, value := range row {
			rowslice[j] = map[string]float64{
				"ratio": ratios[i][j],
				"score": scores[i][j],
				"value": value,
			}
		}
		data[i] = map[string]interface{}{
			"name": rows[i],
			"data": rowslice,
		}
	}

	// Get json.
	json = interactive.Heatmap(data, columns, params, url)
	return
}
