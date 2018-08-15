package main

// Create dummy slices for column and row names
func Dummy(colNum, rowNum int) ([]string, []string) {
	cols := make([]string, colNum)
	rows := make([]string, rowNum)

	// Create dummy columns.
	for i, _ := range cols {
		cols[i] = "column"
	}

	// Create dummy rows.
	for i, _ := range rows {
		rows[i] = "row"
	}
	return cols, rows
}
