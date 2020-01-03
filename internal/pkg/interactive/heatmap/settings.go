package heatmap

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func parseSettings(imageType string, settings types.Settings) string {
	imageSettings := map[string]map[string]map[string]interface{}{
		"main": map[string]map[string]interface{}{
			"current": map[string]interface{}{
				"abundanceCap":  settings.AbundanceCap,
				"fillColor":     settings.FillColor,
				"imageType":     imageType,
				"invertColor":   settings.InvertColor,
				"minAbundance":  settings.MinAbundance,
				"primaryFilter": settings.PrimaryFilter,
			},
		},
	}
	if imageType == "dotplot" {
		imageSettings["main"]["current"]["edgeColor"] = settings.EdgeColor
		imageSettings["main"]["current"]["secondaryFilter"] = settings.SecondaryFilter
	}

	jsonString, _ := json.Marshal(imageSettings)
	return fmt.Sprintf("\"settings\": %s", string(jsonString))
}
