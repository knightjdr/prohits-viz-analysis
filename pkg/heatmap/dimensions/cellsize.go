package dimensions

import "math"

func calculateCellSize(dims *Heatmap) {
	dims.CellSize = int(math.Floor(dims.Ratio * float64(idealCellSize)))
}
