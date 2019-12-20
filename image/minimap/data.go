package minimap

import (
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Data defines the type and variables required for generating a minimap
type Data struct {
	Filename   string
	ImageType  string
	Matrices   *typedef.Matrices
	Parameters typedef.Parameters
}