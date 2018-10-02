package transform

// Normalization normalizes readout abundances between conditions. There are two variants:
// 1) sum the readout abundances for each condition, take the median of that and use the
// median / total readout abundance for each condition as the readout multplier; 2) normalize
// against a specific readout. In the second variant the requested readout's abundance
// is used in place of the total readout abundance. If a condition is missing the requested
// readout, it will use the median as it's default.
func Normalization(
	data []map[string]interface{},
	normalization, normalizationReadout string,
) (transformed []map[string]interface{}) {
	transformed = data
	if normalization == "readout" {
		// Normalize by readout.
		return ReadoutNormalization(transformed, normalizationReadout)
	} else if normalization == "total" {
		// Normalize by summed abundance.
		return SummedNormalization(transformed)
	}
	// Skip if no normalization is required.
	return
}
