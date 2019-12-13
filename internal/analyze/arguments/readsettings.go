package arguments

import (
	"encoding/json"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/spf13/afero"
)

func readSettings(analysisType, filename string) (settings interface{}) {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)

	bytes, err := afero.ReadAll(file)
	log.CheckError(err, true)

	if analysisType == "dotplot" {
		settings = &types.Dotplot{}
		json.Unmarshal(bytes, settings)
	}

	return settings
}
