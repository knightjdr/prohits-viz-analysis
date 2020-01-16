// Package minimap creates a "small" png for dotplots and heatmaps.
package minimap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Data defines the type and variables required for generating a minimap
type Data struct {
	DownsampleThreshold int
	Filename            string
	ImageType           string
	Matrices            *types.Matrices
	Settings            types.Settings
}

// Create a minimap for a dotplot or heatmap.
func Create(data *Data) {
	if !downsample.Should(data.Matrices.Abundance, data.DownsampleThreshold) && data.ImageType == "dotplot" {
		createDotplot(data)
	} else {
		createHeatmap(data)
	}
}
