package interactive

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// ParseHeatmap creates vars for passing to the interactive file maker
// add passes those vars.
func ParseHeatmap(
	imageType string,
	abundance, ratios, scores [][]float64,
	columns, rows []string,
	invertColor bool,
	userParams typedef.Parameters,
	pngfilename string,
) (json string) {
	// User parameters.
	parameters := map[string]interface{}{
		"abundanceColumn":    userParams.Abundance,
		"analysisType":       userParams.AnalysisType,
		"conditionColumn":    userParams.Condition,
		"clustering":         userParams.Clustering,
		"clusteringMethod":   userParams.ClusteringMethod,
		"clusteringOptimize": userParams.ClusteringOptimize,
		"controlColumn":      userParams.Control,
		"distance":           userParams.Distance,
		"files":              helper.Filename(userParams.Files),
		"imageType":          imageType,
		"logBase":            userParams.LogBase,
		"normalization":      userParams.Normalization,
		"readoutColumn":      userParams.Readout,
		"scoreColumn":        userParams.Score,
		"scoreType":          userParams.ScoreType,
	}

	// Needed settings.
	settings := map[string]interface{}{
		"abundanceCap":    userParams.AbundanceCap,
		"fillColor":       userParams.FillColor,
		"imageType":       imageType,
		"invertColor":     invertColor,
		"minAbundance":    userParams.MinAbundance,
		"primaryFilter":   userParams.PrimaryFilter,
		"secondaryFilter": userParams.SecondaryFilter,
	}
	if imageType == "dotplot" {
		settings["edgeColor"] = userParams.EdgeColor
	}

	// Convert png to url.
	url := Pngurl(pngfilename)

	data := rowData(imageType, abundance, ratios, scores, columns, rows)

	// Get json.
	json = Heatmap(data, columns, parameters, settings, url)
	return
}
