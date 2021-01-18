package cc

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeImages(data []types.ScatterPoint, settings types.Settings) {
	filehandle := fmt.Sprintf("%s-%s", settings.ConditionX, settings.ConditionY)
	createSVG(data, settings, filehandle)

	if settings.Png {
		svg.ConvertToPNG(fmt.Sprintf("svg/%s.svg", filehandle), fmt.Sprintf("png/%s.png", filehandle), "white")
	}
}

func createSVG(data []types.ScatterPoint, settings types.Settings, filehandle string) {
	filename := fmt.Sprintf("svg/%s.svg", filehandle)

	scatter := svg.InitializeScatter()
	scatter.LogBase = settings.LogBase
	scatter.Plot = data
	scatter.XLabel = settings.ConditionX
	scatter.YLabel = settings.ConditionY

	scatter.Draw(filename)
}
