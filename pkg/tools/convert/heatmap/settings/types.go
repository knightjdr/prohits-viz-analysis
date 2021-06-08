package settings

type jsonSettings struct {
	AbundanceColumn string  `json:"abundance,omitempty"`
	InvertColor     int     `json:"invert,omitempty"`
	PrimaryFilter   float64 `json:"primary,omitempty"`
	SecondaryFilter float64 `json:"secondary,omitempty"`
	ScoreColumn     string  `json:"score"`
	ScoreType       int     `json:"filterType,omitempty"`
	Type            string  `json:"type,omitempty"`
	XLabel          string  `json:"xAxis,omitempty"`
	YLabel          string  `json:"yAxis,omitempty"`
}
