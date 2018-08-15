package main

import (
	"encoding/json"
	"flag"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/spf13/afero"
)

type Data struct {
	EdgeColor        string  `json:"edgeColor,omitempty"`
	FillColor        string  `json:"fillColor"`
	ImageType        string  `json:"imageType"`
	Invert           bool    `json:"invertColor"`
	MaximumAbundance float64 `json:"abundanceCap"`
	Rows             [][]Row `json:"rows"`
	PrimaryFilter    float64 `json:"primaryFilter,omitempty"`
	ScoreType        string  `json:'scoreType'`
	SecondaryFilter  float64 `json:"secondaryFilter,omitempty"`
}

type Row struct {
	Ratio float64 `json:"ratio,omitempty"`
	Score float64 `json:"score,omitempty"`
	Value float64 `json:"value"`
}

// ParseFlags parses the command line arguments
func ParseJson() (data *Data, err error) {
	// Command line flag arguments.
	jsonFile := flag.String("json", "", "JSON file")
	flag.Parse()

	// Open JSON.
	byteValue, err := afero.ReadFile(fs.Instance, *jsonFile)

	// Parse JSON.
	data = &Data{
		Rows: [][]Row{},
	}
	err = json.Unmarshal(byteValue, data)
	return
}
