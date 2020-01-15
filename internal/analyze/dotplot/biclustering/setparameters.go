package biclustering

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/spf13/afero"
)

func setParameters(noConditions int, settings types.Settings) {
	var biclustParameters string

	if settings.BiclusteringApprox {
		biclustParameters = fmt.Sprintf(
			"np 10\n"+
				"nb %d\n"+
				"a 1.0\n"+
				"b 1.0\n"+
				"lambda 0.0\n"+
				"nu 25.0\n"+
				"alpha 1.0\n"+
				"rho 1.0\n"+
				"gamma 1.0\n"+
				"nburn 50\n"+
				"niter 500\n", noConditions,
		)
	} else {
		biclustParameters = fmt.Sprintln(
			"np 10\n" +
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
