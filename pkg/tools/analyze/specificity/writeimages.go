package specificity

import (
	"fmt"
	"sync"

	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeImages(data map[string][]types.ScatterPoint, settings types.Settings) {
	var wg sync.WaitGroup
	wg.Add(len(data))

	for condition, conditionData := range data {
		go func(condition string, conditionData []types.ScatterPoint) {
			defer wg.Done()
			filehandle := fmt.Sprintf("specificity-%s", condition)
			createSVG(conditionData, settings, filehandle)

			if settings.Png {
				svg.ConvertToPNG(fmt.Sprintf("svg/%s.svg", filehandle), fmt.Sprintf("png/%s.png", filehandle), "white")
			}
		}(condition, conditionData)
	}

	wg.Wait()
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
