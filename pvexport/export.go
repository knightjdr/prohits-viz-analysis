package main

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Export draws a bait prey heatmap/dotplot.
func Export(
	imageType string,
	abundance, ratios, scores [][]float64,
	annotations []typedef.Annotation,
	markers []typedef.Marker,
	columns, rows []string,
	userParams typedef.Parameters,
) {
	// Define dotplot parameters.
	params := map[string]interface{}{
		"annotationFontSize": userParams.AnnotationFontSize,
		"colLabel":           "Baits",
		"edgeColor":          userParams.EdgeColor,
		"fillColor":          userParams.FillColor,
		"invert":             userParams.Invert,
		"markerColor":        userParams.MarkerColor,
		"maximumAbundance":   userParams.MaximumAbundance,
		"primary":            userParams.PrimaryFilter,
		"rowLabel":           "Preys",
		"secondary":          userParams.SecondaryFilter,
		"scoreType":          userParams.ScoreType,
	}
	var content string
	if imageType == "dotplot" {
		content = svg.Dotplot(abundance, ratios, scores, annotations, markers, columns, rows, params)
	} else {
		content = svg.Heatmap(abundance, annotations, markers, columns, rows, params)
	}
	filename := fmt.Sprintf("svg/%s.svg", imageType)
	afero.WriteFile(fs.Instance, filename, []byte(content), 0644)
	return
}
