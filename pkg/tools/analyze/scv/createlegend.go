package scv

import (
	"fmt"
	"math"

	heatmapColor "github.com/knightjdr/prohits-viz-analysis/pkg/heatmap/color"
	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/circheatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createLegend(data map[string]map[string]map[string]float64, settings types.Settings) types.CircHeatmapLegend {
	elements := make([]types.CircHeatmapLegendElement, 0)
	addFileMetrics(&elements, data, settings)
	addProteinExpression(&elements, settings.ProteinTissues)
	addRNAExpression(&elements, settings.RnaTissues)

	writeLegend(elements, settings)

	return elements
}

func addProteinExpression(elements *[]types.CircHeatmapLegendElement, proteinTissues []string) {
	for _, metric := range proteinTissues {
		*elements = append(
			*elements,
			types.CircHeatmapLegendElement{
				Attribute: fmt.Sprintf("Protein expression - %s", metric),
				Color:     "red",
				Filter:    0,
				Max:       7,
				Min:       0,
			},
		)
	}
}

func addRNAExpression(elements *[]types.CircHeatmapLegendElement, rnaTissues []string) {
	for _, metric := range rnaTissues {
		*elements = append(
			*elements,
			types.CircHeatmapLegendElement{
				Attribute: fmt.Sprintf("RNA expression - %s", metric),
				Color:     "green",
				Filter:    0,
				Max:       50,
				Min:       0,
			},
		)
	}
}

func addFileMetrics(elements *[]types.CircHeatmapLegendElement, data map[string]map[string]map[string]float64, settings types.Settings) {
	metrics := []string{settings.Abundance}
	metrics = append(metrics, settings.OtherAbundance...)

	if settings.Specificity {
		metrics = append(metrics, "Specificity")
	}

	for _, metric := range metrics {
		minAbundance := settings.MinAbundance
		if metric == "Specificity" {
			minAbundance = 0
		}

		metricSettings := types.Settings{
			AbundanceCap:         settings.AbundanceCap,
			AbundanceType:        defineValues(data, metric),
			AutomaticallySetFill: true,
			MinAbundance:         minAbundance,
		}
		heatmapColor.SetFillLimits(&metricSettings)
		heatmapColor.AdjustFillColor(&metricSettings)

		*elements = append(
			*elements,
			types.CircHeatmapLegendElement{
				Attribute: metric,
				Color:     metricSettings.FillColor,
				Filter:    math.Abs(minAbundance),
				Max:       metricSettings.FillMax,
				Min:       metricSettings.FillMin,
			},
		)
	}
}

func defineValues(data map[string]map[string]map[string]float64, metric string) string {
	values := make([]float64, 0)
	for _, conditionData := range data {
		for _, readoutData := range conditionData {
			values = append(values, readoutData[metric])
		}
	}

	max := customMath.MaxSliceFloat(values)
	min := customMath.MinSliceFloat(values)

	if max > 0 && min < 0 {
		return "bidirectional"
	}
	if max <= 0 && min < 0 {
		return "negative"
	}
	return "positive"
}

func writeLegend(elements types.CircHeatmapLegend, settings types.Settings) {
	legendData := circheatmap.Legend{
		Elements: elements,
		Filename: "svg/scv-legend.svg",
		Known:    settings.Known,
		Title:    "single condition visualization",
	}
	circheatmap.CreateLegend(legendData)

	if settings.Png {
		svg.ConvertToPNG("svg/scv-legend.svg", "png/scv-legend.png", "white")
	}
}
