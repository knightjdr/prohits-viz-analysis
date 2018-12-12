// Package svg creates svg files for various image types.
package svg

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Heatmap creates a dotplot or heatmap from input matrices of abundance, abundance
// ratio and score.
func Heatmap(
	imageType string,
	heatmap typedef.Matrices,
	annotations typedef.Annotations,
	markers typedef.Markers,
	columns, rows []string,
	minimap bool,
	parameters typedef.Parameters,
) string {
	svg := make([]string, 0)
	dims := HeatmapDimensions(heatmap.Abundance, columns, rows, minimap)
	dotplotparameters := DotplotParameters(dims)
	svg = append(svg, HeatmapHeader(dims))
	if !minimap {
		svg = append(svg, HeatmapColumnNames(dims, columns))
		svg = append(svg, HeatmapRowNames(dims, rows))
	}
	if imageType == "dotplot" {
		svg = append(svg, DotplotRows(heatmap.Abundance, heatmap.Ratio, heatmap.Score, dims, dotplotparameters, parameters))
		if !minimap {
			svg = append(svg, BoundingBox(dims))
		}
	} else {
		svg = append(svg, HeatmapRows(heatmap.Abundance, dims, parameters))
	}
	if !minimap {
		svg = append(svg, HeatmapMarkers(markers, dims))
		svg = append(svg, HeatmapAnnotations(annotations, dims))
		svg = append(svg, HeatmapHeadings(dims, parameters))
	}
	// Add end element wrapper for svg.
	svg = append(svg, "</svg>\n")
	return helper.StringConcat(svg)
}
