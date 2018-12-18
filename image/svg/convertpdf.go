package svg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// ConvertPdf converts svg files to PDFs in the current working directory
// in a subfolder called "pdf". The conversion is done using inkscape
// which requires the full path to svg files and output files. Because
// inkscape needs full paths, this can't be unit tested by mocking
// with afero.
func ConvertPdf(list []string) {
	// Get current working directory.
	workingDir, err := os.Getwd()
	// Return if err but not fatal. We just won't convert anything.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Convert to PDF using rsvg-convert.
	for _, svg := range list {
		extension := filepath.Ext(svg)
		filename := svg[0 : len(svg)-len(extension)]
		fileArg := fmt.Sprintf("%s/svg/%s", workingDir, svg)
		exportArg := fmt.Sprintf("--output=%s/pdf/%s.pdf", workingDir, filename)
		cmd := exec.Command("rsvg-convert", "--format=pdf", exportArg, "--unlimited", fileArg)
		// Using inkscape
		// exportArg := fmt.Sprintf("--export-pdf=%s/pdf/%s.pdf", workingDir, filename)
		// cmd := exec.Command("inkscape", fileArg, exportArg, "--without-gui")
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()

		// Continue to next file if err.
		logmessage.CheckError(err, false)
		if err != nil {
			logmessage.Write(stderr.String())
			continue
		}
	}
	return
}
