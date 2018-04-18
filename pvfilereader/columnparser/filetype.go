package columnparser

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/logmessage"
)

// returns the mime type of a file
func Filetype(filename string, logFile string) string {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		// return unknown if file cannot be opened
		logmessage.Write(logFile, fmt.Sprintf("%s: could not be opened", filename))
		return "unknown"
	}
	defer file.Close()

	// read file and determine type
	buffer := make([]byte, 512)
	bufferLength, _ := file.Read(buffer)

	// split mimetype at first semi colon
	mimetype := strings.Split(http.DetectContentType(buffer[:bufferLength]), ";")[0]
	return mimetype
}
