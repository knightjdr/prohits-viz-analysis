package scatter

type jsonSettings struct {
	Bait            string  `json:"bait,omitempty"`
	Plot            string  `json:"tool,omitempty"`
	PrimaryFilter   float64 `json:"primary,omitempty"`
	Score           string  `json:"score,omitempty"`
	ScoreType       int     `json:"filter,omitempty"`
	SecondaryFilter float64 `json:"secondary,omitempty"`
	XLabel          string  `json:"xAxis,omitempty"`
	YLabel          string  `json:"yAxis,omitempty"`
}
