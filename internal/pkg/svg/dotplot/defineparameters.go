package dotplot

import (
	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"math"
)

const idealCircleSpace float64 = 1.5
const idealEdgeWidth int = 2

type dotplotparameters struct {
	cellSizeHalf int
	edgeWidth    float64
	maxRadius    float64
}

func defineParameters(d *Dotplot) dotplotparameters {
	parameters := dotplotparameters{}

	parameters.cellSizeHalf = int(math.Round(float64(d.CellSize) / float64(2)))

	// Determine the amount to subtract for the maximum circle radius. I do this
	// so that there is some space between circles. For the ideal cell size of 20
	// (half size 10), I subtract 1 pixel from the radius to create 2 pixels of
	// padding between circles, but this can go down as the size of the cell
	// decreases.
	circleSpace := d.Ratio * idealCircleSpace

	// Adjust edgeWidth as cell size gets smaller.
	parameters.edgeWidth = customMath.Round(d.Ratio*float64(idealEdgeWidth), 0.01)

	// Determine the maxium circle radius. Will be 9px for the default cell size
	// of 20px.
	parameters.maxRadius = customMath.Round(float64(parameters.cellSizeHalf)-circleSpace, 0.01)

	return parameters
}
