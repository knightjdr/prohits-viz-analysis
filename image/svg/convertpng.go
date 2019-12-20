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
// in a subfolder called "png". The conversion is done using rsvg-convert.
func ConvertPng(list []string) {
	// Get current working directory.
	workingDir, err := os.Getwd()
	// Return if err but not fatal. We just won't convert anything.
	logmessage.CheckError(err, false)
	if err != nil {
		return
	}

	// Convert to PNG using rsvg-convert.
	for _, svg := range list {
		extension := filepath.Ext(svg)
		filename := svg[0 : len(svg)-len(extension)]
		fileArg := fmt.Sprintf("%s/svg/%s", workingDir, svg)
		exportArg := fmt.Sprintf("--output=%s/png/%s.png", workingDir, filename)
		cmd := exec.Command(
			"rsvg-convert",
			"--format=png",
			exportArg,
			"--background-color=white",
			"--unlimited",
			fileArg,
		)
		// Using inkscape
		// exportArg := fmt.Sprintf("--export-png=%s/png/%s.png", workingDir, filename)
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