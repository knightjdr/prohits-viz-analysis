package circheatmap

import (
	"fmt"
)

func writeHeader(c *CircHeatmap, writeString func(string)) {
	str := fmt.Sprintf(
		"<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\""+
			" xml:space=\"preserve\" width=\"%[2]d\" height=\"%[2]d\" viewBox=\"-%[1]d -%[1]d %[2]d %[2]d\">\n",
		int(c.Dimensions.Center),
		int(c.Dimensions.PlotSize),
	)
	writeString(str)
}
