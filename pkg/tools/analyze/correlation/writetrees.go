package correlation

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

func writeTrees(conditionData, readoutData *correlationData, settings types.Settings) {
	writeTree(conditionData.tree.Newick, settings.Condition)
	writeTree(readoutData.tree.Newick, settings.Readout)
}

func writeTree(tree string, filehandle string) {
	if len(tree) > 0 {
		afero.WriteFile(fs.Instance, fmt.Sprintf("other/%s-dendrogram.txt", filehandle), []byte(tree), 0644)
	}
}
