// Package heatmap exports images in png or svg format.
package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// Settings for heatmap export.
type Settings struct {
	DownsampleThreshold int
	FontPath            string
	Format              string
}

// Export image.
func Export(filename string, settings Settings) {
	data := ReadJSON(filename)

	matrix := createMatrix(data)
	columns, rows := GetColumnsAndRows(data)

	matrices := &types.Matrices{
		Abundance:  matrix,
		Conditions: columns,
		Readouts:   rows,
	}

	if settings.Format == "png" || downsample.Should(matrix, settings.DownsampleThreshold) {
		createPNG(data, matrices, settings)
	} else {
		createSVG(data, matrices)
	}
}
