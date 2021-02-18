package geneid

// MapByColumn maps identifiers from one column to another. If an id is missing in the map
// column, the source id is mapped to itself.
func MapByColumn(data []map[string]string, sourceColumn, mapColumn string) map[string]string {
	mapped := make(map[string]string, 0)

	for _, row := range data {
		source := row[sourceColumn]
		if _, ok := mapped[source]; !ok {
			target := row[mapColumn]
			if target != "" {
				mapped[source] = target
			} else {
				mapped[source] = source
			}
		}
	}

	return mapped
}
