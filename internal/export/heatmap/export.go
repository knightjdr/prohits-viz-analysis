// Package heatmap exports images in png or svg format.
package heatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/matrix/frontend"
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

	matrices := &types.Matrices{
		Abundance:  frontend.CreateHeatmapMatrix(data.RowDB, map[string][]int{"columns": data.ColumnOrder, "rows": data.RowOrder}),
		Conditions: frontend.GetColumnNames(data.ColumnDB, data.ColumnOrder),
		Readouts:   frontend.GetRowNames(data.RowDB, data.RowOrder),
	}

	if settings.Format == "png" || downsample.Should(matrices.Abundance, settings.DownsampleThreshold) {
		createPNG(data, matrices, settings)
	} else {
		createSVG(data, matrices)
	}
}
