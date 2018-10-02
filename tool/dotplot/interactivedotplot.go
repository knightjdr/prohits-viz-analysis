package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
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
	// User parameters.
	parameters := map[string]interface{}{
		"abundanceColumn":  userParams.Abundance,
		"analysisType":     userParams.AnalysisType,
		"conditionColumn":  userParams.Condition,
		"clustering":       userParams.Clustering,
		"clusteringMethod": userParams.ClusteringMethod,
		"controlColumn":    userParams.Control,
		"distance":         userParams.Distance,
		"files":            helper.Filename(userParams.Files),
		"imageType":        "dotplot",
		"logBase":          userParams.LogBase,
		"normalization":    userParams.Normalization,
		"readoutColumn":    userParams.Readout,
		"scoreColumn":      userParams.Score,
		"scoreType":        userParams.ScoreType,
	}

	// Needed settings.
	settings := map[string]interface{}{
		"abundanceCap":    userParams.AbundanceCap,
		"edgeColor":       userParams.EdgeColor,
		"fillColor":       userParams.FillColor,
		"imageType":       "dotplot",
		"invertColor":     false,
		"minAbundance":    userParams.MinAbundance,
		"primaryFilter":   userParams.PrimaryFilter,
		"secondaryFilter": userParams.SecondaryFilter,
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
	json = interactive.Heatmap(data, columns, parameters, settings, url)
	return
}
