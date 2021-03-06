package heatmap

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/png"
)

func parseMinimap(file string) string {
	url := png.ConvertToURI(file)

	return fmt.Sprintf("\"minimap\": {\"main\":{\"image\":\"%s\"}}", url)
}
