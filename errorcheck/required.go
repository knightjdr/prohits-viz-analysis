// Package errorcheck ensures that the input file (passed as slice map) has no errors
//
/* Common errors:
** 1) no entries pass the required filters (this is already checked during the file
** reading step but is repeated here for completeness)
** 2) not enough baits for the type of anlaysis
** 3) missing prey names
** 4) score column must be numeric (also checked during file reading step)
** 5) prey length column (if specified) must be numeric
** 6) control column (if specified) must be a pipe-separated list of numeric values
 */
package errorcheck

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// Required is the entry point for error checking the input data
func Required(
	data []map[string]interface{},
	analysisType string,
	control string,
	preyLength string,
	logFile string,
) error {
	errs := 0 // track error number
	// there is data
	err := MinData(data)
	if err != nil {
		logmessage.Write(logFile, err.Error())
		errs++
		return err // skip nest tests since there is no data
	}

	// minimum baits met
	err = MinBait(data, analysisType)
	if err != nil {
		logmessage.Write(logFile, err.Error())
		errs++
	}

	// all preys have names
	err = PreyName(data)
	if err != nil {
		logmessage.Write(logFile, err.Error())
		errs++
	}

	// abundance column is a pipe-separated list
	err = AbundanceColumn(data)
	if err != nil {
		logmessage.Write(logFile, err.Error())
		errs++
	}

	// score column is a float
	err = ScoreFloat(data)
	if err != nil {
		logmessage.Write(logFile, err.Error())
		errs++
	}

	// prey length column can be parsed as an integer
	err = PreyLengthInt(data, preyLength)
	if err != nil {
		logmessage.Write(logFile, err.Error())
		errs++
	}

	// control column is a pipe-separated list
	err = ControlColumn(data, control)
	if err != nil {
		logmessage.Write(logFile, err.Error())
		errs++
	}
	var formatErr error
	if errs > 0 {
		formatErr = errors.New("Format error")
	}
	return formatErr
}
