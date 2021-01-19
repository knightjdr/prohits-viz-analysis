package specificity

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeImages(data map[string][]types.ScatterPoint, settings types.Settings) {
	for condition, conditionData := range data {
		filehandle := fmt.Sprintf("specificity-%s", condition)
		createSVG(conditionData, settings, filehandle)

		if settings.Png {
			svg.ConvertToPNG(fmt.Sprintf("svg/%s.svg", filehandle), fmt.Sprintf("png/%s.png", filehandle), "white")
		}
	}
}

func createSVG(data []types.ScatterPoint, settings types.Settings, filehandle string) {
	filename := fmt.Sprintf("svg/%s.svg", filehandle)

	scatter := svg.InitializeScatter()
	scatter.LogBase = settings.LogBase
	scatter.Plot = data
	scatter.XLabel = settings.Abundance
	scatter.YLabel = "Specificity"

	scatter.Draw(filename)
}
