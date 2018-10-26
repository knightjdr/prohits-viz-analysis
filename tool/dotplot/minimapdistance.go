package dotplot

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// MinimapDistance draws a distance heatmap.
func MinimapDistance(
	abundance [][]float64,
	sorted []string,
	fillColor string,
	imageName string,
) {
	// Define dotplot parameters.
	parameters := map[string]interface{}{
		"abundanceCap": float64(1),
		"colLabel":     "Conditions",
		"fillColor":    fillColor,
		"invertColor":  true,
		"minimap":      true,
		"rowLabel":     "Conditions",
	}
	minimap := svg.Heatmap(abundance, typedef.Annotations{}, typedef.Markers{}, sorted, sorted, true, parameters)
	fileName := fmt.Sprintf("minimap/%s.svg", imageName)
	afero.WriteFile(fs.Instance, fileName, []byte(minimap), 0644)
	return
}
