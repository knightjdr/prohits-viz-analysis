package scatter

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func parsePlots(plots []types.ScatterPlot) string {
	jsonString, _ := json.Marshal(plots)
	return fmt.Sprintf("\"plots\": %s", string(jsonString))
}
