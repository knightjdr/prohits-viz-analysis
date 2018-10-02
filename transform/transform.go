// Package transform will adjust prey values to the user's requirements.
//
// Preys can be adjusted by (in this order): 1) control values (must be a
// pipe-separated list); 2) prey length; 3) normalized across baits; and 4)
// log transformed.
package transform

import "github.com/knightjdr/prohits-viz-analysis/typedef"

// Preys is the entry point for prey transformations.
func Preys(
	dataset typedef.Dataset,
) (transformed []map[string]interface{}) {
	// Control subtraction.
	transformed = ControlSubtraction(dataset.Data, dataset.Parameters.Control)

	// Prey length normalization.
	transformed = PreyLength(transformed, dataset.Parameters.PreyLength)

	// Bait normalization.
	transformed = Normalization(transformed, dataset.Parameters.Normalization, dataset.Parameters.NormalizationPrey)

	// Log transformation
	transformed = LogTransform(transformed, dataset.Parameters.LogBase)

	return transformed
}
