package circheatmap

import "github.com/knightjdr/prohits-viz-analysis/pkg/types"

// Data defines the type and variables required for generating an interactive circheatmap
type Data struct {
	Filename   string
	Legend     types.CircHeatmapLegend
	Parameters types.Settings
	Plots      []types.CircHeatmap
	Settings   map[string]interface{}
}
