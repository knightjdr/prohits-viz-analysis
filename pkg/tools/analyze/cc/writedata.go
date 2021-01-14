package cc

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeData(data []types.ScatterPoint, settings types.Settings) {
	file, err := fs.Instance.Create("other/x-y-data.txt")
	log.CheckError(err, false)
	if err != nil {
		return
	}
	defer file.Close()

	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf("%s\tx\ty\n", settings.Readout))

	for _, row := range data {
		buffer.WriteString(fmt.Sprintf("%s\t%0.2f\t%0.2f\n", row.Label, row.X, row.Y))
	}
	file.WriteString(buffer.String())
}
