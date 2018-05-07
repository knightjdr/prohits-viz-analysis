package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	testHSL := []map[string]float64{
		{"h": float64(225) / float64(360), "s": 1, "l": 0.4},
		{"h": float64(115) / float64(360), "s": 0, "l": 0.67},
		{"h": float64(183) / float64(360), "s": 0.23, "l": 0.67},
		{"h": float64(324) / float64(360), "s": 0.52, "l": 0.77},
		{"h": float64(28) / float64(360), "s": 0.52, "l": 0.19},
	}

	// TEST1: test several colors conversions. This test will test both HSLtoHex
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

	// TEST2: test generation of color gradient.
	want := []string{
		"#e6ecff",
		"#b3c6ff",
		"#809fff",
		"#4d79ff",
		"#1a53ff",
		"#0039e6",
		"#002db3",
		"#002080",
		"#00134d",
		"#00061a",
	}
	assert.Equal(t, want, ColorGradient("blueBlack", 10), "Blue color gradient is not correct")
}
