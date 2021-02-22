package scv

import (
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/mapf"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeMaps(idMaps map[string]map[string]string, settings types.Settings) {
	writeMapFile(idMaps["condition"], fmt.Sprintf("other/map-%s.txt", settings.Condition))
	writeMapFile(idMaps["readout"], fmt.Sprintf("other/map-%s.txt", settings.Readout))
}

func writeMapFile(idMap map[string]string, filename string) {
	file, err := fs.Instance.Create(filename)
	log.CheckError(err, false)
	if err != nil {
		return
	}
	defer file.Close()

	ids := mapf.KeysStringString(idMap)
	sort.Strings(ids)

	var buffer strings.Builder
	for _, id := range ids {
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%s\n",
				id,
				idMap[id],
			),
		)
	}

	file.WriteString(buffer.String())
}
