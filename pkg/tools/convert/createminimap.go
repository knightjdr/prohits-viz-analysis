package convert

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/minimap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createMinimap(matrices *types.Matrices, settings types.Settings) {
	mapData := &minimap.Data{
		Filename:  "./minimap/minimap.png",
		ImageType: settings.Type,
		Matrices:  matrices,
		Settings:  settings,
	}
	minimap.Create(mapData)
}
