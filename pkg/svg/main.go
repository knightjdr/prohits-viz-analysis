// Package svg has functions for generating and converting svg files.
package svg

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/convert"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/scatter"
)

// ConvertToPNG converts an svg to a png using rsvg.
var ConvertToPNG = convert.RSVG

// InitializeDotplot creates a dotplot. Draw the dotplot by calling the Draw() method
var InitializeDotplot = dotplot.Initialize

// InitializeHeatmap creates a heatmap. Draw the heatmap by calling the Draw() method
var InitializeHeatmap = heatmap.Initialize

// InitializeScatter creates a scatter plot. Draw the scatter plot by calling the Draw() method
var InitializeScatter = scatter.Initialize
