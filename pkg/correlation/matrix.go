// Package correlation can calculate the correlation between columns or rows of a matrix.
package correlation

// Matrix correlation settings.
type Matrix struct {
	Data                    *[][]float64
	Dimension               string
	IgnoreSourceTargetPairs bool
	Labels                  []string
	Method                  string
}

// Correlate calculates the correlation between the columns or rows.
func (m *Matrix) Correlate() [][]float64 {
	return [][]float64{}
}
