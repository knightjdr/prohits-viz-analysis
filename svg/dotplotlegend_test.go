package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDotplotLegend(t *testing.T) {
	// TEST1: draw legend svg with lte as score type.
	want := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"220\" viewBox=\"0 0 200 220\">\n" +
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
		"\t<text y=\"65\" x=\"175\" font-size=\"12\" text-anchor=\"middle\">50</text>\n" +
		"\t<text y=\"65\" x=\"25\" font-size=\"12\" text-anchor=\"middle\">0</text>\n" +
		"\t<g>\n" +
		"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"60\" r=\"6\" />\n" +
		"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"135\" r=\"12\" />\n" +
		"\t\t<line fill=\"none\" stroke=\"#000000\" stroke-width=\"1\" x1=\"70\" y1=\"100\" x2=\"119\" y2=\"100\"/>\n" +
		"\t\t<polygon fill=\"#000000\" points=\"110,96 112,100 110,104 119,100\"/>\n" +
		"\t\t<text y=\"130\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Relative abundance</text>\n" +
		"\t</g>\n" +
		"\t<g>\n" +
		"\t\t<circle fill=\"none\" cy=\"165\" cx=\"50\" r=\"12\" stroke=\"#000000\" stroke-width=\"2\" />\n" +
		"\t\t<text y=\"195\" x=\"50\" font-size=\"12\" text-anchor=\"middle\">≤ 0.01</text>\n" +
		"\t\t<circle fill=\"none\" cy=\"165\" cx=\"100\" r=\"12\" stroke=\"#0040ff\" stroke-width=\"2\" />\n" +
		"\t\t<text y=\"195\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">≤ 0.05</text>\n" +
		"\t\t<circle fill=\"none\" cy=\"165\" cx=\"150\" r=\"12\" stroke=\"#99b3ff\" stroke-width=\"2\" />\n" +
		"\t\t<text y=\"195\" x=\"150\" font-size=\"12\" text-anchor=\"middle\">> 0.05</text>\n" +
		"\t</g>\n" +
		"</svg>\n"
	assert.Equal(
		t,
		want,
		DotplotLegend("blueBlack", "Title - test", 11, 0, 50, 0.01, 0.05, "lte"),
		"Dotplot legend is not correct for lte score type",
	)

	// TEST2: draw legend svg with gte as score type.
	want = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"200\" height=\"220\" viewBox=\"0 0 200 220\">\n" +
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
		"\t<text y=\"65\" x=\"175\" font-size=\"12\" text-anchor=\"middle\">50</text>\n" +
		"\t<text y=\"65\" x=\"25\" font-size=\"12\" text-anchor=\"middle\">0</text>\n" +
		"\t<g>\n" +
		"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"60\" r=\"6\" />\n" +
		"\t\t<circle fill=\"#000000\" cy=\"100\" cx=\"135\" r=\"12\" />\n" +
		"\t\t<line fill=\"none\" stroke=\"#000000\" stroke-width=\"1\" x1=\"70\" y1=\"100\" x2=\"119\" y2=\"100\"/>\n" +
		"\t\t<polygon fill=\"#000000\" points=\"110,96 112,100 110,104 119,100\"/>\n" +
		"\t\t<text y=\"130\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">Relative abundance</text>\n" +
		"\t</g>\n" +
		"\t<g>\n" +
		"\t\t<circle fill=\"none\" cy=\"165\" cx=\"50\" r=\"12\" stroke=\"#000000\" stroke-width=\"2\" />\n" +
		"\t\t<text y=\"195\" x=\"50\" font-size=\"12\" text-anchor=\"middle\">≥ 0.01</text>\n" +
		"\t\t<circle fill=\"none\" cy=\"165\" cx=\"100\" r=\"12\" stroke=\"#0040ff\" stroke-width=\"2\" />\n" +
		"\t\t<text y=\"195\" x=\"100\" font-size=\"12\" text-anchor=\"middle\">≥ 0.05</text>\n" +
		"\t\t<circle fill=\"none\" cy=\"165\" cx=\"150\" r=\"12\" stroke=\"#99b3ff\" stroke-width=\"2\" />\n" +
		"\t\t<text y=\"195\" x=\"150\" font-size=\"12\" text-anchor=\"middle\">< 0.05</text>\n" +
		"\t</g>\n" +
		"</svg>\n"
	assert.Equal(
		t,
		want,
		DotplotLegend("blueBlack", "Title - test", 11, 0, 50, 0.01, 0.05, "gte"),
		"Dotplot legend is not correct for gte score type",
	)
}
