// Package errorcheck ensures that the input file (passed as slice map) has no errors.
//
// Common errors: 1) no entries pass the required filters (this is already checked
// during the file reading step but is repeated here for completeness); 2) not
// enough baits for the type of anlaysis; 3) missing prey names; 4) score column
// must be numeric (also checked during file reading step); 5) prey length column
// (if specified) must be numeric; 6) control column (if specified) must be a
// pipe-separated list of numeric values.
package errorcheck

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/types"
)

// Required is the entry point for error checking the input data.
func Required(
	dataset types.Dataset,
) (err error) {
	errs := 0 // Track error number.
	// Check there is data and panic if not.
	err = MinData(dataset.Data)
	logmessage.CheckError(err, true)

	// Minimum baits met (don't panic so that all errors have a change to log).
	err = MinBait(dataset.Data, dataset.Params.AnalysisType)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// All preys have names.
	err = PreyName(dataset.Data)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// Abundance column is a pipe-separated list.
	err = AbundanceColumn(dataset.Data)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// Score column is a float.
	err = ScoreFloat(dataset.Data)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// Prey length column can be parsed as an integer.
	err = PreyLengthInt(dataset.Data, dataset.Params.PreyLength)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// Control column is a pipe-separated list.
	err = ControlColumn(dataset.Data, dataset.Params.Control)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}
	var formatErr error
	if errs > 0 {
		formatErr = errors.New("Format error")
	}
	return formatErr
}
