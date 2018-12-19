package biclustering

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Parameters creates a file with the biclustering parameters to use
func parameters(fileData []map[string]string, params typedef.Parameters) {
	var biclustParameters string
	if params.BiclusteringApprox {
		// Get number of conditions.
		uniqueConditions := make(map[string]bool)
		for _, row := range fileData {
			condition := row["condition"]
			if _, ok := uniqueConditions[condition]; !ok {
				uniqueConditions[condition] = true
			}
		}
		nb := len(uniqueConditions)

		// Create optimized param file content.
		biclustParameters = fmt.Sprintf("np 10\n"+
			"nb %d\n"+
			"a 1.0\n"+
			"b 1.0\n"+
			"lambda 0.0\n"+
			"nu 25.0\n"+
			"alpha 1.0\n"+
			"rho 1.0\n"+
			"gamma 1.0\n"+
			"nburn 50\n"+
			"niter 500\n", nb,
		)
	} else {
		// Create default param file content.
		biclustParameters = fmt.Sprintln("np 10\n" +
			"nb 100\n" +
			"a 1.0\n" +
			"b 1.0\n" +
			"lambda 0.0\n" +
			"nu 25.0\n" +
			"alpha 1.0\n" +
			"rho 1.0\n" +
			"gamma 1.0\n" +
			"nburn 5000\n" +
			"niter 10000\n",
		)
	}
	afero.WriteFile(fs.Instance, "biclustering/parameters.txt", []byte(biclustParameters), 0644)
}
