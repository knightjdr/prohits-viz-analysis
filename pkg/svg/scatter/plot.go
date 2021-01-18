package scatter

func writePlot(s *Scatter, axisLength float64, writeString func(string)) {
	writeString("\t<g transform=\"translate(0 50)\">\n")
	writePoints(s, axisLength, writeString)
	writeAxes(s, axisLength, writeString)
	writeString("\t</g>\n")
}
