package heatmap

// GetColumnsAndRows parses column and row names from input data.
func GetColumnsAndRows(data *Heatmap) ([]string, []string) {
	columns := make([]string, len(data.ColumnOrder))
	for i, columnIndex := range data.ColumnOrder {
		columns[i] = data.ColumnDB[columnIndex]
	}

	rows := make([]string, len(data.RowOrder))
	for i, rowIndex := range data.RowOrder {
		rows[i] = data.RowDB[rowIndex].Name
	}

	return columns, rows
}
