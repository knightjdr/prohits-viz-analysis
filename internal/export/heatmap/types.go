package heatmap

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"

type heatmap struct {
	Annotations types.Annotations
	ColumnDB    []string
	ColumnOrder []int
	Markers     types.Markers
	RowDB       []rows
	RowOrder    []int
	Settings    types.Settings
}

type rows struct {
	Data []cell
	Name string
}

type cell struct {
	Value float64
}
