package scatter

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

func writePoints(s *Scatter, axisLength float64, writeString func(string)) {
	str := fmt.Sprintf("\t\t<g transform=\"translate(100 0)\">\n"+
		"\t\t\t<defs>\n"+
		"\t\t\t\t<clipPath id=\"plot_points_clip\">\n"+
		"\t\t\t\t\t<rect height=\"%[1]s\" width=\"%[1]s\" x=\"0\" y=\"0\" />\n"+
		"\t\t\t\t</clipPath>\n"+
		"\t\t\t</defs>\n"+
		"\t\t\t<g clip-path=\"url(#plot_points_clip)\">\n",
		float.RemoveTrailingZeros(axisLength),
	)
	writeString(str)

	for _, point := range s.Plot {
		y := axisLength - point.Y
		str = fmt.Sprintf("\t\t\t\t<circle cx=\"%s\" cy=\"%s\" fill=\"%s\" r=\"%d\"><title>%s</title></circle>\n",
			float.RemoveTrailingZeros(math.Round(point.X, 0.01)),
			float.RemoveTrailingZeros(math.Round(y, 0.01)),
			point.Color,
			4,
			point.Label,
		)
		writeString(str)
	}

	writeString("\t\t\t</g>\n" +
		"\t\t</g>\n",
	)
}
