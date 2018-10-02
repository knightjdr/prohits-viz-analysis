package svg

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// DParameters conatins dotplot parameters.
type DParameters struct {
	cellSizeHalf int
	edgeWidth    float64
	maxRadius    float64
}

// DotplotParameters generates dot dimensions for an SVG.
func DotplotParameters(dims HDimensions) (parameters DParameters) {
	parameters.cellSizeHalf = int(math.Round(float64(dims.cellSize) / float64(2)))

	// Determine the amount to subtract for the maximum circle radius. We do this
	// so that there is some space between circles. For the ideal cell size of 20
	// (half size 10), I subtract 1 pixel from the radius to create 2 pixels of
	// padding between circles, but this can go down as the size of the cell
	// decreases.
	circleSpace := dims.ratio * idealCircleSpace

	// Adjust edgeWidth as cell size gets smaller.
	parameters.edgeWidth = helper.Round(dims.ratio*float64(idealEdgeWidth), 0.01)

	// Determine the maxium circle radius. Will be 9px for the default cell size
	// of 20px.
	parameters.maxRadius = helper.Round(float64(parameters.cellSizeHalf)-circleSpace, 0.01)

	return
}
