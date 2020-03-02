package heatmap

// WriteLabels writes column and row labels.
func WriteLabels(h *Heatmap, writeString func(string)) {
	if len(h.Columns) > 0 {
		writeColumnNames(h, writeString)
	}
	if len(h.Rows) > 0 {
		writeRowNames(h, writeString)
	}
}
