// Package sync will create a minimap from settings.
package sync

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	heatmapColor "github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/color"
	createMinimap "github.com/knightjdr/prohits-viz-analysis/pkg/minimap"
)

// Minimap synchronizes a minimap with supplied settings.
func Minimap() {
	jsonFile := parseArguments()

	data := readJSON(jsonFile)
	heatmapColor.SetFillLimits(&data.Settings)
	matrices := createMatrices(data)

	files.CreateFolders([]string{"minimap"})
	minimapData := &createMinimap.Data{
		DownsampleThreshold: 1000,
		Filename:            "minimap/minimap.png",
		ImageType:           data.ImageType,
		Matrices:            matrices,
		Settings:            data.Settings,
	}

	createMinimap.Create(minimapData)
}
