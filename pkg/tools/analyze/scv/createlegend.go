package scv

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/svg"
	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/circheatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createLegend(settings types.Settings) []circheatmap.LegendElement {
	metrics := map[string][]string{
		"file":              defineFileMetrics(settings),
		"proteinExpression": settings.ProteinTissues,
		"rnaExpression":     settings.RnaTissues,
	}

	elements := createElements(metrics, settings)

	writeLegend(elements, settings)

	return elements
}

func defineFileMetrics(settings types.Settings) []string {
	metrics := []string{settings.Abundance}
	metrics = append(metrics, settings.OtherAbundance...)

	if settings.Specificity {
		metrics = append(metrics, "Specificity")
	}

	return metrics
}

func createElements(metrics map[string][]string, settings types.Settings) []circheatmap.LegendElement {
	colorScheme := defineColorScheme(metrics)

	elements := make([]circheatmap.LegendElement, 0)
	for index, metric := range metrics["file"] {
		elements = append(
			elements,
			circheatmap.LegendElement{
				Attribute: metric,
				Color:     colorScheme("file", index),
				Max:       settings.AbundanceCap,
				Min:       settings.MinAbundance,
			},
		)
	}

	for index, metric := range metrics["proteinExpression"] {
		elements = append(
			elements,
			circheatmap.LegendElement{
				Attribute: fmt.Sprintf("Protein expression - %s", metric),
				Color:     colorScheme("proteinExpression", index),
				Max:       7,
				Min:       0,
			},
		)
	}

	for index, metric := range metrics["rnaExpression"] {
		elements = append(
			elements,
			circheatmap.LegendElement{
				Attribute: fmt.Sprintf("RNA expression - %s", metric),
				Color:     colorScheme("rnaExpression", index),
				Max:       50,
				Min:       0,
			},
		)
	}

	return elements
}

func defineColorScheme(metrics map[string][]string) func(string, int) string {
	if len(metrics["rnaExpression"]) > 0 || len(metrics["proteinExpression"]) > 0 {
		return func(metricType string, index int) string {
			if metricType == "proteinExpression" {
				return "red"
			}
			if metricType == "rnaExpression" {
				return "green"
			}
			return "blue"
		}
	}
	return func(metricType string, index int) string {
		colors := []string{"blue", "red", "green", "grey"}
		if index <= 2 {
			return colors[index]
		}
		return colors[3]
	}
}

func writeLegend(elements []circheatmap.LegendElement, settings types.Settings) {
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
