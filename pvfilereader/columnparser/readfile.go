// Package columnparser reads csv formatted files and returns specified columns
package columnparser

func Readfile(files []string, columnMap map[string]string, logFile string) {
	// get mime type for each file
	fileno := len(files)
	filetype := make([]string, fileno)
	for i := 0; i < fileno; i++ {
		filetype[i], _ = Filetype(files[i], logFile)
	}

	// read files to structs
	// parsedfiles := Parsecsv(files, filetype, columnMap, logFile)
}
