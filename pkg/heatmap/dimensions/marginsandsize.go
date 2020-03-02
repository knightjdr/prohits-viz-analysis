package dimensions

import "math"

func calculateMarginsAndSize(dims *Heatmap, columns, rows []string) {
	dims.FontSize = int(math.Floor(dims.Ratio * float64(idealFontSize)))

	// Calculate required top margin. Find the longest column name and assume it
	// is made entirely of the "W" character (which has a width of 11.33px
	// in arial with a 12pt fontsize).
	longestColumnName := getLongestName(columns)
	dims.TopMargin = int(math.Round(float64(longestColumnName) * 11.33 * dims.Ratio))

	longestRowName := getLongestName(rows)
	dims.LeftMargin = int(math.Round(float64(longestRowName) * 11.33 * dims.Ratio))

	dims.PlotHeight = dims.RowNumber * dims.CellSize
	dims.PlotWidth = dims.ColumnNumber * dims.CellSize
	dims.SvgHeight = dims.TopMargin + dims.PlotHeight
	dims.SvgWidth = dims.LeftMargin + dims.PlotWidth
}

func getLongestName(names []string) int {
	longestName := 0
	for _, name := range names {
		nameLength := len([]rune(name))
		if nameLength > longestName {
			longestName = nameLength
		}
	}

	return longestName
}
