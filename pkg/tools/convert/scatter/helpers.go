package scatter

import (
	"encoding/csv"

	"github.com/spf13/afero"
)

func createReader(file afero.File) *csv.Reader {
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	return reader
}
