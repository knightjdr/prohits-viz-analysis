package specificity

import (
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeData(data map[string]map[string]map[string]float64, settings types.Settings) {
	file, err := fs.Instance.Create("other/specificity-data.txt")
	log.CheckError(err, false)
	if err != nil {
		return
	}
	defer file.Close()

	var buffer strings.Builder
	buffer.WriteString(
		fmt.Sprintf(
			"%s\t%s\t%s\tSpecificity\t%s\n",
			settings.Condition,
			settings.Readout,
			settings.Abundance,
			settings.Score,
		),
	)

	conditions := make([]string, len(data))
	i := 0
	for condition := range data {
		conditions[i] = condition
		i++
	}
	sort.Strings(conditions)

	for _, condition := range conditions {
		readouts := make([]string, len(data[condition]))
		i := 0
		for readout := range data[condition] {
			readouts[i] = readout
			i++
		}
		sort.Strings(readouts)

		for _, readout := range readouts {
			buffer.WriteString(
				fmt.Sprintf(
					"%s\t%s\t%.2f\t%.2f\t%.2f\n",
					condition,
					readout,
					data[condition][readout]["abundance"],
					data[condition][readout]["specificity"],
					data[condition][readout]["score"],
				),
			)
		}
	}
	file.WriteString(buffer.String())
}
