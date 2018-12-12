package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHSLtoHex(t *testing.T) {
	testHSL := []map[string]float64{
		{"h": float64(225) / float64(360), "s": 1, "l": 0.4},
		{"h": float64(115) / float64(360), "s": 0, "l": 0.67},
		{"h": float64(183) / float64(360), "s": 0.23, "l": 0.67},
		{"h": float64(324) / float64(360), "s": 0.52, "l": 0.77},
		{"h": float64(28) / float64(360), "s": 0.52, "l": 0.19},
	}

	// TEST: test several colors conversions. This test will test both HSLtoHex
	// and HuetoRGB.
	wantHex := []string{
		"#0033cc",
		"#ababab",
		"#97bcbe",
		"#e3a6ca",
		"#4a2f17",
	}
	for i := range testHSL {
		assert.Equal(t, wantHex[i], HSLtoHex(testHSL[i]), "HSL color not converted correctly")
	}
}

func TestColor(t *testing.T) {
	// TEST: test generation of blueBlack (default) color gradient.
	want := []string{
		"#ffffff",
		"#ccd9ff",
		"#99b3ff",
		"#668cff",
		"#3366ff",
		"#0040ff",
		"#0033cc",
		"#002699",
		"#001966",
		"#000d33",
		"#000000",
	}
	assert.Equal(t, want, colorGradient("blueBlack", 11, false), "Blue (default) color gradient is not correct")

	// TEST: test generation of greenBlack color gradient.
	want = []string{
		"#ffffff",
		"#ccffcc",
		"#99ff99",
		"#66ff66",
		"#33ff33",
		"#00ff00",
		"#00cc00",
		"#009900",
		"#006600",
		"#003300",
		"#000000",
	}
	assert.Equal(t, want, colorGradient("greenBlack", 11, false), "Green color gradient is not correct")

	// TEST: test generation of greyscale color gradient.
	want = []string{
		"#ffffff",
		"#e6e6e6",
		"#cccccc",
		"#b3b3b3",
		"#999999",
		"#808080",
		"#666666",
		"#4d4d4d",
		"#333333",
		"#1a1a1a",
		"#000000",
	}
	assert.Equal(t, want, colorGradient("greyscale", 11, false), "Grey color gradient is not correct")

	// TEST: test generation of redBlack color gradient.
	want = []string{
		"#ffffff",
		"#ffcccc",
		"#ff9999",
		"#ff6666",
		"#ff3333",
		"#ff0000",
		"#cc0000",
		"#990000",
		"#660000",
		"#330000",
		"#000000",
	}
	assert.Equal(t, want, colorGradient("redBlack", 11, false), "Red color gradient is not correct")

	// TEST: test generation of yellowBlack color gradient.
	want = []string{
		"#ffffff",
		"#ffffcc",
		"#ffff99",
		"#ffff66",
		"#ffff33",
		"#ffff00",
		"#cccc00",
		"#999900",
		"#666600",
		"#333300",
		"#000000",
	}
	assert.Equal(t, want, colorGradient("yellowBlack", 11, false), "Yellow color gradient is not correct")

	// TEST: test generation of blueYellow color gradient.
	want = []string{
		"#0040ff",
		"#3366ff",
		"#668cff",
		"#99b3ff",
		"#ccd9ff",
		"#ffffff",
		"#ffffcc",
		"#ffff99",
		"#ffff66",
		"#ffff33",
		"#ffff00",
	}
	assert.Equal(t, want, colorGradient("blueYellow", 11, false), "Blue-yellow color gradient is not correct")

	// TEST: test generation of blueRed color gradient.
	want = []string{
		"#0040ff",
		"#3366ff",
		"#668cff",
		"#99b3ff",
		"#ccd9ff",
		"#ffffff",
		"#ffcccc",
		"#ff9999",
		"#ff6666",
		"#ff3333",
		"#ff0000",
	}
	assert.Equal(t, want, colorGradient("blueRed", 11, false), "Blue-red color gradient is not correct")

	// TEST: gradient inversion.
	want = []string{
		"#000000",
		"#000d33",
		"#001966",
		"#002699",
		"#0033cc",
		"#0040ff",
		"#3366ff",
		"#668cff",
		"#99b3ff",
		"#ccd9ff",
		"#ffffff",
	}
	assert.Equal(t, want, colorGradient("blueBlack", 11, true), "Gradient inversion is not correct")
}
