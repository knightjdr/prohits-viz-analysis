package main

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Export draws a condition readout heatmap/dotplot.
func Export(
	imageType string,
	abundance, ratios, scores [][]float64,
	annotations typedef.Annotations,
	markers typedef.Markers,
	columns, rows []string,
	userParams typedef.Parameters,
) {
	parameters := userParams
	if parameters.Condition == "" {
		parameters.Condition = "Conditions"
	}
	if parameters.Readout == "" {
		parameters.Readout = "Readouts"
	}

	var content string
	if imageType == "dotplot" {
		content = svg.Dotplot(abundance, ratios, scores, annotations, markers, columns, rows, false, parameters)
	} else {
		content = svg.Heatmap(abundance, annotations, markers, columns, rows, false, parameters)
	}
	filename := fmt.Sprintf("svg/%s.svg", imageType)
	afero.WriteFile(fs.Instance, filename, []byte(content), 0644)
	return
}
