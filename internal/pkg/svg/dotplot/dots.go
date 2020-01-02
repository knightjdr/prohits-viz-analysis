package dotplot

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

func writeDots(d *Dotplot, writeString func(string)) {
	fillGradient, edgeGradient := createGradients(d)
	parameters := defineParameters(d)

	convertScoreToIndex := defineScoreConverter(d, d.NumColors)
	convertValueToIndex := float.GetRange(d.MinAbundance, d.AbundanceCap, 0, 100)

	writeString(fmt.Sprintf("\t<g id=\"minimap\" transform=\"translate(%d, %d)\">\n", d.LeftMargin, d.TopMargin))
	for i, row := range d.Matrices.Abundance {
		x := (i * d.CellSize) + parameters.cellSizeHalf

		for j, value := range row {
			fillIndex := int(convertValueToIndex(value))
			fill := fillGradient[fillIndex].Hex

			edgeIndex := convertScoreToIndex(d.Matrices.Score[i][j])
			edge := edgeGradient[edgeIndex].Hex

			radius := math.Round(d.Matrices.Ratio[i][j]*parameters.maxRadius, 0.01)

			circle := fmt.Sprintf(
				"\t\t<circle fill=\"%s\" cy=\"%d\" cx=\"%d\" r=\"%f\" stroke=\"%s\" stroke-width=\"%f\"/>\n",
				fill, x, (j*d.CellSize)+parameters.cellSizeHalf, radius, edge, parameters.edgeWidth,
			)
			writeString(circle)
		}
	}
	writeString("\t</g>\n")
}
