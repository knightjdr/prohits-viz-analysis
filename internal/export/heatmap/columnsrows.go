package heatmap

func getColumnsAndRows(data *heatmap) ([]string, []string) {
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
