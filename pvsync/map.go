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
	parameters := map[string]interface{}{
		"colLabel":     "Conditions",
		"edgeColor":    userParams.EdgeColor,
		"fillColor":    userParams.FillColor,
		"invertColor":  userParams.InvertColor,
		"abundanceCap": userParams.AbundanceCap,
		"primary":      userParams.PrimaryFilter,
		"rowLabel":     "Readouts",
		"secondary":    userParams.SecondaryFilter,
		"scoreType":    userParams.ScoreType,
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