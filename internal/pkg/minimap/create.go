// Package minimap creates a "small" png for dotplots and heatmaps.
package minimap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/heatmap/dimensions"
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
	imageType := defineImageType(data)
	downsampleIfNeeded(data)

	dims := dimensions.Calculate(data.Matrices.Abundance, []string{}, []string{}, true)

	if imageType == "dotplot" {
		createDotplot(data, dims)
	} else {
		createHeatmap(data, dims)
	}
}
