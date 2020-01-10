package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
)

func createFolders(createPNG bool) {
	folders := make([]string, 0)
	folders = append(folders, []string{"cytoscape", "interactive", "minimap", "other", "svg", "treeview"}...)
	if createPNG {
		folders = append(folders, "png")
	}

	files.CreateFolders(folders)
}
