package circheatmap

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func parseParameters(settings types.Settings) string {
	parameters := map[string]interface{}{
		"abundanceColumn": settings.Abundance,
		"analysisType":    settings.Type,
		"conditionColumn": settings.Condition,
		"controlColumn":   settings.Control,
		"files":           files.ParseBaseNames(settings.Files),
		"imageType":       "circheatmap",
		"normalization":   settings.Normalization,
		"readoutColumn":   settings.Readout,
		"scoreColumn":     settings.Score,
		"scoreType":       settings.ScoreType,
	}

	jsonString, _ := json.Marshal(parameters)
	return fmt.Sprintf("\"parameters\": %s", string(jsonString))
}
