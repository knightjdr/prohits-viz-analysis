package arguments

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

func readSettings(filename string) *types.Analysis {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)

	bytes, err := afero.ReadAll(file)
	log.CheckError(err, true)

	settings := &types.Settings{}
	json.Unmarshal(bytes, settings)

	return &types.Analysis{
		Settings: *settings,
	}
}
