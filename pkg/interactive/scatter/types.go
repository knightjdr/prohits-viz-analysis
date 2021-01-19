package scatter

import (
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Data defines the type and variables required for generating an interactive scatter plot
type Data struct {
	AnalysisType string
	Filename     string
	Legend       []map[string]string
	Parameters   types.Settings
	Plots        []types.ScatterPlot
	Settings     map[string]interface{}
}

// PlotRounded contains data for a scatter plot
type PlotRounded struct {
	Labels types.ScatterAxesLabels `json:"labels"`
	Name   string                  `json:"name"`
	Points []PointRounded          `json:"points"`
}

// PointRounded contains data for a point on a scatter plot
type PointRounded struct {
	Color string              `json:"color"`
	Label string              `json:"label"`
	X     RoundedScatterPoint `json:"x"`
	Y     RoundedScatterPoint `json:"y"`
}

// RoundedScatterPoint rounds scatter points to a precision of 2
type RoundedScatterPoint float64

// MarshalJSON rounds scatter points to a precision of 2 when exporting to JSON
func (r RoundedScatterPoint) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(r), 'f', 2, 64)), nil
}
