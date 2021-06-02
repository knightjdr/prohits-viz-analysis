package scv

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeImages(plots []types.CircHeatmap, legend types.CircHeatmapLegend, settings types.Settings) {
	for _, plot := range plots {
		filehandle := plot.Name

		createSVG(plot, legend, settings, filehandle)

		if settings.Png {
			svg.ConvertToPNG(fmt.Sprintf("svg/%s.svg", filehandle), fmt.Sprintf("png/%s.png", filehandle), "white")
		}
	}
}

func createSVG(plot types.CircHeatmap, legend types.CircHeatmapLegend, settings types.Settings, filehandle string) {
	filename := fmt.Sprintf("svg/%s.svg", filehandle)

	circHeatmap := svg.InitializeCircHeatmap()
	circHeatmap.Legend = legend
	circHeatmap.Plot = plot
	circHeatmap.ShowKnown = shouldDrawKnown(settings.Known)

	circHeatmap.Draw(filename)
}

func shouldDrawKnown(known string) bool {
	return known != ""
}
