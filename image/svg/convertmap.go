package svg

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// ConvertMap converts an svg to a minimap. A minimap is a png with height 1000px
// that is used by the interactive viewer at prohis-viz.
func ConvertMap(svg string) (err error) {
	extension := filepath.Ext(svg)
	filename := svg[0 : len(svg)-len(extension)]
	exportArg := fmt.Sprintf("--output=%s.png", filename)
	cmd := exec.Command("rsvg-convert", "--format=png", exportArg, "--height=1000", "--unlimited", svg)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()

	logmessage.CheckError(err, false)
	if err != nil {
		logmessage.Write(stderr.String())
	}
	return
}
