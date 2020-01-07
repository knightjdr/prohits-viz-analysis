// Package dotplot exports images as a dotplot in svg or png format.
package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/export/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
)

// Settings for dotplot export.
type Settings struct {
	DownsampleThreshold int
	FontPath            string
	Format              string
}

// Export image.
func Export(filename string, settings Settings) {
	data := heatmap.ReadJSON(filename)

	matrices := createMatrices(data)
	matrices.Conditions, matrices.Readouts = heatmap.GetColumnsAndRows(data)

	if settings.Format == "png" || downsample.Should(matrices.Abundance, settings.DownsampleThreshold) {
		createPNG(data, matrices, settings)
	} else {
		createSVG(data, matrices)
	}
}
