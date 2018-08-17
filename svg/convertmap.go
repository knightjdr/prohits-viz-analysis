package svg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// ConvertMap converts svgs to a minimaps. A minimap is a png with height 600px
// that is used by the interactive viewer at prohis-viz. The svg must have a g
// element with id="minimap" to use for exporting.
func ConvertMap(list []string) {
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
		exportArg := fmt.Sprintf("--export-png=%s/minimap/%s.png", workingDir, filename)
		cmd := exec.Command("inkscape", fileArg, exportArg, "--export-height=1000", "--export-id=minimap", "--without-gui")
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