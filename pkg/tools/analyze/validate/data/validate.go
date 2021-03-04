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
	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Validate is the entry point for error checking the input data.
func Validate(analysis *types.Analysis, toValidate []string) {
	errs := make([]error, 0)

	if slice.ContainsString("data", toValidate) {
		addError(&errs, confirmParsedData(analysis.Data))
	}
	if slice.ContainsString("minConditions", toValidate) {
		addError(&errs, confirmMinimumConditions(analysis.Data, analysis.Settings.Type))
	}
	if slice.ContainsString("readout", toValidate) {
		addError(&errs, confirmReadoutsHaveNames(analysis.Data))
	}
	if slice.ContainsString("readoutLength", toValidate) {
		addError(&errs, confirmReadLengthIsInt(analysis.Data, analysis.Settings.ReadoutLength))
	}
	if slice.ContainsString("score", toValidate) {
		addError(&errs, confirmScoreIsFloat(analysis.Data))
	}

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
