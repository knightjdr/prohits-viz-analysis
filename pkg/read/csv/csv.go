// Package csv reads csv files.
package csv

import (
	gocsv "encoding/csv"
	"io"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/spf13/afero"
)

// Read a csv file to a slice with header as map keys.
func Read(filename string, sep rune) []map[string]string {
	file := openFile(filename)
	reader := createReader(file, sep)

	header := readHeader(reader)
	return readLines(reader, header)
}

func openFile(filename string) afero.File {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)

	return file
}

func createReader(file afero.File, sep rune) *gocsv.Reader {
	reader := gocsv.NewReader(file)
	reader.Comma = sep
	reader.LazyQuotes = true

	return reader
}

func readHeader(reader *gocsv.Reader) []string {
	header, err := reader.Read()
	log.CheckError(err, true)

	return header
}

func readLines(reader *gocsv.Reader, header []string) []map[string]string {
	data := make([]map[string]string, 0)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		log.CheckError(err, true)

		parsedLine := make(map[string]string, len(header))
		for i, field := range header {
			parsedLine[field] = line[i]
		}

		data = append(data, parsedLine)
	}

	return data
}
