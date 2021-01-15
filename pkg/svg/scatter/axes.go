package scatter

import "fmt"

func writeAxes(s *Scatter, axisLength float64, writeString func(string)) {
	writeString("\t\t<g>\n")
	writeXAxis(s, axisLength, writeString)
	writeYAxis(s, axisLength, writeString)
	writeString("\t\t</g>\n")
}

func writeXAxis(s *Scatter, axisLength float64, writeString func(string)) {
	str := fmt.Sprintf("\t\t\t<g transform=\"translate(0 %0.2[1]f)\">\n"+
		"\t\t\t\t<rect width=\"%0.2[2]f\" height=\"100\" fill=\"white\" />\n"+
		"\t\t\t\t<g transform=\"translate(100 0)\">\n"+
		"\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"0\" x2=\"%0.2[1]f\" y1=\"0\" y2=\"0\"/>\n"+
		"\t\t\t\t\t<g>\n",
		axisLength,
		s.PlotSize,
	)
	writeString(str)

	for i, tick := range s.Ticks.X {
		str = fmt.Sprintf(
			"\t\t\t\t\t\t<g>\n"+
				"\t\t\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"%0.2[1]f\" x2=\"%0.2[1]f\" y1=\"0\" y2=\"10\" />\n"+
				"\t\t\t\t\t\t\t<text font-size=\"12\" text-anchor=\"middle\" x=\"%0.2[1]f\" y=\"28\">%[2]s</text>\n"+
				"\t\t\t\t\t\t</g>\n",
			tick,
			s.Ticks.XLabel[i],
		)
		writeString(str)
	}

	str = fmt.Sprintf(
		"\t\t\t\t\t</g>\n"+
			"\t\t\t\t\t<text text-anchor=\"middle\" x=\"%0.2[1]f\" y=\"70\">%[2]s</text>\n"+
			"\t\t\t\t</g>\n"+
			"\t\t\t</g>\n",
		axisLength/2,
		s.XLabel,
	)
	writeString(str)
}

func writeYAxis(s *Scatter, axisLength float64, writeString func(string)) {
	str := fmt.Sprintf("\t\t\t<g>\n"+
		"\t\t\t\t<rect width=\"100\" height=\"%0.2[1]f\" fill=\"white\" />\n"+
		"\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"100\" x2=\"100\" y1=\"0\" y2=\"%0.2[2]f\"/>\n"+
		"\t\t\t\t<g>\n",
		s.PlotSize-150,
		axisLength,
	)
	writeString(str)

	for i, tick := range s.Ticks.Y {
		yPosition := axisLength - tick
		str = fmt.Sprintf(
			"\t\t\t\t\t<g>\n"+
				"\t\t\t\t\t\t<line stroke=\"#333333\" stroke-width=\"2\" x1=\"100\" x2=\"90\" y1=\"%0.2[1]f\" y2=\"%0.2[1]f\" />\n"+
				"\t\t\t\t\t\t<text dy=\"4\" font-size=\"12\" text-anchor=\"end\" x=\"87\" y=\"%0.2[1]f\">%[2]s</text>\n"+
				"\t\t\t\t\t</g>\n",
			yPosition,
			s.Ticks.YLabel[i],
		)
		writeString(str)
	}

	str = fmt.Sprintf(
		"\t\t\t\t</g>\n"+
			"\t\t\t\t<text dy=\"15\" text-anchor=\"middle\" transform=\"rotate(-90, 0, %0.2[1]f)\" x=\"0\" y=\"%0.2[1]f\">%[2]s</text>\n"+
			"\t\t\t</g>\n",
		axisLength/2,
		s.YLabel,
	)
	writeString(str)
}
