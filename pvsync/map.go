package main

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Map draws a bait prey heatmap/dotplot.
func Map(
	imageType string,
	abundance, ratios, scores [][]float64,
	columns, rows []string,
	userParams typedef.Parameters,
) {
	// Define dotplot parameters.
	params := map[string]interface{}{
		"annotationFontSize": 12,
		"colLabel":           "Baits",
		"edgeColor":          userParams.EdgeColor,
		"fillColor":          userParams.FillColor,
		"invert":             userParams.Invert,
		"maximumAbundance":   userParams.MaximumAbundance,
		"primary":            userParams.PrimaryFilter,
		"rowLabel":           "Preys",
		"secondary":          userParams.SecondaryFilter,
		"scoreType":          userParams.ScoreType,
	}
	var content string
	if imageType == "dotplot" {
		content = svg.Dotplot(abundance, ratios, scores, []typedef.Annotation{}, []typedef.Marker{}, columns, rows, params)
	} else {
		content = svg.Heatmap(abundance, []typedef.Annotation{}, []typedef.Marker{}, columns, rows, params)
	}
	filename := fmt.Sprintf("svg/%s.svg", imageType)
	afero.WriteFile(fs.Instance, filename, []byte(content), 0644)
	return
}
