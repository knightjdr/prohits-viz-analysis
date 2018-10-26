// Package svg creates svg files for various image types.
package svg

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// ScoreColorFunc returns a function for determining the gradient index to use
// for the score color.
func ScoreColorFunc(scoretype string, primary, secondary float64, numColors int) func(score float64) int {
	if scoretype == "gte" {
		return func(score float64) int {
			if score >= primary {
				return numColors
			} else if score < primary && score >= secondary {
				return numColors / 2
			}
			return numColors / 4
		}
	}
	return func(score float64) int {
		if score <= primary {
			return numColors
		} else if score > primary && score <= secondary {
			return numColors / 2
		}
		return numColors / 4
	}
}

// Dotplot creates a dotplot from an input matrices of abundance, abundance
// ratios and score.
func Dotplot(
	matrix, ratios, scores [][]float64,
	annotations typedef.Annotations,
	markers typedef.Markers,
	columns, rows []string,
	minimap bool,
	options map[string]interface{},
) string {
	svg := make([]string, 0)
	dims := HeatmapDimensions(matrix, columns, rows, minimap)
	parameters := DotplotParameters(dims)
	svg = append(svg, HeatmapHeader(dims))
	if !minimap {
		svg = append(svg, HeatmapColumnNames(dims, columns))
		svg = append(svg, HeatmapRowNames(dims, rows))
	}
	svg = append(svg, DotplotRows(matrix, ratios, scores, dims, parameters, options))
	if !minimap {
		svg = append(svg, HeatmapMarkers(markers, dims))
		svg = append(svg, HeatmapAnnotations(annotations, dims))
		svg = append(svg, BoundingBox(dims))
		svg = append(svg, HeatmapHeadings(dims, options))
	}
	// Add end element wrapper for svg.
	svg = append(svg, "</svg>\n")
	return helper.StringConcat(svg)
}
