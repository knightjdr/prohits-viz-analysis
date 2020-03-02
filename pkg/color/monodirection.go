package color

import "github.com/knightjdr/prohits-viz-analysis/pkg/math"

type monoDirectionSettings struct {
	hue        float64
	saturation float64
}

// createMonoDirectionGradiant creates a single direction hex gradient. The color scale is set using the hue
// and saturation components of HSL. The gradient is then defined by changing the lightness
// from 1 (light) to 0 (dark). HSL values are on a 0-1 scale.
// The maximum hue value of 1 equals 360 so all values are relative to that.
func createMonoDirectionGradiant(settings *Gradient) []Space {
	gradient := make([]Space, settings.NumColors)
	increment := 1.00 / float64(settings.NumColors-1)
	hslSettings := defineMonoDirectionSettings(settings.ColorSpace)

	startL := 1.00
	for i := 0; i < settings.NumColors; i++ {
		lightness := math.Round(startL-(float64(i)*increment), 0.0001)
		gradient[i] = convertHSLtoSpace(HSL{h: hslSettings.hue, s: hslSettings.saturation, l: lightness})
	}
	return gradient
}

func defineMonoDirectionSettings(colorSpace string) monoDirectionSettings {
	switch colorSpace {
	case "green":
		// Middle HSL value = (120, 100%, 50%).
		return monoDirectionSettings{
			hue:        float64(120) / float64(360),
			saturation: 1,
		}
	case "grey":
		// Middle HSL value = (0, 0%, 50%).
		return monoDirectionSettings{
			hue:        0,
			saturation: 0,
		}
	case "red":
		// Middle HSL value = (0, 100%, 50%).
		return monoDirectionSettings{
			hue:        0,
			saturation: 1,
		}
	case "yellow":
		// Middle HSL value = (60, 100%, 50%).
		return monoDirectionSettings{
			hue:        float64(60) / float64(360),
			saturation: 1,
		}
	default:
		// Middle blue (HSL value = (225, 100%, 50%).
		return monoDirectionSettings{
			hue:        0.625,
			saturation: 1,
		}
	}
}
