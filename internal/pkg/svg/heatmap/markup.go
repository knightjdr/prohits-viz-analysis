package heatmap

// WriteMarkup writes annotations, markers and column/row headings.
func WriteMarkup(h *Heatmap, writeString func(string)) {
	writeMarkers(h, writeString)
	writeAnnotations(h, writeString)
	writeHeadings(h, writeString)
}
