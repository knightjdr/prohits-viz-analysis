package minimap

// Matrices with a dimension greater than 500 will always generate a heatmap,
// otherwise it will output whatever is requested.
func defineImageType(data *Data) string {
	if data.ImageType == "dotplot" &&
		len(data.Matrices.Abundance) <= 500 &&
		len(data.Matrices.Abundance[0]) <= 500 {
		return "dotplot"
	}

	return "heatmap"
}
