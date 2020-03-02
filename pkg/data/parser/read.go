// Package parser reads csv formatted files and returns specified columns.
package parser

import (
	"errors"

	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Read will read a csv file(s) into a slice. The columnMap specifies
// how to map the columns to interface keys. "ignoreMissing" declares if missing
// interface keys can be safely ignored.
func Read(analysis *types.Analysis, ignoreMissing bool) {
	columnMap := analysis.Columns
	files := analysis.Settings.Files

	mimeTypes := getMimeTypes(files)

	analysis.Data = parseCSV(files, mimeTypes, columnMap, ignoreMissing)
	checkParsedData(analysis.Data)
}

func checkParsedData(parsed []map[string]string) {
	if len(parsed) == 0 {
		log.CheckError(errors.New("no parsed results"), true)
	}
}
