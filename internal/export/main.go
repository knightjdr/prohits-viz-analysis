// Package export creates images in png or svg format.
package export

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/export/heatmap"
)

// Image exports images in png or svg format.
func Image() {
	params := parseArguments()

	switch params.imageType {
	case "heatmap":
		settings := heatmap.Settings{
			FontPath: params.fontPath,
			Format:   params.format,
		}
		heatmap.Export(params.jsonFile, settings)
	}
}
