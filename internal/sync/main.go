// Package sync will create a minimap from settings.
package sync

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
	createMinimap "github.com/knightjdr/prohits-viz-analysis/internal/pkg/minimap"
)

// Minimap synchronizes a minimap with supplied settings.
func Minimap() {
	jsonFile := parseArguments()

	data := readJSON(jsonFile)
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
