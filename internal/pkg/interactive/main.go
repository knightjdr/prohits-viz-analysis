// Package interactive generates files for the interactive viewer.
package interactive

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/interactive/heatmap"
)

// CreateHeatmap generates a file for the interactive heatmap/dotplot viewer.
var CreateHeatmap = heatmap.Create

// HeatmapData contains data for generating a heatmap.
type HeatmapData = heatmap.Data
