package scatter

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createInteractive(plots []types.ScatterPlot, settings types.Settings, legend []map[string]string, filename string) {
	interactiveData := &interactive.ScatterData{
		AnalysisType: settings.Type,
		Filename:     fmt.Sprintf("interactive/%s.json", filename),
		Legend:       legend,
		Parameters:   settings,
		Plots:        plots,
		Settings: map[string]interface{}{
			"xFilter": 0,
			"yFilter": 0,
		},
	}

	interactive.CreateScatter(interactiveData)
}
