package heatmap

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func parseParameters(imageType string, settings types.Settings) string {
	var parameters map[string]interface{}

	switch imageType {
	case "correlation":
		parameters = parseCorrelationParameters(settings)
	default:
		parameters = parseHeatmapParamaters(imageType, settings)
	}

	jsonString, _ := json.Marshal(parameters)
	return fmt.Sprintf("\"parameters\": %s", string(jsonString))
}

func parseCorrelationParameters(settings types.Settings) map[string]interface{} {
	return map[string]interface{}{
		"abundanceColumn":           settings.Abundance,
		"analysisType":              settings.Type,
		"clustering":                settings.Clustering,
		"clusteringMethod":          settings.ClusteringMethod,
		"clusteringOptimize":        settings.ClusteringOptimize,
		"conditionAbundanceFilter":  settings.ConditionAbundanceFilter,
		"conditionColumn":           settings.Condition,
		"conditionScoreFilter":      settings.ConditionScoreFilter,
		"controlColumn":             settings.Control,
		"correlation":               settings.Correlation,
		"distance":                  settings.Distance,
		"files":                     files.ParseBaseNames(settings.Files),
		"imageType":                 "heatmap",
		"IgnoreSourceTargetMatches": settings.IgnoreSourceTargetMatches,
		"logBase":                   settings.LogBase,
		"mockConditionAbundance":    settings.MockConditionAbundance,
		"minConditions":             settings.MinConditions,
		"normalization":             settings.Normalization,
		"parsimoniousReadouts":      settings.ParsimoniousReadoutFiltering,
		"readoutAbundanceFilter":    settings.ReadoutAbundanceFilter,
		"readoutColumn":             settings.Readout,
		"readoutScoreFilter":        settings.ReadoutScoreFilter,
		"scoreColumn":               settings.Score,
		"scoreType":                 settings.ScoreType,
		"useReplicates":             settings.UseReplicates,
		"xLabel":                    settings.XLabel,
		"yLabel":                    settings.YLabel,
	}
}

func parseHeatmapParamaters(imageType string, settings types.Settings) map[string]interface{} {
	return map[string]interface{}{
		"abundanceColumn":        settings.Abundance,
		"analysisType":           settings.Type,
		"clustering":             settings.Clustering,
		"clusteringMethod":       settings.ClusteringMethod,
		"clusteringOptimize":     settings.ClusteringOptimize,
		"conditionColumn":        settings.Condition,
		"controlColumn":          settings.Control,
		"distance":               settings.Distance,
		"files":                  files.ParseBaseNames(settings.Files),
		"imageType":              imageType,
		"logBase":                settings.LogBase,
		"minConditions":          settings.MinConditions,
		"mockConditionAbundance": settings.MockConditionAbundance,
		"normalization":          settings.Normalization,
		"parsimoniousReadouts":   settings.ParsimoniousReadoutFiltering,
		"readoutColumn":          settings.Readout,
		"scoreColumn":            settings.Score,
		"scoreType":              settings.ScoreType,
		"xLabel":                 settings.XLabel,
		"yLabel":                 settings.YLabel,
	}
}
