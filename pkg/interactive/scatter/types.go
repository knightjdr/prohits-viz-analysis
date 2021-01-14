package scatter

import "github.com/knightjdr/prohits-viz-analysis/pkg/types"

// Data defines the type and variables required for generating an interactive scatter plot
type Data struct {
	AnalysisType string
	Filename     string
	Legend       []map[string]string
	Parameters   types.Settings
	Plots        []types.ScatterPlot
	Settings     map[string]interface{}
}
