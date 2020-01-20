package convert

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createInteractive(matrices *types.Matrices, parameters types.Settings) {
	settings := defineSettings(parameters)
	data := &interactive.HeatmapData{
		AnalysisType: parameters.Type,
		Filename:     fmt.Sprintf("interactive/%s.json", parameters.Type),
		Matrices:     matrices,
		Minimap:      "./minimap/minimap.png",
		Parameters:   parameters,
		Settings:     settings,
	}
	interactive.CreateHeatmap(data)
}

func defineSettings(parameters types.Settings) map[string]interface{} {
	if parameters.Type == "dotplot" {
		return map[string]interface{}{
			"abundanceCap":    parameters.AbundanceCap,
			"edgeColor":       parameters.EdgeColor,
			"fillColor":       parameters.FillColor,
			"imageType":       "dotplot",
			"invertColor":     parameters.InvertColor,
			"minAbundance":    parameters.MinAbundance,
			"primaryFilter":   parameters.PrimaryFilter,
			"secondaryFilter": parameters.SecondaryFilter,
		}
	}
	return map[string]interface{}{
		"abundanceCap":  parameters.AbundanceCap,
		"fillColor":     parameters.FillColor,
		"imageType":     "heatmap",
		"invertColor":   parameters.InvertColor,
		"minAbundance":  parameters.MinAbundance,
		"primaryFilter": parameters.PrimaryFilter,
	}
}
