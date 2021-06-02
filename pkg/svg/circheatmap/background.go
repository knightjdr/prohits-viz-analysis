package circheatmap

import "fmt"

func writeBackground(c *CircHeatmap, writeString func(string)) {
	str := fmt.Sprintf(
		"\t<rect width=\"100%%\" height=\"100%%\" fill=\"white\" transform=\"translate(-%[1]d -%[1]d)\"/>\n",
		int(c.Dimensions.Center),
	)
	writeString(str)
}
