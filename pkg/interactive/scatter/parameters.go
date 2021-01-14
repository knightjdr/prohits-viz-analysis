package scatter

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func parseParameters(imageType string, settings types.Settings) string {
	var parameters map[string]interface{}

	if imageType == "condition-condition" {
		parameters = parseCCParameters(settings)
	}
	if imageType == "specificity" {
		parameters = parseSpecificityParameters(settings)
	}

	jsonString, _ := json.Marshal(parameters)
	return fmt.Sprintf("\"parameters\": %s", string(jsonString))
}

func parseCCParameters(settings types.Settings) map[string]interface{} {
	return map[string]interface{}{
		"abundanceColumn":        settings.Abundance,
		"analysisType":           settings.Type,
		"conditionColumn":        settings.Condition,
		"controlColumn":          settings.Control,
		"files":                  files.ParseBaseNames(settings.Files),
		"imageType":              "scatter",
		"mockConditionAbundance": settings.MockConditionAbundance,
		"normalization":          settings.Normalization,
		"readoutColumn":          settings.Readout,
		"scoreColumn":            settings.Score,
		"scoreType":              settings.ScoreType,
	}
}

func parseSpecificityParameters(settings types.Settings) map[string]interface{} {
	return map[string]interface{}{
		"abundanceColumn":        settings.Abundance,
		"analysisType":           settings.Type,
		"conditionColumn":        settings.Condition,
		"controlColumn":          settings.Control,
		"files":                  files.ParseBaseNames(settings.Files),
		"imageType":              "scatter",
		"mockConditionAbundance": settings.MockConditionAbundance,
		"normalization":          settings.Normalization,
		"primaryFilter":          settings.PrimaryFilter,
		"readoutColumn":          settings.Readout,
		"scoreColumn":            settings.Score,
		"scoreType":              settings.ScoreType,
		"specificityMetric":      settings.SpecificityMetric,
	}
}
