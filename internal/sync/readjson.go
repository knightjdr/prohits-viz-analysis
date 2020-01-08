package sync

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/spf13/afero"
)

func readJSON(filename string) *minimap {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)

	bytes, err := afero.ReadAll(file)
	log.CheckError(err, true)

	data := &minimap{}
	json.Unmarshal(bytes, data)

	return data
}
