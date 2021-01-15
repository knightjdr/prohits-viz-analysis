package heatmap

import (
	"fmt"
)

// WriteHeader writes an opening svg tag with dimensions.
func WriteHeader(h *Heatmap, writeString func(string)) {
	str := fmt.Sprintf(
		"<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\""+
			" xml:space=\"preserve\" width=\"%[1]d\" height=\"%[2]d\" viewBox=\"0 0 %[1]d %[2]d\">\n",
		h.SvgWidth, h.SvgHeight,
	)
	writeString(str)
}
