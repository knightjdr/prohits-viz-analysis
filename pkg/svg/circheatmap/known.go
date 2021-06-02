package circheatmap

import (
	"fmt"
	"math"

	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

type pathParameters struct {
	Arc    int
	Circle bool
	X      float64
	Y      float64
}

func writeKnown(c *CircHeatmap, writeString func(string)) {
	if c.ShowKnown {
		path := definePath(c.Plot.Readouts, c.Dimensions.Radius)

		str := ""
		if path.Circle {
			str = fmt.Sprintf(
				"\t\t<circle cx=\"0\" cy=\"0\" fill=\"none\" r=\"%.2f\" stroke=\"#333\" stroke-width=\"5\" transform=\"scale(0.9 0.9)\"/>\n",
				c.Dimensions.Radius,
			)
		} else {
			str = fmt.Sprintf(
				"\t\t<path d=\"M %.2[1]f 0 A %.2[1]f %.2[1]f 0 %[2]d 1 %.2[3]f %.2[4]f\" fill=\"none\" stroke=\"#333\" stroke-width=\"5\" transform=\"scale(0.9 0.9)\"/>\n",
				c.Dimensions.Radius,
				path.Arc,
				path.X,
				path.Y,
			)
		}
		writeString(str)
	}
}

func definePath(readouts []types.CircHeatmapReadout, radius float64) pathParameters {
	known := 0
	for _, readout := range readouts {
		if readout.Known {
			known += 1
		}
	}

	numReadouts := len(readouts)
	percent := customMath.Round(float64(known)/float64(numReadouts), 0.0001)

	if percent == 1 {
		return pathParameters{
			Circle: true,
		}
	}

	arc := 0
	if percent > 0.5 {
		arc = 1
	}
	point := percentToCoordinate(percent, radius)
	return pathParameters{
		Arc:    arc,
		Circle: false,
		X:      point[0],
		Y:      point[1],
	}
}

func percentToCoordinate(percent float64, radius float64) []float64 {
	x := customMath.Round(radius*math.Cos(2*math.Pi*percent), 0.00001)
	y := customMath.Round(radius*math.Sin(2*math.Pi*percent), 0.00001)
	return []float64{x, y}
}
