package sync

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/frontend"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

type minimap struct {
	ColumnDB    []string
	ColumnOrder []int
	ImageType   string
	RowDB       []frontend.Row
	RowOrder    []int
	Settings    types.Settings
}
