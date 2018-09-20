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
	annotations []typedef.Annotation,
	markers []typedef.Marker,
	columns, rows []string,
	options map[string]interface{},
) string {
	svg := make([]string, 0)
	dims := HeatmapDimensions(matrix, columns, rows)
	params := DotplotParameters(dims)
	svg = append(svg, HeatmapHeader(dims))
	svg = append(svg, HeatmapColumnNames(dims, columns))
	svg = append(svg, HeatmapRowNames(dims, rows))
	svg = append(svg, DotplotRows(matrix, ratios, scores, dims, params, options))
	svg = append(svg, HeatmapMarkers(markers, dims, options["markerColor"].(string)))
	svg = append(svg, HeatmapAnnotations(annotations, dims, options["annotationFontSize"].(int)))
	svg = append(svg, BoundingBox(dims))
	svg = append(svg, HeatmapHeadings(dims, options))
	// Add end element wrapper for svg.
	svg = append(svg, "</svg>\n")
	return helper.StringConcat(svg)
}
