package interactive

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// ParseHeatmap creates vars for passing to the interactive file maker
// add passes those vars.
func ParseHeatmap(data *Data) {
	// User parameters.
	parameters := map[string]interface{}{
		"abundanceColumn":    data.Parameters.Abundance,
		"analysisType":       data.Parameters.AnalysisType,
		"conditionColumn":    data.Parameters.Condition,
		"clustering":         data.Parameters.Clustering,
		"clusteringMethod":   data.Parameters.ClusteringMethod,
		"clusteringOptimize": data.Parameters.ClusteringOptimize,
		"controlColumn":      data.Parameters.Control,
		"distance":           data.Parameters.Distance,
		"files":              helper.Filename(data.Parameters.Files),
		"imageType":          data.ImageType,
		"logBase":            data.Parameters.LogBase,
		"normalization":      data.Parameters.Normalization,
		"readoutColumn":      data.Parameters.Readout,
		"scoreColumn":        data.Parameters.Score,
		"scoreType":          data.Parameters.ScoreType,
		"xLabel":             data.Parameters.XLabel,
		"yLabel":             data.Parameters.YLabel,
	}

	// Needed settings.
	settings := map[string]interface{}{
		"abundanceCap":    data.Parameters.AbundanceCap,
		"fillColor":       data.Parameters.FillColor,
		"imageType":       data.ImageType,
		"invertColor":     data.Parameters.InvertColor,
		"minAbundance":    data.Parameters.MinAbundance,
		"primaryFilter":   data.Parameters.PrimaryFilter,
		"secondaryFilter": data.Parameters.SecondaryFilter,
	}
	if data.ImageType == "dotplot" {
		settings["edgeColor"] = data.Parameters.EdgeColor
	}

	// Convert png to url.
	url := Pngurl(data.Minimap)

	rowData := rowData(data.ImageType, data.Matrices)

	// Write json.
	Heatmap(data.Matrices.Conditions, rowData, parameters, settings, url, data.Filename)
	return
}
