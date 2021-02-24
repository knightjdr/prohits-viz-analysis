package circheatmap

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func parsePlots(plots []types.CircHeatmap) string {
	jsonString, _ := json.Marshal(plots)
	return fmt.Sprintf("\"plots\": %s", string(jsonString))
}
