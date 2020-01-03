package heatmap

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func parseParameters(imageType string, settings types.Settings) string {
	parameters := map[string]interface{}{
		"abundanceColumn":    settings.Abundance,
		"analysisType":       settings.Type,
		"conditionColumn":    settings.Condition,
		"clustering":         settings.Clustering,
		"clusteringMethod":   settings.ClusteringMethod,
		"clusteringOptimize": settings.ClusteringOptimize,
		"controlColumn":      settings.Control,
		"distance":           settings.Distance,
		"files":              files.ParseBaseNames(settings.Files),
		"imageType":          imageType,
		"logBase":            settings.LogBase,
		"normalization":      settings.Normalization,
		"readoutColumn":      settings.Readout,
		"scoreColumn":        settings.Score,
		"scoreType":          settings.ScoreType,
		"xLabel":             settings.XLabel,
		"yLabel":             settings.YLabel,
	}

	jsonString, _ := json.Marshal(parameters)
	return fmt.Sprintf("\"parameters\": %s", string(jsonString))
}
