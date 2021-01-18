package specificity

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/scatter"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createLegend(settings types.Settings) {
	legendPoints := []map[string]string{
		{
			"color": "#6e97ff",
			"text":  "Infinite specificity",
		},
		{
			"color": "#dfcd06",
			"text":  "Finite specificity",
		},
	}

	legendData := scatter.Legend{
		Filename: "svg/specificity-legend.svg",
		Points:   legendPoints,
		Title:    "specificity",
	}
	scatter.CreateLegend(legendData)

	if settings.Png {
		svg.ConvertToPNG("svg/specificity-legend.svg", "png/specificity-legend.png", "white")
	}
}
