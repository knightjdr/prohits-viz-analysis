package geneid

import "strings"

// MapByColumn maps identifiers from one column to another. If an id is missing in the map
// column, the source id is mapped to itself. When the source is mapped to itself, only
// the characters leading to the first underscore are used. This can also be done with
// the target by making trimTarget true.
func MapByColumn(data []map[string]string, sourceColumn, mapColumn string, trimTarget bool) map[string]string {
	mapped := make(map[string]string, 0)

	parse := defineTargetParser(trimTarget)
	for _, row := range data {
		source := row[sourceColumn]
		if _, ok := mapped[source]; !ok {
			target := row[mapColumn]
			if target != "" {
				mapped[source] = parse(target)
			} else {
				mapped[source] = strings.Split(source, "_")[0]
			}
		}
	}

	return mapped
}

func defineTargetParser(trimTarget bool) func(string) string {
	if trimTarget {
		return func(target string) string {
			return strings.Split(target, "_")[0]
		}
	}

	return func(target string) string {
		return target
	}
}
