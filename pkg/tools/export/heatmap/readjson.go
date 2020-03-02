package heatmap

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/spf13/afero"
)

// ReadJSON heatmap information for exporting as image.
func ReadJSON(filename string) *Heatmap {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)

	bytes, err := afero.ReadAll(file)
	log.CheckError(err, true)

	data := &Heatmap{}
	json.Unmarshal(bytes, data)

	return data
}
