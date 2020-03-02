// Package color has functions for create color scales and color transformations.
package color

const onethird float64 = float64(1) / float64(3)
const sixth float64 = float64(1) / float64(6)
const twothirds float64 = float64(2) / float64(3)

// Space to store Hex and RGB values.
type Space struct {
	Hex string
	RGB []int
}

// HSL color representation.
type HSL struct {
	h float64
	s float64
	l float64
}

// RGB tuple.
type RGB struct {
	b float64
	g float64
	r float64
}
