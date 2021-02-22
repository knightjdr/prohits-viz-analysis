package geneid

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/read/csv"
)

// MapByFile maps identifiers from one column to mapping specifed in a file. If an id is missing in the map
// file, the source id is mapped to itself. When the source is mapped to itself, only
// the characters leading to the first underscore are used. This can also be done with
// the target by making trimTarget true.
func MapByFile(data []map[string]string, sourceColumn, mapFile string) map[string]string {
	mapped := make(map[string]string, 0)

	fileMapping := csv.ReadToMap(mapFile, '\t')

	for _, row := range data {
		source := row[sourceColumn]
		if _, ok := mapped[source]; !ok {
			target := fileMapping[source]
			if target != "" {
				mapped[source] = target
			} else {
				mapped[source] = strings.Split(source, "_")[0]
			}
		}
	}

	return mapped
}
