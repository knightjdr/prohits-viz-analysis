package transform

// Normalization normalizes prey abundances between baits. There are two variants:
// 1) sum the prey abundances for each bait, take the median of that and use the
// median / total prey abundance for each bait as the prey multplier; 2) normalize
// against a specific prey. In the second variant the requested prey's abundance
// is used in place of the total prey abundance. If a bait is missing the requested
// prey, it will use the median as it's default.
func Normalization(
	data []map[string]interface{},
	normalization, normalizationPrey string,
) (transformed []map[string]interface{}) {
	transformed = data
	if normalization == "prey" {
		// Normalize by prey.
		return PreyNormalization(transformed, normalizationPrey)
	} else if normalization == "total" {
		// Normalize by summed abundance.
		return SummedNormalization(transformed)
	}
	// Skip if no normalization is required.
	return
}
