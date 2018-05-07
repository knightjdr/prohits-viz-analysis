package svg

import (
	"fmt"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// ColorGradient defines a color gradient to use for fill values. It defines
// the colors to use via HSL and then converts those to HEX.
func ColorGradient(colorSpace string, numColors int) (gradient []string) {

	// Create hex gradient.
	gradient = make([]string, numColors)
	if colorSpace == "blueBlack" {
		// Middle (blue) HSL value = (225, 100%, 50%), but using 0-1 scale for each
		increment := 0.9 / float64(numColors-1)
		startL := 0.95
		for i := 0; i < numColors; i++ {
			lightness := helper.Round(startL-(float64(i)*increment), 0.0001)
			gradient[i] = HSLtoHex(map[string]float64{"h": 0.625, "s": 1, "l": lightness})
		}
	}
	return
}

// HSLtoHex converts hsl colors to rgb to hex. Takes HSL values between 0 - 1 and
// converts to range from 0 - 255, then converts to hex.
const onethird float64 = float64(1) / float64(3)

func HSLtoHex(hsl map[string]float64) (hex string) {
	var r, g, b float64

	if hsl["s"] == 0 { // Achromatic.
		r = hsl["l"]
		g = hsl["l"]
		b = hsl["l"]
	} else {
		var q float64
		if hsl["l"] < 0.5 {
			q = hsl["l"] * (float64(1) + hsl["s"])
		} else {
			q = hsl["l"] + hsl["s"] - (hsl["l"] * hsl["s"])
		}
		p := (float64(2) * hsl["l"]) - q
		r = HuetoRGB(p, q, hsl["h"]+onethird)
		g = HuetoRGB(p, q, hsl["h"])
		b = HuetoRGB(p, q, hsl["h"]-onethird)
	}
	rgb := map[string]int64{
		"r": int64(math.Round(r * 255)),
		"g": int64(math.Round(g * 255)),
		"b": int64(math.Round(b * 255)),
	}
	hex = fmt.Sprintf("#%02x%02x%02x", rgb["r"], rgb["g"], rgb["b"])
	return
}

// HuetoRGB converts a hue to rgb.
const sixth float64 = float64(1) / float64(6)
const twothirds float64 = float64(2) / float64(3)

func HuetoRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	} else if t > 1 {
		t -= 1
	}
	if t < sixth {
		return p + ((q - p) * float64(6) * t)
	} else if t < 0.5 {
		return q
	} else if t < twothirds {
		return p + ((q - p) * (twothirds - t) * float64(6))
	}
	return p
}
