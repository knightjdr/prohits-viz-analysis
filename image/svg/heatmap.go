// Package svg creates svg files for various image types.
package svg

import (
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/image/file"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/spf13/afero"
)

func write(svg *[]string, file afero.File) func(str string) {
	if file != nil {
		return func(str string) {
			file.WriteString(str)
		}
	}
	return func(str string) {
		*svg = append(*svg, str)
	}
}

// Heatmap creates a dotplot or heatmap from input matrices of abundance, abundance
// ratio and score.
func Heatmap(data *Data) string {
	svg := make([]string, 0)
	dims := Dimensions(data.Matrices, data.Minimap)
	dotplotparameters := DotplotParametersDefine(dims)

	// Open file for writing if requested
	file, err := file.Create(data.Path)
	logmessage.CheckError(err, true)
	if file != nil {
		defer file.Close()
	}

	writeString := write(&svg, file)

	Header(dims, writeString)
	if !data.Minimap {
		ColumnNames(dims, data.Matrices.Conditions, writeString)
		RowNames(dims, data.Matrices.Readouts, writeString)
	}
	if data.ImageType == "dotplot" {
		DotplotRows(data.Matrices, dims, dotplotparameters, data.Parameters, writeString)
		if !data.Minimap {
			BoundingBox(dims, writeString)
		}
	} else {
		HeatmapRows(data.Matrices.Abundance, dims, data.Parameters, writeString)
	}
	if !data.Minimap {
		Markers(data.Markers, dims, writeString)
		Annotations(data.Annotations, dims, writeString)
		Headings(data.Parameters, dims, writeString)
	}
	// Add end element wrapper for svg.
	writeString("</svg>\n")
	return helper.StringConcat(svg)
}
