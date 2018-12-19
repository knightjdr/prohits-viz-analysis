/*Package minimap creates a png image from a 2D data matrix.*/
package minimap

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/image/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Write first creates an svg for that input datat matrix and then uses
// rsvg to convert that to a png. If the matrix has a dimension greater
// than 500, it will always generate a heatmap, otherwise it will output
// whatever is requested. Matrices larger than 1000 in either dimension
// are downsampled as the output png will be max 1000x1000.
func Write(data *Data) {
	imageType := "heatmap"
	if data.ImageType == "dotplot" && len(data.Matrices.Abundance) <= 500 && len(data.Matrices.Abundance[0]) <= 500 {
		imageType = "dotplot"
	}

	svgName := fmt.Sprintf("%s.svg", data.Filename)
	svgData := svg.Data{
		Annotations: typedef.Annotations{},
		Filename:    svgName,
		ImageType:   imageType,
		Markers:     typedef.Markers{},
		Matrices:    data.Matrices,
		Minimap:     true,
		Parameters:  data.Parameters,
	}
	svg.Heatmap(&svgData)

	// Generate minimap
	svg.ConvertMap(svgName)
}
