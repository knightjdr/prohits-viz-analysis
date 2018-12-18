package file

import (
	"github.com/spf13/afero"
)

// Write a string to a file or append it to a pointer
func Write(outstring *[]string, writestring string, file afero.File) {
	if file != nil {
		file.WriteString(writestring)
	} else {
		*outstring = append(*outstring, writestring)
	}
}
