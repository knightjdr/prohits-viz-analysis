package hierarchical

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/spf13/afero"
)

func writeTrees(clusteredData hclustData, settings types.Settings) {
	writeTree(clusteredData.tree["condition"].Newick, settings.Condition)
	writeTree(clusteredData.tree["readout"].Newick, settings.Readout)
}

func writeTree(tree string, filehandle string) {
	afero.WriteFile(fs.Instance, fmt.Sprintf("other/%s-dendrogram.txt", filehandle), []byte(tree), 0644)
}
