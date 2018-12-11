package main

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Map draws a condition readout heatmap/dotplot.
func Map(
	imageType string,
	abundance, ratios, scores [][]float64,
	columns, rows []string,
	userParams typedef.Parameters,
) {
	// Define dotplot parameters.
	parameters := userParams
	if parameters.Condition == "" {
		parameters.Condition = "Conditions"
	}
	if parameters.Readout == "" {
		parameters.Readout = "Readouts"
	}
	var content string
	if imageType == "dotplot" {
		content = svg.Dotplot(abundance, ratios, scores, typedef.Annotations{}, typedef.Markers{}, columns, rows, true, parameters)
	} else {
		content = svg.Heatmap(abundance, typedef.Annotations{}, typedef.Markers{}, columns, rows, true, parameters)
	}
	filename := fmt.Sprintf("minimap/%s.svg", imageType)
	afero.WriteFile(fs.Instance, filename, []byte(content), 0644)
	return
}
