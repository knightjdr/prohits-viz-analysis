package frontend

// GetColumnNames parses column names based on input order.
func GetColumnNames(db []string, order []int) []string {
	columns := make([]string, len(order))
	for i, columnIndex := range order {
		columns[i] = db[columnIndex]
	}

	return columns
}

// GetRowNames parses rows names based on input order.
func GetRowNames(db []Row, order []int) []string {
	rows := make([]string, len(order))
	for i, rowIndex := range order {
		rows[i] = db[rowIndex].Name
	}

	return rows
}
