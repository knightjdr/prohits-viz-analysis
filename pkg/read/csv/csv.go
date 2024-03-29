// Package csv reads csv files.
package csv

import (
	gocsv "encoding/csv"
	"io"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/spf13/afero"
)

// ReadToMap reads a two column csv file to a map.
func ReadToMap(filename string, sep rune) map[string]string {
	file := openFile(filename)
	reader := createReader(file, sep)

	data := make(map[string]string, 0)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		log.CheckError(err, true)

		data[line[0]] = line[1]
	}

	return data
}

// ReadToSliceMap reads a two column csv file to a map of slices.
func ReadToSliceMap(filename string, sep rune) map[string][]string {
	file := openFile(filename)
	reader := createReader(file, sep)

	data := make(map[string][]string, 0)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		log.CheckError(err, true)

		if _, ok := data[line[0]]; !ok {
			data[line[0]] = make([]string, 0)
		}
		data[line[0]] = append(data[line[0]], line[1])
	}

	return data
}

// ReadToSliceViaHeader reads a csv file to a slice. Using the columnMap to map from column name to
// row field.
func ReadToSliceViaHeader(filename string, sep rune, columnMap map[string]string) []map[string]string {
	file := openFile(filename)
	reader := createReader(file, sep)

	header := readHeader(reader)
	headerMap := mapHeaders(header, columnMap)
	return readLines(reader, header, headerMap)
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

func mapHeaders(header []string, columnMap map[string]string) map[string]string {
	headerMap := make(map[string]string, len(header))

	for _, column := range header {
		if _, ok := columnMap[column]; ok {
			headerMap[column] = columnMap[column]
		} else {
			headerMap[column] = column
		}
	}

	return headerMap
}

func readLines(reader *gocsv.Reader, header []string, headerMap map[string]string) []map[string]string {
	data := make([]map[string]string, 0)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		log.CheckError(err, true)

		parsedLine := make(map[string]string, len(header))
		for i, column := range header {
			field := headerMap[column]
			parsedLine[field] = line[i]
		}

		data = append(data, parsedLine)
	}

	return data
}
