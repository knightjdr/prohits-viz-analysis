package circheatmap

func writePlot(c *CircHeatmapSVG, writeString func(string)) {
	writeString("\t<g transform=\"rotate(-90)\">\n")
	writeKnown(c, writeString)
	writeCircles(c, writeString)
	writeString("\t</g>\n")
}
