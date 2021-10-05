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
	str := fmt.Sprintf("\t\t\t<g transform=\"translate(100 0)\">\n"+
		"\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"%s\" x2=\"%s\" y1=\"%s\" y2=\"%s\"/>\n",
		float.RemoveTrailingZeros(s.Axes.X.X1),
		float.RemoveTrailingZeros(s.Axes.X.X2),
		float.RemoveTrailingZeros(s.Axes.X.Y1),
		float.RemoveTrailingZeros(s.Axes.X.Y2),
	)
	writeString(str)

	tickLabelPosition := s.Axes.X.Y1 + 28
	tickEnd := s.Axes.X.Y1 + 10
	for i, tick := range s.Ticks.X {
		if s.Ticks.XLabel[i] != "0" {
			str = fmt.Sprintf("\t\t\t\t<g>\n"+
				"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"%[1]s\" x2=\"%[1]s\" y1=\"%[2]s\" y2=\"%[3]s\" />\n"+
				"\t\t\t\t\t<text font-size=\"12\" text-anchor=\"middle\" x=\"%[1]s\" y=\"%[4]s\">%[5]s</text>\n"+
				"\t\t\t\t</g>\n",
				float.RemoveTrailingZeros(math.Round(tick, 0.01)),
				float.RemoveTrailingZeros(s.Axes.X.Y1),
				float.RemoveTrailingZeros(tickEnd),
				float.RemoveTrailingZeros(tickLabelPosition),
				s.Ticks.XLabel[i],
			)
			writeString(str)
		}
	}

	str = fmt.Sprintf("\t\t\t\t<g transform=\"translate(0 %s)\">\n"+
		"\t\t\t\t\t<text text-anchor=\"middle\" x=\"%s\" y=\"70\">%s</text>\n"+
		"\t\t\t\t</g>\n"+
		"\t\t\t</g>\n",
		float.RemoveTrailingZeros(axisLength),
		float.RemoveTrailingZeros(math.Round(axisLength/2, 0.01)),
		s.XLabel,
	)
	writeString(str)
}

func writeYAxis(s *Scatter, axisLength float64, writeString func(string)) {
	str := fmt.Sprintf("\t\t\t<g transform=\"translate(100 0)\">\n"+
		"\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"%s\" x2=\"%s\" y1=\"%s\" y2=\"%s\"/>\n",
		float.RemoveTrailingZeros(s.Axes.Y.X1),
		float.RemoveTrailingZeros(s.Axes.Y.X2),
		float.RemoveTrailingZeros(s.Axes.Y.Y1),
		float.RemoveTrailingZeros(s.Axes.Y.Y2),
	)
	writeString(str)

	tickLabelPosition := s.Axes.Y.X1 - 13
	tickEnd := s.Axes.Y.X1 - 10
	for i, tick := range s.Ticks.Y {
		if s.Ticks.YLabel[i] != "0" {
			yPosition := axisLength - tick
			str = fmt.Sprintf("\t\t\t\t<g>\n"+
				"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"%[1]s\" x2=\"%[2]s\" y1=\"%[3]s\" y2=\"%[3]s\" />\n"+
				"\t\t\t\t\t<text dy=\"4\" font-size=\"12\" text-anchor=\"end\" x=\"%[4]s\" y=\"%[3]s\">%[5]s</text>\n"+
				"\t\t\t\t</g>\n",
				float.RemoveTrailingZeros(s.Axes.Y.X1),
				float.RemoveTrailingZeros(tickEnd),
				float.RemoveTrailingZeros(math.Round(yPosition, 0.01)),
				float.RemoveTrailingZeros(tickLabelPosition),
				s.Ticks.YLabel[i],
			)
			writeString(str)
		}
	}

	str = fmt.Sprintf("\t\t\t</g>\n"+
		"\t\t\t<text dy=\"15\" text-anchor=\"middle\" transform=\"rotate(-90, 0, %[1]s)\" x=\"0\" y=\"%[1]s\">%[2]s</text>\n",
		float.RemoveTrailingZeros(math.Round(axisLength/2, 0.01)),
		s.YLabel,
	)
	writeString(str)
}
