package color

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

type biDirectionSettings struct {
	hueEnd     float64
	hueStart   float64
	saturation float64
}

// createBiDirectionGradient creates a two color hex gradient. The starting color is set using the
// hue and the first half of the gradient is populated with colors by increasing the
// lightness from 0.5 to 1. The middle color of the gradient is set to white and
// the remaining portion of the gradient is created for the end color (also
// defined by its hue) by decreasing the lightness from 1 to 0.5.
func createBiDirectionGradient(settings *Gradient) []Space {
	gradient := make([]Space, settings.NumColors)
	halfColors := (settings.NumColors - 1) / 2
	increment := 1.00 / float64(settings.NumColors-1)
	hslSettings := defineBiDirectionSettings(settings.ColorSpace)

	startL := .50
	for i := 0; i < halfColors; i++ {
		lightness := math.Round(startL+(float64(i)*increment), 0.0001)
		gradient[i] = convertHSLtoSpace(HSL{h: hslSettings.hueStart, s: hslSettings.saturation, l: lightness})
	}
	gradient[halfColors] = convertHSLtoSpace(HSL{h: 0, s: hslSettings.saturation, l: 1})

	startL = 1.00
	startIndex := halfColors + 1
	for i := 0; i < halfColors; i++ {
		lightness := math.Round(startL-(float64(i+1)*increment), 0.0001)
		gradient[i+startIndex] = convertHSLtoSpace(HSL{h: hslSettings.hueEnd, s: hslSettings.saturation, l: lightness})
	}

	return gradient
}

func defineBiDirectionSettings(colorSpace string) biDirectionSettings {
	if colorSpace == "blueYellow" {
		// Start (HSL value = (225, 100%, 50%).
		// End (HSL value = (60, 100%, 50%).
		return biDirectionSettings{
			hueEnd:     float64(60) / float64(360),
			hueStart:   0.625,
			saturation: 1,
		}
	}
	// Default blueRed scale
	// Start (HSL value = (225, 100%, 50%).
	// End (HSL value = (0, 100%, 50%).
	return biDirectionSettings{
		hueEnd:     0,
		hueStart:   0.625,
		saturation: 1,
	}
}
