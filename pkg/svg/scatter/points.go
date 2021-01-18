package scatter

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

func writePoints(s *Scatter, axisLength float64, writeString func(string)) {
	writeString("\t\t<g transform=\"translate(100 0)\">\n")

	for _, point := range s.Plot {
		y := axisLength - point.Y
		str := fmt.Sprintf("\t\t\t<circle cx=\"%s\" cy=\"%s\" fill=\"%s\" r=\"%d\"><title>%s</title></circle>\n",
			float.RemoveTrailingZeros(math.Round(point.X, 0.01)),
			float.RemoveTrailingZeros(math.Round(y, 0.01)),
			point.Color,
			4,
			point.Label,
		)
		writeString(str)
	}

	writeString("\t\t</g>\n")
}
