package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradient(t *testing.T) {
	// TEST1: create svg.
	want := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"80\" viewBox=\"0 0 200 80\">\n" +
		"\t<text y=\"20\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Title - test</text>\n" +
		"\t<g>\n" +
		"\t\t<rect fill=\"#ffffff\" y=\"30\" x=\"25.000000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"30\" x=\"38.640000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#99b3ff\" y=\"30\" x=\"52.280000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#668cff\" y=\"30\" x=\"65.920000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#3366ff\" y=\"30\" x=\"79.560000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#0040ff\" y=\"30\" x=\"93.200000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#0033cc\" y=\"30\" x=\"106.840000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#002699\" y=\"30\" x=\"120.480000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#001966\" y=\"30\" x=\"134.120000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000d33\" y=\"30\" x=\"147.760000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"30\" x=\"161.400000\" width=\"13.640000\" height=\"20\" />\n" +
		"\t</g>\n" +
		"\t<rect fill=\"none\" y=\"30\" x=\"25\" width=\"150\" height=\"20\" stroke=\"#000000\" stroke-width=\"1\"/>\n" +
		"\t<text y=\"65\" x=\"175\" font-size=\"12\" text-anchor=\"middle\">1</text>\n" +
		"\t<text y=\"65\" x=\"25\" font-size=\"12\" text-anchor=\"middle\">0</text>\n" +
		"</svg>\n"
	assert.Equal(t, want, Gradient("blueBlack", "Title - test", 11, 0, 1, false), "Gradient svg is not correct")
}
