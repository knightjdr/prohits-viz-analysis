package heatmap

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"

// Data defines the type and variables required for generating a minimap
type Data struct {
	AnalysisType string
	Filename     string
	Matrices     *types.Matrices
	Minimap      string
	Parameters   types.Settings
	Settings     map[string]interface{}
}
