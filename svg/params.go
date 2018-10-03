package svg

const idealCellSize int = 20
const idealCircleSpace float64 = 1.5
const idealEdgeWidth int = 2
const idealFontSize int = 12

// Min ratio for shrinking an image is based on the idealCellSize.
// At 0.05, it will ensure cells are never less than 1px.
const minRatio float64 = 0.05

// Max image height in pixels assuming 11inch height with 1 inch of margins.
const maxImageHeight int = 3000

// Max image width in pixels assuming 8.5inch width with 1 inch of margins.
const maxImageWidth int = 2250
