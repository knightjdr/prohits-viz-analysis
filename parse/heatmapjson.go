// Package parse contains functions for parsing JSON objects.
package parse

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

// Data is the struct for the parsed JSON of heatmap/dotplot objects.
type Data struct {
	Annotations      typedef.Annotations `json:"annotations,omitempty"`
	Columns          []string            `json:"columns,omitempty"`
	EdgeColor        string              `json:"edgeColor,omitempty"`
	FillColor        string              `json:"fillColor"`
	ImageType        string              `json:"imageType"`
	Invert           bool                `json:"invertColor"`
	Markers          typedef.Markers     `json:"markers,omitempty"`
	MaximumAbundance float64             `json:"abundanceCap"`
	Rows             []Row               `json:"rows"`
	PrimaryFilter    float64             `json:"primaryFilter,omitempty"`
	ScoreType        string              `json:"scoreType"`
	SecondaryFilter  float64             `json:"secondaryFilter,omitempty"`
}

// Row is the parsed row structure.
type Row struct {
	Data []Column `json:"data"`
	Name string   `json:"name"`
}

// Column is the parsed columns structure.
type Column struct {
	Ratio float64 `json:"ratio,omitempty"`
	Score float64 `json:"score,omitempty"`
	Value float64 `json:"value"`
}

// HeatmapJSON parses the command line arguments.
func HeatmapJSON(jsonFile string) (data *Data, err error) {
	// Open JSON.
	byteValue, err := afero.ReadFile(fs.Instance, jsonFile)

	// Parse JSON.
	data = &Data{
		Rows: []Row{},
	}
	err = json.Unmarshal(byteValue, data)
	return
}
