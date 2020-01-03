// Package svg has functions for generating and converting svg files.
package svg

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/svg/convert"

// ConvertToPNG converts an svg to a png using rsvg.
var ConvertToPNG = convert.RSVG
