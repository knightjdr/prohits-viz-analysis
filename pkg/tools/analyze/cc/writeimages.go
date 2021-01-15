package cc

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeImages(data []types.ScatterPoint, settings types.Settings) {
	createSVG(data, settings)

	if settings.Png {

	}
}

func createSVG(data []types.ScatterPoint, settings types.Settings) {
	filename := fmt.Sprintf("svg/%s-%s.svg", settings.ConditionX, settings.ConditionY)

	scatter := svg.InitializeScatter()
	scatter.LogBase = settings.LogBase
	scatter.Plot = data
	scatter.XLabel = settings.ConditionX
	scatter.YLabel = settings.ConditionY

	scatter.Draw(filename)
}
