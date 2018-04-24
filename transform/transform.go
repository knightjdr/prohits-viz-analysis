// Package transform will adjust prey values to the user's requirements.
//
// Preys can be adjusted by (in this order): 1) control values (must be a
// pipe-separated list); 2) prey length; 3) normalized across baits; and 4)
// log transformed.
package transform

// Preys is the entry point for prey transformations.
func Preys(
	data []map[string]interface{},
	control, preyLength, normalization, normalizationPrey, logBase string,
) (transformed []map[string]interface{}, err error) {
	// Control subtraction.
	transformed = ControlSubtraction(data, control)

	// Prey length normalization.
	transformed = PreyLength(transformed, preyLength)

	// Bait normalization.
	transformed = Normalization(transformed, normalization, normalizationPrey)

	// Log transformation
	transformed = LogTransform(transformed, logBase)

	return transformed, err
}
