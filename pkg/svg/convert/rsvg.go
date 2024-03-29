// Package convert turns an SVG into a PNG.
package convert

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

// RSVG converts an svg to PNG using rsvg-convert.
func RSVG(svg, targetFileName, bgColor string) {
	cmdStr := fmt.Sprintf(
		"docker run --rm -v $(pwd):/files/ --user $(id -u):$(id -g) rsvg --format=png --output=%s --background-color=%s --unlimited %s",
		targetFileName,
		bgColor,
		svg,
	)

	cmd := exec.Command(
		"/bin/sh",
		"-c",
		cmdStr,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	log.CheckError(err, true)
}
