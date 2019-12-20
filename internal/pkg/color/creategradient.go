package color

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
)

// Gradient type.
type Gradient struct {
	ColorSpace string
	ColorType  string // Hex or RGB
	Invert     bool
	NumColors  int
}

// InitializeGradient initializes a color gradient.
func InitializeGradient() *Gradient {
	gradient := &Gradient{
		ColorSpace: "blue",
		ColorType:  "Hex",
		Invert:     false,
		NumColors:  101,
	}

	return gradient
}

// CreateColorGradient creates a color gradient.
func (g *Gradient) CreateColorGradient() (gradient []Space) {
	var twoColor = []string{"blueYellow", "blueRed"}

	if slice.ContainsString(g.ColorSpace, twoColor) {
		gradient = createBiDirectionGradient(g)
	} else {
		gradient = createMonoDirectionGradiant(g)
	}

	if g.Invert {
		gradient = reverseGradient(gradient)
	}

	return
}

func reverseGradient(gradient []Space) []Space {
	reversed := gradient

	for i, j := 0, len(gradient)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return reversed
}
