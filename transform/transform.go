// Package transform will adjust readout values to the user's requirements.
//
// Readouts can be adjusted by (in this order): 1) control values (must be a
// pipe-separated list); 2) readout length; 3) normalized across conditions; and 4)
// log transformed.
package transform

import "github.com/knightjdr/prohits-viz-analysis/typedef"

// Readouts is the entry point for readout transformations.
func Readouts(
	dataset typedef.Dataset,
) (transformed []map[string]string) {
	// Control subtraction.
	transformed = ControlSubtraction(dataset.FileData, dataset.Parameters.Control)

	// Readout length normalization.
	transformed = ReadoutLength(transformed, dataset.Parameters.ReadoutLength)

	// Condition normalization.
	transformed = Normalization(transformed, dataset.Parameters.Normalization, dataset.Parameters.NormalizationReadout)

	// Log transformation
	transformed = LogTransform(transformed, dataset.Parameters.LogBase)

	return transformed
}
