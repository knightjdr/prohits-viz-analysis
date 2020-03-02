// Package data ensures that the input file (passed as slice map) has no errors in formatting.
//
// Common errors:
// 1) no entries pass the required filters
// 2) not enough conditions for the type of analysis
// 3) missing readout names
// 4) score column is numeric
// 5) readout length column (if specified) must be numeric
package data

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Validate is the entry point for error checking the input data.
func Validate(analysis *types.Analysis) {
	errs := make([]error, 0)

	addError(&errs, confirmParsedData(analysis.Data))
	addError(&errs, confirmMinimumConditions(analysis.Data, analysis.Settings.Type))
	addError(&errs, confirmReadoutsHaveNames(analysis.Data))
	addError(&errs, confirmScoreIsFloat(analysis.Data))
	addError(&errs, confirmReadLengthIsInt(analysis.Data, analysis.Settings.ReadoutLength))

	if len(errs) > 0 {
		log.WriteAndExit(joinErrors(errs))
	}
}

func addError(errs *[]error, err error) {
	if err != nil {
		*errs = append(*errs, err)
	}
}

func joinErrors(errs []error) string {
	var message strings.Builder

	for _, err := range errs {
		message.WriteString(fmt.Sprintf("%s\n", err.Error()))
	}

	return message.String()
}
