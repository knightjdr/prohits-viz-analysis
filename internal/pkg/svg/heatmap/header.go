package heatmap

import (
	"fmt"
)

// WriteHeader writes an opening svg tag with dimensions.
func WriteHeader(h *Heatmap, writeString func(string)) {
	str := fmt.Sprintf(
		"<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\""+
			" xml:space=\"preserve\" width=\"%d\" height=\"%d\" viewBox=\"0 0 %d %d\">\n",
		h.SvgWidth, h.SvgHeight, h.SvgWidth, h.SvgHeight,
	)
	writeString(str)
}
