// Package svg creates svg files for various image types.
package svg

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Heatmap creates a heatmap from an input matrix.
func Heatmap(
	matrix [][]float64,
	annotations typedef.Annotations,
	markers typedef.Markers,
	columns, rows []string,
	options map[string]interface{},
) string {
	svg := make([]string, 0)
	dims := HeatmapDimensions(matrix, columns, rows)
	svg = append(svg, HeatmapHeader(dims))
	svg = append(svg, HeatmapColumnNames(dims, columns))
	svg = append(svg, HeatmapRowNames(dims, rows))
	svg = append(svg, HeatmapRows(matrix, dims, options))
	svg = append(svg, HeatmapMarkers(markers, dims))
	svg = append(svg, HeatmapAnnotations(annotations, dims))
	svg = append(svg, HeatmapHeadings(dims, options))
	// Add end element wrapper for svg.
	svg = append(svg, "</svg>\n")
	return helper.StringConcat(svg)
}
