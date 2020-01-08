// Package dotplot exports images as a dotplot in svg or png format.
package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/export/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/matrix/frontend"
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

	matrices := frontend.CreateDotplotMatrices(data.RowDB, map[string][]int{"columns": data.ColumnOrder, "rows": data.RowOrder})
	matrices.Conditions = frontend.GetColumnNames(data.ColumnDB, data.ColumnOrder)
	matrices.Readouts = frontend.GetRowNames(data.RowDB, data.RowOrder)

	if settings.Format == "png" || downsample.Should(matrices.Abundance, settings.DownsampleThreshold) {
		createPNG(data, matrices, settings)
	} else {
		createSVG(data, matrices)
	}
}
