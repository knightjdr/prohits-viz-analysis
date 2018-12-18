package svg

import (
	"fmt"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/helper"
)

// Define bidirectional color scales
var twoColor = []string{"blueYellow", "blueRed"}

// biDirection creates a two color hex gradient. The starting color is set using the
// hue and the first half of the gradient is populated with colors by increasing the
// lightness from 0.5 to 1. The middle color of the gradient is set to white and
// the remaining portion of the gradient is created for the end color (also
// defined by its hue) by decreasing the lightness from 1 to 0.5.
func biDirection(colorSpace string, numColors int) (gradient []string) {
	var hEnd, hStart, s float64
	if colorSpace == "blueYellow" {
		// Start (HSL value = (225, 100%, 50%).
		// End (HSL value = (60, 100%, 50%).
		hEnd = float64(60) / float64(360)
		hStart = 0.625
		s = 1
	} else {
		// Default blueRed scale
		// Start (HSL value = (225, 100%, 50%).
		// End (HSL value = (0, 100%, 50%).
		hEnd = 0
		hStart = 0.625
		s = 1
	}
	halfColors := (numColors - 1) / 2
	increment := 1.00 / float64(numColors-1)
	startL := .50
	gradient = make([]string, numColors)
	for i := 0; i < halfColors; i++ {
		lightness := helper.Round(startL+(float64(i)*increment), 0.0001)
		gradient[i] = HSLtoHex(map[string]float64{"h": hStart, "s": s, "l": lightness})
	}
	gradient[halfColors] = HSLtoHex(map[string]float64{"h": 0, "s": s, "l": 1})
	startL = 1.00
	startIndex := halfColors + 1
	for i := 0; i < halfColors; i++ {
		lightness := helper.Round(startL-(float64(i+1)*increment), 0.0001)
		gradient[i+startIndex] = HSLtoHex(map[string]float64{"h": hEnd, "s": s, "l": lightness})
	}
	return
}

// monoDirection creates a single direction hex gradient. The color scale is set using the hue
// and saturation components of HSL. The gradient is then defined by changing the lightness
// from 1 (light) to 0 (dark). HSL values are on a 0-1 scale.
// The maximum hue value of 1 equals 360 so all values are relative to that.
func monoDirection(colorSpace string, numColors int) (gradient []string) {
	var h, s float64
	if colorSpace == "greenBlack" {
		// Middle HSL value = (120, 100%, 50%).
		h = float64(120) / float64(360)
		s = 1
	} else if colorSpace == "greyscale" {
		// Middle HSL value = (0, 0%, 50%).
		h = 0
		s = 0
	} else if colorSpace == "redBlack" {
		// Middle HSL value = (0, 100%, 50%).
		h = 0
		s = 1
	} else if colorSpace == "yellowBlack" {
		// Middle HSL value = (60, 100%, 50%).
		h = float64(60) / float64(360)
		s = 1
	} else { // default blueBlack
		// Middle (HSL value = (225, 100%, 50%).
		h = 0.625
		s = 1
	}
	increment := 1.00 / float64(numColors-1)
	startL := 1.00
	gradient = make([]string, numColors)
	for i := 0; i < numColors; i++ {
		lightness := helper.Round(startL-(float64(i)*increment), 0.0001)
		gradient[i] = HSLtoHex(map[string]float64{"h": h, "s": s, "l": lightness})
	}
	return
}

// colorGradient defines a color gradient to use for fill values. It defines
// the colors to use via HSL and then converts those to HEX.
func colorGradient(colorSpace string, numColors int, invertColor bool) (gradient []string) {
	if helper.SliceContains(colorSpace, twoColor) {
		gradient = biDirection(colorSpace, numColors)
	} else {
		gradient = monoDirection(colorSpace, numColors)
	}
	// Invert gradient if requested.
	if invertColor {
		for i, j := 0, numColors-1; i < j; i, j = i+1, j-1 {
			gradient[i], gradient[j] = gradient[j], gradient[i]
		}
	}
	return
}

const onethird float64 = float64(1) / float64(3)

// HSLtoHex converts hsl colors to rgb to hex. Takes HSL values between 0 - 1 and
// converts to range from 0 - 255, then converts to hex.
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

const sixth float64 = float64(1) / float64(6)
const twothirds float64 = float64(2) / float64(3)

// HuetoRGB converts a hue to rgb.
func HuetoRGB(p, q, t float64) float64 {
	u := t
	if u < 0 {
		u++
	} else if u > 1 {
		u--
	}
	if u < sixth {
		return p + ((q - p) * float64(6) * u)
	} else if u < 0.5 {
		return q
	} else if u < twothirds {
		return p + ((q - p) * (twothirds - u) * float64(6))
	}
	return p
}
