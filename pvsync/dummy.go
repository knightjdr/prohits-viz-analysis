package main

// Dummy creates slices filled with bogus column and row names.
func Dummy(colNum, rowNum int) ([]string, []string) {
	cols := make([]string, colNum)
	rows := make([]string, rowNum)

	// Create dummy columns.
	for i := range cols {
		cols[i] = "column"
	}

	// Create dummy rows.
	for i := range rows {
		rows[i] = "row"
	}
	return cols, rows
}
