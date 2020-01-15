package hierarchical

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/spf13/afero"
)

// WriteTrees for newick format.
func WriteTrees(clusteredData HclustData, settings types.Settings) {
	writeTree(clusteredData.Tree["condition"].Newick, settings.Condition)
	writeTree(clusteredData.Tree["readout"].Newick, settings.Readout)
}

func writeTree(tree string, filehandle string) {
	if len(tree) > 0 {
		afero.WriteFile(fs.Instance, fmt.Sprintf("other/%s-dendrogram.txt", filehandle), []byte(tree), 0644)
	}
}
