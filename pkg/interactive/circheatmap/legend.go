package circheatmap

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func parseLegend(legend types.CircHeatmapLegend) string {
	jsonString, _ := json.Marshal(legend)
	return fmt.Sprintf(
		"\"circles\": {"+
			"\"order\": %s"+
			"}",
		string(jsonString))
}
