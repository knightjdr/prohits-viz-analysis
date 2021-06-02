package circheatmap

func writePlot(c *CircHeatmap, writeString func(string)) {
	writeString("\t<g transform=\"rotate(-90)\">\n")
	writeKnown(c, writeString)
	writeString("\t</g>\n")
}
