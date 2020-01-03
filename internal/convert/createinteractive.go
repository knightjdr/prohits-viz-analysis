package convert

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/interactive"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createInteractive(matrices *types.Matrices, settings types.Settings) {
	data := &interactive.HeatmapData{
		Filename:  fmt.Sprintf("interactive/%s.json", settings.Type),
		ImageType: settings.Type,
		Matrices:  matrices,
		Minimap:   "./minimap/minimap.png",
		Settings:  settings,
	}
	interactive.CreateHeatmap(data)
}
