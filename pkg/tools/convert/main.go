// Package convert takes a file from ProHits-viz V1 and converts it to V2 JSON.
package convert

import "github.com/knightjdr/prohits-viz-analysis/pkg/tools/convert/heatmap"

// File converts a file from v1 format to v2.
func File() {
	conversionSettings := parseArguments()

	if conversionSettings.imageType == "dotplot" || conversionSettings.imageType == "heatmap" {
		heatmap.Convert(conversionSettings.file)
	}
}
