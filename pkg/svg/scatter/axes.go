package scatter

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

func writeAxes(s *Scatter, axisLength float64, writeString func(string)) {
	writeString("\t\t<g>\n")
	writeXAxis(s, axisLength, writeString)
	writeYAxis(s, axisLength, writeString)
	writeString("\t\t</g>\n")
}

func writeXAxis(s *Scatter, axisLength float64, writeString func(string)) {
	str := fmt.Sprintf("\t\t\t<g transform=\"translate(0 %[1]s)\">\n"+
		"\t\t\t\t<rect width=\"%[2]s\" height=\"100\" fill=\"white\" />\n"+
		"\t\t\t\t<g transform=\"translate(100 0)\">\n"+
		"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"%[1]s\" y1=\"0\" y2=\"0\"/>\n"+
		"\t\t\t\t\t<g>\n",
		float.RemoveTrailingZeros(axisLength),
		float.RemoveTrailingZeros(s.PlotSize),
	)
	writeString(str)

	for i, tick := range s.Ticks.X {
		str = fmt.Sprintf(
			"\t\t\t\t\t\t<g>\n"+
				"\t\t\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"%[1]s\" x2=\"%[1]s\" y1=\"0\" y2=\"10\" />\n"+
				"\t\t\t\t\t\t\t<text font-size=\"12\" text-anchor=\"middle\" x=\"%[1]s\" y=\"28\">%[2]s</text>\n"+
				"\t\t\t\t\t\t</g>\n",
			float.RemoveTrailingZeros(math.Round(tick, 0.01)),
			s.Ticks.XLabel[i],
		)
		writeString(str)
	}

	str = fmt.Sprintf(
		"\t\t\t\t\t</g>\n"+
			"\t\t\t\t\t<text text-anchor=\"middle\" x=\"%[1]s\" y=\"70\">%[2]s</text>\n"+
			"\t\t\t\t</g>\n"+
			"\t\t\t</g>\n",
		float.RemoveTrailingZeros(math.Round(axisLength/2, 0.01)),
		s.XLabel,
	)
	writeString(str)
}

func writeYAxis(s *Scatter, axisLength float64, writeString func(string)) {
	str := fmt.Sprintf("\t\t\t<g>\n"+
		"\t\t\t\t<rect width=\"100\" height=\"%[1]s\" fill=\"white\" />\n"+
		"\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"100\" x2=\"100\" y1=\"0\" y2=\"%[1]s\"/>\n"+
		"\t\t\t\t<g>\n",
		float.RemoveTrailingZeros(axisLength),
	)
	writeString(str)

	for i, tick := range s.Ticks.Y {
		yPosition := axisLength - tick
		str = fmt.Sprintf(
			"\t\t\t\t\t<g>\n"+
				"\t\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"100\" x2=\"90\" y1=\"%[1]s\" y2=\"%[1]s\" />\n"+
				"\t\t\t\t\t\t<text dy=\"4\" font-size=\"12\" text-anchor=\"end\" x=\"87\" y=\"%[1]s\">%[2]s</text>\n"+
				"\t\t\t\t\t</g>\n",
			float.RemoveTrailingZeros(math.Round(yPosition, 0.01)),
			s.Ticks.YLabel[i],
		)
		writeString(str)
	}

	str = fmt.Sprintf(
		"\t\t\t\t</g>\n"+
			"\t\t\t\t<text dy=\"15\" text-anchor=\"middle\" transform=\"rotate(-90, 0, %[1]s)\" x=\"0\" y=\"%[1]s\">%[2]s</text>\n"+
			"\t\t\t</g>\n",
		float.RemoveTrailingZeros(math.Round(axisLength/2, 0.01)),
		s.YLabel,
	)
	writeString(str)
}
