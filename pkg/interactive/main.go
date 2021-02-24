// Package interactive generates files for the interactive viewer.
package interactive

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive/circheatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/interactive/scatter"
)

// CreateCircHeatmap generates a file for the interactive circheatmap viewer.
var CreateCircHeatmap = circheatmap.Create

// CreateHeatmap generates a file for the interactive heatmap/dotplot viewer.
var CreateHeatmap = heatmap.Create

// CreateScatter generates a file for the interactive scatter plot viewer.
var CreateScatter = scatter.Create

// CircHeatmapData contains data for generating a heatmap.
type CircHeatmapData = circheatmap.Data

// HeatmapData contains data for generating a heatmap.
type HeatmapData = heatmap.Data

// ScatterData contains data for generating a heatmap.
type ScatterData = scatter.Data
