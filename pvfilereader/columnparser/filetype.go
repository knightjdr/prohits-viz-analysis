package columnparser

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/logmessage"
	"github.com/spf13/afero"
)

var appFs = afero.NewOsFs()

// returns the mime type of a file
func Filetype(filename string, logFile string) (string, error) {
	// open file
	file, err := appFs.Open(filename)
	if err != nil {
		// return unknown if file cannot be opened
		logmessage.Write(logFile, fmt.Sprintf("%s: could not be opened", filename))
		return "unknown", err
	}
	defer file.Close()

	// read file and determine type
	buffer := make([]byte, 512)
	bufferLength, err := file.Read(buffer)
	if err != nil {
		// return unknown if file cannot be read
		logmessage.Write(logFile, fmt.Sprintf("%s: could not read", filename))
		return "unknown", err
	}

	// split mimetype at first semi colon
	mimetype := strings.Split(http.DetectContentType(buffer[:bufferLength]), ";")[0]
	return mimetype, err
}
