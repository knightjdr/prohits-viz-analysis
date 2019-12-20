package color

import (
	"fmt"
	"math"
)

// convertHSLtoSpace take an HSL value between 0 - 1 and
// converts to range from 0 - 255 and to hexadecimal.
func convertHSLtoSpace(hsl HSL) Space {
	rgb := convertHSLtoRGB(hsl)

	r := int(math.Round(rgb.r * 255))
	g := int(math.Round(rgb.g * 255))
	b := int(math.Round(rgb.b * 255))
	return Space{
		Hex: fmt.Sprintf("#%02x%02x%02x", r, g, b),
		RGB: []int{r, g, b},
	}
}

func convertHSLtoRGB(hsl HSL) RGB {
	if hsl.s == 0 { // Achromatic.
		return RGB{
			r: hsl.l,
			g: hsl.l,
			b: hsl.l,
		}
	}

	var q float64
	if hsl.l < 0.5 {
		q = hsl.l * (float64(1) + hsl.s)
	} else {
		q = hsl.l + hsl.s - (hsl.l * hsl.s)
	}
	p := (float64(2) * hsl.l) - q
	return RGB{
		r: createRGBComponentFromHSL(p, q, hsl.h+onethird),
		g: createRGBComponentFromHSL(p, q, hsl.h),
		b: createRGBComponentFromHSL(p, q, hsl.h-onethird),
	}
}

func createRGBComponentFromHSL(p, q, t float64) float64 {
	u := t
	if u < 0 {
		u++
	} else if u > 1 {
		u--
	}

	if u < sixth {
		return p + ((q - p) * float64(6) * u)
	}
	if u < 0.5 {
		return q
	}
	if u < twothirds {
		return p + ((q - p) * (twothirds - u) * float64(6))
	}
	return p
}
