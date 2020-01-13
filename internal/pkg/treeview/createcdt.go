package treeview

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/spf13/afero"
)

func createCDT(data Data) {
	file, err := fs.Instance.Create(fmt.Sprintf("%s.cdt", data.Filename))
	log.CheckError(err, true)
	defer file.Close()

	writeCDTHeader(file, data.Names.Columns)
	writeArrayIDs(file, len(data.Names.Columns))
	writeCDTEweight(file, len(data.Names.Columns))
	writeCDTMatrix(file, data)
}

func writeCDTHeader(file afero.File, columns []string) {
	var buffer strings.Builder
	buffer.WriteString("GID\tUNIQID\tNAME\tGWEIGHT")

	for _, column := range columns {
		buffer.WriteString(fmt.Sprintf("\t%s", column))
	}
	buffer.WriteString("\n")

	file.WriteString(buffer.String())
}

func writeArrayIDs(file afero.File, noColumns int) {
	var buffer strings.Builder
	buffer.WriteString("AID\t\t\t")

	for i := 0; i < noColumns; i++ {
		buffer.WriteString(fmt.Sprintf("\tARRY%dX", i))
	}
	buffer.WriteString("\n")

	file.WriteString(buffer.String())
}

func writeCDTEweight(file afero.File, noColumns int) {
	var buffer strings.Builder
	buffer.WriteString("EWEIGHT\t\t\t")

	for i := 0; i < noColumns; i++ {
		buffer.WriteString("\t1")
	}
	buffer.WriteString("\n")

	file.WriteString(buffer.String())
}

func writeCDTMatrix(file afero.File, data Data) {
	var buffer strings.Builder

	for rowIndex, rowName := range data.Names.Rows {
		buffer.WriteString(fmt.Sprintf("GENE%dX\t%s\t%s\t1", rowIndex, rowName, rowName))
		for _, value := range data.Matrix[rowIndex] {
			buffer.WriteString(fmt.Sprintf("\t%0.5f", value))
		}
		buffer.WriteString("\n")
	}

	file.WriteString(buffer.String())
}
