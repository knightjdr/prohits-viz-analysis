package heatmap

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createInteractive(matrices *types.Matrices, parameters types.Settings, fileid string) {
	settings := defineSettings(parameters)
	data := &interactive.HeatmapData{
		AnalysisType: parameters.Type,
		Filename:     fmt.Sprintf("%s.json", fileid),
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
			"abundanceType":   parameters.AbundanceType,
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
		"abundanceType": parameters.AbundanceType,
		"fillColor":     parameters.FillColor,
		"imageType":     "heatmap",
		"invertColor":   parameters.InvertColor,
		"minAbundance":  parameters.MinAbundance,
		"primaryFilter": parameters.PrimaryFilter,
	}
}
