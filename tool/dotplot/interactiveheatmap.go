package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/interactive"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// InteractiveHeatmap creates vars for passing to the interactive file maker
// add passes those vars.
func InteractiveHeatmap(
	abundance [][]float64,
	columns, rows []string,
	invertColor bool,
	userParams typedef.Parameters,
	pngfilename string,
) (json string) {
	// Run parameters.
	parameters := map[string]interface{}{
		"abundanceColumn":  userParams.Abundance,
		"analysisType":     userParams.AnalysisType,
		"baitColumn":       userParams.Bait,
		"clustering":       userParams.Clustering,
		"clusteringMethod": userParams.ClusteringMethod,
		"controlColumn":    userParams.Control,
		"distance":         userParams.Distance,
		"files":            userParams.Files,
		"imageType":        "heatmap",
		"logBase":          userParams.LogBase,
		"normalization":    userParams.Normalization,
		"preyColumn":       userParams.Prey,
		"scoreColumn":      userParams.Score,
		"scoreType":        userParams.ScoreType,
	}

	// Needed settings.
	settings := map[string]interface{}{
		"abundanceCap":    userParams.AbundanceCap,
		"fillColor":       userParams.FillColor,
		"imageType":       "heatmap",
		"invertColor":     invertColor,
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
