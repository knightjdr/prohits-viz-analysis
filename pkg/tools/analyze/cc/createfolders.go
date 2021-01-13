package cc

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func createFolders(settings types.Settings) {
	folders := make([]string, 0)
	folders = append(folders, []string{"interactive", "other", "svg"}...)
	if settings.Png {
		folders = append(folders, "png")
	}

	files.CreateFolders(folders)
}
