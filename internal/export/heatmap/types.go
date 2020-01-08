package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/matrix/frontend"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Heatmap data for exporting image.
type Heatmap struct {
	Annotations types.Annotations
	ColumnDB    []string
	ColumnOrder []int
	Markers     types.Markers
	RowDB       []frontend.Row
	RowOrder    []int
	Settings    types.Settings
}
