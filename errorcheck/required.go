// Package errorcheck ensures that the input file (passed as slice map) has no errors.
//
// Common errors: 1) no entries pass the required filters (this is already checked
// during the file reading step but is repeated here for completeness); 2) not
// enough conditions for the type of anlaysis; 3) missing readout names; 4) score column
// must be numeric (also checked during file reading step); 5) readout length column
// (if specified) must be numeric; 6) control column (if specified) must be a
// pipe-separated list of numeric values.
package errorcheck

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Required is the entry point for error checking the input data.
func Required(
	dataset typedef.Dataset,
) (err error) {
	errs := 0 // Track error number.
	// Check if there is data and panic if not.
	err = MinData(dataset.FileData)
	logmessage.CheckError(err, true)

	// Minimum conditions met (don't panic so that all errors have a change to log).
	err = MinCondition(dataset.FileData, dataset.Parameters.AnalysisType)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// All readouts have names.
	err = ReadoutName(dataset.FileData)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// Abundance column is a pipe-separated list.
	err = AbundanceColumn(dataset.FileData)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// Score column is a float.
	err = ScoreFloat(dataset.FileData)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// Readout length column can be parsed as an integer.
	err = ReadoutLengthInt(dataset.FileData, dataset.Parameters.ReadoutLength)
	logmessage.CheckError(err, false)
	if err != nil {
		errs++
	}

	// Control column is a pipe-separated list.
	err = ControlColumn(dataset.FileData, dataset.Parameters.Control)
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
