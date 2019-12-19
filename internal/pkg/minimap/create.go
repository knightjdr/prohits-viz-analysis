// Package minimap creates a "small" png for dotplots and heatmaps.
package minimap

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Data defines the type and variables required for generating a minimap
type Data struct {
	Filename  string
	ImageType string
	Matrices  *types.Matrices
	Settings  types.Settings
}

// Create a minimap for a dotplot or heatmap.
func Create(data *Data) {
	imageType := defineImageType(data)

	if downsample.Should(data.Matrices.Abundance) {
		data.Matrices.Abundance = downsample.Matrix(data.Matrices.Abundance, 1000)
	}

	fmt.Println(imageType)
}
