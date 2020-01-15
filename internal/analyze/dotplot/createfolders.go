package dotplot

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func createFolders(settings types.Settings) {
	folders := make([]string, 0)
	folders = append(folders, []string{"cytoscape", "interactive", "minimap", "other", "svg", "treeview"}...)
	if settings.Clustering == "biclustering" {
		folders = append(folders, "biclustering")
	}
	if settings.Png {
		folders = append(folders, "png")
	}

	files.CreateFolders(folders)
}
