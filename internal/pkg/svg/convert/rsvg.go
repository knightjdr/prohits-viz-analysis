// Package convert turns an SVG into a PNG.
package convert

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
)

// RSVG converts an svg to PNG using rsvg-convert.
func RSVG(svg, targetFileName, bgColor string) {
	workingDir, err := os.Getwd()
	log.CheckError(err, true)

	fileArg := fmt.Sprintf("%s/%s", workingDir, svg)
	exportArg := fmt.Sprintf("--output=%s/%s", workingDir, targetFileName)

	cmd := exec.Command(
		"rsvg-convert",
		"--format=png",
		exportArg,
		fmt.Sprintf("--background-color=%s", bgColor),
		"--unlimited",
		fileArg,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()

	log.CheckError(err, true)
}
