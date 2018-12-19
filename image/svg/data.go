package svg

import (
	"github.com/knightjdr/prohits-viz-analysis/typedef"
)

// Data defines the type and variables required for generating a minimap
type Data struct {
	Annotations typedef.Annotations
	Filename    string
	ImageType   string
	Markers     typedef.Markers
	Matrices    *typedef.Matrices
	Minimap     bool
	Parameters  typedef.Parameters
}
