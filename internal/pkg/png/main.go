// Package png has functions for generating and converting png images.
package png

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/png/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/png/uri"
)

// ConvertToURI converts a png file to a uri.
var ConvertToURI = uri.Convert

// InitializeHeatmap creates a heatmap. Draw the heatmap by calling the Draw() method.
var InitializeHeatmap = heatmap.Initialize
