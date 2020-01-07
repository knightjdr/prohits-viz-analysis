package heatmap

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"

// Heatmap data for exporting image.
type Heatmap struct {
	Annotations types.Annotations
	ColumnDB    []string
	ColumnOrder []int
	Markers     types.Markers
	RowDB       []Rows
	RowOrder    []int
	Settings    types.Settings
}

// Rows contains heatmap row information.
type Rows struct {
	Data []Cell
	Name string
}

// Cell contains individual heatmap cell information.
type Cell struct {
	Ratio float64
	Score float64
	Value float64
}
