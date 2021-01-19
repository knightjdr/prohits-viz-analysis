package scatter

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func parsePlots(plots []types.ScatterPlot) string {
	roundedPlots := make([]PlotRounded, len(plots))
	for i, plot := range plots {
		roundedPlots[i] = PlotRounded{
			Labels: plot.Labels,
			Name:   plot.Name,
			Points: make([]PointRounded, len(plot.Points)),
		}
		for j, points := range plot.Points {
			roundedPlots[i].Points[j] = PointRounded{
				Color: points.Color,
				Label: points.Label,
				X:     RoundedScatterPoint(points.X),
				Y:     RoundedScatterPoint(points.Y),
			}
		}
	}

	jsonString, _ := json.Marshal(roundedPlots)
	return fmt.Sprintf("\"plots\": %s", string(jsonString))
}
