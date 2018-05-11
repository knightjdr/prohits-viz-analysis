package svg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// ConvertPng converts svg files to PNGs in the current working directory
// in a subfolder called "png". The conversion is done using inkscape
// which requires the full path to svg files and output files. Because
// inkscape needs full paths, this can't be unit tested by mocking
// with afero.
func ConvertPng(list []string) {
	// Get current working directory.
	workingDir, err := os.Getwd()
	// Return if err but not fatal. We just won't convert anything.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Convert to PNG using inkscape.
	for _, svg := range list {
		extension := filepath.Ext(svg)
		filename := svg[0 : len(svg)-len(extension)]
		fileArg := fmt.Sprintf("%s/svg/%s", workingDir, svg)
		exportArg := fmt.Sprintf("--export-png=%s/png/%s.png", workingDir, filename)
		cmd := exec.Command("inkscape", fileArg, exportArg, "--export-dpi=96", "--without-gui")
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
