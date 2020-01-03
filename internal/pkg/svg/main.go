// Package svg has functions for generating and converting svg files.
package svg

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/convert"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/heatmap"
)

// ConvertToPNG converts an svg to a png using rsvg.
var ConvertToPNG = convert.RSVG

// InitializeDotplot creates a dotplot. Draw the dotplot by calling the Draw() method
var InitializeDotplot = dotplot.Initialize

// InitializeHeatmap creates a heatmap. Draw the heatmap by calling the Draw() method
var InitializeHeatmap = heatmap.Initialize
