package convert

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/minimap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
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
