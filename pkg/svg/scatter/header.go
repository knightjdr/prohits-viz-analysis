package scatter

import (
	"fmt"
)

func writeHeader(s *Scatter, writeString func(string)) {
	str := fmt.Sprintf(
		"<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\""+
			" xml:space=\"preserve\" width=\"%[1]d\" height=\"%[1]d\" viewBox=\"0 0 %[1]d %[1]d\">\n",
		int(s.PlotSize),
	)
	writeString(str)
}
