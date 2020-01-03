package heatmap

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"

// Data defines the type and variables required for generating a minimap
type Data struct {
	Filename  string
	ImageType string
	Matrices  *types.Matrices
	Minimap   string
	Settings  types.Settings
}
